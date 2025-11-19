package main


func splitWithQuotes(s string) []string {
	var result []string
	var current string
	inSingleQuote := false
	inDoubleQuote := false

	for i := 0; i < len(s); i++ {

		switch s[i] {
		case '"':
			inDoubleQuote = !inDoubleQuote
		case '\'':
			if inDoubleQuote {
				result = append(result, string(s[i]))
			}else{
				inSingleQuote = !inSingleQuote
			}
		case ' ':
			if inDoubleQuote {
				current += string(s[i])
			}else{
				if !inSingleQuote {
					if current != "" {
						result = append(result, current)
						current = ""
					}
				}
			}
		default:
			current += string(s[i])
		}
	}

	if current != "" {
		result = append(result, current)
	}
	
	return result
}