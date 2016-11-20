// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Test of echo command.  Run with: go test gopl.io/ch11/echo

//!+
package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
		{true, ",", []string{"a", "b", "c"}, "a b c\n"}, // NOTE: wrong expectation!
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v, %q, %q)",
			test.newline, test.sep, test.args)

		out = new(bytes.Buffer) // captured output
		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		// 注意，这里使用到了类型断言
		got := out.(*bytes.Buffer).String()
		// 因为out在echo.go申明的类型是io.Writer(该接口只声明了Write方法，并没有声明String()方法)，因此，这里必须要显示转成具体类型（concrete type）*bytes.Buffer后才能调String()函数, 如下的调用会报错
		// got := out.String()
		// 必须调用String()方法才行，如下这种强制转换会报错
		// got := string(out.(*bytes.Buffer))
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

//!-
