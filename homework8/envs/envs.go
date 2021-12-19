package envs

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"net/url"
	"strconv"
)

//port: 8080
//db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
//jaeger_url: http://jaeger:16686
//sentry_url: http://sentry:9000
//kafka_broker: kafka:9092
//some_app_id: testid
//some_app_key: testkey
type Config struct {
	Port         Port    `envconfig:"PORT" default:"8080"`
	Db_url       UrlType `envconfig:"DB_URL" default:"postgres://user:password@localhost:5432/petstore?sslmode=disable" required:"true"`
	Jaeger_url   UrlType `default:"http://jaeger:16686"`
	Sentry_url   UrlType `default:"http://sentry:9000"`
	Kafka_broker UrlType `default:"kafka:9092"`
	Some_app_id  string  `default:"testid"`
	Some_app_key string  `default:"testkey"`
}

type Port int

type UrlType string

func (p *Port) Decode(value string) error {
	portInt, err := strconv.Atoi(value)
	if err != nil {
		fmt.Errorf("Не удалось преобразовать порт в Int: %w", err)
		return err
	}
	*p = Port(portInt)
	return nil
}

func (u *UrlType) Decode(value string) error {
	_, err := url.ParseRequestURI(value)
	if err != nil {
		fmt.Errorf("Не удалось распознать URL: %w", err)
		return err
	}
	*u = UrlType(value) // // сохраним исходное значение URL, если парсинг прошел успешно
	return nil
}

func (c *Config) Read() error {
	err := envconfig.Process("", c)
	if err != nil {
		fmt.Errorf("Не удалось прочитать конфигурацию из переменных окружения: %w", err)
	}
	return err
}
