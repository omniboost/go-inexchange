package inexchange

import (
	"io"
	"net/url"
)

type Form interface {
	IsMultiPart() bool
	Values() url.Values
	Files() map[string]FormFile
}

type FormFile struct {
	Filename string
	Content  io.Reader
}
