package utils

func checkError(err error) {
	if err != nil {
		panic(err.Error)
	}
}
