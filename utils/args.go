package utils

func IsFlagInArgs(args []string, flag string) bool {
	for _, v := range args {
		if v == flag {
			return true
		}
	}
	return false
}

func FilterFlags(args []string) []string {
	result := []string{}
	for _, arg := range args {
		if arg[0] != '-' {
			result = append(result, arg)
		}
	}
	return result
}
