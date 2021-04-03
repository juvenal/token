package token

import (
	"fmt"
	"log"
)

func RequestToken(message string) string {
	log.Println("Got a call, providing a response...")
	return fmt.Sprintf("Sample answer\n%s\n", message)
}
