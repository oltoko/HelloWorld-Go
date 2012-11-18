package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

    fn2 := 0
    fn1 := 1

    return func() int {

        fib := fn2

        fn2, fn1 = fn1, fn2 + fn1

        return fib
    }

}

func main() {
    f := fibonacci()
    for i := 0; i < 49; i++ {
        fmt.Println(i, ":", f())
    }
}
