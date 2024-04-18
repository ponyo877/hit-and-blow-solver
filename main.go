package main

import (
	"math"
	"math/rand"
	"slices"
	"time"
)

func main() {
}

type Numbers []int

func (n Numbers) Feedback(estimate Numbers) *Feedback {
	f := NewFeedback(0, 0)
	for i, num := range estimate {
		if n[i] == num {
			f.IncHit()
		}
		if slices.Contains(n, num) {
			f.IncBlow()
		}
	}
	return f
}

type Feedback struct {
	hit  int
	blow int
}

func NewFeedback(hit int, blow int) *Feedback {
	return &Feedback{hit, blow}
}

func (f *Feedback) IncHit() {
	f.hit++
}

func (f *Feedback) IncBlow() {
	f.blow++
}

type Hint struct {
	numbers  Numbers
	feedback *Feedback
}

type Solver struct {
	history []Hint
}

func (s *Solver) IsEmpty() bool {
	return len(s.history) == 0
}

func (s *Solver) Push(h Hint) {
	s.history = append(s.history, h)
}

func (s *Solver) Undo() {
	s.history = s.history[:len(s.history)-1]
}

func combinations(list []int, choose, buf int) (c chan Numbers) {
	c = make(chan Numbers, buf)
	go func() {
		defer close(c)
		switch {
		case choose == 0:
			c <- Numbers{}
		case choose == len(list):
			c <- list
		case len(list) < choose:
			return
		default:
			for i := 0; i < len(list); i++ {
				for subComb := range combinations(list[i+1:], choose-1, buf) {
					c <- append(Numbers{list[i]}, subComb...)
				}
			}
		}
	}()
	return
}

func allNumbers() []Numbers {
	numberList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	digits := 3
	buf := 2
	allNumbers := []Numbers{}
	for number := range combinations(numberList, digits, buf) {
		allNumbers = append(allNumbers, number)
	}
	return allNumbers
}

func (s *Solver) candidate() []Numbers {
	candidate := []Numbers{}
	for _, numbers := range allNumbers() {
		var isBreak bool
		for _, hint := range s.history {
			if hint.feedback != numbers.Feedback(hint.numbers) {
				isBreak = true
				break
			}
		}
		if !isBreak {
			candidate = append(candidate, numbers)
		}
	}
	return candidate
}

func (s *Solver) BestEstimate() Numbers {
	mine := math.MaxFloat64
	minNumbers := Numbers{}
	allNumbers := allNumbers()
	if s.IsEmpty() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return allNumbers[r.Intn(len(allNumbers))]
	}

	for _, estimate := range allNumbers {
		h := map[*Feedback]int{}
		for _, candidate := range s.candidate() {
			f := candidate.Feedback(estimate)
			h[f]++
		}
		e := entropy(h)
		if e < mine {
			mine = e
			minNumbers = estimate
		}
	}
	return minNumbers
}

func entropy(h map[*Feedback]int) float64 {
	var e float64
	for _, val := range h {
		absv := math.Abs(float64(val))
		e += absv * math.Log(1+absv)
	}
	return e
}
