package main

import "testing"

var testNumbers = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

var testBoards = []board{
	{
		rows: [][]*coordinate{
			{&coordinate{number: 22}, &coordinate{number: 13}, &coordinate{number: 17}, &coordinate{number: 11}, &coordinate{number: 0}},
			{&coordinate{number: 8}, &coordinate{number: 2}, &coordinate{number: 23}, &coordinate{number: 4}, &coordinate{number: 24}},
			{&coordinate{number: 21}, &coordinate{number: 9}, &coordinate{number: 14}, &coordinate{number: 16}, &coordinate{number: 7}},
			{&coordinate{number: 6}, &coordinate{number: 10}, &coordinate{number: 3}, &coordinate{number: 18}, &coordinate{number: 5}},
			{&coordinate{number: 1}, &coordinate{number: 12}, &coordinate{number: 20}, &coordinate{number: 15}, &coordinate{number: 19}},
		},
	},
	{
		rows: [][]*coordinate{
			{&coordinate{number: 3}, &coordinate{number: 15}, &coordinate{number: 0}, &coordinate{number: 2}, &coordinate{number: 22}},
			{&coordinate{number: 9}, &coordinate{number: 18}, &coordinate{number: 13}, &coordinate{number: 17}, &coordinate{number: 5}},
			{&coordinate{number: 19}, &coordinate{number: 8}, &coordinate{number: 7}, &coordinate{number: 25}, &coordinate{number: 23}},
			{&coordinate{number: 20}, &coordinate{number: 11}, &coordinate{number: 10}, &coordinate{number: 24}, &coordinate{number: 4}},
			{&coordinate{number: 14}, &coordinate{number: 21}, &coordinate{number: 16}, &coordinate{number: 12}, &coordinate{number: 6}},
		},
	},
	{
		rows: [][]*coordinate{
			{&coordinate{number: 14}, &coordinate{number: 21}, &coordinate{number: 17}, &coordinate{number: 24}, &coordinate{number: 4}},
			{&coordinate{number: 10}, &coordinate{number: 16}, &coordinate{number: 15}, &coordinate{number: 9}, &coordinate{number: 19}},
			{&coordinate{number: 18}, &coordinate{number: 8}, &coordinate{number: 23}, &coordinate{number: 26}, &coordinate{number: 20}},
			{&coordinate{number: 22}, &coordinate{number: 11}, &coordinate{number: 13}, &coordinate{number: 6}, &coordinate{number: 5}},
			{&coordinate{number: 2}, &coordinate{number: 0}, &coordinate{number: 12}, &coordinate{number: 3}, &coordinate{number: 7}},
		},
	},
}

func TestExercise1(t *testing.T) {
	expected := 4512
	result, err := Exercise1(testBoards, testNumbers)

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
	expected := 1924
	result, err := Exercise2(testBoards, testNumbers)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if result != expected {
		t.Logf("result should %d but got %d", expected, result)
		t.Fail()
	}
}
