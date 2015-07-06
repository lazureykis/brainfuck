package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const STACK_SIZE = 1024

func usage() {
	fmt.Printf("USAGE: %s script.bf\r\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	files := os.Args[1:len(os.Args)]
	for _, file := range files {
		program, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		execute(program)
	}
}

func execute(program []byte) {
	s := make([]byte, STACK_SIZE) // stack
	p := 0                        // stack pointer

	for i := 0; i < len(program); i++ {
		switch program[i] {
		case '>':
			p += 1
			if len(s) <= p {
				s = append(s, 0)
			}

		case '<':
			if p > 0 {
				p -= 1
			}

		case '+':
			s[p]++

		case '-':
			s[p]--

		case '.':
			fmt.Print(string(s[p]))

		case ',':
			b := make([]byte, 1)
			os.Stdin.Read(b)
			s[p] = b[0]

		case '[':
			if s[p] == 0 {
				i++
				loop_depth := 0
				for loop_depth > 0 || program[i] != ']' {
					if program[i] == '[' {
						loop_depth++
					} else if program[i] == ']' {
						loop_depth--
					}
					i++
				}
			}

		case ']':
			i--
			loop_depth := 0
			for loop_depth > 0 || program[i] != '[' {
				if program[i] == ']' {
					loop_depth++
				} else if program[i] == '[' {
					loop_depth--
				}
				i--
			}
			i--
		}
	}
}
