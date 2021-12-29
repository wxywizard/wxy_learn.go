package test

import "testing"

func TestFatRateCase1(t *testing.T) {
	var expectedFatRate float64 = 0.1784
	fatRate, err := CalcFatRate(23, 28, "男")
	if err != nil {
		t.Error("should be error, but err returned is nil")
	}
	if fatRate != expectedFatRate {
		t.Error("expected is different from calculated")
	}
}

//case2 第一种情况
func TestFatRateCase2(t *testing.T) {
	_, err := CalcFatRate(0, 30, "男")
	if err == nil {
		t.Error("should be error, but err returned is nil")
	}
}

//case2 第二种情况
func TestFatRateCase3(t *testing.T) {
	_, err := CalcFatRate(23, 300, "男")
	if err == nil {
		t.Error("should be error, but err returned is nil")
	}
}

//case2 第三种情况
func TestFatRateCase4(t *testing.T) {
	_, err := CalcFatRate(23, 30, "")
	if err == nil {
		t.Error("should be error, but err returned is nil")
	}
}

//case3
func TestFatRateCase5(t *testing.T) {
	var expectedSuggest string = "目前是：偏重。吃完饭多散散步，消化消化"
	bmi, err2 := BMI(80, 1.83)
	if err2 != nil {
		t.Error("bmi should be error, but err returned is nil")
	}

	fatRate, err := CalcFatRate(bmi, 28, "男")
	if err != nil {
		t.Error("fatRate should be error, but err returned is nil")
	}
	man := switchMan(fatRate, 28)

	if man != expectedSuggest {
		t.Error("expected is different from suggest")
	}
}
