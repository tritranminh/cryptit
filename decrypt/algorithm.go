package decrypt

func Nimbus(str string) string {
	descryptedString := ""

	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		descryptedString += character
	}

	return descryptedString
}