package letter

func getMiddle(word string) string {
	var answer string

	stringLength := len(word)
	if stringLength%2 == 0 {

	} else {
		for index, letter := range word {
			if index == stringLength/2 {
				answer = string(letter)
			}
		}
	}
	return answer
}
