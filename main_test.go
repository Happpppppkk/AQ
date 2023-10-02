package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/montanaflynn/stats"
)

func TestOLS(t *testing.T) {
	data := stats.Series{{X: 1, Y: 2}, {X: 2, Y: 3}, {X: 3, Y: 4}}

	expectedSlope := 1
	expectedIntercept := 1

	slope, intercept := OLS(data)

	if slope != 1 {
		fmt.Printf("Failed function, Ecpected %v, got %v\n", expectedSlope, slope)
	} else {
		t.Log("Pass")
	}

	if intercept != 1 {
		fmt.Printf("Failed function, Ecpected %v, got %v\n", expectedIntercept, intercept)
	} else {
		t.Log("Pass")
	}

}

func TestCreateSeries(t *testing.T) {
	x := []float64{1.0, 2.0, 3.0}
	y := []float64{2.0, 3.0, 4.0}

	expectedResult := stats.Series{{X: 1, Y: 2}, {X: 2, Y: 3}, {X: 3, Y: 4}}

	result := createSeries(x, y)

	assert.Equal(t, expectedResult, result, "test Pass")
}
