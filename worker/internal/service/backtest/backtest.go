package backtest

import (
	"sync"
	"worker/internal/entity"
)

type BackTest struct {
	strategy TradingStrategy

	pnl           float64
	calculateTax  entity.TaxFn
	numberDeal    int
	strategyInCh  chan entity.Candle
	strategyOutCh chan entity.Event

	inCh chan entity.Candle
}

type backTestConfig struct {
	inCh           chan entity.Candle
	strategy       TradingStrategy
	calculateTaxFn entity.TaxFn
}

func NewBackTest(cfg backTestConfig) BackTest {
	return BackTest{
		strategy:      cfg.strategy,
		calculateTax:  cfg.calculateTaxFn,
		strategyInCh:  cfg.strategy.InChannel(),
		strategyOutCh: cfg.strategy.OutChannel(),

		inCh: cfg.inCh,
	}
}

func (s *BackTest) Do() entity.BackTestResult {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		s.strategy.Do()
	}()

	go func() {
		wg.Wait()

		close(s.strategyOutCh)
	}()

	go func() {
		for candle := range s.inCh {
			s.strategyInCh <- candle
		}

		close(s.strategyInCh)
	}()

	var (
		dials       []entity.Dial
		numberDials int
		currentDial entity.Dial
		pnl         float64
	)

	for event := range s.strategyOutCh {
		numberDials++

		if event.Typ == entity.Buy {
			currentDial = entity.Dial{
				Buy: event.Price,
			}

			continue
		}

		currentDial.Sell = event.Price
		currentDial.Period = event.Period
		currentDial.CalculatePNL()
		pnl += currentDial.PNL

		dials = append(dials, currentDial)
	}

	return entity.BackTestResult{
		NumberDial: numberDials,
		PNL:        pnl,
		Dials:      dials,
	}
}
