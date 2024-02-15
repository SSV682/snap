package broker

import (
	"context"
	"errors"
	"fmt"
	"solver/internal/entity"

	"github.com/russianinvestments/invest-api-go-sdk/investgo"
	investapi "github.com/russianinvestments/invest-api-go-sdk/proto"
	log "github.com/sirupsen/logrus"
)

// Client is the client for the broker.
type Client struct {
	client         *investgo.Client
	accountsCli    *investgo.SandboxServiceClient
	instrumentsCli *investgo.InstrumentsServiceClient
	ordersCli      *investgo.OrdersServiceClient
	operationsCli  *investgo.OperationsServiceClient

	accountUID string

	logger investgo.Logger
}

var (
	orderDirections = map[entity.EventType]investapi.OrderDirection{
		entity.Sell: investapi.OrderDirection_ORDER_DIRECTION_SELL,
		entity.Buy:  investapi.OrderDirection_ORDER_DIRECTION_BUY,
	}
)

// NewClient creates a new broker client.
func NewClient(ctx context.Context, config investgo.Config, logger investgo.Logger) (*Client, error) {
	client, err := investgo.NewClient(ctx, config, logger)
	if err != nil {
		return nil, fmt.Errorf("client creating error %v", err.Error())
	}

	accounts := client.NewSandboxServiceClient()
	instruments := client.NewInstrumentsServiceClient()

	sandboxAccount, err := accounts.OpenSandboxAccount()
	if err != nil {
		return nil, fmt.Errorf("sandbox account opening error %v", err.Error())
	}

	accountID := sandboxAccount.GetAccountId()
	log.Infof("Account ID: %s", accountID)

	if _, err = accounts.SandboxPayIn(&investgo.SandboxPayInRequest{
		AccountId: accountID,
		Unit:      10000,
		Currency:  "RUB",
	}); err != nil {
		return nil, fmt.Errorf("sandbox account pay: %v", err.Error())
	}

	return &Client{
		client:         client,
		accountsCli:    accounts,
		instrumentsCli: instruments,

		accountUID: sandboxAccount.GetAccountId(),

		logger: logger,
	}, nil
}

// Stop stops the client
func (c *Client) Stop() {
	c.logger.Infof("Closing client connection")

	if _, err := c.accountsCli.CloseSandboxAccount(c.accountUID); err != nil {
		log.Errorf("sandbox account closing: %v", err.Error())
	}

	err := c.client.Stop()
	if err != nil {
		c.logger.Errorf("client shutdown: %v", err.Error())
	}
}

// GetFreeMoney returns the amount of free money in the account.
func (c *Client) GetFreeMoney() (int64, error) {
	portfolio, err := c.operationsCli.GetPortfolio(c.accountUID, 0)
	if err != nil {
		return 0, err
	}

	money := portfolio.GetTotalAmountCurrencies()

	return money.GetUnits(), nil

}

func (c *Client) GetQuantityAvailabilityInstrument(ticker string) (int64, error) {
	instrument, err := c.getInstrument(ticker)
	if err != nil {
		return 0, fmt.Errorf("find instrument: %v", err)
	}

	portfolio, err := c.operationsCli.GetPortfolio(c.accountUID, 0)
	if err != nil {
		return 0, err
	}

	positions := portfolio.GetPositions()
	for i := range positions {
		if positions[i].Figi != instrument.Figi {
			continue
		}

		return positions[i].GetQuantity().GetUnits(), nil
	}

	return 0, nil
}

// PostOrder posts an order to the broker.
func (c *Client) PostOrder(order entity.Order) error {
	instrument, err := c.getInstrument(order.Ticker)
	if err != nil {
		return err
	}

	direction, ok := orderDirections[order.EventType]
	if !ok {
		return errors.New("invalid order type")
	}

	//TODO: do something with order
	if _, err = c.ordersCli.PostOrder(
		&investgo.PostOrderRequest{
			AccountId:    c.accountUID,
			InstrumentId: instrument.GetUid(),
			Direction:    direction,
			Quantity:     order.Quantity,
			OrderType:    investapi.OrderType_ORDER_TYPE_BESTPRICE,
			OrderId:      investgo.CreateUid(),
		},
	); err != nil {
		return err
	}

	return nil
}

// getInstrument returns the instrument by ticker
func (c *Client) getInstrument(ticker string) (*investapi.InstrumentShort, error) {
	instrResp, err := c.instrumentsCli.FindInstrument(ticker)
	if err != nil {
		return nil, fmt.Errorf("find instrument: %v", err)
	}

	ins := instrResp.GetInstruments()

	if len(ins) > 0 {
		return ins[0], nil
	}

	return nil, errors.New("not found instrument")
}
