package parse

func IsIdent(s string) bool {
	for i, c := range s {
		if c == '_' {
			continue
		}
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			continue
		}
		if c >= '0' && c <= '9' {
			if i == 0 {
				return false
			}
			continue
		}
		return false
	}

	return true
}
