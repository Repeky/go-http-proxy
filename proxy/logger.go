package proxy

import (
	"io"
	"log"
	"os"
)

var logFile *os.File

func InitLogger(filename string) {
	var err error
	logFile, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Ошибка открытие лог-файла %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
}

func LogRequest(method, url string, headers map[string][]string) {
	log.Printf("Запрос: %s %s", method, url)
	for k, v := range headers {
		log.Printf("%s, %v", k, v)
	}
	log.Println("---------------------")
}

func CloseLogger() {
	logFile.Close()
}
