package main

import "fmt"

type Service interface {
	ServiceName() string
}


type Config struct {
	Endpoint string
}

var CfgMap map[string]*Config
func init() {
	CfgMap = make(map[string]*Config, 4)
	CfgMap["hello"] = &Config{
		Endpoint: "http://localhost:8083/",
	}
	initRead()
}

func initRead() {
	fmt.Println("the info of config: ")
	for key,val := range CfgMap {
		fmt.Printf("%s => %s \n", key, val)
	}
}