package jsons

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//port: 8080
//db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
//jaeger_url: http://jaeger:16686
//sentry_url: http://sentry:9000
//kafka_broker: kafka:9092
//some_app_id: testid
//some_app_key: testkey
type Config struct {
	Port        Port    `json:"port"`
	DbURL       UrlType `json:"db_url"`
	JaegerURL   UrlType `json:"jaeger_url"`
	SentryURL   UrlType `json:"sentry_url"`
	KafkaBroker UrlType `json:"kafka_broker"`
	SomeAppID   string  `json:"some_app_id"`
	SomeAppKey  string  `json:"some_app_key"`
	Users       []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"users"`
}

type Port int

type UrlType string

func (p *Port) UnmarshalJSON(b []byte) error {
	portInt, err := strconv.Atoi(strings.Trim(string(b), `""`))
	if err != nil {
		fmt.Errorf("Не удалось преобразовать порт в Int: %w", err)
		return err
	}
	*p = Port(portInt)
	return nil
}

func (u *UrlType) UnmarshalJSON(b []byte) error {
	urlString := strings.Trim(string(b), `""`)
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Errorf("Не удалось распознать URL: %w", err)
		return err
	}
	*u = UrlType(urlString) // сохраним исходное значение URL, если парсинг прошел успешно
	return nil
}

func (c *Config) Read() error {
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Errorf("Не удалось прочитать json-файл: %w", err)
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		fmt.Errorf("Не удалось прочитать конфигурацию из json-файла: %w", err)
	}
	return err
}
