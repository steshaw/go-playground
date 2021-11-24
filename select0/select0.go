package main

import (
	"fmt"
)

func single() {
	fmt.Println("Single channel")
	{
		// Determine what happens with receive from closed channel.
		var c = make(chan int)
		close(c)
		item, ok := <-c
		fmt.Printf("1: item=%v ok=%v\n", item, ok)
		fmt.Printf("2: item=%v ok=%v\n", item, ok)
		fmt.Printf("3: item=%v ok=%v\n", item, ok)
		select {
		case item, ok := <-c:
			fmt.Printf("select/case: item=%v ok=%v\n", item, ok)
		}
		if ok == false {
			c = nil
			fmt.Printf("nil: item=%v ok=%v\n", item, ok)
		}
	}
	{
		// Range over channel.
		var c = make(chan string)
		go func() {
			defer close(c)
			for i := 0; i < 3; i++ {
				c <- "boo"
			}
		}()
		for i := range c {
			fmt.Println("range:", i)
		}
	}
	{
		// Forever loop over channel with explicit break.
		var c = make(chan uint8)
		go func() {
			defer close(c)
			for i := 0; i < 3; i++ {
				c <- uint8(i)
			}
		}()
		for {
			msg, status := <-c
			if !status {
				break
			}
			fmt.Printf("loop: msg=%v status=%v\n", msg, status)
		}
	}
	fmt.Println("---")
}

func multi0() {
	fmt.Println("multi0")
	var c1 = make(chan uint8)
	go func() {
		defer close(c1)
		for i := uint8(0); i < 9; i++ {
			c1 <- i
		}
	}()

	var c2 = make(chan uint8)
	go func() {
		defer close(c2)
		for i := uint8(0); i < 3; i++ {
			c2 <- i
		}
	}()

	// This would loop forever if not limited to n iterations.
	for i := 0; i < 15; i++ {
		fmt.Println("iteration", i)
		select {
		case item, ok := <-c1:
			fmt.Printf("c1: item=%v ok=%v\n", item, ok)
		case item, ok := <-c2:
			fmt.Printf("c2: item=%v ok=%v\n", item, ok)
		}
	}
	fmt.Println("---")
}

func multi1() {
	fmt.Println("multi1")
	var c1 = make(chan uint8)
	close(c1)
	var c2 = make(chan uint8)
	close(c2)

	// Handle the panic, print and continue.
	// Hmmm, seems that you cannot recover from a deadlock.
	// Seems it is a _fatal error_ rather than a _panic_.
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err =", err)
		}
	}()

	// Provokes panic "all goroutines are asleep - deadlock!".
	i := 0
	for {
		i++
		fmt.Println("iteration", i)
		select {
		case item, ok := <-c1:
			fmt.Printf("c1: item=%v ok=%v\n", item, ok)
			if !ok {
				c1 = nil
			}
		case item, ok := <-c2:
			fmt.Printf("c2: item=%v ok=%v\n", item, ok)
			if !ok {
				c2 = nil
			}
		}
	}
	fmt.Println("---")
}

// Reading from a nil channel provokes an unrecoverable deadlock.
func multi2() {
	fmt.Println("multi2")
	var c1 chan uint8 = nil
	fmt.Println("c1", c1)
	i := <-c1 // Provokes deadlock
	fmt.Println(i)
	fmt.Println("---")
}

func multi3() {
	fmt.Println("multi3")
	var c1 = make(chan uint8)
	go func() {
		defer close(c1)
		for i := uint8(0); i < 9; i++ {
			c1 <- i
		}
	}()

	var c2 = make(chan uint8)
	go func() {
		defer close(c2)
		for i := uint8(0); i < 3; i++ {
			c2 <- i
		}
	}()

	// Reassign channel to nil so that it is not read from after a read
	// indicates that the channel is closed (ok == false). Bailing out of the
	// forever loop when both channels are nil is necesssary to prevent panic
	// ("fatal error: all goroutines are asleep - deadlock!").
	i := 0
	for {
		i++
		fmt.Println("iteration", i)
		select {
		case item, ok := <-c1:
			fmt.Printf("c1: item=%v ok=%v\n", item, ok)
			if ok {
			} else {
				c1 = nil
			}
		case item, ok := <-c2:
			fmt.Printf("c2: item=%v ok=%v\n", item, ok)
			if ok {
			} else {
				c2 = nil
			}
		}
		// Break when all channels drained.
		if c1 == nil && c2 == nil {
			break
		}
	}
	fmt.Println("---")
}

func multi3Quiet() {
	fmt.Println("multi3Quiet")
	var c1 = make(chan uint8)
	go func() {
		defer close(c1)
		for i := uint8(1); i <= 9; i++ {
			c1 <- i
		}
	}()

	var c2 = make(chan string)
	go func() {
		defer close(c2)
		for _, s := range []string{"one", "two", "three"} {
			c2 <- s
		}
	}()

	for {
		select {
		case item, ok := <-c1:
			if ok {
				fmt.Println("c1:", item)
			} else {
				c1 = nil
			}
		case item, ok := <-c2:
			if ok {
				fmt.Println("c2:", item)
			} else {
				c2 = nil
			}
		}
		// Break when all channels drained.
		if c1 == nil && c2 == nil {
			break
		}
	}
	fmt.Println("---")
}

func main() {
	single()
	multi0()
	if false {
		multi1() // deadlocks
	}
	if false {
		multi2() // deadlocks
	}
	multi3()
	multi3Quiet()
}
