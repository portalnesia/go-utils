package utils

import (
	"testing"
)

func TestTruncate(t *testing.T) {
	text := "Hello world, this is from testing"

	truncate := Truncate(text, 11)

	if truncate != "Hello world..." {
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

	if truncate != "Hello world..." {
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

func TestNumberSize(t *testing.T) {
	size := 50486525485

	parse := NumberSize(float64(size), 2)

	if parse != "5.05 GB" {
		t.Errorf("Invalid NumberSize. Get: %s", parse)
	}
}
