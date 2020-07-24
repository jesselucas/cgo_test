package main

// #include "./libadd/add.h"
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
	fmt.Println("%d", sum)
	return sum
}
