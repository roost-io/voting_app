package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func httpClientRequest(operation, hostAddr, command string, params io.Reader) (int, []byte, error) {

	url := "http://" + hostAddr + command
	if strings.Contains(hostAddr, "http://") {
		url = hostAddr + command
	}

	req, err := http.NewRequest(operation, url, params)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("Failed to create HTTP request." + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer resp.Body.Close()

	body, ioErr := ioutil.ReadAll(resp.Body)
	if hBit := resp.StatusCode / 100; hBit != 2 && hBit != 3 {
		if ioErr != nil {
			ioErr = fmt.Errorf("status code error %d", resp.StatusCode)
		}
	}
	return resp.StatusCode, body, ioErr
}

func TestHttpClientRequestdae1710fc2(t *testing.T) {
	t.Run("SuccessCase", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		}))
		defer ts.Close()

		status, body, err := httpClientRequest("GET", ts.URL, "", nil)
		if err != nil {
			t.Error("Expected no error, got:", err)
		}
		if status != http.StatusOK {
			t.Error("Expected status 200, got:", status)
		}
		if string(body) != "OK" {
			t.Error("Expected body 'OK', got:", string(body))
		}
	})

	t.Run("FailureCase", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		}))
		defer ts.Close()

		status, body, err := httpClientRequest("POST", ts.URL, "", bytes.NewBuffer([]byte(`{"key": "value"}`)))
		if err == nil {
			t.Error("Expected error, got nil")
		}
		if status != http.StatusInternalServerError {
			t.Error("Expected status 500, got:", status)
		}
		if string(body) != "Internal Server Error" {
			t.Error("Expected body 'Internal Server Error', got:", string(body))
		}
	})
}