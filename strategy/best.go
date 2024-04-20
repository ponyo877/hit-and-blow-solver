package strategy

import (
	"github.com/ponyo877/hit-and-blow-solver.git/entity"
)

type BestStrategy struct {
	history *entity.History
}

func NewBestStrategy(h *entity.History) *BestStrategy {
	return &BestStrategy{h}
}

func (l *BestStrategy) Estimate() entity.Numbers {
	return nil
}
