package service

import (
	"math/rand"
	"time"
)

func GetRandomIndices(max, numIndices int) []int {
	result := make([]int, 0)
	existingIndices := make(map[int]struct{}, 0)

	for i := 0; i < numIndices; i++ {
		randomIndex := randomIndex(max, existingIndices)
		result = append(result, randomIndex)
	}

	return result
}

func randomIndex(size int, existingIndices map[int]struct{}) int {
	rand.Seed(time.Now().UnixNano())

	for {
		randomIndex := rand.Intn(size)

		if _, exists := existingIndices[randomIndex]; !exists {
			existingIndices[randomIndex] = struct{}{}
			return randomIndex
		}
	}
}
