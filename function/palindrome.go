package function

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Palindrome() {
	fmt.Println("Apakah Palindrome?")
	reader := bufio.NewReader(os.Stdin)
	var newWord string

	fmt.Print("Masukkan Kata : ")
	word, _ := reader.ReadString('\n')

	word = strings.TrimSpace(word)
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, ",", "")

	for i := len(word); i > 0; i-- {
		newWord += string(word[i-1])
	}

	if word == newWord {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
