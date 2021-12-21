package yamls

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"net/url"
	"os"
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
	Port        Port    `yaml:"port"`
	DbURL       UrlType `yaml:"db_url"`
	JaegerURL   UrlType `yaml:"jaeger_url"`
	SentryURL   UrlType `yaml:"sentry_url"`
	KafkaBroker UrlType `yaml:"kafka_broker"`
	SomeAppID   string  `yaml:"some_app_id"`
	SomeAppKey  string  `yaml:"some_app_key"`
	Users       []struct {
		Name string `yaml:"name"`
		ID   int    `yaml:"id"`
	} `yaml:"users"`
}

type Port int

type UrlType string

func (p *Port) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var portStr string
	err := unmarshal(&portStr)
	if err != nil {
		return fmt.Errorf("Не удалось прочитать поле с типом Port: %w", err)
	}
	portInt, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("Не удалось преобразовать порт в Int: %w", err)
	}
	*p = Port(portInt)
	return nil
}

func (u *UrlType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var urlString string
	err := unmarshal(&urlString)
	if err != nil {
		return fmt.Errorf("Не удалось прочитать поле с типом UrlType: %w", err)
	}
	_, err = url.ParseRequestURI(urlString)
	if err != nil {
		return fmt.Errorf("Не удалось распознать URL: %w", err)
	}
	*u = UrlType(urlString) // сохраним исходное значение URL, если парсинг прошел успешно
	return nil
}

func (c *Config) Read() error {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return fmt.Errorf("Не удалось прочитать yaml-файл: %w", err)
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		err = fmt.Errorf("Не удалось прочитать конфигурацию из yaml-файла: %w", err)
	}
	return err
}
