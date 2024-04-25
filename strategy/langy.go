package strategy

import (
	"math"

	"github.com/ponyo877/hit-and-blow-solver/entity"
)

type LandyStrategy struct {
	solver  *entity.Solver
	history *entity.History
}

func NewLandyStrategy(s *entity.Solver, h *entity.History) *LandyStrategy {
	return &LandyStrategy{s, h}
}

func (l *LandyStrategy) Estimate(init entity.Numbers) entity.Numbers {
	mine := math.MaxFloat64
	minNumbers := entity.Numbers{}
	if l.history.IsEmpty() {
		return init
	}
	ca := l.history.Candidate()
	// Victory conditions
	if len(ca) == 1 {
		return ca[0]
	}
	if len(ca) == 2 {
		return ca[0]
	}
	for _, estimate := range l.solver.AllPatterns() {
		e := entity.NewHistgram(l.solver, estimate, ca).Entropy()
		if e < mine {
			mine = e
			minNumbers = estimate
		}
	}
	return minNumbers
}
