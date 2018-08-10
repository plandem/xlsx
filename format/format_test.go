package format

func createAndFill(callback func(*StyleFormat)) *StyleFormat {
	f := New()
	callback(f)
	return f
}
