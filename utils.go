package utils

import (
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strings"

	"github.com/gosimple/slug"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Truncate string
func Truncate(s string, max int) string {
	return s[:max] + "..."
}

// Clean html format in string
func Clean(s string) string {
	p := bluemonday.NewPolicy()
	return p.Sanitize(s)
}

// Clean HTML format and truncate string
func CleanAndTruncate(s string, max int) string {
	str := Clean(s)
	return str[:max] + "..."
}

// Parse raw URL to clean URL
func ParseUrl(s string) (string, error) {
	url, err := url.Parse(s)
	query := url.RawQuery

	if query != "" {
		query = fmt.Sprintf("?%s", query)
	}

	parser := fmt.Sprintf("%s%s%s", url.Host, url.Path, query)
	parser = strings.Replace(parser, "www.", "", 1)
	return parser, err
}

func Ucwords(s string) string {
	caser := cases.Title(language.Indonesian)
	return caser.String(strings.ToLower(s))
}

// Parse string to first letter uppercase
//
// Ex. Hello world => HM
//
// If max is less than 1, function return all first letter of strings
func FirstLetter(s string, max int) string {
	regex, _ := regexp.Compile(`\b[A-Z]`)

	str := regex.FindAllString(strings.ToUpper(s), -1)

	output := strings.Join(str, "")

	if max > 0 {
		return output[0:max]
	}

	return output
}

func Slug(s string) string {
	return slug.Make(s)
}

func NumberSize(bytes float64, precision int) string {
	if precision <= 0 {
		precision = 2
	}

	if bytes <= 0 {
		return "-"
	}
	units := []string{"B", "KB", "MB", "GB", "TB"}

	bytes = math.Max(bytes, 0)
	pow := math.Floor(math.Log10(bytes)) / math.Log10(1024)
	pow = math.Min(pow, float64(len(units)-1))

	bytes /= math.Pow(1024, pow)

	factorOfTen := math.Pow(10, float64(precision))
	parsed := math.Round(bytes*factorOfTen) / factorOfTen
	p := int(math.Round(pow))
	result := fmt.Sprintf("%.2f %s", parsed, units[p])
	return result
}
