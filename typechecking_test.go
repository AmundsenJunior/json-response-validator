package main

import "testing"

func TestInput_MatcherSuccess(t *testing.T) {
	for _, v := range matches {
		err := Matcher(Input(v))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestInput_MatcherError(t *testing.T) {
	for _, v := range nonmatches {
		err := Matcher(Input(v))
		if err == nil {
			t.Error("Incorrectly matched on", v)
		}
	}
}

