package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(`Guess four digit number of unique digits.
   A correct digit but not in the correct place shown with -.
   A correct digit in the correct place is shown with +`)

	num := make([]byte, 4)
	rand.Seed(time.Now().Unix())
	r := rand.Perm(9)
	for i := range num {
		num[i] = '1' + byte(r[i])
	}

	// accept and score guesses
	valid := []byte("123456789")
guess:
	for in := bufio.NewReader(os.Stdin); ; {
		fmt.Print("Guess: ")
		guess, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("\n not valid.")
			return
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 4 {
			fmt.Errorf("please guess a four digit number")
			continue
		}
		var wrongPos, rightPos int
		for ig, cg := range guess {
			if strings.IndexRune(guess[:ig], cg) >= 0 {
				// malformed:  repeated digit
				fmt.Printf("Repeated digit: %c\n", cg)
				continue guess
			}
			switch bytes.IndexByte(num, byte(cg)) {
			case -1:
				if bytes.IndexByte(valid, byte(cg)) == -1 {
					// malformed:  not a digit
					fmt.Printf("Invalid digit: %c\n", cg)
					continue guess
				}
			default: // I just think wrongPos should go first
				wrongPos--
			case ig:
				rightPos++
			}
		}
		fmt.Printf("rightPos: %d, wrongPos: %d\n", rightPos, wrongPos)
		if rightPos == 4 {
			fmt.Println("You got it.")
			return
		}
	}
}
