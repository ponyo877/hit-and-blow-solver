package entity

import "slices"

type Numbers []int

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

type NumbersList []Numbers

func (n NumbersList) Contains(target Numbers) bool {
	for _, num := range n {
		if num.Equals(target) {
			return true
		}
	}
	return false
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

func AllNumbers(digits int) []Numbers {
	numberList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return permutations(numberList, digits)
}
