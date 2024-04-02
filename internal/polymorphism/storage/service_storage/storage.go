package service_storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"test/internal/polymorphism/storage"
)

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

type getResponseBody struct {
	Value string `json:"value"`
}

func (s *Storage) GetValue(key string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/find?key=%s", key))
	if err != nil {
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == http.StatusNotFound {
		return "", storage.ErrNotFound
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error from server: %s", string(respBody))
	}

	body := getResponseBody{}
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return "", err
	}

	return body.Value, nil
}

type RequestBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *Storage) SavePair(key, value string) error {
	requestBody := RequestBody{
		Key:   key,
		Value: value,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:3000/save"), bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error from server: %s", string(respBody))
	}

	return nil
}

func (s *Storage) Close() error {
	return nil
}
