package main

import (
	"go-http-proxy/proxy"
	"log"
)

func main() {
	config, err := proxy.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}

	proxy.InitLogger(config.LogFile)
	defer proxy.CloseLogger()

	proxy.StartProxyServer(config.ProxyPort, config.TargetURL)
}
