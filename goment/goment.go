package goment

import (
	"fmt"
	"regexp"

	moment "github.com/nleeper/goment"
)

type PortalnesiaGoment struct {
	moment.Goment
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
				return &PortalnesiaGoment{*out}, err
			}
			out, err = moment.New(out)
		case int32:
			out, err = moment.Unix(int64(v))
			if err != nil {
				return &PortalnesiaGoment{*out}, err
			}
			out, err = moment.New(out)
		case int64:
			out, err = moment.Unix(v)
			if err != nil {
				return &PortalnesiaGoment{*out}, err
			}
			out, err = moment.New(out)
		default:
			out, err = moment.New(v)
		}
	default:
		out, err = moment.New(args)
	}

	a := &PortalnesiaGoment{*out}
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
