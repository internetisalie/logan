package logan

import (
	"bytes"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type HttpJsonHook struct {
	endpoint  string
	formatter logrus.Formatter
	extras    map[string]interface{}
}

func (h *HttpJsonHook) WithExtras(extras map[string]interface{}) *HttpJsonHook {
	h.extras = extras
	return h
}

func (h *HttpJsonHook) Fire(entry *logrus.Entry) error {
	for k, v := range h.extras {
		entry.Data[k] = v
	}
	bodyBytes, err := h.bytes(entry)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewBuffer(bodyBytes)
	resp, err := http.Post(h.endpoint, "application/json", bodyReader)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func (h *HttpJsonHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *HttpJsonHook) bytes(entry *logrus.Entry) ([]byte, error) {
	serialized, err := h.formatter.Format(entry)
	if err != nil {
		return nil, err
	}
	return serialized, nil
}

func NewHttpJsonHook(endpoint string) *HttpJsonHook {
	return &HttpJsonHook{
		endpoint: endpoint,
		formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap: logrus.FieldMap{
				"time": "@timestamp",
			},
		},
	}
}
