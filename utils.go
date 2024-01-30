package utils

import (
	"errors"
	"fmt"
	"github.com/oklog/ulid"
	"math"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Truncate string
//
// Example: "lorem ipsum lorem ipsum lorem ipsum" => "lorem ipsum lor..."
func Truncate(s string, max int) string {
	if max > len(s) {
		max = len(s)
		return s
	} else {
		max -= 3
		return s[:max] + "..."
	}
}

// Clean html format in string
func Clean(s string) string {
	p := bluemonday.NewPolicy()
	return p.Sanitize(s)
}

// Clean HTML format and truncate string
func CleanAndTruncate(s string, max int) string {
	str := Clean(s)
	return Truncate(str, max)
}

// Parse raw URL to clean URL
//
// Example: "https://portalnesia.com/contact" => "portalnesia.com/contact"
func ParseUrl(s string) (string, error) {
	if len(s) < 4 || len(s) >= 4 && s[0:4] != "http" {
		return "", errors.New("invalid url")
	}

	url, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	query := url.RawQuery

	if query != "" {
		query = fmt.Sprintf("?%s", query)
	}

	parser := fmt.Sprintf("%s%s%s", url.Host, url.Path, query)
	parser = strings.Replace(parser, "www.", "", 1)
	return parser, nil
}

// Capitalize each words in string
//
// Example: "Hello world" => "Hello World"
func Ucwords(s string) string {
	caser := cases.Title(language.Indonesian)
	return caser.String(strings.ToLower(s))
}

// Parse string to first letter uppercase
//
// Example: "Hello world" => "HM"
//
// If max is less than 1, function return all first letter of strings
func FirstLetter(s string, max int) string {
	regex, err := regexp.Compile(`\b[A-Z]`)
	if err != nil {
		return s
	}

	str := regex.FindAllString(strings.ToUpper(s), -1)

	output := strings.Join(str, "")

	if max > 0 {
		return output[0:max]
	}

	return output
}

// Slugify format of string
//
// Example: "hello world" => "hello-world"
func Slug(s string) string {
	return slug.Make(s)
}

// Format bytes to human readable string
//
// Example: 50486525485 => "5.05 GB"
func NumberSize(bytes float64, precision int) string {
	if precision <= 0 {
		precision = 2
	}

	if bytes <= 0 {
		return "-"
	}
	units := []string{"B", "KB", "MB", "GB", "TB"}

	bytes = math.Max(bytes, 0)
	pow := float64(int64(math.Floor(math.Log10(bytes)) / math.Log10(1024)))
	pow = math.Min(pow, float64(len(units)-1))

	bytes /= math.Pow(1024, pow)

	factorOfTen := math.Pow(10, float64(precision))
	parsed := math.Round(bytes*factorOfTen) / factorOfTen
	p := int(math.Round(pow))
	result := fmt.Sprintf("%.2f %s", parsed, units[p])
	return result
}

// Generate random ID
func NanoId(length ...int) string {
	return nanoid.Must(length...)
}

// Generate UUID
func UUID() string {
	return uuid.NewString()
}

func NanoIdStr(str string, length int) string {
	if str == "" {
		str = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if length == 0 {
		length = 20
	}
	return nanoid.MustGenerate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
}

func Ulid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

// Comma separate integer
//
// Example: 5000 => "5,000"
func SeparateNumber(number float64, tags ...language.Tag) string {
	tag := language.English
	if len(tags) == 1 {
		tag = tags[0]
	}
	p := message.NewPrinter(tag)
	str := p.Sprintf("%1.f", number)
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
// Example: "hello world" => "Hello world"
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
// Example: 64768456 => "64.77 M"
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
	switch t := value.(type) {
	case string:
		if t == "1" || strings.ToLower(t) == "true" {
			return true
		}
	case int:
		if t == 1 {
			return true
		}
	case int8:
		if t == 1 {
			return true
		}
	case int16:
		if t == 1 {
			return true
		}
	case int32:
		if t == 1 {
			return true
		}
	case int64:
		if t == 1 {
			return true
		}
	case uint:
		if t == 1 {
			return true
		}
	case uint8:
		if t == 1 {
			return true
		}
	case uint16:
		if t == 1 {
			return true
		}
	case uint32:
		if t == 1 {
			return true
		}
	case uint64:
		if t == 1 {
			return true
		}
	case float32:
		if t == float32(1) {
			return true
		}
	case float64:
		if t == float64(1) {
			return true
		}
	case bool:
		return t
	default:
		return false
	}
	return false
}

func Ternary[D any](cond bool, ifTrue D, ifFalse D) D {
	var response D
	if cond {
		response = ifTrue
	} else {
		response = ifFalse
	}
	return response
}
