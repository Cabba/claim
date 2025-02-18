//go:build !unclaim

// claim is a simple package to implement defensive programming and handle unrecoverable errors.
package claim

func init() {
	Pre = prePanic
	Post = postPanic
}

// Specify a precondition.
// If the condition fails the application will panic.
// Use preconditions to express conditions that must hold true,
// prefer returning errors for recoverable errors.
//
// Example:
//
//	func divide(a, b int) {
//		claim.Pre(b != 0)
//	}
func prePanic(c bool) {
	if !c {
		panic("precondition not satisfied")
	}
}

// Specify a postcondition.
// If the condition fails the application will panic.
// Use postconditions to express conditions that must hold true,
// prefer returning errors for recoverable errors.
//
// Example:
//
//	func (d *Door) Open() {
//		defer claim.Post(d.open)
//	}
func postPanic(c bool) {
	if !c {
		panic("postcondition not satisfied")
	}
}
