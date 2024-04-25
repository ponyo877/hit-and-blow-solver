package entity

import (
	"github.com/mowshon/iterium"
)

type Solver struct {
	digit   int
	numbers []int
	cache   []Numbers
}

func NewSolver(d int, n []int) *Solver {
	return &Solver{
		digit:   d,
		numbers: n,
	}
}

func (s *Solver) Digit() int {
	return s.digit
}

func (s *Solver) Numbers() []int {
	return s.numbers
}

func (s *Solver) AllPatterns() []Numbers {
	if len(s.cache) > 0 {
		return s.cache
	}
	permutations := iterium.Permutations(s.numbers, s.Digit())
	numbersList, _ := permutations.Slice()
	for _, numbers := range numbersList {
		s.cache = append(s.cache, Numbers(numbers))
	}
	return s.cache
}
