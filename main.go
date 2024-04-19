package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/ponyo877/hit-and-blow-solver.git/entity"
	"github.com/ponyo877/hit-and-blow-solver.git/strategy"
)

func main() {
	ansHist := map[int]int{}
	sum := 0
	d := 3
	initNumbers := entity.Numbers{0, 1, 2}
	all := entity.AllNumbers(d)
	for _, answer := range all {
		estimate := initNumbers
		history := entity.NewHistory(d)
		var i int
		for i = 1; i <= 10; i++ {
			f := answer.Feedback(estimate)
			if reflect.DeepEqual(f, entity.NewFeedback(3, 0)) {
				break
			}
			h := entity.NewHint(estimate, f)
			history.Push(h)
			estimate = strategy.NewLandyStrategy(history).Estimate()
			// ca := s.Candidate()
			// fs := s.FeedbackSelect(estimate)
			// if !ca.Contains(be) && len(ca) <= 4 {
			// 	fmt.Println(ca)
			// }
			// fmt.Println(be)
			// for _, f := range fs {
			// 	fmt.Println(f)
			// }
		}
		sum += i
		ansHist[i]++
	}
	s := []int{}
	for k := range ansHist {
		s = append(s, k)
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	for _, k := range s {
		fmt.Printf("%2d: %3d\n", k, ansHist[k])
	}
	fmt.Printf("avg: %6.4f\n", float64(sum)/float64(len(all)))
}

type EstimateStrategy interface {
	Estimate() entity.Numbers
}
