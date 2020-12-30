package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lack-io/vine/service/config"
)

type Host struct {
	Address string `json:"address"`
	Port int `json:"port"`
}

type Config struct {
	Hosts map[string]Host `json:"hosts"`
}

func main() {
	pwd, _ := os.Getwd()
	err := config.LoadFile(pwd + "/config/" + "config.json")
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	config.Scan(&cfg)

	fmt.Println(cfg.Hosts["database"].Address)
}
