package c2api

import "math/rand"

func chunkSlice[T any](input []T, numChunks int) [][]T {
	// Create the 2D slice to hold the chunks
	chunks := make([][]T, numChunks)

	// Iterate over the input slice and divide it into chunks
	for i := range input {
		chunkIndex := i % numChunks
		chunks[chunkIndex] = append(chunks[chunkIndex], input[i])
	}

	return chunks
}

func randElem[T any](slice *[]T) T {
	return (*slice)[rand.Intn(len(*slice))]
}
