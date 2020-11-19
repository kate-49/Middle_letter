package letter

func getMiddle(word string) string {
	var answer string

	stringLength := len(word)
	if stringLength%2 == 0 {
		for index, letter := range word {
			if index == (stringLength/2 - 1) {
				answer += string(letter)
			}
			if index == (stringLength / 2) {
				answer += string(letter)
			}
		}
	} else {
		for index, letter := range word {
			if index == stringLength/2 {
				answer = string(letter)
			}
		}
	}
	return answer
}
