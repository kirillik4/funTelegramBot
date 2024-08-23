package main

import (
	"Projects/dailyUpdate"
	"Projects/startBot"
	"fmt"
)

func main() {
	var token string
	var dailyAttempt bool
	fmt.Scan(&token, &dailyAttempt)
	if dailyAttempt {
		dailyUpdate.Update()
		startBot.StartBot(token)
	} else {
		startBot.StartBot(token)
	}
}
