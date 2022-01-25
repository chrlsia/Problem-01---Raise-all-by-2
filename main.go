package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Give me some integers with space in between->")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	elements := strings.Split(strings.TrimSpace(line), " ")

	for _, value := range elements {
		result,_ := strconv.Atoi(value)
		result *=result
		fmt.Printf("%d ",result)
	}
	fmt.Println("\nFinished  ...")
	// fmt.Println(elements)
}
