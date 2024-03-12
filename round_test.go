package round

import (
	"testing"
)

func TestRound(t *testing.T) {
	r := genRandomFloat(1, 9)
	t.Logf("Generate random float64 number(1-9): %f\n", r)

	t.Logf("after round: %.3f", Ceil(r))
	t.Logf("rand: %.3f", r)

	if Round(r) == r {
		t.Errorf("Test Round Failed: %f  is equal %f", r, Round(r))
	} else {
		t.Log("Test Round Passed!!!")
	}
}

func TestCeil(t *testing.T) {
	r := genRandomFloat(1, 9)
	t.Logf("Generate random float64 number(1-9): %f\n", r)

	t.Logf("after ceil: %f", Ceil(r))
	t.Logf("rand: %f", r)

	if Ceil(r) == r {
		t.Errorf("Test Ceil Failed: %f  is equal %f", r, Ceil(r))
	} else {
		t.Log("Test Ceil Passed!!!")
	}
}

func TestFloor(t *testing.T) {
	r := genRandomFloat(1, 9)
	t.Logf("Generate random float64 number(1-9): %f\n", r)

	t.Logf("after floor: %f", Ceil(r))
	t.Logf("rand: %f", r)

	if Floor(r) == r {
		t.Errorf("Test Floor Failed: %f  is equal %f", r, Floor(r))
	} else {
		t.Log("Test Floor Passed!!!")
	}
}
