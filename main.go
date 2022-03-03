package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Введите целое число: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf("Вы ввели число: %s\n", txt)
	}
}
