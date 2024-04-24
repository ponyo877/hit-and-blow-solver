package strategy

import (
	"math"
	"math/rand"
	"time"

	"github.com/ponyo877/hit-and-blow-solver.git/entity"
)

type LandyStrategy struct {
	solver  *entity.Solver
	history *entity.History
}

func NewLandyStrategy(s *entity.Solver, h *entity.History) *LandyStrategy {
	return &LandyStrategy{s, h}
}

func (l *LandyStrategy) Estimate() entity.Numbers {
	mine := math.MaxFloat64
	minNumbers := entity.Numbers{}
	allPatterns := l.solver.AllPatterns()
	if l.history.IsEmpty() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return allPatterns[r.Intn(len(allPatterns))]
	}
	ca := l.history.Candidate()
	// Victory conditions
	if len(ca) == 1 {
		return ca[0]
	}
	if len(ca) == 2 {
		return ca[0]
	}
	for _, estimate := range allPatterns {
		e := entity.NewHistgram(l.solver, estimate, ca).Entropy()
		if e < mine {
			mine = e
			minNumbers = estimate
		}
	}
	return minNumbers
}
