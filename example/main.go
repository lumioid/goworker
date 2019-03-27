// manually add job:
// 		redis-cli -r 100 `
// run the example executables:
// 		./example -uri=redis://localhost:6379/db -queues="myqueue"
package main

import (
	"fmt"

	"github.com/benmanns/goworker"
)

func myFunc(queue string, args ...interface{}) error {
	fmt.Printf("From %s, %v\n", queue, args)
	return nil
}

func init() {
	goworker.Register("MyClass", myFunc)
}

func main() {
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}
