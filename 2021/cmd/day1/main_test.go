package main

import (
	"testing"
)

var testInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func TestExercise1(t *testing.T) {
	expected := 7
	result := Exercise1(testInput)

	if result != expected {
		t.Logf("result should %d but got %d", expected, result)
		t.Fail()
	}
}

func TestExercise2(t *testing.T) {
	expected := 5
	result := Exercise2(testInput)

	if result != expected {
		t.Logf("result should %d but got %d", expected, result)
		t.Fail()
	}
}
