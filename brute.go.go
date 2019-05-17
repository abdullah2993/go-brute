package brute

import "sync"

// Brute returns a channel of all possible combinations of a characterset.
// charset is the characterset to be used, minLen is the minimum length of
// combinations and should be greater than 0 it will panic otherwise, maxLen is the
// maxmimum length of combinations and should be greater than or equal to minLen, otherwise it will panic
// buffer specifies the buffer size of the returned channel,
// closer can be called to terminte the generator
func Brute(charset []rune, minLen, maxLen, buffer int) (combos <-chan string, closer func()) {

	results := make(chan string, buffer)
	done := make(chan struct{})
	charlen := len(charset)
	once := new(sync.Once)

	closer = func() {
		once.Do(func() {
			close(done)
		})
	}

	if minLen == 0 {
		minLen = 1
	}

	go func() {
		defer close(results)
		defer closer()
		for k := minLen; k <= maxLen; k++ {
			carry := 0
			indices := make([]int, k)
			for {
				select {
				case <-done:
					return
				default:
					out := ""
					for j := 0; j < k; j++ {
						out += string(charset[indices[j]])
					}
					results <- out
				}
				carry = 1
				for i := k - 1; i >= 0; i-- {
					if carry == 0 {
						break
					}

					indices[i] += carry
					carry = 0

					if indices[i] == charlen {
						carry = 1
						indices[i] = 0
					}
				}
				if carry == 1 {
					break
				}
			}
		}

	}()
	return results, closer
}
