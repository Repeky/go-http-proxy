package main

import (
	"fmt"
	"go-http-proxy/proxy"
	"log"
	"path/filepath"
)

func main() {
	configPath, err := filepath.Abs("./config.yaml")
	if err != nil {
		fmt.Println("Ошибка получения абсолютного пути:", err)
		return
	}

	config, err := proxy.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}

	proxy.InitLogger(config.LogFile)
	defer proxy.CloseLogger()

	proxy.StartProxyServer(config.ProxyPort, config.TargetURL)
}
