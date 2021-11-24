package main

type Day int

const (
	sunday Day = iota + 1
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

func (d Day) show() string {
	return []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}[d-1] // panic?
}

func (d Day) dayOfWeek() int {
	return int(d)
}

func (d Day) PrintAll() {
	println(d)
	println(d.dayOfWeek())
	println(d.show())
}

func main() {
	println("Print all")
	for dayI := range days {
		var day = days[dayI]
		day.PrintAll()
	}

	println("Print Thursday")
	thursday.PrintAll()

	println("Print out of bounds")
	const outOfBounds Day = 12
	outOfBounds.PrintAll()
}
