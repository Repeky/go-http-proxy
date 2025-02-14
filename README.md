# Go HTTP Proxy

**Go HTTP Proxy** — это прокси-сервер на Go, который логирует HTTP-запросы и пересылает их на целевой сервер.

---

## Возможности
- Логирование HTTP-запросов (метод, заголовки, путь).
- Перенаправление запросов на указанный сервер.
- Конфигурация через `config.yaml`.
- Автоматическое тестирование.

---

## Установка и запуск

### 1. Клонируйте репозиторий
```bash
git clone https://github.com/Repeky/go-http-proxy.git
cd go-http-proxy
```

### 2. Установите зависимости
```bash
go mod tidy
```

### 3. Запустите сервер
```bash
go run main.go
```

### 4. Отправьте тестовый запрос
```bash
curl -X GET http://localhost:8080/posts/1
```

---

## Конфигурация (config.yaml)

Файл `config.yaml` позволяет задать параметры работы сервера:

```yaml
proxy_port: "8080"
target_url: "https://jsonplaceholder.typicode.com"
log_file: "requests.log"
```

---

## Логирование

Запросы логируются в файл `requests.log` и в терминал.

Пример логов:

```bash
Запрос: GET /posts/1
User-Agent: curl/7.68.0
Accept: */*
---------------------
```

---

## Автоматическое тестирование

Запустить тесты:
```bash
go test ./proxy