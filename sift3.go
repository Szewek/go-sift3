// Package sift3 contains Sift3 (a "Super Fast and Accurate string distance algorithm") implementation.
// For more details, see http://siderite.blogspot.com/2007/04/super-fast-and-accurate-string-distance.html
package sift3

type (
	// DistImpl is a string distance implementation.
	DistImpl interface {
		// Distance calculates distance between two strings.
		Distance(string, string) float64
	}
	// ByteSift measures distance between strings using bytes.
	ByteSift int
	// RuneSift measures distance between strings using runes.
	RuneSift int
)

// DefaultOffset used for Sift3.
const DefaultOffset int = 5

// Distance uses SiftBytes to get distance.
func (bs ByteSift) Distance(a, b string) float64 {
	return SiftBytes([]byte(a), []byte(b), int(bs))
}

// Distance uses SiftRunes to get distance.
func (rs RuneSift) Distance(a, b string) float64 {
	return SiftRunes([]rune(a), []rune(b), int(rs))
}

// SiftBytes measures distance between two arrays of bytes.
// Use this function if provided strings aren't Unicode (e-mails, URLs, etc...).
//
// maxOffset represents a range where values are swapped. By default it should be set to "sift3.DefaultOffset".
func SiftBytes(a, b []byte, maxOffset int) float64 {
	la := len(a)
	lb := len(b)
	if la == 0 {
		return float64(lb)
	}
	if lb == 0 {
		return float64(la)
	}
	var ca, cb, lcs int
	for ca < la && cb < lb {
		if a[ca] == b[cb] {
			lcs++
		} else {
			for i := 1; i < maxOffset; i++ {
				if ca+i < la && a[ca+i] == b[cb] {
					ca += i
					break
				}
				if cb+i < lb && a[ca] == b[cb+i] {
					cb += i
					break
				}
			}
		}
		ca++
		cb++
	}
	return float64(la+lb)/2.0 - float64(lcs)
}

// SiftRunes measures distance between two arrays of runes.
// Use this function if strings contain Unicode symbols.
//
// maxOffset represents a range where values are swapped. By default it should be set to "sift3.DefaultOffset".
func SiftRunes(a, b []rune, maxOffset int) float64 {
	la := len(a)
	lb := len(b)
	if la == 0 {
		return float64(lb)
	}
	if lb == 0 {
		return float64(la)
	}
	var ca, cb, lcs int
	for ca < la && cb < lb {
		if a[ca] == b[cb] {
			lcs++
		} else {
			for i := 1; i < maxOffset; i++ {
				if ca+i < la && a[ca+i] == b[cb] {
					ca += i
					break
				}
				if cb+i < lb && a[ca] == b[cb+i] {
					cb += i
					break
				}
			}
		}
		ca++
		cb++
	}
	return float64(la+lb)/2.0 - float64(lcs)
}
