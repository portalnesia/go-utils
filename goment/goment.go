package goment

import (
	"fmt"
	"regexp"

	moment "github.com/nleeper/goment"
)

type PortalnesiaGoment struct {
	*moment.Goment
}

// New Custom Goment Instance
//
// Different with original goment package is:
//
// If args type is `int` or `int32` or `int64`, it will return goment.Unix() instead of goment.fromUnixNanoseconds()
func New(args ...interface{}) (*PortalnesiaGoment, error) {
	var err error
	var out *moment.Goment

	switch len(args) {
	case 0:
		out, err = moment.New()
	case 1:
		switch v := args[0].(type) {
		case int:
			out, err = moment.Unix(int64(v))
			if err != nil {
				return &PortalnesiaGoment{out}, err
			}
			out, err = moment.New(out)
		case int32:
			out, err = moment.Unix(int64(v))
			if err != nil {
				return &PortalnesiaGoment{out}, err
			}
			out, err = moment.New(out)
		case int64:
			out, err = moment.Unix(v)
			if err != nil {
				return &PortalnesiaGoment{out}, err
			}
			out, err = moment.New(out)
		default:
			out, err = moment.New(v)
		}
	default:
		out, err = moment.New(args)
	}

	a := &PortalnesiaGoment{out}
	a.UTC()

	return a, err
}

type TimeAgoResult struct {
	Format    string `json:"format"`
	Timestamp int64  `json:"timestamp"`
}

// Get Relative Time from now with custom Portalnesia output struct
//
// If args is true, years different will be formatted to:
//   - `1 year ago`
//
// Else:
//   - `DD MMM YYYY` format
func (g *PortalnesiaGoment) TimeAgo(args ...bool) TimeAgoResult {
	format := g.FromNow()
	if len(args) == 1 {
		if args[0] {
			return TimeAgoResult{Format: format, Timestamp: g.ToUnix()}
		}
	}
	if regexp.MustCompile(`years? ago$`).MatchString(format) {
		format = g.Format("DD MMM YYYY")
	}
	return TimeAgoResult{Format: format, Timestamp: g.ToUnix()}
}

// Utility format for Portalnesia
//
// t args:
//   - minimal: `DD MMM YYYY`
//   - fulldate: `DD MMMM YYYY`
//   - full: `DD MMMM YYYY, HH:mm`
//   - default: `YYYY-MM-DD HH:mm:ss` ISO8601
func (g *PortalnesiaGoment) PNformat(t ...interface{}) string {
	var format string

	switch len(t) {
	case 1:
		switch v := t[0].(type) {
		case string:
			switch v {
			case "minimal":
				format = "DD MMM YYYY"
			case "fulldate":
				format = "DD MMMM YYYY"
			case "full":
				format = "DD MMMM YYYY, HH:mm"
			}
		}
	default:
		format = "YYYY-MM-DD HH:mm:ss"
	}
	return g.Format(format)
}

// Get range format from g to d
//
// Example:
//   - `15:00 - 18:00, 02 January 2020`
//   - `02 - 05 January 2020`
//   - `02 January - 05 February 2020`
//   - `2 January 2020 - 5 January 2021`
func (g *PortalnesiaGoment) RangeFormat(d *PortalnesiaGoment) string {
	tanggal1 := g.Format("YYYY-MM-DD")
	tanggal2 := d.Format("YYYY-MM-DD")
	waktu1 := g.Format("HH:mm")
	waktu2 := d.Format("HH:mm")
	bln1 := g.Month()
	bln2 := d.Month()
	thn1 := g.Year()
	thn2 := d.Year()

	if tanggal1 == tanggal2 && waktu1 != waktu2 {
		return fmt.Sprintf("%s - %s, %s", waktu1, waktu2, g.Format("DD MMMM YYYY"))
	}
	if thn1 != thn2 {
		return fmt.Sprintf("%s - %s", g.Format("DD MMMM YYYY"), d.Format("DD MMMM YYYY"))
	}
	if bln1 != bln2 {
		return fmt.Sprintf("%s - %s %d", g.Format("DD MMMM"), d.Format("DD MMMM"), thn1)
	}
	return fmt.Sprintf("%s - %s %s", g.Format("DD"), d.Format("DD"), g.Format("MMMM YYYY"))
}

func (g *PortalnesiaGoment) Clone() *PortalnesiaGoment {
	return &PortalnesiaGoment{g.Goment.Clone()}
}

// Subtract mutates the original PortalnesiaGoment by subtracting time.
func (g *PortalnesiaGoment) Subtract(v int, unit string) *PortalnesiaGoment {
	g.Goment.Subtract(v, unit)
	return g
}

// Add mutates the original PortalnesiaGoment by adding time.
func (g *PortalnesiaGoment) Add(v int, unit string) *PortalnesiaGoment {
	g.Goment.Add(v, unit)
	return g
}

// EndOf mutates the original PortalnesiaGoment by setting it to the end of a unit of time.
func (g *PortalnesiaGoment) EndOf(units string) *PortalnesiaGoment {
	g.Goment.EndOf(units)
	return g
}

// Local will set the PortalnesiaGoment to use local time.
func (g *PortalnesiaGoment) Local() *PortalnesiaGoment {
	g.Goment.Local()
	return g
}

// Set is a generic setter, accepting units as the first argument, and value as the second.
func (g *PortalnesiaGoment) Set(units string, value int) *PortalnesiaGoment {
	g.Goment.Set(units, value)
	return g
}

// SetDate sets the day of the month. If the date passed in is greater than the number of days in the month, then the day is set to the last day of the month.
func (g *PortalnesiaGoment) SetDate(date int) *PortalnesiaGoment {
	g.Goment.SetDate(date)
	return g
}

// SetDay sets the day of the week (Sunday = 0...).
//
// A day name is also supported. This is parsed in the Goment object's locale.
//
//	Example: SetDay("sunday")
func (g *PortalnesiaGoment) SetDay(args interface{}) *PortalnesiaGoment {
	g.Goment.SetDay(args)
	return g
}

// SetDayOfYear sets the day of the year. For non-leap years, 366 is treated as 365.
func (g *PortalnesiaGoment) SetDayOfYear(doy int) *PortalnesiaGoment {
	g.Goment.SetDayOfYear(doy)
	return g
}

// SetHour sets the hour.
func (g *PortalnesiaGoment) SetHour(hours int) *PortalnesiaGoment {
	g.Goment.SetHour(hours)
	return g
}

// SetISOWeek sets the ISO week of the year.
func (g *PortalnesiaGoment) SetISOWeek(week int) *PortalnesiaGoment {
	g.Goment.SetISOWeek(week)
	return g
}

// SetISOWeekYear sets the ISO week-year.
func (g *PortalnesiaGoment) SetISOWeekYear(weekYear int) *PortalnesiaGoment {
	g.Goment.SetISOWeekYear(weekYear)
	return g
}

// SetISOWeekday sets the ISO day of the week with 1 being Monday and 7 being Sunday.
func (g *PortalnesiaGoment) SetISOWeekday(weekday int) *PortalnesiaGoment {
	g.Goment.SetISOWeekday(weekday)
	return g
}

// SetMillisecond sets the milliseconds.
func (g *PortalnesiaGoment) SetMillisecond(milliseconds int) *PortalnesiaGoment {
	g.Goment.SetMillisecond(milliseconds)
	return g
}

// SetMinute sets the minutes.
func (g *PortalnesiaGoment) SetMinute(minutes int) *PortalnesiaGoment {
	g.Goment.SetMinute(minutes)
	return g
}

// SetMonth sets the month (January = 1...). If new month has less days than current month, the date is pinned to the end of the target month.
func (g *PortalnesiaGoment) SetMonth(month int) *PortalnesiaGoment {
	g.Goment.SetMonth(month)
	return g
}

// SetNanosecond sets the nanoseconds.
func (g *PortalnesiaGoment) SetNanosecond(nanoseconds int) *PortalnesiaGoment {
	g.Goment.SetNanosecond(nanoseconds)
	return g
}

// SetQuarter sets the quarter (1 to 4).
func (g *PortalnesiaGoment) SetQuarter(quarter int) *PortalnesiaGoment {
	g.Goment.SetQuarter(quarter)
	return g
}

// SetSecond sets the seconds.
func (g *PortalnesiaGoment) SetSecond(seconds int) *PortalnesiaGoment {
	g.Goment.SetSecond(seconds)
	return g
}

// SetUTCOffset sets the UTC offset in minutes. If the offset is less than 16 and greater than -16, the value is treated as hours.
func (g *PortalnesiaGoment) SetUTCOffset(offset int) *PortalnesiaGoment {
	g.Goment.SetUTCOffset(offset)
	return g
}

// SetWeek sets the week of the year according to the locale.
func (g *PortalnesiaGoment) SetWeek(week int) *PortalnesiaGoment {
	g.Goment.SetWeek(week)
	return g
}

// SetWeekYear sets the week-year according to the locale.
func (g *PortalnesiaGoment) SetWeekYear(weekYear int) *PortalnesiaGoment {
	g.Goment.SetWeekYear(weekYear)
	return g
}

// SetWeekday sets the day of the week according to the locale.
func (g *PortalnesiaGoment) SetWeekday(weekday int) *PortalnesiaGoment {
	g.Goment.SetWeekday(weekday)
	return g
}

// SetYear sets the year.
func (g *PortalnesiaGoment) SetYear(year int) *PortalnesiaGoment {
	g.Goment.SetYear(year)
	return g
}

// StartOf mutates the original Goment by setting it to the start of a unit of time.
func (g *PortalnesiaGoment) StartOf(units string) *PortalnesiaGoment {
	g.Goment.StartOf(units)
	return g
}

// UTC will set the Goment to use UTC time.
func (g *PortalnesiaGoment) UTC() *PortalnesiaGoment {
	g.Goment.UTC()
	return g
}
