package iteration

func Repeat(character string, count int) string {
	// declares a variable without initializing it
	// := is shorthand for declaring + initializing
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
