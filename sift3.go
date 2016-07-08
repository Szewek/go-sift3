// Package sift3 contains Sift3 (a "Super Fast and Accurate string distance algorithm") implementation
// For more details, see http://siderite.blogspot.com/2007/04/super-fast-and-accurate-string-distance.html
package sift3

// DefaultOffset
const DefaultOffset int = 5

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
