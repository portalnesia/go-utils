package goment

import (
	"fmt"
	"testing"
)

var (
	stringFormat string = "2020-02-02 12:00:00" // UTC 0000
	unixFormat   int64  = 1580644800
)

func TestNewStringFormat(t *testing.T) {
	a, err := New(stringFormat) // UTC 0000

	if err != nil {
		t.Errorf("[New] %+v", err)
	}
	f := a.Format("YYYY-MM-DD HH:mm:ss")
	if f != stringFormat {
		t.Errorf("[New] Error Format Date. Get: %s", f)
	}

	u := a.ToUnix()
	if u != unixFormat {
		t.Errorf("[New] Error Unix Date. Get: %v", u)
	}
}

func TestUnixFormat(t *testing.T) {
	a, err := New(unixFormat)

	if err != nil {
		t.Errorf("[New] %+v", err)
	}
	f := a.Format("YYYY-MM-DD HH:mm:ss")
	if f != stringFormat {
		t.Errorf("[New] Error Format Date. Get: %s", f)
	}

	u := a.ToUnix()
	if u != unixFormat {
		t.Errorf("[New] Error Unix Date. Get: %v", u)
	}
}

func TestTimeagoYear(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Errorf("[TimeagoYear] %+v", err)
	}

	a.Subtract(2, "years")
	b := a.TimeAgo()

	if b.Format != a.Format("DD MMM YYYY") {
		t.Errorf("[TimeagoYear] Timestamp. Get: %v", b.Format)
	}
}

func TestTimeagoMonth(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Errorf("[TimeagoMonth] %+v", err)
	}
	a.Subtract(4, "month")
	b := a.TimeAgo()

	if b.Format != "4 months ago" {
		t.Errorf("[TimeagoMonth] Timestamp. Get: %v", b.Format)
	}
}

func TestTimeagoDay(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Errorf("[TimeagoDay] %+v", err)
	}
	a.Subtract(7, "days")
	b := a.TimeAgo()

	if b.Format != "7 days ago" {
		t.Errorf("[TimeagoDay] Timestamp. Get: %v", b.Format)
	}
}

func TestTimeagoMinute(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Errorf("[TimeagoDay] %+v", err)
	}
	a.Subtract(7, "minutes")
	b := a.TimeAgo()

	if b.Format != "7 minutes ago" {
		t.Errorf("[TimeagoDay] Timestamp. Get: %v", b.Format)
	}
}

func TestTimeagoSecond(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Errorf("[TimeagoSecond] %+v", err)
	}
	a.Subtract(7, "seconds")
	b := a.TimeAgo()

	if b.Format != "a few seconds ago" {
		t.Errorf("[TimeagoSecond] Timestamp. Get: %v", b.Format)
	}
}

func TestFormatMinimal(t *testing.T) {
	a, err := New(stringFormat)
	if err != nil {
		t.Errorf("[FormatMinimal] %+v", err)
	}
	b := a.PNformat("minimal")
	if b != "02 Feb 2020" {
		t.Errorf("[FormatMinimal]. Get: %v", b)
	}
}

func TestFormatFulldate(t *testing.T) {
	a, err := New(stringFormat)
	if err != nil {
		t.Errorf("[FormatFulldate] %+v", err)
	}
	b := a.PNformat("fulldate")
	if b != "02 February 2020" {
		t.Errorf("[FormatFulldate]. Get: %v", b)
	}
}

func TestFormatFull(t *testing.T) {
	a, err := New(stringFormat)
	if err != nil {
		t.Errorf("[FormatFull] %+v", err)
	}
	b := a.PNformat("full")
	if b != "02 February 2020, 12:00" {
		t.Errorf("[FormatFull]. Get: %v", b)
	}
}

func TestFormatISO8601(t *testing.T) {
	a, err := New(stringFormat)
	if err != nil {
		t.Errorf("[FormatISO8601] %+v", err)
	}
	b := a.PNformat()
	if b != stringFormat {
		t.Errorf("[FormatISO8601]. Get: %v", b)
	}
}

func TestRangeFormatYear(t *testing.T) {
	a, err := New(stringFormat)
	if err != nil {
		t.Errorf("[RangeFormatYear] %+v", err)
	}
	b, err := New()
	if err != nil {
		t.Errorf("[RangeFormatYear] %+v", err)
	}

	r := a.RangeFormat(b)

	if r != fmt.Sprintf("%s - %s", a.Format("DD MMMM YYYY"), b.Format("DD MMMM YYYY")) {
		t.Errorf("[RangeFormatYear]. Get: %s", r)
	}
}

func TestRangeFormatMonth(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Errorf("[RangeFormatMonth] %+v", err)
	}
	a.Subtract(2, "months")
	b, err := New()
	if err != nil {
		t.Errorf("[RangeFormatMonth] %+v", err)
	}

	r := a.RangeFormat(b)

	if r != fmt.Sprintf("%s - %s %d", a.Format("DD MMMM"), b.Format("DD MMMM"), a.Year()) {
		t.Errorf("[RangeFormatMonth]. Get: %s", r)
	}
}

func TestRangeFormatDay(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Errorf("[RangeFormatDay] %+v", err)
	}
	a.Subtract(2, "days")
	b, err := New()
	if err != nil {
		t.Errorf("[RangeFormatDay] %+v", err)
	}

	r := a.RangeFormat(b)

	if r != fmt.Sprintf("%s - %s %s", a.Format("DD"), b.Format("DD"), a.Format("MMMM YYYY")) {
		t.Errorf("[RangeFormatDay]. Get: %s", r)
	}
}

func TestRangeFormatTime(t *testing.T) {
	a, err := New()
	if err != nil {
		t.Errorf("[RangeFormatTime] %+v", err)
	}
	a.Subtract(2, "hours")
	b, err := New()
	if err != nil {
		t.Errorf("[RangeFormatTime] %+v", err)
	}

	r := a.RangeFormat(b)

	if r != fmt.Sprintf("%s - %s, %s", a.Format("HH:mm"), b.Format("HH:mm"), a.Format("DD MMMM YYYY")) {
		t.Errorf("[RangeFormatTime]. Get: %s", r)
	}
}
