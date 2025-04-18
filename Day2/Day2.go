package main

import (
    "os"
    "fmt"
    "flag"
    // "strconv"
)

type SolutionFunc func([]byte) (int)

func first_solution(data []byte) (int) {
    var dataString string = string(data)

    println("File input:", dataString)
    panic("Please implement first solution")
    return 0
}

func second_solution(data []byte) (int) {
    panic("Please implement first solution")
    return 0
}

func main() {
    // Set default help printout.
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [options] <filepath> \n", os.Args[0])
        flag.PrintDefaults()
    }

    var solution_index *int = flag.Int("sol-index", 1, "Select which solution to use. 1: first half, 2: second half. Default to 1.")

    flag.Parse()

    // Check non-flag arguments.
    if flag.NArg() != 1 {
        fmt.Fprintf(os.Stderr, "Invalid number of arguments\n")
        fmt.Fprintf(os.Stderr, "Try '%s --help' for help\n", os.Args[0])
        os.Exit(64)
    }

    // Check solution index argument.
    var sol_func SolutionFunc

    switch *solution_index {
        case 1:
            sol_func = first_solution
        case 2:
            sol_func = second_solution
        default:
            fmt.Fprintf(os.Stderr, "Invalid value for sol-index\n")
            fmt.Fprintf(os.Stderr, "Try '%s --help' for help\n", os.Args[0])
            os.Exit(64)
    }

    var filePath string = flag.Args()[0]

    println("Reading from file:", filePath)
    var data []byte
    var err error

    data, err = os.ReadFile(filePath)

    // Exit if error.
    if err != nil {
        panic(err)
    }

    // Remove EOF from data array.
    data = data[:len(data) - 1]

    println("Solution:", sol_func(data))
}
