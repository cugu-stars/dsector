// Package input provides various Input implementations for reading from []byte or os.File.
package input

import (
	"io"
)

// Input interface provides the minimum methods needed to parse a binary file.
// The interface is changed slightly, by ensuring the Read methods always try to do a full read.
// If they are unable to read the full amount, then a ShortReadError is returned, as well as as much
// as could be read.
type Input interface {
	io.Seeker

	io.Reader
	io.ReaderAt
	io.ByteReader

	Tell() (int64, error) // Here for convenience, perhaps remove.
}
