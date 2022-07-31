package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	awb := generate_random_seed(0)

	if len(os.Args) == 3 {
		number_of_awbs, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("First argument needs to be an Number")
		}
		if number_of_awbs < 1 {
			fmt.Println("Number of AWBs must be one or more")
		}
		print_awbs(number_of_awbs, os.Args[2], awb)
	} else {
		awb_prefix := prefixInput()
		print_awbs(10, awb_prefix, awb)
	}
}

func print_awbs(number_of_awbs int, awb_prefix string, awb int) {
	for i := 1; i <= number_of_awbs; i += 1 {
		fmt.Printf("%02v: %v\n", i, generate_awb(awb_prefix, awb+i))
	}
}
func input() (string, error) {
	stdin := bufio.NewReader(os.Stdin)
	var awb_prefix string
	_, err := fmt.Fscanln(stdin, &awb_prefix)
	return awb_prefix, err
}

func prefixInput() string {
	for {
		fmt.Print("Please enter an AWB-Prefix: ")
		awb_prefix, err := input()

		if err != nil {
			fmt.Printf(": %+v\n", err)
			continue
		}
		if len(awb_prefix) != 3 {
			fmt.Println("The prefix needs to be three charactres long, not only ", len(awb_prefix))
			continue
		}
		return awb_prefix
	}
}

func generate_awb(prefix string, number int) string {
	return fmt.Sprintf("%v-%v%v", prefix, number, number%5)
}

func generate_random_seed(seed int64) int {
	var s1 rand.Source
	if seed == 0 {
		s1 = rand.NewSource(time.Now().UnixNano())
	} else {
		s1 = rand.NewSource(seed)
	}
	r1 := rand.New(s1)
	return r1.Intn(10000000)
}
