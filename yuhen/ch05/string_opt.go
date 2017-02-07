package main

// string优化策略
// go build -gcflags "-N -l" // 阻止优化

func main() {
	m := map[string]int{
		"abc": 123,
	}

	key := []byte("abc")
	x, ok := m[string(key)]

	println(x, ok)
}
