package fstab

import (
	"fmt"
	"strconv"
	"strings"
)

type FstabEntry struct {
	fs_spec    string
	fs_file    string
	fs_vfstype string
	fs_mntopts string
	fs_freq    int
	fs_passno  int
}

func remove_trailing(s string, ch byte) string {
	if len(s) > 1 && s[len(s)-1] == ch {
		return s[0 : len(s)-1]
	}
	return s
}

// Creates a new fstab entry with given parameters.
func NewFstabEntry(spec, file, vfstype, mntopts string, freq, passno int) *FstabEntry {
	return &FstabEntry{remove_trailing(spec, '/'), remove_trailing(file, '/'), vfstype, mntopts, freq, passno}
}

// Returns the device or server of the entry.
func (f *FstabEntry) FsSpec() string {
	return f.fs_spec
}

// Returns the mount point of the entry.
func (f *FstabEntry) FsFile() string {
	return f.fs_file
}

// Returns the virtual filesystem type of the entry.
func (f *FstabEntry) FsVfsType() string {
	return f.fs_vfstype
}

// Returns the mount options of the entry.
func (f *FstabEntry) FsMntOpts() string {
	return f.fs_mntopts
}

// Returns the flag which controls if the filesystem is a candidate of
// the dump(8) command.
func (f *FstabEntry) FsFreq() int {
	return f.fs_freq
}

// Returns the flag which controls the order how the fsck(8) program
// checks the filesystems.
func (f *FstabEntry) FsPassNo() int {
	return f.fs_passno
}

// Predicate to determine if the entry is of the given virtual
// filesystem type.
func (f *FstabEntry) IsVfsType(vfstype ...string) bool {
	for _, vfst := range vfstype {
		if f.fs_vfstype == vfst {
			return true
		}
	}
	return false
}

// Returns a string representation of entry.
func (f *FstabEntry) String() string {
	return fmt.Sprintf("%s %s %s %s %d %d", f.fs_spec, f.fs_file, f.fs_vfstype,
		f.fs_mntopts, f.fs_freq, f.fs_passno)
}

// Attempts to parse a FstabEntry from the given line. The lineNum
// pararameters is for better error messages.
func ParseFstabEntry(line string, lineNum int) (*FstabEntry, error) {
	var parts = strings.Fields(line)
	if len(parts) != 6 {
		return nil, fmt.Errorf("WARN: Too few fields in fstab line %d", lineNum)
	}

	var fs_spec = strings.TrimSpace(parts[0])
	var fs_file = strings.TrimSpace(parts[1])
	var fs_vfstype = strings.TrimSpace(parts[2])
	var fs_mntopts = strings.TrimSpace(parts[3])
	var fs_freq, fs_passno int

	var toInt32 = func(s, id string, lineNum int) (int, error) {
		if n, e := strconv.ParseInt(s, 10, 16); e != nil {
			return 666, fmt.Errorf("Failed parsing %s %s in line %d: %v",
				id, s, lineNum, e)
		} else {
			return int(n), nil
		}
	}
	if num, e := toInt32(parts[4], "fs_freq", lineNum); e != nil {
		return nil, e
	} else {
		fs_freq = num
	}
	if num, e := toInt32(parts[5], "fs_passno", lineNum); e != nil {
		return nil, e
	} else {
		fs_passno = num
	}
	return NewFstabEntry(fs_spec, fs_file, fs_vfstype, fs_mntopts, fs_freq, fs_passno), nil
}

// EOF
