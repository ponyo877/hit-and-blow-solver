package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"slices"
	"time"
)

func main() {
	d := 3
	s := NewSolver(d)
	h := NewHist(Numbers{0, 1, 2}, &Feedback{2, 0})
	s.Push(h)
	ca := s.Candidate()
	be := s.BestEstimate()
	fs := s.FeedbackSelect(be)
	if !ca.Contains(be) && len(ca) <= 4 {
		fmt.Println(ca)
	}
	fmt.Println(be)
	for _, f := range fs {
		fmt.Println(f)
	}

	// 2nd
	h = NewHist(be, &Feedback{1, 2})
	s.Push(h)
	ca = s.Candidate()
	be = s.BestEstimate()
	fs = s.FeedbackSelect(be)
	if !ca.Contains(be) && len(ca) <= 4 {
		fmt.Println(ca)
	}
	fmt.Println(be)
	for _, f := range fs {
		fmt.Println(f)
	}
}

func permutations(iterable []int, r int) []Numbers {
	n := len(iterable)
	if r > n || r <= 0 {
		return nil
	}
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	perms := []Numbers{}
	perm := Numbers{}
	for _, idx := range indices[:r] {
		perm = append(perm, iterable[idx])
	}
	perms = append(perms, perm)
	for {
		var i int
		for i = r - 1; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				indices = append(indices[:i], append(indices[i+1:], indices[i:i+1]...)...)
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]
				perm := Numbers{}
				for _, idx := range indices[:r] {
					perm = append(perm, iterable[idx])
				}
				perms = append(perms, perm)
				break
			}
		}
		if i < 0 {
			break
		}
	}
	return perms
}

type Numbers []int

func allNumbers(digits int) []Numbers {
	numberList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return permutations(numberList, digits)
}

func (n Numbers) Feedback(estimate Numbers) *Feedback {
	f := NewFeedback(0, 0)
	for i, num := range estimate {
		if n[i] == num {
			f.IncHit()
			continue
		}
		if slices.Contains(n, num) {
			f.IncBlow()
		}
	}
	return f
}

type NumbersList []Numbers

func (n NumbersList) Contains(target Numbers) bool {
	for _, num := range n {
		if num.Equals(target) {
			return true
		}
	}
	return false
}

func (n Numbers) Equals(other Numbers) bool {
	if len(n) != len(other) {
		return false
	}
	for i := range n {
		if n[i] != other[i] {
			return false
		}
	}
	return true
}

type Feedback struct {
	hit  int
	blow int
}

func NewFeedback(hit int, blow int) *Feedback {
	return &Feedback{hit, blow}
}

func (f Feedback) String() string {
	return fmt.Sprintf("%dH%dB", f.hit, f.blow)
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

func NewHist(n Numbers, f *Feedback) *Hint {
	return &Hint{n, f}
}

type Solver struct {
	digit   int
	history []*Hint
}

func NewSolver(d int) *Solver {
	return &Solver{d, []*Hint{}}
}

func (s *Solver) IsEmpty() bool {
	return len(s.history) == 0
}

func (s *Solver) Push(h *Hint) {
	s.history = append(s.history, h)
}

func (s *Solver) Pop() {
	s.history = s.history[:len(s.history)-1]
}

func (s *Solver) Candidate() NumbersList {
	candidate := NumbersList{}
	for _, numbers := range allNumbers(s.digit) {
		var isBreak bool
		for _, hint := range s.history {
			if !reflect.DeepEqual(hint.feedback, numbers.Feedback(hint.numbers)) {
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

func (s *Solver) FeedbackSelect(e Numbers) []Feedback {
	return NewHistgram(e, s.Candidate()).Feedbacks()
}

func (s *Solver) BestEstimate() Numbers {
	mine := math.MaxFloat64
	minNumbers := Numbers{}
	allNumbers := allNumbers(s.digit)
	if s.IsEmpty() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return allNumbers[r.Intn(len(allNumbers))]
	}
	for _, estimate := range allNumbers {
		h := NewHistgram(estimate, s.Candidate())
		// 勝利条件
		if len(h.Feedbacks()) == 1 && reflect.DeepEqual(h.Feedbacks()[0], *NewFeedback(3, 0)) {
			return estimate
		}
		e := h.Entropy()
		if e < mine {
			mine = e
			minNumbers = estimate
		}
	}
	return minNumbers
}

type Histogram struct {
	mp map[Feedback]int
}

func NewHistgram(e Numbers, c []Numbers) *Histogram {
	mp := map[Feedback]int{}
	for _, candidate := range c {
		f := candidate.Feedback(e)
		mp[*f]++
	}
	return &Histogram{mp}
}

func (h *Histogram) Feedbacks() []Feedback {
	fs := []Feedback{}
	for f := range h.mp {
		fs = append(fs, f)
	}
	return fs
}

func (h *Histogram) Entropy() float64 {
	var e float64
	for _, val := range h.mp {
		absv := math.Abs(float64(val))
		e += absv * math.Log(1+absv)
	}
	return e
}
