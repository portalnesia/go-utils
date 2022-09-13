package utils

import (
	"testing"
	"time"
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

func TestSlug(t *testing.T) {
	text := "Hello World, This Is From Testing"
	parse := Slug(text)

	if parse != "hello-world-this-is-from-testing" {
		t.Errorf("Invalid TestSlug. Get: %s", parse)
	}
}

func TestNumberSize(t *testing.T) {
	size := 50486525485

	parse := NumberSize(float64(size), 2)

	if parse != "5.05 GB" {
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

func TestTimeAgo(t *testing.T) {
	times := time.Now()
	date1 := times.Add(-11 * time.Second)
	date2 := times.Add(-11 * time.Minute)
	date3 := times.Add(-5 * time.Hour)
	date4 := times.AddDate(0, 0, -6)
	date5 := times.AddDate(0, -2, 0)
	date6 := times.AddDate(-2, 0, 0)

	parse1 := TimeAgo(date1.Unix())
	parse2 := TimeAgo(date2.Unix())
	parse3 := TimeAgo(date3.Unix())
	parse4 := TimeAgo(date4.Unix())
	parse5 := TimeAgo(date5.Unix())
	parse6 := TimeAgo(date6.Unix())

	if parse1 != "a few seconds ago" {
		t.Errorf("Invalid TimeAgo less minutes. Get: %s", parse1)
	}
	if parse2 != "11 minutes ago" {
		t.Errorf("Invalid TimeAgo minutes ago. Get: %s", parse2)
	}
	if parse3 != "5 hours ago" {
		t.Errorf("Invalid TimeAgo hours ago. Get: %s", parse3)
	}
	if parse4 != "6 days ago" {
		t.Errorf("Invalid TimeAgo days ago. Get: %s", parse4)
	}
	if parse5 != "2 months ago" {
		t.Errorf("Invalid TimeAgo months ago. Get: %s", parse5)
	}
	if parse6 != "2 years ago" {
		t.Errorf("Invalid TimeAgo years ago. Get: %s", parse6)
	}
}

func TestSeparateNumber(t *testing.T) {
	var number int64 = 25000

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
	var n int64 = 5025
	parse := NumberFormatShort(n)

	if parse.Format != "5.03 K" {
		t.Errorf("Invalid NumberFormatShort 5025. Get: %s", parse.Format)
	}

	n = 64768456
	parse = NumberFormatShort(n)

	if parse.Format != "64.77 M" {
		t.Errorf("Invalid NumberFormatShort 64768456. Get: %s", parse.Format)
	}

	n = 1065201025
	parse = NumberFormatShort(n)

	if parse.Format != "1.07 B" {
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
	d := "true"
	parse1 := IsTrue(d)
	d = "sagasg"
	parse2 := IsTrue(d)

	e := 5
	parse3 := IsTrue(e)
	e = 1
	parse4 := IsTrue(e)

	f := false
	parse5 := IsTrue(f)
	f = true
	parse6 := IsTrue(f)

	if !parse1 {
		t.Errorf("Invalid IsTrue `true`. Get: %v", parse1)
	}
	if parse2 {
		t.Errorf("Invalid IsTrue `sagasg`. Get: %v", parse2)
	}
	if parse3 {
		t.Errorf("Invalid IsTrue 5. Get: %v", parse3)
	}
	if !parse4 {
		t.Errorf("Invalid IsTrue 1. Get: %v", parse4)
	}
	if parse5 {
		t.Errorf("Invalid IsTrue false. Get: %v", parse5)
	}
	if !parse6 {
		t.Errorf("Invalid IsTrue true. Get: %v", parse6)
	}
}

func TestNewGoment(t *testing.T) {
	var (
		stringFormat string = "2020-02-02 12:00:00" // UTC 0000
		unixFormat   int64  = 1580644800
	)

	a := NewGoment(unixFormat)

	f := a.Format("YYYY-MM-DD HH:mm:ss")
	if f != stringFormat {
		t.Errorf("[NewGoment] Error Format Date. Get: %s", f)
	}

	u := a.ToUnix()
	if u != unixFormat {
		t.Errorf("[NewGoment] Error Unix Date. Get: %v", u)
	}
}
