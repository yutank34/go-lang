package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Host = "https://api.github.com"

func main() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		inputs := strings.Split(in.Text(), " ")
		switch inputs[0] {
		case "exit":
			os.Exit(0)
		case "issues":
			if len(inputs) == 1 {
				fmt.Println("レポジトリを選択してください。例) repo:hoge/fuga")
				continue
			}
			getIssues(inputs[1:])
		case "create":
			if len(inputs) < 2 {
				fmt.Println("titleを入力してください。例) create title")
			}
			// cmd := exec.Command("vim")
			// cmd.Stdin = os.Stdin
			// err := cmd.Run()
			createIssue(inputs[1:])
		case "hoge":
			fmt.Println("hoge")
		default:
			fmt.Println("def")
		}
	}
}
