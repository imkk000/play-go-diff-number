package main

import (
	"fmt"
	"sort"
)

func main() {
	// a, b := generateData(5)
	a := []string{"11", "11", "12"}
	b := []string{"11", "12", "13"}
	sort.Strings(a)
	sort.Strings(b)
	d := diff(a, b)
	sort.Strings(d)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("d:", d)
}

func diff(a, b []string) (d []string) {
	l := len(a[0])
	m := countMap(a, b)
	for i := range a {
		if m[a[i]] > 0 {
			d = append(d, a[i])
			m[a[i]]--
			a[i] = "XXXXXX"[:l]
		}
		if m[b[i]] < 0 {
			d = append(d, b[i])
			m[b[i]]++
			b[i] = "XXXXXX"[:l]
		}
	}
	return d
}

func countMap(a, b []string) map[string]int {
	m := make(map[string]int)
	for i := range a {
		m[a[i]]++
		m[b[i]]--
	}
	return m
}
