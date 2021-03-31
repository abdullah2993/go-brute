package brute

import (
	"fmt"
	"strconv"
	"testing"
)

func TestABFixed(t *testing.T) {
	res := []string{"AA", "AB", "BA", "BB"}

	b, c := Brute([]rune("AB"), 2, 2, 4)
	defer c()
	i := 0
	for c := range b {
		t.Log(c)
		if res[i] != c {
			t.Fail()
		}
		i++
	}
}

func TestABNotFixed(t *testing.T) {
	res := []string{"A", "B", "AA", "AB", "BA", "BB"}

	b, c := Brute([]rune("AB"), 1, 2, 4)
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
	b, c := Brute([]rune("0123456789"), 6, 6, 1000)
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

func ExampleBrute() {
	characterSer := []rune("AB")
	minLen := 1
	maxLen := 2
	b, _ := Brute(characterSer, minLen, maxLen, 4)
	for combination := range b {
		fmt.Println(combination)
	}
	// Output:
	// A
	// B
	// AA
	// AB
	// BA
	// BB
}
