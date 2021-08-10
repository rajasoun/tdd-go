// Licensed under the Creative Commons License.

package greet

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	t.Run("Test Greet", func(t *testing.T) {
		want := "Weclome to the Team"
		got := Say(want)
		assert.Equal(t, got, want)
	})
	t.Run("Test Greet Dependency Injection", func(t *testing.T) {
		want := "Weclome to the Team"
		outputBuffer := bytes.Buffer{}
		SayDependnecyInjection(&outputBuffer, want)
		got := outputBuffer.String()
		assert.Equal(t, got, want)
	})
}
