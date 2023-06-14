package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Vote struct {
	VoterID     string `json:"voter_id"`
	CandidateID string `json:"candidate_id"`
}

type Status struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func TestServeRootf5dfb6ba89(t *testing.T) {
	t.Run("Test GET method", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatalf("error creating request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `{"candidate1":0,"candidate2":0,"candidate3":0}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("Test POST method with valid vote", func(t *testing.T) {
		vote := Vote{
			VoterID:     "voter1",
			CandidateID: "candidate1",
		}

		voteJSON, err := json.Marshal(vote)
		if err != nil {
			t.Fatalf("error marshaling vote: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(voteJSON))
		if err != nil {
			t.Fatalf("error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		expected := `{"message":"Vote saved successfully"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("Test POST method with invalid vote", func(t *testing.T) {
		vote := Vote{
			VoterID:     "voter2",
			CandidateID: "",
		}

		voteJSON, err := json.Marshal(vote)
		if err != nil {
			t.Fatalf("error marshaling vote: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(voteJSON))
		if err != nil {
			t.Fatalf("error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := `{"message":"Vote is not valid. Vote cannot be saved"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("Test unsupported method", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/", nil)
		if err != nil {
			t.Fatalf("error creating request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(serveRoot)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusMethodNotAllowed)
		}
	})
}

func countVote() (map[string]int, error) {
	// TODO: Implement vote counting logic
	return map[string]int{
		"candidate1": 0,
		"candidate2": 0,
		"candidate3": 0,
	}, nil
}

func saveVote(vote Vote) error {
	// TODO: Implement vote saving logic
	return nil
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status.Code)
	json.NewEncoder(w).Encode(status)
}