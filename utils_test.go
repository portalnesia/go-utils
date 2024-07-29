/*
 * Copyright (c) Portalnesia - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package utils

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/google/uuid"
)

func TestTruncate(t *testing.T) {
	text := "Hello world, this is from testing"

	truncate := Truncate(text, 11)

	if truncate != "Hello wo..." {
		t.Errorf("Invalid truncate. Get: %s", truncate)
	}
}

func TestClean(t *testing.T) {
	text := "<p>Hello world</p>"

	truncate := Clean(text)

	if truncate != "Hello world" {
		t.Errorf("Invalid clean. Get: %s", truncate)
	}
}

func TestCleanAndTruncate(t *testing.T) {
	text := "<p>Hello world, this is from testing</p>"

	truncate := CleanAndTruncate(text, 11)

	if truncate != "Hello wo..." {
		t.Errorf("Invalid clean and truncate. Get: %s", truncate)
	}
}

func TestParseUrl(t *testing.T) {
	url := "https://www.portalnesia.com/news?foo=bar"

	parser, err := ParseUrl(url)

	if err != nil {
		t.Error(err)
	}

	if parser != "portalnesia.com/news?foo=bar" {
		t.Errorf("Invalid ParseUrl. Get: %s", parser)
	}

	if parser, err = ParseUrl("error url"); err == nil {
		t.Errorf("This url must be error: %s", parser)
	}

	if parser, err = ParseUrl("https://err. https://"); err == nil {
		t.Errorf("This url must be error: %s", parser)
	}
}

func TestUcwords(t *testing.T) {
	text := "hello world, this is from testing"

	parse := Ucwords(text)

	if parse != "Hello World, This Is From Testing" {
		t.Errorf("Invalid Ucwords. Get: %s", parse)
	}
}

func TestFirstLetter(t *testing.T) {
	text := "hello world, this is from testing"
	parse := FirstLetter(text, 0)

	if parse != "HWTIFT" {
		t.Errorf("Invalid FirstLetter All Word. Get: %s", parse)
	}

	parse = FirstLetter(text, 3)

	if parse != "HWT" {
		t.Errorf("Invalid FirstLetter 3 Word. Get: %s", parse)
	}
}

func TestSlug(t *testing.T) {
	text := "Hello World, This Is From Testing"
	parse := Slug(text)

	if parse != "hello-world-this-is-from-testing" {
		t.Errorf("Invalid TestSlug. Get: %s", parse)
	}
}

func TestNumberSize(t *testing.T) {
	size := 50486525485

	if parse := NumberSize(float64(size), 2); parse != "47.02 GB" {
		t.Errorf("Invalid NumberSize. Get: %s", parse)
	}

	if parse := NumberSize(float64(18037807), 2); parse != "17.20 MB" {
		t.Errorf("Invalid NumberSize. Get: %s", parse)
	}

	if parse := NumberSize(float64(18037807), 0); parse != "17.20 MB" {
		t.Errorf("Invalid NumberSize. Get: %s", parse)
	}

	if parse := NumberSize(float64(0), 0); parse != "-" {
		t.Errorf("Invalid NumberSize. Get: %s", parse)
	}
}

func TestNanoId(t *testing.T) {
	parse := NanoId()

	if parse == "" {
		t.Errorf("Invalid NanoId. Get: %s", parse)
	}
	if len(parse) != 21 {
		t.Errorf("Invalid NanoId length. Get: %s", parse)
	}

	parse = NanoId(50)
	if parse == "" {
		t.Errorf("Invalid NanoId. Get: %s", parse)
	}
	if len(parse) != 50 {
		t.Errorf("Invalid NanoId length. Get: %s", parse)
	}
}

func TestUUID(t *testing.T) {

	uid := UUID()

	if _, err := uuid.Parse(uid); err != nil {
		t.Error(err)
	}
}

func TestNanoIdStr(t *testing.T) {
	parse := NanoIdStr("ABCDEFGHIJKLMN")

	if parse == "" {
		t.Errorf("Invalid NanoIdStr. Get: %s", parse)
	}
	if len(parse) != 20 {
		t.Errorf("Invalid NanoId length. Get: %s", parse)
	}

	if regexp.MustCompile(`[0-9a-z]]`).MatchString(parse) {
		t.Errorf("Invalid NanoId alphabet. Get: %s", parse)
	}

	parse = NanoIdStr("123456789", 50)
	if parse == "" {
		t.Errorf("Invalid NanoId. Get: %s", parse)
	}
	if len(parse) != 50 {
		t.Errorf("Invalid NanoId length. Get: %s", parse)
	}

	if regexp.MustCompile(`[a-zA-Z]]`).MatchString(parse) {
		t.Errorf("Invalid NanoId alphabet. Get: %s", parse)
	}
}

func TestSeparateNumber(t *testing.T) {
	var number float64 = 25000

	parse := SeparateNumber(number)
	if parse != "25,000" {
		t.Errorf("Invalid SeparateNumber 25000. Get: %s", parse)
	}

	number = 50252574753
	parse = SeparateNumber(number)
	if parse != "50,252,574,753" {
		t.Errorf("Invalid SeparateNumber 50252574753. Get: %s", parse)
	}
}

func TestIsUrl(t *testing.T) {
	url := "astgikaewlog"
	parse := IsUrl(url)

	if parse {
		t.Errorf("Invalid IsUrl astgikaewlog. Get: %t", parse)
	}

	url = "http://portalnesia.com/contact"
	parse = IsUrl(url)

	if !parse {
		t.Errorf("Invalid IsUrl http://portalnesia.com/contact. Get: %t", parse)
	}

	url = "https://portalnesia.com/contact"
	parse = IsUrl(url)

	if !parse {
		t.Errorf("Invalid IsUrl https://portalnesia.com/contact. Get: %t", parse)
	}
}

func TestIsTwitterUrl(t *testing.T) {
	url := "astgikaewlog"
	parse := IsTwitterUrl(url)

	if parse {
		t.Errorf("Invalid IsTwitterUrl astgikaewlog. Get: %t", parse)
	}

	url = "http://portalnesia.com/contact"
	parse = IsTwitterUrl(url)

	if parse {
		t.Errorf("Invalid IsTwitterUrl http://portalnesia.com/contact. Get: %t", parse)
	}

	url = "http://portalnesia.com/twitter.com/contact"
	parse = IsTwitterUrl(url)

	if parse {
		t.Errorf("Invalid IsTwitterUrl http://portalnesia.com/twitter.com/contact. Get: %t", parse)
	}

	url = "https://twitter.com/Portalnesia1"
	parse = IsTwitterUrl(url)

	if !parse {
		t.Errorf("Invalid IsTwitterUrl https://twitter.com/Portalnesia1. Get: %t", parse)
	}
}

func TestFirstToUpper(t *testing.T) {
	text := "hello world, this is from testing"
	parse := FirstToUpper(text)

	if parse != "Hello world, this is from testing" {
		t.Errorf("Invalid FirstLetter All Word. Get: %s", parse)
	}
}

func TestNumberFormatShort(t *testing.T) {
	if parse := NumberFormatShort(50); parse.Format != "50" {
		t.Errorf("Invalid NumberFormatShort 50. Get: %s", parse.Format)
	}

	if parse := NumberFormatShort(5025); parse.Format != "5.03 K" {
		t.Errorf("Invalid NumberFormatShort 5025. Get: %s", parse.Format)
	}

	if parse := NumberFormatShort(64768456); parse.Format != "64.77 M" {
		t.Errorf("Invalid NumberFormatShort 64768456. Get: %s", parse.Format)
	}

	if parse := NumberFormatShort(1065201025); parse.Format != "1.07 B" {
		t.Errorf("Invalid NumberFormatShort 1065201025. Get: %s", parse.Format)
	}

	if parse := NumberFormatShort(6065201025456); parse.Format != "6.07 T" {
		t.Errorf("Invalid NumberFormatShort 1065201025. Get: %s", parse.Format)
	}
}

func TestValidateEmail(t *testing.T) {
	e := "support@portalnesia"
	parse := ValidateEmail(e)

	if parse {
		t.Errorf("Invalid ValidateEmail support@portalnesia. Get: %v", parse)
	}

	e = "support@portalnesia.com"
	parse = ValidateEmail(e)

	if !parse {
		t.Errorf("Invalid ValidateEmail support@portalnesia.com. Get: %v", parse)
	}
}

func TestIsTrue(t *testing.T) {
	tests := []struct {
		value  interface{}
		expect bool
	}{
		{"true", true},
		{"random", false},
		{5, false},
		{1, true},
		{false, false},
		{true, true},
		{int8(1), true},
		{int8(0), false},
		{int16(1), true},
		{int16(0), false},
		{int32(1), true},
		{int32(0), false},
		{int64(1), true},
		{int64(0), false},
		{uint(1), true},
		{uint(0), false},
		{uint8(1), true},
		{uint8(0), false},
		{uint16(1), true},
		{uint16(0), false},
		{uint32(1), true},
		{uint32(0), false},
		{uint64(1), true},
		{uint64(0), false},
		{float32(1), true},
		{float32(0), false},
		{float64(1), true},
		{float64(0), false},
		{struct{ custom string }{custom: "string"}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("TestIsTrue for value: %v", tt.value), func(t *testing.T) {
			if isTrue := IsTrue(tt.value); isTrue != tt.expect {
				t.Errorf("Invalid return for value: %v", tt.value)
			}
		})
	}
}

func TestTernary(t *testing.T) {
	shouldTrue := Ternary(true, "true", "false")
	if shouldTrue != "true" {
		t.Errorf("Invalid ternary; get :%v", shouldTrue)
	}

	shouldFalse := Ternary(false, "true", "false")
	if shouldFalse != "false" {
		t.Errorf("Invalid ternary; get :%v", shouldFalse)
	}
}
