package main

import "errors"

type StaticInvokeExample struct{}

func (s *StaticInvokeExample) MaxForThree(x int, y int16, z int8) int {
	max := s.maxForTwo(x, y)

	if max > int(z) {
		return max
	} else {
		return int(z)
	}
}

func (s *StaticInvokeExample) maxForTwo(x int, y int16) int {
	if x > int(y) {
		return x
	} else {
		return int(y)
	}
}

type SimpleFormulaExample struct{}

func (s *SimpleFormulaExample) mult(a, b int) int {
	return a * b
}

func (s *SimpleFormulaExample) half(a int) int {
	return a / 2
}

func (s *SimpleFormulaExample) SimpleFormula(fst, snd int) (int, error) {
	if fst < 100 {
		return 0, errors.New("IllegalArgumentException: fst < 100")
	} else if snd < 100 {
		return 0, errors.New("IllegalArgumentException: snd < 100")
	}

	x := fst + 5
	y := s.half(snd)

	return s.mult(x, y), nil
}
