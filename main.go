package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/ponyo877/hit-and-blow-solver.git/entity"
	"github.com/ponyo877/hit-and-blow-solver.git/strategy"
)

func main() {
	d := 3
	history := entity.NewHistory(d)
	var i int
	for i = 0; i <= 10; i++ {
		estimate := strategy.NewLandyStrategy(history).Estimate()
		ca := history.Candidate()
		fs := history.FeedbackSelect(estimate)
		if len(ca) == 1 {
			fmt.Println("You win!: ", ca[0])
			break
		}
		if !ca.Contains(estimate) && len(ca) <= 4 {
			fmt.Println("candidates: ", ca)
		}
		fmt.Println("estimate: ", estimate)
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "👉 {{ .Hit | cyan }}H{{ .Blow | red }}B",
			Inactive: "  {{ .Hit | cyan }}H{{ .Blow | red }}B",
			Selected: "👉 {{ .Hit | cyan }}H{{ .Blow | red }}B",
		}

		searcher := func(input string, index int) bool {
			feedback := fs[index]
			name := strings.Replace(strings.ToLower(feedback.String()), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		}

		prompt := promptui.Select{
			Label:     "Feedbacks",
			Items:     fs,
			Templates: templates,
			Size:      14,
			Searcher:  searcher,
		}
		i, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose Feedback: %s\n", fs[i].String())
		h := entity.NewHint(estimate, &fs[i])
		history.Push(h)
	}
	// measure()
}

func measure() {
	ansHist := map[int]int{}
	sum := 0
	d := 3
	all := entity.AllNumbers(d)
	for _, answer := range all {
		// fmt.Printf("%d, ", answer)
		history := entity.NewHistory(d)
		var i int
		for i = 1; i <= 10; i++ {
			estimate := strategy.NewLandyStrategy(history).Estimate()
			f := answer.Feedback(estimate)
			if reflect.DeepEqual(f, entity.NewFeedback(d, 0)) {
				break
			}
			h := entity.NewHint(estimate, f)
			history.Push(h)
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
