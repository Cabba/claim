//go:build unclaim

package claim

import "runtime/debug"

func init() {
	Pre = preSoft
	Post = postSoft
}

func preSoft(c bool) {
	if !c {
		UnclaimWriter.Write([]byte("precondition not satisfied\n"))
		UnclaimWriter.Write(debug.Stack())
	}
}

func postSoft(c bool) {
	if !c {
		UnclaimWriter.Write([]byte("postcondition not satisfied\n"))
		UnclaimWriter.Write(debug.Stack())
	}
}
