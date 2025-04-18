package main

import (
    "os"
    "fmt"
    "flag"
    "strings"
    "strconv"
)

type SolutionFunc func([]byte) (int)

func parse_data(data []byte) ([][]int) {
    var dataLines []string = strings.Split(string(data), "\n")
    inputArray := make([][]int, len(dataLines))

    // Iterate through each line in the input.
    for i := 0; i < len(dataLines); i++ {
        var line []string = strings.Split(dataLines[i], "\t")
        inputArray[i] = make([]int, len(line))

        // Iterate though each element in the line.
        for j := 0; j < len(line); j++ {
            var err error
            inputArray[i][j], err = strconv.Atoi(line[j])

            if err != nil {
                panic(err)
            }
        }
    }

    return inputArray
}

// Iterate through number list and return
// division of two numbers that divide.
func find_divisor(arr []int) (int) {
    for i := 0; i < len(arr); i++ {
        var firstNo int = arr[i]

        for j := (i + 1); j < len(arr); j++ {
            var secondNo int = arr[j]

            if firstNo % secondNo == 0 {
                return firstNo / secondNo
            }

            if secondNo % firstNo == 0 {
                return secondNo / firstNo
            }
        }
    }

    return 0
}

func first_solution(data []byte) (int) {
    var totalChecksum int = 0
    var inputData [][]int = parse_data(data)

    for i := 0; i < len(inputData); i++ {
        var lowestValue int = inputData[i][0]
        var highestValue int = inputData[i][0]

        // Iterate lines and find lowest and highest.
        for j := 0; j < len(inputData[i]); j++ {
            if inputData[i][j] < lowestValue {
                lowestValue = inputData[i][j]
            }

            if inputData[i][j] > highestValue {
                highestValue = inputData[i][j]
            }
        }

        // Calculate checksum of line.
        totalChecksum += (highestValue - lowestValue)
    }

    return totalChecksum
}

func second_solution(data []byte) (int) {
    var inputArray [][]int = parse_data(data)

    var totalDivisor int = 0
    for i := 0; i < len(inputArray); i++ {
        totalDivisor += find_divisor(inputArray[i])
    }

    return totalDivisor
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
