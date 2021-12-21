package main

import (
	"flag"
	"fmt"
	envs "homework9/envs"
	myjsons "homework9/jsons"
	myyamls "homework9/yamls"
	"log"
)

var configType = flag.String("conf", "e", "Тип config-файла: 'e' - переменные окружения, 'j' - json, 'y' - yaml")

var configMap = map[string]func() error{
	"e": getConfigEnv,
	"j": getConfigJson,
	"y": getConfigYaml,
}

func getConfigEnv() error {
	conf := envs.Config{}
	err := conf.Read()
	if err == nil {
		fmt.Printf("%#v\n", conf)
	}
	return err
}

func getConfigJson() error {
	confJson := myjsons.Config{}
	err := confJson.Read()
	if err == nil {
		fmt.Printf("%+v\n", confJson)
	}
	return err
}

func getConfigYaml() error {
	confYaml := myyamls.Config{}
	err := confYaml.Read()
	if err == nil {
		fmt.Printf("%+v\n", confYaml)
	}
	return err
}

func main() {
	flag.Parse()
	println(*configType)

	f, ok := configMap[*configType]
	if !ok {
		fmt.Println("Не найден обработчик для configType = ", *configType)
		return
	}

	err := f()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
