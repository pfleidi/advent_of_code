package main

import "testing"

var testInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestExercise1(t *testing.T) {
	expected := 198
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
	expected := 230
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
