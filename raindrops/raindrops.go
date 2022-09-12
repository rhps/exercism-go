package raindrops

import (
	"bytes"
	"fmt"
)

func Convert(number int) string {
	var result bytes.Buffer
	if number%3 == 0 {
		result.WriteString("Pling")
	}

	if number%5 == 0 {
		result.WriteString("Plang")
	}

	if number%7 == 0 {
		result.WriteString("Plong")
	}

	if result.String() == "" {
		return fmt.Sprintf("%d", number)
	} else {
		return result.String()
	}
}
