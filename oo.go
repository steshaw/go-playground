//
// See http://www.stanford.edu/class/ee380/Abstracts/100428-pike-stanford.pdf
//

package main

import "fmt"
import "os"

type Day int

var dayName = []string {"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

func (d Day) String() string {
   if 0 <= d && int(d) < len(dayName) { return dayName[d] }
   return "NoSuchDay"
}

type Fahrenheit float32

func (f Fahrenheit) String() string {
   return fmt.Sprintf("%.1f°F", float32(f))
}

type Stringer interface {
   String() string
}

func print(args ...Stringer) {
   for i, s := range args {
      if i > 0 { os.Stdout.WriteString(" ") }
      os.Stdout.WriteString(s.String())
   }
   os.Stdout.WriteString("\n")
}

func main() {
   print(Day(1), Day(3), Day(0), Fahrenheit(26.45), Fahrenheit(72.29), Day(-1), Day(7))
}
