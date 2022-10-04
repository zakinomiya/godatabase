package godb

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Builtin func()

func printPrompt(inMiddle bool) {
	if inMiddle {
		fmt.Print("    -> ")
		return
	}
	fmt.Print("mydb > ")
}

func exit() {
	fmt.Println("Bye")
	os.Exit(0)
}

func tables() {
	fmt.Println("show tables")
}

func getMetaCommand(statement string) Builtin {
	switch statement {
	case ".exit":
		return exit
	case ".tables":
		return tables
	}

	return func() {
		fmt.Printf("Unknown command %s\n", statement)
	}
}

func Run(user string, password string) error {
	inMiddle := false
	var stmtBuilder strings.Builder
	stmtBuilder.Grow(100)
	scanner := bufio.NewReader(os.Stdin)

	for {
		printPrompt(inMiddle)

		for {
			l, isPrefix, err := scanner.ReadLine()
			if err != nil {
				log.Fatal(err)
			}

			stmtBuilder.Write(l)

			if isPrefix {
				continue
			} else {
				break
			}
		}

		stmt := stmtBuilder.String()
		if stmt[len(stmt)-1] != ';' {
			stmtBuilder.WriteByte(' ')
			inMiddle = true
			continue
		} else {
			inMiddle = false
		}

		stmtBuilder.Reset()
		if strings.HasPrefix(stmt, ".") {
			getMetaCommand(stmt)()
			continue
		}

		if stmt == ".exit" {
			fmt.Println("Bye")
			return nil
		}

		tokenizer := NewTokenizer(stmt)
		for _, t := range tokenizer.Tokenize() {
			fmt.Println(t)
		}
	}
}
