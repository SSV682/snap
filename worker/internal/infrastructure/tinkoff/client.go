package tinkoff

import (
	"context"
	"errors"
	"fmt"
	"time"
	"worker/internal/entity"

	"github.com/russianinvestments/invest-api-go-sdk/investgo"
	investapi "github.com/russianinvestments/invest-api-go-sdk/proto"
)

type Client struct {
	client      *investgo.Client
	instruments *investgo.InstrumentsServiceClient
	marketData  *investgo.MarketDataServiceClient
	logger      investgo.Logger
}

func NewClient(ctx context.Context, config investgo.Config, logger investgo.Logger) (*Client, error) {
	client, err := investgo.NewClient(ctx, config, logger)
	if err != nil {
		return nil, fmt.Errorf("client creating error %v", err.Error())
	}

	return &Client{
		client:      client,
		instruments: client.NewInstrumentsServiceClient(),
		marketData:  client.NewMarketDataServiceClient(),
		logger:      logger,
	}, nil
}

func (c *Client) Stop() {
	c.logger.Infof("Closing client connection")
	err := c.client.Stop()
	if err != nil {
		c.logger.Errorf("client shutdown error %v", err.Error())
	}
}

func (c *Client) getInstrument(ticker string) (*investapi.InstrumentShort, error) {
	instrResp, err := c.instruments.FindInstrument(ticker)
	if err != nil {
		return nil, fmt.Errorf("find instrument: %v", err)
	}

	ins := instrResp.GetInstruments()

	if len(ins) > 0 {
		return ins[0], nil
	}

	return nil, errors.New("not found instrument")
}

func (c *Client) HistoricCandles(ticker string, timeFrom, timeTo time.Time) ([]entity.Candle, error) {
	instrument, err := c.getInstrument(ticker)
	if err != nil {
		return nil, fmt.Errorf("get instrument: %v", err)
	}

	candles, err := c.marketData.GetHistoricCandles(&investgo.GetHistoricCandlesRequest{
		Instrument: instrument.Uid,
		Interval:   investapi.CandleInterval_CANDLE_INTERVAL_1_MIN,
		From:       timeFrom,
		To:         timeTo,
		File:       false,
		FileName:   "",
	})
	if err != nil {
		return nil, fmt.Errorf("get historic candles: %v", err)
	}

	result := make([]entity.Candle, len(candles))

	for i, candle := range candles {
		result[i] = entity.Candle{
			Open:   candle.GetOpen().ToFloat(),
			High:   candle.GetHigh().ToFloat(),
			Low:    candle.GetLow().ToFloat(),
			Close:  candle.GetClose().ToFloat(),
			Volume: candle.Volume,
			Time:   candle.Time.AsTime(),
		}
	}

	return result, nil
}

func (c *Client) GetTaxFn() entity.TaxFn {
	return func(price float64) float64 { return price * 0.05 / 100 }
}

func (c *Client) GetCurrencies() ([]entity.Instrument, error) {
	resp, err := c.instruments.Currencies(investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE)
	if err != nil {
		return nil, fmt.Errorf("get currencies: %v", err)
	}

	currencies := resp.GetInstruments()
	result := make([]entity.Instrument, len(currencies))

	for i := range currencies {
		result[i] = entity.Instrument{
			Name:   currencies[i].GetName(),
			Figi:   currencies[i].GetFigi(),
			Ticker: currencies[i].GetTicker(),
		}
	}

	return result, nil
}

func (c *Client) GetStocks() ([]entity.Instrument, error) {
	resp, err := c.instruments.Shares(investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE)
	if err != nil {
		return nil, fmt.Errorf("get stocks: %v", err)
	}

	bonds := resp.GetInstruments()
	result := make([]entity.Instrument, len(bonds))

	for i := range bonds {
		result[i] = entity.Instrument{
			Name:   bonds[i].GetName(),
			Figi:   bonds[i].GetFigi(),
			Ticker: bonds[i].GetTicker(),
		}
	}

	return result, nil
}

func (c *Client) GetFutures() ([]entity.Instrument, error) {
	resp, err := c.instruments.Futures(investapi.InstrumentStatus_INSTRUMENT_STATUS_BASE)
	if err != nil {
		return nil, fmt.Errorf("get futures: %v", err)
	}

	instruments := resp.GetInstruments()
	result := make([]entity.Instrument, len(instruments))

	for i := range instruments {
		result[i] = entity.Instrument{
			Name:   instruments[i].GetName(),
			Figi:   instruments[i].GetFigi(),
			Ticker: instruments[i].GetTicker(),
		}
	}

	return result, nil
}
