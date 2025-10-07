package colors

import "fmt"

// Reset color
func Reset(s string) string {
	return fmt.Sprintf("\033[0m%s\033[0m", s)
}

// Basic colors
func Red(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s)
}

func Green(s string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", s)
}

func Yellow(s string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", s)
}

func Blue(s string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", s)
}

func Magenta(s string) string {
	return fmt.Sprintf("\033[35m%s\033[0m", s)
}

func Cyan(s string) string {
	return fmt.Sprintf("\033[36m%s\033[0m", s)
}

func White(s string) string {
	return fmt.Sprintf("\033[37m%s\033[0m", s)
}

func Black(s string) string {
	return fmt.Sprintf("\033[30m%s\033[0m", s)
}

func Orange(s string) string {
	return fmt.Sprintf("\033[38;5;208m%s\033[0m", s)
}

func Pink(s string) string {
	return fmt.Sprintf("\033[38;5;201m%s\033[0m", s)
}

// Bright colors (fixed codes)
func BrightRed(s string) string {
	return fmt.Sprintf("\033[91m%s\033[0m", s)
}

func BrightGreen(s string) string {
	return fmt.Sprintf("\033[92m%s\033[0m", s)
}

func BrightYellow(s string) string {
	return fmt.Sprintf("\033[93m%s\033[0m", s)
}

func BrightBlue(s string) string {
	return fmt.Sprintf("\033[94m%s\033[0m", s)
}

func BrightMagenta(s string) string {
	return fmt.Sprintf("\033[95m%s\033[0m", s)
}

func BrightCyan(s string) string {
	return fmt.Sprintf("\033[96m%s\033[0m", s)
}

func BrightWhite(s string) string {
	return fmt.Sprintf("\033[97m%s\033[0m", s)
}

// Gray colors (fixed codes)
func Gray(s string) string {
	return fmt.Sprintf("\033[90m%s\033[0m", s)
}

func LightGray(s string) string {
	return fmt.Sprintf("\033[37m%s\033[0m", s)
}

// Text styles
func Bold(s string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", s)
}

func Dim(s string) string {
	return fmt.Sprintf("\033[2m%s\033[0m", s)
}

func Italic(s string) string {
	return fmt.Sprintf("\033[3m%s\033[0m", s)
}

func Underline(s string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", s)
}

func Blink(s string) string {
	return fmt.Sprintf("\033[5m%s\033[0m", s)
}

func Reverse(s string) string {
	return fmt.Sprintf("\033[7m%s\033[0m", s)
}

// Background colors
func BgRed(s string) string {
	return fmt.Sprintf("\033[41m%s\033[0m", s)
}

func BgGreen(s string) string {
	return fmt.Sprintf("\033[42m%s\033[0m", s)
}

func BgYellow(s string) string {
	return fmt.Sprintf("\033[43m%s\033[0m", s)
}

func BgBlue(s string) string {
	return fmt.Sprintf("\033[44m%s\033[0m", s)
}

func BgMagenta(s string) string {
	return fmt.Sprintf("\033[45m%s\033[0m", s)
}

func BgCyan(s string) string {
	return fmt.Sprintf("\033[46m%s\033[0m", s)
}

func BgWhite(s string) string {
	return fmt.Sprintf("\033[47m%s\033[0m", s)
}

func BgBlack(s string) string {
	return fmt.Sprintf("\033[40m%s\033[0m", s)
}

// Combined styles for log types
func LogInfo(s string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", s) // Green
}

func LogWarning(s string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", s) // Yellow
}

func LogError(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s) // Red
}

func LogDebug(s string) string {
	return fmt.Sprintf("\033[36m%s\033[0m", s) // Cyan
}

func LogFatal(s string) string {
	return fmt.Sprintf("\033[1m\033[31m%s\033[0m", s) // Bold Red
}

// Utility function to combine colors and styles
func Style(s string, color string, style ...string) string {
	result := s

	// Apply color
	switch color {
	case "red":
		result = Red(result)
	case "green":
		result = Green(result)
	case "yellow":
		result = Yellow(result)
	case "blue":
		result = Blue(result)
	case "magenta":
		result = Magenta(result)
	case "cyan":
		result = Cyan(result)
	case "white":
		result = White(result)
	case "black":
		result = Black(result)
	case "gray":
		result = Gray(result)
	case "bright-red":
		result = BrightRed(result)
	case "bright-green":
		result = BrightGreen(result)
	case "bright-yellow":
		result = BrightYellow(result)
	case "bright-blue":
		result = BrightBlue(result)
	case "bright-magenta":
		result = BrightMagenta(result)
	case "bright-cyan":
		result = BrightCyan(result)
	case "bright-white":
		result = BrightWhite(result)
	}

	// Apply styles
	for _, s := range style {
		switch s {
		case "bold":
			result = Bold(result)
		case "dim":
			result = Dim(result)
		case "italic":
			result = Italic(result)
		case "underline":
			result = Underline(result)
		case "blink":
			result = Blink(result)
		case "reverse":
			result = Reverse(result)
		}
	}

	return result
}
