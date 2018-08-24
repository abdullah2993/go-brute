package brute

import (
	"strconv"
	"testing"
)

func TestABFixed(t *testing.T) {
	res := []string{"AA", "AB", "BA", "BB"}

	b, c := Brute([]rune("AB"), 2, 4, true)
	defer c()
	i := 0
	for c := range b {
		if res[i] != c {
			t.Fail()
		}
		i++
	}
}

func TestABNotFixed(t *testing.T) {
	res := []string{"A", "B", "AA", "AB", "BA", "BB"}

	b, c := Brute([]rune("AB"), 2, 4, false)
	defer c()
	i := 0
	for c := range b {
		if res[i] != c {
			t.Fail()
		}
		i++
	}
}
func TestNumberFixed(t *testing.T) {
	b, c := Brute([]rune("0123456789"), 6, 1000, true)
	defer c()
	i := 0
	for c := range b {
		if n, _ := strconv.Atoi(c); i != n {
			t.Errorf("Expected: %v\t Got: %v", n, c)
			t.Fail()
		}
		i++
	}
}
