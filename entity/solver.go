package entity

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
	s.cache = permutations(s.numbers, s.Digit())
	return s.cache
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
