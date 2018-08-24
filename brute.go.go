package brute

// Brute returns a channel of all possible combinations of a characterset.
// charset is the characterset to be used, length is the maximum length of
// combinations, buffer specifies the buffer size of the returned channel,
// fixed indicates whether or not the cominations whould of the fixed size i.e. length
// if false then it will retuen all combinations of size 1 upto length.
// closer can be called to terminte the generator
func Brute(charset []rune, length, buffer int, fixed bool) (combos <-chan string, closer func()) {
	results := make(chan string, buffer)
	done := make(chan struct{})
	charlen := len(charset)

	m := 1
	if fixed {
		m = length
	}

	go func() {
		defer close(results)
		for k := m; k <= length; k++ {
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

	return results, func() {
		close(done)
	}
}
