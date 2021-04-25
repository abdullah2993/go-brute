# go-brute [![Go Reference](https://pkg.go.dev/badge/github.com/abdullah2993/go-brute.svg)](https://pkg.go.dev/github.com/abdullah2993/go-brute)
A bruteforce combination generator for a specific character set

## Example
```
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
```
