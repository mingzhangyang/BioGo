package sequence

import (
	"fmt"
)

// Validate the input string
func Validate(s string) (string, error) {
	res := make([]byte, 0)
	for _, c := range s {
		switch {
		case c < 65:
			fmt.Printf("An illegal character \"%c\" (ASCII: %v) removed.\n", c, c)
		case c > 64 && c < 91:
			res = append(res, byte(c))
		case c > 96 && c < 123:
			res = append(res, byte(c))
		default:
			return "", fmt.Errorf("only [^A-Za-z] are allowed. But \"%c\" found", c)
		}
	}
	return string(res), nil
}
