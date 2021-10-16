package main

import (
	"fmt"
)

type Ordered interface {
    type
       int, int8, int16, int32, int64,
       uint, uint8, uint16, uint32, uint64,
       uintptr,
       float32, float64,
       string
}

func Max[A Ordered](as ...A) (A, error) {
  if len(as) == 0 {
    var zero A
    return zero, fmt.Errorf("Cannot find max of zero elements")
  } else {
    max := as[0]
    for _, a := range as[1:] {
      if a > max {
        max = a
      }
    }
    return max, nil
  }
}

func main() {
  fmt.Println(Max(1,6,4,2))
  fmt.Println(Max[int]())
}
