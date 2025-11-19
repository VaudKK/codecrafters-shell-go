package main


func splitWithQuotes(s string) []string {
	var result []string
	var current string
	inSingleQuote := false
	inDoubleQuote := false

	for i := 0; i < len(s); i++ {

		if s[i] == '"' {
			inDoubleQuote = !inDoubleQuote
			continue
		}

		if s[i] == '\'' {
			if inDoubleQuote {
				result = append(result, string(s[i]))
			}else{
				inSingleQuote = !inSingleQuote
			}
		} else if s[i] == ' ' && !inSingleQuote && !inDoubleQuote {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(s[i])
		}
	}

	if current != "" {
		result = append(result, current)
	}
	
	return result
}