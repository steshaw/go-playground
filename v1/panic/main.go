package main

import (
	"fmt"
	"runtime"
)

func baabaa(goexit bool) {
	if goexit {
		runtime.Goexit()
	} else {
		panic("baa baa!")
	}
}

func foo(goexit bool) {
	defer func() {
		fmt.Println("foo:deferred")
	}()
	fmt.Println("foo:start")

	defer func() {
		fmt.Println("foo:recover.1")
		v := recover()
		fmt.Println("foo:recover.1: recover() -> ", v)
	}()

	baabaa(goexit)

	defer func() {
		fmt.Println("foo:recover.2")
		v := recover()
		fmt.Println("foo:recover.2: recover() -> ", v)
	}()
}

func doIt(finished chan bool) {
	defer func() {
		fmt.Println("main:deferred.first")
	}()
	fmt.Println("main:start")

	defer func() {
		fmt.Println("main:deferred.before.foo.false")
	}()
	foo(false)
	defer func() {
		fmt.Println("main:deferred.after.foo.false")
	}()

	defer func() {
		fmt.Println("main:deferred.before.foo.true")
	}()
	foo(true)
	defer func() {
		fmt.Println("main:deferred.after.foo.true")
	}()

	fmt.Println("main:end")
	defer func() {
		fmt.Println("main:deferred.last")
		finished <- false // Any bool will do
	}()
}

func main() {
	finished := make(chan bool)
	go doIt(finished)
	b := <-finished
	fmt.Printf("b = %v\n", b)
}
