package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/berrylradianh/cmlabs-backend-fulltime-test/function"
)

func main() {
	fmt.Println("Pilih fungsi:")
	fmt.Println("1. Fizzbuzz")
	fmt.Println("2. Palindrome")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan pilihan: ")
	pilihan, _ := reader.ReadString('\n')
	pilihan = strings.TrimSpace(pilihan)

	if pilihan == "1" {
		function.Fizzbuzz()
	} else if pilihan == "2" {
		function.Palindrome()
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}
