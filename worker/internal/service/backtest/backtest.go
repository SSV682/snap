package backtest

import (
	"worker/internal/entity"
)

type BackTest struct {
	candles  []entity.Candle
	strategy TradingStrategy

	pnl           float64
	calculateTax  entity.TaxFn
	numberDeal    int
	strategyInCh  chan entity.Candle
	strategyOutCh chan entity.Event
}

type backTestConfig struct {
	candles        []entity.Candle
	strategy       TradingStrategy
	calculateTaxFn entity.TaxFn
	strategyInCh   chan entity.Candle
	strategyOutCh  chan entity.Event
}

func NewBackTest(cfg *backTestConfig) BackTest {
	return BackTest{
		candles:       cfg.candles,
		strategy:      cfg.strategy,
		calculateTax:  cfg.calculateTaxFn,
		strategyInCh:  cfg.strategyInCh,
		strategyOutCh: cfg.strategyOutCh,
	}
}

func (s *BackTest) Do() entity.BackTestResult {
	s.strategy.Run()

	go func() {
		for i := range s.candles {
			s.strategyInCh <- s.candles[i]
		}

		if len(s.strategyInCh) == 0 {
			s.strategy.Close()
		}
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
