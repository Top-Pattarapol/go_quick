package main

func main() {
	for i := 0; i <= 100; i++ {
		println(Roman(i))
	}
}

// Roman naja
func Roman(number int) string {
	number, result := romanLoop(number, "")
	return result
}

func romanLoop(input int, result string) (int, string) {
	if input == 0 {
		return 0, result
	}

	key1, key2, word1, word2 := getKey(input)
	if value := input - key1; value >= -key2 {
		if value < 0 && value >= -key2 {
			result = result + word2
			input += key2
		}
		result = result + word1
		input = input - key1
	}
	return romanLoop(input, result)
}

func getKey(data int) (int, int, string, string) {
	if data >= 90 {
		return 100, 10, "C", "X"
	} else if data >= 40 {
		return 50, 10, "L", "X"
	} else if data >= 9 {
		return 10, 1, "X", "I"
	} else if data >= 4 {
		return 5, 1, "V", "I"
	} else {
		return 1, 0, "I", ""
	}
}
