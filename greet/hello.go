// Licensed under the Creative Commons License.

package greet

import (
	"fmt"
	"io"
)

func Say(message string) string {
	return message
}

func SayDependnecyInjection(writer io.Writer, message string) {
	fmt.Fprintf(writer, "%s", message)
}
