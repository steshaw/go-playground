package main

func main() {
	fn := func() { go print(1) }
	defer fn()
	fn = func() { go print(2) }
	defer fn()
}
