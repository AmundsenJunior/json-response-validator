package main

import (
	"encoding/json"
	"errors"
)

type CirclePattern struct {
	x    float64
	y    float64
	r    float64
	text string
}

type SquareMetadata struct {
	text    string
	colors  []string
	friends SquareFriends
}

type SquareFriends struct {
	sally bool
	bob   bool
}

type SquarePattern struct {
	coordinates []float64
	metadata    SquareMetadata
}

type Pattern interface{}

type Patterns map[string]Pattern

type Input string

// unmarshals raw JSON into a Pattern type instance
func Receiver(input Input, patternType string, output *Pattern) error {
	var circleOutput CirclePattern
	var squareOutput SquarePattern
	var err error

	switch patternType {
	case "circle":
		err = json.Unmarshal([]byte(input), &circleOutput)
		*output = &circleOutput
	case "square":
		err = json.Unmarshal([]byte(input), &squareOutput)
		*output = &squareOutput
	default:
		err = errors.New("input does not match a pattern")
	}

	return err
}

// evaluates if input values match the pattern values
func Matcher(input Input) error {
	var output Pattern
	var err error
	for patternType, patternValues := range patterns {
		err = Receiver(input, patternType, &output)
		if err != nil {
			continue
		}

		switch patternValues.(type) {
		case CirclePattern:

		case SquarePattern:

		}
	}

	return err
}

var	patterns = Patterns{
		"circle": CirclePattern{
			x:    150.0,
			y:    150.0,
			r:    100.0,
			text: "Hello circle!",
		},
		"square": SquarePattern{
			coordinates: []float64{1.0, 2.0, 3.0, 4.0},
			metadata: SquareMetadata{
				text:   "silly",
				colors: []string{"green", "blue"},
				friends: SquareFriends{
					sally: true,
					bob:   false,
				},
			},
		},
	}
