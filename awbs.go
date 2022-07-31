package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	awb := generate_random_seed(0)
	var amount_of_awbs int
	flag.IntVar(&amount_of_awbs, "a", 10, "Number of AWBs to generate")
	var awb_prefix string
	flag.StringVar(&awb_prefix, "p", "", "AWB prefix of the ABWs that will be generated, must be a string")
	flag.Parse()

	if len(awb_prefix) != 3 {
		fmt.Println("AWB must be a string with three characters, for example \"020\"")
		awb_prefix = prefixInput()
	}
	print_awbs(amount_of_awbs, awb_prefix, awb)

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
