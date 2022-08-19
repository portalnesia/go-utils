package utils

import (
	"fmt"
	"math"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/gosimple/slug"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Truncate string
//
// Ex. lorem ipsum lorem ipsum lorem ipsum => lorem ipsum lor...
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
//
// Ex. https://portalnesia.com/contact => portalnesia.com/contact
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

// Capitalize each words in string
//
// Ex. Hello world => Hello World
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

// Slugify format of string
//
// Ex. hello world => hello-world
func Slug(s string) string {
	return slug.Make(s)
}

// Format bytes to human readable string
//
// Ex. 50486525485 => 5.05 GB
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

// Generate random ID
func NanoId() string {
	return nanoid.Must()
}

// Format second integer to human readable `timeago` format
//
// Ex. 11 minutes ago
func TimeAgo(seconds int64) string {
	interval := int(math.Floor(float64(seconds) / 31536000))

	if interval > 1 {
		return fmt.Sprintf("%d years ago", interval)
	}

	interval = int(math.Floor(float64(seconds) / 2592000))
	if interval > 1 {
		return fmt.Sprintf("%d months ago", interval)
	}

	interval = int(math.Floor(float64(seconds) / 86400))
	if interval > 1 {
		return fmt.Sprintf("%d days ago", interval)
	}

	interval = int(math.Floor(float64(seconds) / 3600))
	if interval > 1 {
		return fmt.Sprintf("%d hours ago", interval)
	}

	interval = int(math.Floor(float64(seconds) / 60))
	if interval > 1 {
		return fmt.Sprintf("%d minutes ago", interval)
	}

	return "less minutes ago"
}

// Comma separate integer
//
// Ex. 5000 => 5,000
func SeparateNumber(number int64) string {
	p := message.NewPrinter(language.English)
	str := p.Sprintf("%d", number)
	return str
}

// ValidateURL
func IsUrl(stringUrl string) bool {
	_, err := url.ParseRequestURI(stringUrl)
	return err == nil
}

// Validate Twitter URL
func IsTwitterUrl(twitterUrl string) bool {
	isUrl := IsUrl(twitterUrl)

	if !isUrl {
		return false
	}

	regex := regexp.MustCompile(`^https?\:\/\/(www.)?twitter\.com`)

	return regex.MatchString(twitterUrl)
}

// Capitalize first characters in word
//
// Ex. hello world => Hello world
func FirstToUpper(text string) string {
	a := text[0:1]
	a = strings.ToUpper(a)
	return fmt.Sprintf("%s%s", a, text[1:])
}

type NumberFormatType struct {
	Number int64  `json:"number"`
	Format string `json:"format"`
}

// Format integer to K,M,B,T format
//
// Ex. 64768456 => 64.77 M
func NumberFormatShort(n int64) NumberFormatType {
	num := "0"

	if n < 900 { // 0 - 900
		num = fmt.Sprintf("%d", n)
	} else if n < 900000 { // 0.9 K - 850 K
		num = fmt.Sprintf("%.2f K", (float64(n) / 1000))
	} else if n < 900000000 { // 0.9 M - 850 M
		num = fmt.Sprintf("%.2f M", (float64(n) / 1000000))
	} else if n < 900000000000 { // 0.9 B - 850 B
		num = fmt.Sprintf("%.2f B", (float64(n) / 1000000000))
	} else { // 0.9 T +
		num = fmt.Sprintf("%.2f T", (float64(n) / 1000000000000))
	}

	return NumberFormatType{
		Number: n,
		Format: num,
	}
}

// Validate email
func ValidateEmail(e string) bool {
	regex := regexp.MustCompile(`^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$`)
	return regex.MatchString(e)
}

// Check if variable is true
func IsTrue(value interface{}) bool {
	r := reflect.TypeOf(value)
	t := r.String()

	if t == "string" {
		if value == "1" || value == "true" {
			return true
		}
	}

	if t == "int" || t == "int32" || t == "int64" {
		if value == 1 {
			return true
		}
	}

	if t == "bool" {
		if value == true {
			return true
		}
	}

	return false
}
