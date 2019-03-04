package main

import (
	"testing"
)

func Test_kPerm(t *testing.T) {
	tests := []struct {
		k int
		n int
	}{
		{1, 2},
	}
	for _, tt := range tests {
		got, err := kPerm(tt.k, tt.n)

		if err != nil {
			t.Errorf("kPerm(%v, %v) error = %v", tt.k, tt.n, err)
		}
		count := len(got)
		if count != tt.k {
			t.Errorf("kPerm(%v, %v) = %v. But number of elements %v", tt.k, tt.n, got, count)
		}
		uniqueResults := map[int]bool{}
		for _, g := range got {
			if g > tt.n {
				t.Errorf("kPerm(%v, %v) = %v. But %v is out of range", tt.k, tt.n, got, g)
			}
			if _, exists := uniqueResults[g]; exists {
				t.Errorf("kPerm(%v, %v) = %v. But %v is not unique", tt.k, tt.n, got, g)
			}
			uniqueResults[g] = true
		}
	}
}
