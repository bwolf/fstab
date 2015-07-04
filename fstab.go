// Package fstab implements a simple facility for reading unix (5) fstab files.
// A fstab can be read and traversed.
package fstab

import (
	"bytes"
	"fmt"
	"strings"
)

// A fstab is a array of type FstabEntry.
type Fstab []FstabEntry

// Returns a string representation of the fstab.
func (ft *Fstab) String() string {
	var buf bytes.Buffer
	buf.WriteString("fstab[]")
	for n, el := range *ft {
		if n > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%+v", el))
	}
	buf.WriteString("]")
	return buf.String()
}

// Returns the items of the fstab.
func (ft *Fstab) Items() []FstabEntry {
	return *ft
}

// Filter fstab items by the given predicate.
func (ft *Fstab) Filter(pred func(entry *FstabEntry) bool) Fstab {
	var res []FstabEntry
	for _, el := range *ft {
		if pred(&el) {
			res = append(res, el)
		}
	}
	return res
}

// Parse fstab from the given byte representation.
func ParseFstab(data []byte) (Fstab, error) {
	var res []FstabEntry

	for lineNum, l := range strings.Split(string(data), "\n") {
		var line = strings.TrimSpace(l)
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			if entry, err := ParseFstabEntry(line, lineNum); err != nil {
				return nil, err
			} else {
				res = append(res, *entry)
			}
		}
	}
	return res, nil
}

// EOF
