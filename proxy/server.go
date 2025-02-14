package proxy

import (
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
		http.Error(w, "Ошибка создание запроса ", http.StatusInternalServerError)
		return
	}

	req.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Ошибка запроса к целевому серверу", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func StartProxyServer(port string, targetURL string) {
	proxy := NewProxy(targetURL)

	http.HandleFunc("/", proxy.HandleRequest)
	log.Printf("Запуск прокси-сервера на порту %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
