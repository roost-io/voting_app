package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(resp)
}

func TestWriteVoterResponse(t *testing.T) {
	t.Run("TestWriteVoterResponsed4e306ce05_Success", func(t *testing.T) {
		status := Status{Code: 200, Message: "Vote successfully recorded"}

		rr := httptest.NewRecorder()
		writeVoterResponse(rr, status)

		if rr.Code != http.StatusOK {
			t.Errorf("writeVoterResponse returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		}

		var respStatus Status
		err := json.Unmarshal(rr.Body.Bytes(), &respStatus)
		if err != nil {
			t.Error("unable to unmarshal response")
		}

		if respStatus.Code != status.Code || respStatus.Message != status.Message {
			t.Errorf("writeVoterResponse returned unexpected body: got %v want %v", respStatus, status)
		}
	})

	t.Run("TestWriteVoterResponsed4e306ce05_Failure", func(t *testing.T) {
		status := Status{Code: 500, Message: "Internal server error"}

		rr := httptest.NewRecorder()
		writeVoterResponse(rr, status)

		if rr.Code != http.StatusOK {
			t.Errorf("writeVoterResponse returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
		}

		var respStatus Status
		err := json.Unmarshal(rr.Body.Bytes(), &respStatus)
		if err != nil {
			t.Error("unable to unmarshal response")
		}

		if respStatus.Code != status.Code || respStatus.Message != status.Message {
			t.Errorf("writeVoterResponse returned unexpected body: got %v want %v", respStatus, status)
		}
	})
}