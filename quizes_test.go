package main

import (
	"testing"
)

func Test_getQuizes(t *testing.T) {
	got := getQuizes("quizes.csv")
	want := 12
	if len(got) != want {
		t.Errorf("got %v but want %v", len(got), want)
	}
}
