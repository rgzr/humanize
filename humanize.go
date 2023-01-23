package humanize

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	day  = time.Minute * 60 * 24
	year = 365 * day
)

var byteSizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

func DurationSeconds(d time.Duration) string {
	d = d.Truncate(time.Second)
	if d < day {
		return d.String()
	}

	var b strings.Builder

	if d >= year {
		years := d / year
		fmt.Fprintf(&b, "%dy", years)
		d -= years * year
	}

	days := d / day
	d -= days * day
	fmt.Fprintf(&b, "%dd%s", days, d)

	return b.String()
}

func Bytes(s uint64) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), 1000))
	val := math.Floor(float64(s)/math.Pow(1000, e)*10+0.5) / 10
	f := "%.0f%s"
	if val < 10 {
		f = "%.1f%s"
	}

	return fmt.Sprintf(f, val, byteSizes[int(e)])
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func stripTrailingZeros(s string) string {
	if !strings.ContainsRune(s, '.') {
		return s
	}
	offset := len(s) - 1
	for offset > 0 {
		if s[offset] == '.' {
			offset--
			break
		}
		if s[offset] != '0' {
			break
		}
		offset--
	}
	return s[:offset+1]
}

func stripTrailingDigits(s string, digits int) string {
	if i := strings.Index(s, "."); i >= 0 {
		if digits <= 0 {
			return s[:i]
		}
		i++
		if i+digits >= len(s) {
			return s
		}
		return s[:i+digits]
	}
	return s
}

func Float(num float64, digits int) string {
	return stripTrailingZeros(stripTrailingDigits(strconv.FormatFloat(num, 'f', 6, 64), digits))
}
