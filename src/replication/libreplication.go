package main

// TODO: logging (either integrate or log to a different file)

import (
	"C"
	"fmt"
	"sync"
)

// Aside: "export" as a comment (no space) exports the symbol
// TODO: test with GODEBUG=cgocheck=1

//export PrintStringFromGo
func PrintStringFromGo(str *C.char, strlen C.int) {
	fmt.Println("Printing string")
	goString := C.GoStringN(str, strlen)
	fmt.Println(goString)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		fmt.Printf("In goroutine: %s\n", goString)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Done")
}

// Need to have a main() function, but it can be empty.
func main() {}
