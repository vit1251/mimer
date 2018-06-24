package mimer

import "io"

type Attachment struct {
	filename string
	content  io.Reader
	inline   bool
}
