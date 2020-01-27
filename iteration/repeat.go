package iteration

func Repeat(character string, count int) (repeated string) {
	// there's no way to repeat zero or negative times
	if count <= 0 {
		return character
	}

	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
