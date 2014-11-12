package color

import (
	"strconv"
	"strings"
)

const (
	// Normal escape code of console
	Normal = 0
	// Bold escape code of console
	Bold = 1
	// Underlined escape code of console
	Underlined = 4
	// Blinking escape code of console
	Blinking = 5
	// Reverse escape code of console
	Reverse = 8
	// ForegroundBlack escape code of console
	ForegroundBlack = 30
	// ForegroundRed escape code of console
	ForegroundRed = 31
	// ForegroundGreen escape code of console
	ForegroundGreen = 32
	// ForegroundYellow escape code of console
	ForegroundYellow = 33
	// ForegroundBlue escape code of console
	ForegroundBlue = 34
	// ForegroundPurple escape code of console
	ForegroundPurple = 35
	// ForegroundCyan escape code of console
	ForegroundCyan = 36
	// ForegroundWhite escape code of console
	ForegroundWhite = 37
	// BackgroundBlack escape code of console
	BackgroundBlack = 40
	// BackgroundRed escape code of console
	BackgroundRed = 41
	// BackgroundGreen escape code of console
	BackgroundGreen = 42
	// BackgroundYellow escape code of console
	BackgroundYellow = 43
	// BackgroundBlue escape code of console
	BackgroundBlue = 44
	// BackgroundPurple escape code of console
	BackgroundPurple = 45
	// BackgroundCyan escape code of console
	BackgroundCyan = 46
	// BackgroundWhite escape code of console
	BackgroundWhite = 47
	escapePrefix    = "\033["
	escapeSuffix    = "m"
)

var (
	escapeCode = map[int]string{
		0:  "normal",
		1:  "bold",
		4:  "underlined",
		5:  "blinking",
		8:  "reverse",
		30: "foregroundBlack",
		31: "foregroundRed",
		32: "foregroundGreen",
		33: "foregroundYellow",
		34: "foregroundBlue",
		35: "foregroundPurple",
		36: "foregroundCyan",
		37: "foregroundWhite",
		40: "backgroundBlack",
		41: "backgroundRed",
		42: "backgroundGreen",
		43: "backgroundYellow",
		44: "backgroundBlue",
		45: "backgroundPurple",
		46: "backgroundCyan",
		47: "backgroundWhite",
	}
)

func generateEscapeString(codes []int) string {
	var codeList []string
	for _, code := range codes {
		if "" != escapeCode[code] {
			codeList = append(codeList, strconv.Itoa(code))
		}
	}
	return escapePrefix + strings.Join(codeList, ";") + escapeSuffix
}

// GenerateString generagte string whih color
func GenerateString(text string, codes []int) string {
	resetCodes := []int{Normal}
	return generateEscapeString(codes) + text + generateEscapeString(resetCodes)
}
