package iteration

func Repeat(ntimes int, character string) string {
	var repeated string
	for i := 0; i < ntimes; i++ {
		repeated += character
	}
	return repeated
}
