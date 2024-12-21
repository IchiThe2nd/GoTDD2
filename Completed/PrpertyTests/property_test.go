package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{500, "D"},
	{1000, "M"},
	{900, "CM"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestRomanNumeral(t *testing.T) {

	for _, test := range cases {
		t.Run(fmt.Sprintf("Arabic %d gets convertered to Roman %q ", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			//fmt.Println(test.Roman)
			if got != test.Roman {
				t.Errorf("got %q wanted %q", got, test.Roman)
			}
		})

	}
}
func TestArabic(t *testing.T) {
	for _, test := range cases[:4] {

		t.Run(fmt.Sprintf("Roman %q convertered to Arabic %d ", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			//fmt.Println(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d wanted %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {

		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fmt.Println(roman)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed check", err)
	}
}
