package utils

import "unicode/utf8"

func BytesToStrings(bytesList []byte) []string {
	var stringsList []string
	for i := 0; i < len(bytesList); {
		r, size := utf8.DecodeRune(bytesList[i:])
		if r == utf8.RuneError {
			// Handle invalid UTF-8 sequence, e.g., by skipping or replacing it
			i++
			continue
		}
		stringsList = append(stringsList, string(r))
		i += size
	}
	return stringsList
}
