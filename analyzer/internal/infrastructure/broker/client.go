package broker

import (
	"analyzer/internal/entity"
	"context"
	"errors"
	"fmt"
	"time"

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

// Stop stops the client
func (c *Client) Stop() {
	c.logger.Infof("Closing client connection")
	err := c.client.Stop()
	if err != nil {
		c.logger.Errorf("client shutdown error %v", err.Error())
	}
}

// getInstrument returns the instrument by ticker
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

// HistoricCandles returns the candles for the given ticker and time interval
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

// TODO:
func (c *Client) GetTaxFn() entity.TaxFn {
	return func(price float64) float64 { return price * 0.05 / 100 }
}

// GetCurrencies returns the list of currencies
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
			FIGI:   currencies[i].GetFigi(),
			Ticker: currencies[i].GetTicker(),
		}
	}

	return result, nil
}

// GetStocks returns the list of stocks
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
			FIGI:   bonds[i].GetFigi(),
			Ticker: bonds[i].GetTicker(),
		}
	}

	return result, nil
}

// GetFutures returns the list of futures
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
			FIGI:   instruments[i].GetFigi(),
			Ticker: instruments[i].GetTicker(),
		}
	}

	return result, nil
}

// LastCandle returns the last candles for the given ticker
func (c *Client) LastCandle(ticker string) (entity.Candle, error) {
	instrument, err := c.getInstrument(ticker)
	if err != nil {
		return entity.Candle{}, fmt.Errorf("get instrument: %v", err)
	}

	response, err := c.marketData.GetCandles(instrument.Uid, investapi.CandleInterval_CANDLE_INTERVAL_1_MIN, time.Now().Add(-1*time.Minute), time.Now())
	if err != nil {
		return entity.Candle{}, fmt.Errorf("get historic candles: %v", err)
	}

	//TODO: can be nil, panic
	candles := response.GetCandles()
	if len(candles) == 0 {
		return entity.Candle{}, errors.New("not found candles")
	}

	lastCandle := candles[len(response.GetCandles())-1]

	return entity.Candle{
		Ticker: ticker,
		Open:   lastCandle.GetOpen().ToFloat(),
		High:   lastCandle.GetHigh().ToFloat(),
		Low:    lastCandle.GetLow().ToFloat(),
		Close:  lastCandle.GetClose().ToFloat(),
		Volume: lastCandle.Volume,
		Time:   lastCandle.Time.AsTime(),
	}, nil
}
