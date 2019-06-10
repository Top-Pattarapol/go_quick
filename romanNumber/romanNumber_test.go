package main

import (
	"fmt"
	"testing"
)

type romanNum struct {
	input  int
	output string
}

func testNumber(t *testing.T, data romanNum) {
	name := fmt.Sprintf("Number %d", data.input)
	t.Run(name, func(t *testing.T) {
		result := Roman(data.input)
		if result != data.output {
			t.Errorf("Roman number not match %s : %s", result, data.output)
		}
	})
}

func TestRomanNumber(t *testing.T) {

	testNumber(t, romanNum{1, "I"})

	testNumber(t, romanNum{4, "IV"})

	testNumber(t, romanNum{5, "V"})

	testNumber(t, romanNum{8, "VIII"})

	testNumber(t, romanNum{9, "IX"})

	testNumber(t, romanNum{10, "X"})

	testNumber(t, romanNum{11, "XI"})

	testNumber(t, romanNum{15, "XV"})

	testNumber(t, romanNum{18, "XVIII"})

	testNumber(t, romanNum{20, "XX"})

	testNumber(t, romanNum{23, "XXIII"})

	testNumber(t, romanNum{28, "XXVIII"})

	testNumber(t, romanNum{34, "XXXIV"})

	testNumber(t, romanNum{39, "XXXIX"})

	testNumber(t, romanNum{40, "XL"})

	testNumber(t, romanNum{45, "XLV"})

	testNumber(t, romanNum{50, "L"})

	testNumber(t, romanNum{55, "LV"})

	testNumber(t, romanNum{60, "LX"})

	testNumber(t, romanNum{65, "LXV"})

	testNumber(t, romanNum{70, "LXX"})

	testNumber(t, romanNum{75, "LXXV"})

	testNumber(t, romanNum{80, "LXXX"})

	testNumber(t, romanNum{85, "LXXXV"})

	testNumber(t, romanNum{90, "XC"})

	testNumber(t, romanNum{92, "XCII"})

	testNumber(t, romanNum{95, "XCV"})

	testNumber(t, romanNum{99, "XCIX"})

	testNumber(t, romanNum{100, "C"})
}
