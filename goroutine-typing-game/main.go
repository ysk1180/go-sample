package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const timeLimitSecond = 10

func main() {
	wordList := []string{"apple", "soccer", "goroutine", "interface", "awesome", "ruby", "keyboard", "linux"}
	fmt.Printf("タイピングゲーム、スタート！（制限時間%d秒）\n\n", timeLimitSecond)
	ch := input(os.Stdin)
	limit := time.After(timeLimitSecond * time.Second)
	point := 0

Game:
	for _, word := range wordList {
		fmt.Printf("> %s\n\n", word)

		select {
		case inputedWord := <-ch:
			if inputedWord == word {
				fmt.Printf("OK!\n\n")
				point++
			} else {
				fmt.Printf("NG!\n\n")
			}
		case <-limit:
			break Game
		}
	}
	fmt.Printf("\nタイムアップ!!%d問クリア!!\n", point)
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()

	return ch
}
