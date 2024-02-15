package config

import (
	"analyzer/internal/bot/telegram"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/russianinvestments/invest-api-go-sdk/investgo"
)

//type IntervalStrategyConfig struct {
//	// Instruments - Слайс идентификаторов инструментов первичный
//	Instruments []string
//	// PreferredPositionPrice - Предпочтительная стоимость открытия позиции в валюте
//	PreferredPositionPrice float64
//	// MaxPositionPrice - Максимальная стоимость открытия позиции в валюте
//	MaxPositionPrice float64
//	// MinProfit - Минимальный процент выгоды, с которым можно совершать сделки
//	MinProfit float64
//	// IntervalUpdateDelay - Время ожидания для перерасчета интервала цены
//	IntervalUpdateDelay time.Duration
//	// TopInstrumentsQuantity - Топ лучших инструментов по волатильности
//	TopInstrumentsQuantity int
//	// SellOut - Если true, то по достижению дедлайна бот выходит из всех активных позиций
//	SellOut bool
//	// StorageDBPath - Путь к бд sqlite, в которой лежат исторические свечи по инструментам
//	StorageDBPath string
//	// StorageCandleInterval - Интервал для обновления и запроса исторических свечей
//	StorageCandleInterval pb.CandleInterval
//	// StorageFromTime - Время, от которого будет хранилище будет загружать историю для новых инструментов
//	StorageFromTime time.Time
//	// StorageUpdate - Если true, то в хранилище обновятся все свечи до now
//	StorageUpdate bool
//	// DaysToCalculateInterval - Кол-во дней, на которых рассчитывается интервал цен для торговли
//	DaysToCalculateInterval int
//	// StopLossPercent - Процент изменения цены, для стоп-лосс заявки
//	StopLossPercent float64
//	// AnalyseLowPercentile - Нижний процентиль для расчета интервала
//	AnalyseLowPercentile float64
//	// AnalyseHighPercentile - Верхний процентиль для расчета интервала
//	AnalyseHighPercentile float64
//	// Analyse - Тип анализа исторических свечей при расчете интервала
//	Analyse AnalyseType
//}
//
//// AnalyseType - Тип анализа исторических свечей при расчете интервала
//type AnalyseType int
//
//const (
//	// MATH_STAT - Анализ свечей при помощи пакета stats. Интервал для цены это AnalyseLowPercentile-AnalyseHighPercentile
//	// из выборки средних цен свечей за последние DaysToCalculateInterval дней
//	MATH_STAT AnalyseType = iota
//	// BEST_WIDTH - Анализ свечей происходит так:
//	// Вычисляется медиана распределения выборки средних цен свечей за последние DaysToCalculateInterval дней, от нее берется
//	// сначала фиксированный интервал шириной MinProfit процентов от медианы, далее если это выгодно интервал расширяется.
//	// Так же есть возможность задать фиксированный интервал в процентах.
//	BEST_WIDTH
//	// SIMPLEST - Поиск интервала простым перебором
//	SIMPLEST
//)

//type InvestConfig struct {
//	// EndPoint - Для работы с реальным контуром и контуром песочницы нужны разные эндпоинты.
//	// По умолчанию = sandbox-invest-public-api.tinkoff.ru:443
//	// https://tinkoff.github.io/investAPI/url_difference/
//	EndPoint string `yaml:"EndPoint"`
//	// Token - Ваш токен для InvestAPI
//	Token string `yaml:"APIToken"`
//	// AppName - Название вашего приложения, по умолчанию = tinkoff-api-go-sdk
//	AppName string `yaml:"AppName"`
//	// AccountId - Если уже есть аккаунт для апи можно указать напрямую,
//	// по умолчанию откроется новый счет в песочнице
//	AccountId string `yaml:"AccountId"`
//	// DisableResourceExhaustedRetry - Если true, то сдк не пытается ретраить, после получения ошибки об исчерпывании
//	// лимита запросов, если false, то сдк ждет нужное время и пытается выполнить запрос снова. По умолчанию = false
//	DisableResourceExhaustedRetry bool `yaml:"DisableResourceExhaustedRetry"`
//	// DisableAllRetry - Отключение всех ретраев
//	DisableAllRetry bool `yaml:"DisableAllRetry"`
//	// MaxRetries - Максимальное количество попыток переподключения, по умолчанию = 3
//	// (если указать значение 0 это не отключит ретраи, для отключения нужно прописать DisableAllRetry = true)
//	MaxRetries uint `yaml:"MaxRetries"`
//}

type ConnConfig struct {
	Network  string   `yaml:"network"`
	Database string   `yaml:"database"`
	Hosts    []string `yaml:"hosts"`
	Ports    []string `yaml:"ports"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

type SQLConfig struct {
	ConnConfig      `yaml:"conn_config"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

type DatabaseConfig struct {
	Postgres SQLConfig `yaml:"postgres"`
}

type HTTPServerConfig struct {
	Listen       string        `yaml:"listen"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type GRPCClient struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
	Retries int           `yaml:"retries"`
}

type ClientsConfig struct {
	Solver GRPCClient `yaml:"solver"`
}

type BotsConfig struct {
	Telegram telegram.Config `yaml:"telegram"`
}

type Config struct {
	Invest          investgo.Config  `yaml:"invest"`
	HTTPServer      HTTPServerConfig `yaml:"httpserver"`
	Databases       DatabaseConfig   `yaml:"databases"`
	GracefulTimeout time.Duration    `yaml:"graceful_timeout"`
	Clients         ClientsConfig    `yaml:"clients"`
	Bots            BotsConfig       `yaml:"bots"`
}

func ReadConfig(filePath string) (Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadConfig(filePath, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
