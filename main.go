package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите данные: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf("Вы ввели следующие данные: %s\n", txt)
	}
}
