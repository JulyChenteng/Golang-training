package defs

import "strconv"

type Score struct {
	StuID string
	Grade int
}

func (score Score) Hash() int {
	h := 0

	for _, ch := range []byte(score.StuID) {
		h = 5*h + int(ch)
	}

	return h
}

func (score Score) ToString() string {
	grade := strconv.Itoa(score.Grade)

	return string(score.StuID + " " + grade + "\n")
}
