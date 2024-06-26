package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func generateData(l int) ([]string, []string) {
	a := make([]string, l)
	b := make([]string, l)
	for i := range a {
		a[i] = randomNumbers(6)
		switch rand.Intn(4) {
		case 1:
			b[i] = randomNumbers(6)
		case 2:
			bs := []byte(a[i])
			for range rand.Intn(4) {
				i := rand.Intn(len(bs))
				bs[i] = randomDigit()
			}
			b[i] = string(bs)
		case 3:
			j := max(0, i-1)
			a[i] = a[j]
			b[i] = b[j]
		default:
			b[i] = a[i]
		}
	}
	return a, b
}

func randomDigit() byte {
	return byte(rand.Intn(9) + '0')
}

func randomNumbers(l int) string {
	var sb strings.Builder
	for range l {
		sb.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return sb.String()
}
