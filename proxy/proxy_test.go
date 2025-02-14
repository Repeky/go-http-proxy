package proxy

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockTargetServer() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "ok"}`))
	})
	return httptest.NewServer(handler)
}

func TestProxyServer(t *testing.T) {
	targetServer := mockTargetServer()
	defer targetServer.Close()

	proxy := NewProxy(targetServer.URL)

	testServer := httptest.NewServer(http.HandlerFunc(proxy.HandleRequest))
	defer testServer.Close()

	resp, err := http.Get(testServer.URL + "/test")
	if err != nil {
		t.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался код 200, но получили %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Ошибка чтения тела ответа: %v", err)
	}

	expectedBody := `{"message": "ok"}`
	if string(body) != expectedBody {
		t.Errorf("Ожидали %s, получили %s", expectedBody, string(body))
	}
}
