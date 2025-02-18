package proxy

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
)

type ProxyServer struct {
	TargetURL string
}

func NewProxy(target string) *ProxyServer {
	return &ProxyServer{TargetURL: target}
}

func (p ProxyServer) HandleRequest(w http.ResponseWriter, r *http.Request) {
	LogRequest(r.Method, r.URL.String(), r.Header)

	req, err := http.NewRequest(r.Method, p.TargetURL+r.URL.Path, r.Body)
	if err != nil {
		logger.WithError(err).Error("Ошибка создания запроса")
		http.Error(w, "Ошибка создание запроса ", http.StatusInternalServerError)
		return
	}

	req.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.WithError(err).Error("Ошибка запроса к целевому серверу")
		http.Error(w, "Ошибка запроса к целевому серверу", http.StatusBadGateway)
		return
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Println("ошибка закрытия тела ответа:", err)
		}
	}()

	logger.WithFields(logrus.Fields{
		"status": resp.StatusCode,
		"url":    r.URL.String(),
	}).Info("Ответ от целевого сервера")

	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println("ошибка копирования тела ответа:", err)
	}

}

func StartProxyServer(port string, targetURL string) {
	proxy := NewProxy(targetURL)

	http.HandleFunc("/", proxy.HandleRequest)
	logger.Infof("Запуск прокси-сервера на порту %s...", port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))
}
