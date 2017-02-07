package main

import (
	"log"
	"runtime/debug"
)

func simplePanic() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()

	println("start.")
	panic("some thing is wrong")
	println("exit.")
}

func multiPanic() {
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("fatal")
			}
		}
	}()
	/*
		defer func() {
			if err := recover(); err != nil {
				log.Fatalln(err)
			} else {
				log.Fatalln("fatal")
			}
		}()
	*/

	defer func() {
		panic("same thing wrong?")
	}()

	panic("same thing is wrong")
}

func otherPanic() {
	//defer recover()
	defer func() {
		log.Println(recover())
	}()
	defer log.Println(recover())

	panic("some thing is wrong")
}

func debugPanic() {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	panic("test debug")
}

func main() {
	//simplePanic()
	//multiPanic()
	//otherPanic()
	debugPanic()
}
