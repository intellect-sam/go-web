package calc

import "testing"

func TestAdd(t *testing.T) {
	result := Add(10, 20)

	if result != 30 {
		t.Error("Expected 30, got ", result)
	}
}

func TestSub(t *testing.T) {
	result := Subtract(20, 10)

	if result != 10 {
		t.Error("Expected 10, got ", result)
	}
}

func TestVote(t *testing.T) {
	result := VoteEligibility(10)

	if result != "You are not eligible to vote" {
		t.Error("It should return 'You are not eligible to vote'")
	}
}

func TestGetById(t *testing.T) {
	result, err := GetById("105")

	if result != "Task 101" {
		t.Error("Expected 'Task 105', got ", result)
	} else if err != nil {
		t.Error("Expected nil, got ", err)
	}
}
