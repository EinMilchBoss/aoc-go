package main

type pair struct {
	from int
	to   int
}

func (left pair) includesAll(right pair) bool {
	return left.from >= right.from && right.to >= left.to
}

func (left pair) includesAny(right pair) bool {
	for l := left.from; l <= left.to; l++ {
		for r := right.from; r <= right.to; r++ {
			if l == r {
				return true
			}
		}
	}
	return false
}
