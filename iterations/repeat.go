package iterations

// Appends the add character the desired number of times
func Repeat(char string, count int) string {
	var rval string
	for i := 0; i < count; i++ {
		rval += char
	}
	return rval
}
