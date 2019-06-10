package main

import (
	"reflect"
	"testing"
)

func TestPhoneNumberNormalizer(t *testing.T) {
	input := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	output := map[string]int{
		"1234567890": 2,
		"1234567891": 1,
		"1234567892": 3,
		"1234567893": 1,
		"1234567894": 1,
	}

	result := Normalizer(input)

	if !reflect.DeepEqual(result, output) {
		t.Errorf("error is not match % v : % v", result, output)
	}
}
