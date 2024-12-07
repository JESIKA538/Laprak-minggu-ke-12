package main

import (
	"fmt"
	"strings"
)

func main() {
	const correctUsername = "Admin"
	const correctPassword = "Admin"
	var failedAttempts int
	for {
		var username, password string
		fmt.Scan(&username, &password)
		if strings.TrimSpace(username) == correctUsername && strings.TrimSpace(password) == correctPassword {
			break 
		} else {
			failedAttempts++
		}
	}
	fmt.Printf("%d percobaan gagal login\n", failedAttempts)
}
