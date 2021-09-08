package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadStdin reads from the system's `stdin` stream
// and retrieves the input received as a `string`
func ReadStdin(message string) (string, error) {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	fmt.Printf("%s ", message)

	if text, err := reader.ReadString('\n'); err == nil {
		text = strings.TrimSpace(text)

		return text, nil
	} else {
		return "", err
	}
}
