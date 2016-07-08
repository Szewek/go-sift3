package sift3

// SiftBytes measures distance between two arrays of bytes.
// Use this function if provided strings aren't Unicode (e-mails, URLs, etc...).
func SiftBytes(a, b []byte, o int) float64 {
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
			for i := 1; i < o; i++ {
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
func SiftRunes(a, b []rune, o int) float64 {
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
			for i := 1; i < o; i++ {
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
