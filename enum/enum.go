package main

type Day int

const (
	sunday Day = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

var days = []Day{
	sunday,
	monday,
	tuesday,
	wednesday,
	thursday,
	friday,
	saturday,
}

func main() {
	for t := range days {
		println(t)
	}

	const outOfBounds Day = 12
	println(outOfBounds)
}
