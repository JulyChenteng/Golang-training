package defs

import (
	"fmt"
	"testing"
)

func TestScore_ToString(t *testing.T) {
	score := Score{"20180661782", 66}
	fmt.Println(len(score.ToString()))
}
