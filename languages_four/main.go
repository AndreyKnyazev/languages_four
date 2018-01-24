package main

import (
	"fmt"
	"time"
)

type Token struct{
	data string
	recipient int
	ttl int
}

func SendToken(tc chan Token, recipient int) {
	stateOfToken := <-tc
	fmt.Println(stateOfToken.ttl)
	if stateOfToken.ttl != 0 && stateOfToken.recipient != recipient {
		stateOfToken.ttl -= 1
		tc <- stateOfToken
		time.Sleep(time.Second * 1)
	} else if stateOfToken.recipient == recipient {
		fmt.Println(stateOfToken.data)
	} else {
		fmt.Println("Sort Error")
	}
}

func main() {
	var n int
	var token Token
	fmt.Scanf("%s\n", &token.data)
	fmt.Scanf("%d\n", &token.recipient)
	fmt.Scanf("%d\n", &token.ttl)
	fmt.Scanf("%d\n", &n)

	chanOfToken := make(chan Token)
	for i := 0; i < n; i++ {
		go SendToken(chanOfToken, i)
	}

	chanOfToken <- token

	var input string
	fmt.Scanln(&input)
}