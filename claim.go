package claim

// Specify a precondition
func Pre(c bool, msg string) {
	if !c {
		panic(msg)
	}
}

func Post(c bool, msg string) {
	if !c {
		panic(msg)
	}
}
