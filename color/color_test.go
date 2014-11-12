package color

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateString(t *testing.T) {
	green, red, bold := ForegroundGreen, ForegroundRed, Bold
	text1, text2 := "hello", "world"
	fmt.Fprintf(os.Stdout,
		"%s, %s\n",
		GenerateString(text1, []int{green, bold}),
		GenerateString(text2, []int{red, bold}))

	invalidCode := 39
	fmt.Fprintf(os.Stdout,
		"%s\n",
		GenerateString("Hello", []int{invalidCode, bold}))
}

func TestGenerateEscapeString(t *testing.T) {
	colors := []int{ForegroundPurple, Bold}
	resetCodes := []int{Normal}
	fmt.Fprintf(os.Stdout, "%sWorld%s\n",
		generateEscapeString(colors), generateEscapeString(resetCodes))
}
