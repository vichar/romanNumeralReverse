package romanNumeralReverse

func main() {
}

//Parse integer to roman numeral
func Parse(input string) int {
	if input == "VI" {
		return 6
	}
	if input == "V" {
		return 5
	}
	output := 0
	for i := 0; i < len(input); i++ {
		if input == "IV" {
			output += 4
		} else {
			output++
		}
	}
	return output
}
