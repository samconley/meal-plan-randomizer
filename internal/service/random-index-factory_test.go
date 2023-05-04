package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomIndexFactory_GetRandomIndices(t *testing.T) {
	tests := []struct {
		name       string
		max        int
		numIndices int
	}{
		{
			name:       `2 indices, max 10 do not overlap`,
			max:        10,
			numIndices: 2,
		},
		{
			name:       `7 indices, max 100000 do not overlap`,
			max:        10000,
			numIndices: 7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			randomIndices := GetRandomIndices(tc.max, tc.numIndices)
			assert.Equal(t, tc.numIndices, len(randomIndices))
			for i := 0; i < len(randomIndices)-1; i++ {
				for j := i + 1; j < len(randomIndices); j++ {
					assert.NotEqual(t, randomIndices[i], randomIndices[j])
				}
			}
		})
	}
}
