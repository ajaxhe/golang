package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		/*
			{"aha", true},
				{"ajaxhe", false},
				{"zhuzhu", false},
				{"Et se resservir, ivresse reste.", true},
		*/
		{"上海自来水来自海上", true},
		{"上海自测海上", true},
		//{"冷水滩滩滩水冷", true},
	}

	for _, test := range tests {
		if got, want := IsPalindrome(test.input), test.want; got != want {
			t.Errorf("IsPalindrom(%s) = %v, want %v", test.input, got, want)
		}
	}
}

// go test -bench=.
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

/*
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
*/
