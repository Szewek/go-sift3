package sift3

import (
	"fmt"
	"testing"
)

type siftTest struct {
	a, b string
	val  float64
}

var (
	bytesifts = map[string]siftTest{
		"domain": {"www.github.com", "ww.github.cmo", 3.5},
		"email":  {"some@e-mailaddress.com", "some@emailadress.com", 3},
		"us-uk":  {"colorful theater", "colourful theatre", 3.5},
	}
	runesifts = map[string]siftTest{
		"accents-pl": {"złocić", "zlocic", 2},
	}
)

func TestSiftBytes(t *testing.T) {
	var dist float64
	for k, v := range bytesifts {
		dist = SiftBytes([]byte(v.a), []byte(v.b), 5)
		t.Log(k, v.val, dist)
		if dist != v.val {
			t.FailNow()
		}
	}
}
func TestSiftRunes(t *testing.T) {
	var dist float64
	for k, v := range runesifts {
		dist = SiftRunes([]rune(v.a), []rune(v.b), 5)
		t.Log(k, v.val, dist)
		if dist != v.val {
			t.FailNow()
		}
	}
}
func ExampleSiftBytes() {
	a := []byte("Hello world!")
	b := []byte("Hello others!")
	fmt.Println(SiftBytes(a, b, 5))
	// Output: 6.5
}
