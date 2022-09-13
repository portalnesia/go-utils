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
	a := Must(stringFormat) // UTC 0000

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
	a := Must(unixFormat)

	f := a.Format("YYYY-MM-DD HH:mm:ss")
	if f != stringFormat {
		t.Errorf("[New] Error Format Date. Get: %s", f)
	}

	u := a.ToUnix()
	if u != unixFormat {
		t.Errorf("[New] Error Unix Date. Get: %v", u)
	}

	a, err := New(unixFormat)

	if err != nil {
		t.Errorf("[New] %+v", err)
	}

	f = a.Format("YYYY-MM-DD HH:mm:ss")
	if f != stringFormat {
		t.Errorf("[New] Error Format Date. Get: %s", f)
	}

	u = a.ToUnix()
	if u != unixFormat {
		t.Errorf("[New] Error Unix Date. Get: %v", u)
	}
}

func TestTimeagoYear(t *testing.T) {
	a := Must()

	a.Subtract(2, "years")
	b := a.TimeAgo()

	if b.Format != a.Format("DD MMM YYYY") {
		t.Errorf("[TimeagoYear] Timestamp. Get: %v", b.Format)
	}
}

func TestTimeagoMonth(t *testing.T) {
	a := Must()
	a.Subtract(4, "month")
	b := a.TimeAgo()

	if b.Format != "4 months ago" {
		t.Errorf("[TimeagoMonth] Timestamp. Get: %v", b.Format)
	}
}

func TestTimeagoDay(t *testing.T) {
	a := Must()

	a.Subtract(7, "days")
	b := a.TimeAgo()

	if b.Format != "7 days ago" {
		t.Errorf("[TimeagoDay] Timestamp. Get: %v", b.Format)
	}
}

func TestTimeagoMinute(t *testing.T) {
	a := Must()

	a.Subtract(7, "minutes")
	b := a.TimeAgo()

	if b.Format != "7 minutes ago" {
		t.Errorf("[TimeagoDay] Timestamp. Get: %v", b.Format)
	}
}

func TestTimeagoSecond(t *testing.T) {
	a := Must()

	a.Subtract(7, "seconds")
	b := a.TimeAgo()

	if b.Format != "a few seconds ago" {
		t.Errorf("[TimeagoSecond] Timestamp. Get: %v", b.Format)
	}
}

func TestFormatMinimal(t *testing.T) {
	a := Must(stringFormat)

	b := a.PNformat("minimal")
	if b != "02 Feb 2020" {
		t.Errorf("[FormatMinimal]. Get: %v", b)
	}
}

func TestFormatFulldate(t *testing.T) {
	a := Must(stringFormat)

	b := a.PNformat("fulldate")
	if b != "02 February 2020" {
		t.Errorf("[FormatFulldate]. Get: %v", b)
	}
}

func TestFormatFull(t *testing.T) {
	a := Must(stringFormat)

	b := a.PNformat("full")
	if b != "02 February 2020, 12:00" {
		t.Errorf("[FormatFull]. Get: %v", b)
	}
}

func TestFormatISO8601(t *testing.T) {
	a := Must(stringFormat)

	b := a.PNformat()
	if b != stringFormat {
		t.Errorf("[FormatISO8601]. Get: %v", b)
	}
}

func TestRangeFormatYear(t *testing.T) {
	a := Must(stringFormat)

	b := Must()

	r := a.RangeFormat(b)

	if r != fmt.Sprintf("%s - %s", a.Format("DD MMMM YYYY"), b.Format("DD MMMM YYYY")) {
		t.Errorf("[RangeFormatYear]. Get: %s", r)
	}
}

func TestRangeFormatMonth(t *testing.T) {
	a := Must()

	a.Subtract(2, "months")
	b := Must()

	r := a.RangeFormat(b)

	if r != fmt.Sprintf("%s - %s %d", a.Format("DD MMMM"), b.Format("DD MMMM"), a.Year()) {
		t.Errorf("[RangeFormatMonth]. Get: %s", r)
	}
}

func TestRangeFormatDay(t *testing.T) {
	a := Must().Subtract(2, "days")

	b := Must()

	r := a.RangeFormat(b)

	if r != fmt.Sprintf("%s - %s %s", a.Format("DD"), b.Format("DD"), a.Format("MMMM YYYY")) {
		t.Errorf("[RangeFormatDay]. Get: %s", r)
	}
}

func TestRangeFormatTime(t *testing.T) {
	a := Must()
	b := a.Clone().Subtract(2, "hours")
	r := b.RangeFormat(a)

	if r != fmt.Sprintf("%s - %s, %s", b.Format("HH:mm"), a.Format("HH:mm"), b.Format("DD MMMM YYYY")) {
		t.Errorf("[RangeFormatTime]. Get: %s", r)
	}
}

func TestAdd(t *testing.T) {
	a := Must("2022-05-05 12:00:00")

	a.Add(6, "hours")
	f := a.Hour()

	if f != 18 {
		t.Errorf("[TestAdd] Not added. %d", f)
	}
}

func TestSubtract(t *testing.T) {
	a := Must("2022-05-05 12:00:00")

	a.Subtract(6, "hours")
	f := a.Hour()

	if f != 6 {
		t.Errorf("[TestSubtract] Not Subtracted. %d", f)
	}
}

func TestClone(t *testing.T) {
	a := Must()

	b := a.Clone().Subtract(2, "hours")

	if a.Format() == b.Format() {
		t.Error()
	}
}

func TestEndOf(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	b := a.EndOf("month")

	if b.Date() != 30 {
		t.Error()
	}
}

func TestStartOf(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	b := a.StartOf("month")

	if b.Date() != 1 {
		t.Error()
	}
}

func TestSetByUnits(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.Set("y", 2016).Year() != 2016 {
		t.Error()
	}
	if a.Set("year", 2017).Year() != 2017 {
		t.Error()
	}
	if a.Set("years", 2018).Year() != 2018 {
		t.Error()
	}
	if a.Set("M", 9).Month() != 9 {
		t.Error()
	}
	if a.Set("month", 10).Month() != 10 {
		t.Error()
	}
	if a.Set("months", 11).Month() != 11 {
		t.Error()
	}
	if a.Set("D", 8).Date() != 8 {
		t.Error()
	}
	if a.Set("D", 9).Date() != 9 {
		t.Error()
	}
	if a.Set("D", 10).Date() != 10 {
		t.Error()
	}
	if a.Set("h", 14).Hour() != 14 {
		t.Error()
	}
	if a.Set("hour", 15).Hour() != 15 {
		t.Error()
	}
	if a.Set("hours", 16).Hour() != 16 {
		t.Error()
	}
	if a.Set("m", 17).Minute() != 17 {
		t.Error()
	}
	if a.Set("minute", 18).Minute() != 18 {
		t.Error()
	}
	if a.Set("minutes", 19).Minute() != 19 {
		t.Error()
	}
	if a.Set("s", 20).Second() != 20 {
		t.Error()
	}
	if a.Set("second", 21).Second() != 21 {
		t.Error()
	}
	if a.Set("seconds", 22).Second() != 22 {
		t.Error()
	}
	if a.Set("ms", 23000).Millisecond() != 23000 {
		t.Error()
	}
	if a.Set("millisecond", 24000).Millisecond() != 24000 {
		t.Error()
	}
	if a.Set("milliseconds", 25000).Millisecond() != 25000 {
		t.Error()
	}
	if a.Set("ns", 100000).Nanosecond() != 100000 {
		t.Error()
	}
	if a.Set("nanosecond", 100001).Nanosecond() != 100001 {
		t.Error()
	}
	if a.Set("nanoseconds", 100002).Nanosecond() != 100002 {
		t.Error()
	}
}

func TestSetDay(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	a.SetDay(0)
	f := a.Day()

	if f != 0 {
		t.Error()
	}

	a.SetDay("tuesday")
	f = a.Day()

	if f != 2 {
		t.Error()
	}
}

func TestSetYear(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetYear(2016).Year() != 2016 {
		t.Error()
	}
}

func TestSetMonth(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetMonth(5).Month() != 5 {
		t.Error()
	}
}

func TestSetDate(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetDate(5).Date() != 5 {
		t.Error()
	}
}

func TestSetDayOfYear(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetDayOfYear(300).DayOfYear() != 300 {
		t.Error()
	}
}

func TestSetHour(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetHour(12).Hour() != 12 {
		t.Error()
	}
}

func TestSetMinute(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetMinute(12).Minute() != 12 {
		t.Error()
	}
}

func TestSetMillisecond(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetMillisecond(1000).Millisecond() != 1000 {
		t.Error()
	}
}

func TestSetNanosecond(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetNanosecond(1000000).Nanosecond() != 1000000 {
		t.Error()
	}
}

func TestSetSecond(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetSecond(12).Second() != 12 {
		t.Error()
	}
}

func TestSetWeek(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetWeek(12).Week() != 12 {
		t.Error()
	}
}

func TestSetWeekYear(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetWeekYear(2020).WeekYear() != 2020 {
		t.Error()
	}
}

func TestSetWeekDay(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetWeekday(2).Weekday() != 2 {
		t.Error()
	}
}

func TestSetQuarter(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetQuarter(1).Quarter() != 1 {
		t.Error()
	}
}

func TestSetUTCOffset(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetUTCOffset(7).UTCOffset() != 7*60 {
		t.Error()
	}
}

func TestSetISOWeekYear(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetISOWeekYear(2022).ISOWeekYear() != 2022 {
		t.Error()
	}
}

func TestSetISOWeekday(t *testing.T) {
	a := Must("2022-09-12 12:00:00")

	if a.SetISOWeekday(2).ISOWeekday() != 2 {
		t.Error()
	}
}
