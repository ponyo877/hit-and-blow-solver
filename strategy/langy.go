package strategy

import (
	"math"
	"math/rand"
	"time"

	"github.com/ponyo877/hit-and-blow-solver.git/entity"
)

type LandyStrategy struct {
	history *entity.History
}

func NewLandyStrategy(h *entity.History) *LandyStrategy {
	return &LandyStrategy{h}
}

func (l *LandyStrategy) Estimate() entity.Numbers {
	mine := math.MaxFloat64
	minNumbers := entity.Numbers{}
	allNumbers := entity.AllNumbers(l.history.Digit())
	if l.history.IsEmpty() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return allNumbers[r.Intn(len(allNumbers))]
	}
	ca := l.history.Candidate()
	// 勝利条件
	if len(ca) == 1 {
		return ca[0]
	}
	for _, estimate := range allNumbers {
		e := entity.NewHistgram(estimate, ca).Entropy()
		if e < mine {
			mine = e
			minNumbers = estimate
		}
	}
	return minNumbers
}
