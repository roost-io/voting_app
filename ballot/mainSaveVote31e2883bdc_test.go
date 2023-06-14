package main

import (
	"sync"
	"testing"
)

type Vote struct {
	CandidateID int
}

var candidateVotesStore map[int]int
var mu sync.Mutex

func getCandidatesVote() map[int]int {
	if candidateVotesStore == nil {
		mu.Lock()
		defer mu.Unlock()
		if candidateVotesStore == nil {
			candidateVotesStore = make(map[int]int)
		}
	}
	return candidateVotesStore
}

func saveVote(vote Vote) error {
	candidateVotesStore = getCandidatesVote()
	candidateVotesStore[vote.CandidateID]++
	return nil
}

func TestSaveVote31e2883bdc(t *testing.T) {
	// Test case 1: Valid vote
	vote1 := Vote{CandidateID: 1}
	err := saveVote(vote1)
	if err != nil {
		t.Errorf("Failed to save vote: %v", err)
	}

	if candidateVotesStore[vote1.CandidateID] != 1 {
		t.Errorf("Expected vote count for candidate %d to be 1, got %d", vote1.CandidateID, candidateVotesStore[vote1.CandidateID])
	}

	// Test case 2: Multiple votes for the same candidate
	vote2 := Vote{CandidateID: 1}
	err = saveVote(vote2)
	if err != nil {
		t.Errorf("Failed to save vote: %v", err)
	}

	if candidateVotesStore[vote2.CandidateID] != 2 {
		t.Errorf("Expected vote count for candidate %d to be 2, got %d", vote2.CandidateID, candidateVotesStore[vote2.CandidateID])
	}
}