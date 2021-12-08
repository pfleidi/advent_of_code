package main

import "testing"

var testLines = []line{
	{start: point{x: 0, y: 9}, end: point{x: 5, y: 9}},
	{start: point{x: 8, y: 0}, end: point{x: 0, y: 8}},
	{start: point{x: 9, y: 4}, end: point{x: 3, y: 4}},
	{start: point{x: 2, y: 2}, end: point{x: 2, y: 1}},
	{start: point{x: 7, y: 0}, end: point{x: 7, y: 4}},
	{start: point{x: 6, y: 4}, end: point{x: 2, y: 0}},
	{start: point{x: 0, y: 9}, end: point{x: 2, y: 9}},
	{start: point{x: 3, y: 4}, end: point{x: 1, y: 4}},
	{start: point{x: 0, y: 0}, end: point{x: 8, y: 8}},
	{start: point{x: 5, y: 5}, end: point{x: 8, y: 2}},
}

func TestExercise1(t *testing.T) {
	expected := 5
	result, err := Exercise1(testLines)

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
	expected := 12
	result, err := Exercise2(testLines)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if result != expected {
		t.Logf("result should %d but got %d", expected, result)
		t.Fail()
	}
}
