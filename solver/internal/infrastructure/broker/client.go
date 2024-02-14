package broker

import (
	"context"
	"errors"
	"fmt"

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

// getFreeMoney returns the amount of free money in the account.
func (c *Client) getFreeMoney() (int64, error) {
	portfolio, err := c.operationsCli.GetPortfolio(c.accountUID, 0)
	if err != nil {
		return 0, err
	}

	money := portfolio.GetTotalAmountCurrencies()

	return money.GetUnits(), nil

}

// postOrder posts an order to the broker.
func (c *Client) postOrder(ticker string, quantity int64, direction investapi.OrderDirection) error {
	instrument, err := c.getInstrument(ticker)
	if err != nil {
		return err
	}

	//TODO: do something with order
	if _, err = c.ordersCli.PostOrder(
		&investgo.PostOrderRequest{
			AccountId:    c.accountUID,
			InstrumentId: instrument.GetUid(),
			Direction:    direction,
			Quantity:     quantity,
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
