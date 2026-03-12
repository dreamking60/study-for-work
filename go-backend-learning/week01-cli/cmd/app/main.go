package main

func main() {
	err := run()

	if err != nil {
		handleError(err)
	}
}
