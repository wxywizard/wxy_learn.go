package test

import (
	"testing"
)

func TestBMICase1(t *testing.T) {
	var expectedBmi float64 = 2
	bmi, err := BMI(2, 1)
	if err != nil {
		t.Error("should be error, but err returned is nil")
	}
	if bmi != expectedBmi {
		t.Error("expected is different from calculated")
	}
}

func TestBMICase2(t *testing.T) {
	_, err := BMI(1, 0)
	if err == nil {
		t.Error("should be error, but err returned is nil")
	}
}

func TestBMICase3(t *testing.T) {
	_, err := BMI(-1, 1)
	if err == nil {
		t.Error("should be error, but err returned is nil")
	}
}
