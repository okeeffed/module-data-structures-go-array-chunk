package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkThree(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6}
	res := arrayChunk(arr, 3)
	expect := [][]int{}
	rowOne := []int{1, 2, 3}
	rowTwo := []int{4, 5, 6}

	expect = append(expect, rowOne)
	expect = append(expect, rowTwo)
	assert.Equal(expect, res, "Chunked Arr does not match for 3")
}

func TestChunkZero(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6}
	res := arrayChunk(arr, 0)
	var expect [][]int
	assert.Equal(expect, res, "Chunked Arr does not match for 0")
}

func TestChunkTwo(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6}
	res := arrayChunk(arr, 2)
	expect := [][]int{}
	rowOne := []int{1, 2}
	rowTwo := []int{3, 4}
	rowThree := []int{5, 6}

	expect = append(expect, rowOne)
	expect = append(expect, rowTwo)
	expect = append(expect, rowThree)
	assert.Equal(expect, res, "Chunked Arr does not match for 3")
}

func TestChunkSix(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6}
	res := arrayChunk(arr, 6)
	expect := [][]int{}
	rowOne := []int{1, 2, 3, 4, 5, 6}

	expect = append(expect, rowOne)
	assert.Equal(expect, res, "Chunked Arr does not match for 6")
}

func TestChunkSeven(t *testing.T) {
	assert := assert.New(t)
	arr := []int{1, 2, 3, 4, 5, 6}
	res := arrayChunk(arr, 7)
	expect := [][]int{}
	rowOne := []int{1, 2, 3, 4, 5, 6}

	expect = append(expect, rowOne)
	assert.Equal(expect, res, "Chunked Arr does not match for 7")
}
