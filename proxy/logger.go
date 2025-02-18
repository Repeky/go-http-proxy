package proxy

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

var logger = logrus.New()

type WriterHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
}

func InitLogger(filename string) {
	rootDir, err := filepath.Abs("../../")
	if err != nil {
		logrus.Fatalf("Ошибка определения корневой директории проекта: %v", err)
	}

	logFilePath := filepath.Join(rootDir, filename)

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatalf("Ошибка открытия лог-файла: %v", err)
	}

	terminalFormatter := &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}

	fileFormatter := &logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}

	logger.SetOutput(io.Discard)
	logger.SetFormatter(terminalFormatter)

	logger.AddHook(&WriterHook{
		Writer:    os.Stdout,
		Formatter: terminalFormatter,
	})
	logger.AddHook(&WriterHook{
		Writer:    file,
		Formatter: fileFormatter,
	})

	logger.SetLevel(logrus.DebugLevel)

	logger.Info("Логирование инициализировано")
}

func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

func (hook *WriterHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func LogRequest(method, url string, headers map[string][]string) {
	logger.WithFields(logrus.Fields{
		"method": method,
		"url":    url,
	}).Info("Новый HTTP-запрос")

	for k, v := range headers {
		logger.WithField(k, v).Debug("Заголовок запроса")
	}
}

func CloseLogger() {
	logger.Info("Закрытие логирования")
}
