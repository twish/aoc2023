package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOne_Get_56397_AsValueForSolution(t *testing.T) {
	total := One("./input_data")

	assert.Equal(t, 56397, total)
}
