package InputOutput

import (
	"bufio"
	"fmt"
	"os"
)

func InputOutput() {

	var name string
	fmt.Println("Enter your name")

	fmt.Scanln(&name)

	fmt.Println("what's up bro ", name)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter message")
	message, _ := reader.ReadString('\n')

	fmt.Printf("You said %s", message)
}
