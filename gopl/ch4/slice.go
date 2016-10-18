package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, idx int) {
	reverse(s[:idx])
	reverse(s[idx:])
	reverse(s)
}

func noempty(strs []string) []string {
	var i int
	for _, str := range strs {
		if str != "" {
			strs[i] = str
			i++
		}
	}
	return strs[:i]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
	rotate(s, 2)
	fmt.Println(s)
	strs := []string{"ajaxhe", "", "zhuzhu"}
	fmt.Printf("%q\n", strs)
	fmt.Printf("%q\n", noempty(strs))
	fmt.Printf("%q\n", strs)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = remove(slice, 2)
	fmt.Println(slice)
	fmt.Println(remove2(slice, 1))
}
