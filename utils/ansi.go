package utils

func MakeBold(content string) string {
	return "\033[1m" + content + "\033[0m"
}

func MakeItalic(content string) string {
	return "\033[3m" + content + "\033[0m"
}

func MakeUndeline(content string) string {
	return "\033[4m" + content + "\033[0m"
}

func MakeStrikethrough(content string) string {
	return "\033[9m" + content + "\033[0m"
}
