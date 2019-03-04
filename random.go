package main

import (
	"errors"
	"math/rand"
)

func kPerm(k, n int) ([]int, error) {
	if k > n {
		return nil, errors.New("k cannot be greater than n")
	}

	perm := rand.Perm(n)
	return perm[:k], nil
}
