package main

import (
	"math/rand"
	"sort"
	"testing"
)

type CandidateVotes struct {
	CandidateID string
	Votes       int
}

type ResultBoard struct {
	Results    []CandidateVotes
	TotalVotes int
}

func countVote(votes map[string]int) (res ResultBoard, err error) {
	for candidateID, votes := range votes {
		res.Results = append(res.Results, CandidateVotes{candidateID, votes})
		res.TotalVotes += votes
	}

	sort.Slice(res.Results, func(i, j int) bool {
		return res.Results[i].Votes > res.Results[j].Votes
	})
	return res, err
}

func getCandidatesVote() map[string]int {
	return map[string]int{
		"Candidate1": rand.Intn(1000),
		"Candidate2": rand.Intn(1000),
		"Candidate3": rand.Intn(1000),
		"Candidate4": rand.Intn(1000),
		"Candidate5": rand.Intn(1000),
	}
}

func TestCountVote(t *testing.T) {
	votes := map[string]int{
		"Candidate1": 500,
		"Candidate2": 300,
		"Candidate3": 100,
		"Candidate4": 50,
		"Candidate5": 50,
	}
	expectedTotalVotes := 1000

	result, err := countVote(votes)
	if err != nil {
		t.Error("Error occurred while counting votes:", err)
	}

	if result.TotalVotes != expectedTotalVotes {
		t.Errorf("Expected total votes: %d, got: %d", expectedTotalVotes, result.TotalVotes)
	}

	if len(result.Results) != len(votes) {
		t.Errorf("Expected number of candidates: %d, got: %d", len(votes), len(result.Results))
	}

	prevVotes := result.Results[0].Votes
	for _, candidate := range result.Results {
		if candidate.Votes > prevVotes {
			t.Error("Results are not sorted in descending order of votes")
			break
		}
		prevVotes = candidate.Votes
	}
}

func TestCountVoteEmptyInput(t *testing.T) {
	votes := map[string]int{}
	expectedTotalVotes := 0

	result, err := countVote(votes)
	if err != nil {
		t.Error("Error occurred while counting votes:", err)
	}

	if result.TotalVotes != expectedTotalVotes {
		t.Errorf("Expected total votes: %d, got: %d", expectedTotalVotes, result.TotalVotes)
	}

	if len(result.Results) != len(votes) {
		t.Errorf("Expected number of candidates: %d, got: %d", len(votes), len(result.Results))
	}
}