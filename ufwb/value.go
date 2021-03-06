package ufwb

import (
	"bramp.net/dsector/input"
	"encoding/binary"
	"fmt"
	"io"
)

// Value represents one of the parsed elements in the file.
// It doesn't contain the element, just the offset where it starts, and which element it is.
type Value struct {
	Offset  int64 // In bytes from the beginning of the file
	Len     int64 // In bytes
	Element Element
	Extra   interface{} // Extra info defined by the Element

	Children []*Value

	ByteOrder binary.ByteOrder // Only used for Number, TODO, and TODO. Why have this?
}

func (v *Value) Name() string {
	return v.Element.Name()
}

func (v *Value) Description() string {
	return v.Element.Description()
}

// String returns this value's string representation (based on display, etc)
func (v *Value) Format(file io.ReaderAt) (string, error) {
	return v.Element.Format(file, v)
}

func (v *Value) String() string {
	if v == nil {
		return "<nil>"
	}

	elem := "<unknown>"
	if v.Element != nil {
		elem = v.Element.IdString()
	}
	return fmt.Sprintf("[0x%x len:%d] %s", v.Offset, v.Len, elem)
}

func (v *Value) Write(f input.Input) {
	panic("TODO")
}

// find does a depth-first search of this value looking for a value matching the given element.
// This is mainly used for debugging/testing.
func (v *Value) find(e Element) (*Value, bool) {
	if v.Element == e {
		return v, true
	}

	for _, c := range v.Children {
		if vv, found := c.find(e); found {
			return vv, found
		}
	}

	return nil, false
}

// mustValidiate checks if this Value is valid, and panics if not
func (v *Value) mustValidiate() {
	if err := v.validiate(); err != nil {
		panic(err)
	}
}

// validiate checks if this Value is valid, only used for debugging.
func (v *Value) validiate() error {
	//log.Debugf("Checking %s", v)

	if v == nil {
		return fmt.Errorf("value = nil")
	}
	if v.Offset < 0 {
		return fmt.Errorf("%s value.Offset = %d want >= 0", v, v.Offset)
	}
	if v.Len < 0 {
		return fmt.Errorf("%s value.Len = %d want >= 0", v, v.Len)
	}
	if v.Element == nil {
		return fmt.Errorf("%s value.Element = nil want a valid value", v.String())
	}

	if len(v.Children) > 0 {
		switch v.Element.(type) {
		case *Structure, *StructRef, *Grammar: // ok
		default:
			return fmt.Errorf("%v only Structure and Grammar Values can have children, got %T", v, v.Element)
		}

		offset := v.Offset
		for _, child := range v.Children {
			if err := child.validiate(); err != nil {
				return err
			}

			if child.Offset != offset {
				return fmt.Errorf("%v child Value does not start at correct offset: %v>%v want:%v", v, v.Children, child, offset)
			}

			offset += child.Len
		}

		end := v.Offset + v.Len
		if offset != end {
			return fmt.Errorf("child Values does not end at the correct offset: %v>%v want:%v", v, v.Children, end)
		}
	}

	return nil
}
