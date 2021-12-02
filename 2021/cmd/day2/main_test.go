package main

import "testing"

var testInput = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestExercise1(t *testing.T) {
	expected := 150
	result, err := Exercise1(testInput)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if result != expected {
		t.Logf("result should %d but got %d", expected, result)
		t.Fail()
	}
}

func TestExercise2(t *testing.T) {
	expected := 900
	result, err := Exercise2(testInput)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if result != expected {
		t.Logf("result should %d but got %d", expected, result)
		t.Fail()
	}
}
