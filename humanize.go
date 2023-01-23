package humanize

import (
	"fmt"
	"math"
	"strings"
	"time"
)

const (
	day  = time.Minute * 60 * 24
	year = 365 * day
)

var byteSizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

func Duration(d time.Duration) string {
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
