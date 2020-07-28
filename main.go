package main

// #cgo LDFLAGS: -L${SRCDIR}/libadd
// #include "add.h"
import "C"
import (
	"fmt"
	"io"
)

var clientSocket io.WriteCloser

func main() {}

//export AddInt
func AddInt(n1 C.int, n2 C.int) C.int {
	sum := C.add(n1, n2)
	fmt.Printf("%d\n", sum)
	return sum
}
