package main

import (
	"fmt"
	envs "homework8/envs"
	"log"
)

func main() {
	conf := envs.Config{}
	err := (&conf).Read()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%#v", conf)
}
