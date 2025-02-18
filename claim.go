package claim

import (
	"io"
	"os"
)

type claim func(bool)

// Specify a precondition.
// If the condition fails the application will panic or log and error based
// on build tag provided.
// Use preconditions to express conditions that must hold true,
// prefer returning errors for recoverable errors.
//
// Compile with -tag unclaim to avoid panic the application.
//
// Example:
//
//	func divide(a, b int) {
//		claim.Pre(b != 0)
//	}
var Pre claim

// Specify a postcondition.
// If the condition fails the application will panic or log an error based
// on build tag provided.
// Use postconditions to express conditions that must hold true,
// prefer returning errors for recoverable errors.
//
// Compile with -tag unclaim to avoid panic the application.
//
// Example:
//
//	func (d *Door) Open() {
//		defer claim.Post(d.open)
//	}
var Post claim

// Writer used to print fails when unclaim build tag is provided.
var UnclaimWriter io.Writer = os.Stderr
