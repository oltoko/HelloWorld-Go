package main

import (
	"io"
	"os"
	"strings"
)

const pos_A = 65
const pos_Z = 90

const pos_a = 97
const pos_z = 122

type rot13Reader struct {
	r io.Reader
}

func (rot13R *rot13Reader) Read(p []byte) (n int, err error) {

	count, error := rot13R.r.Read(p)

	for i := 0; i < count; i++ {

		if isUpper(p[i]) {
			p[i] = rot13(p[i], pos_A)
		} else if isLower(p[i]) {
			p[i] = rot13(p[i], pos_a)
		}
	}

	return count, error
}

func isUpper(p byte) bool {

	if p >= pos_A && p <= pos_Z {
		return true
	}

	return false
}

func isLower(p byte) bool {

	if p >= pos_a && p <= pos_z {
		return true
	}

	return false
}

func rot13(b byte, base int) byte {

	bbase := byte(base)

	b -= bbase
	b = (b + 13) % 26

	return (b + bbase)
}

func main() {

	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
