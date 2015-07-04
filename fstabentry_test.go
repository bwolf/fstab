package fstab

import "testing"

func TestFstabEntryIsVfstype(t *testing.T) {
	el := NewFstabEntry("/dev/sda1", "/usr", "ext3", "defaults", 0, 0)
	if !el.IsVfsType("ext3", "ext4") {
		t.Errorf("IsVfsType ext3, ext4 failed")
	}
}

func TestParseFstabEntry(t *testing.T) {
	entry, error := ParseFstabEntry("//server/share /mntpoint\t\t sometype foo=1,bar=2	0 1\r\n\t", 0)
	if error != nil {
		t.Errorf("Parsing failed with error: %v", error)
	}
	spec := "//server/share"
	if entry.FsSpec() != spec {
		t.Errorf("Parsing error type is invalid expected '%s' got '%s'", spec, entry.fs_spec)
	}
	file := "/mntpoint"
	if entry.FsFile() != file {
		t.Errorf("Parsing error file is invalid expected '%s' got '%s'", file, entry.fs_file)
	}
	vfstype := "sometype"
	if entry.FsVfsType() != vfstype {
		t.Errorf("Parsing error type is invalid expected '%s' got '%s'", vfstype, entry.fs_vfstype)
	}
	mntopts := "foo=1,bar=2"
	if entry.FsMntOpts() != mntopts {
		t.Errorf("Parsing error opts is invalid expected '%s' got '%s'", mntopts, entry.fs_mntopts)
	}
	freq := 0
	if entry.FsFreq() != freq {
		t.Errorf("Parsing error freq is invalid expected '%d' got '%d'", freq, entry.fs_freq)
	}
	passno := 1
	if entry.FsPassNo() != passno {
		t.Errorf("Parsing error passno is invalid expected '%d' got '%d'", passno, entry.fs_passno)
	}
}

func TestRemoveTrailing(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"/", "/"},
		{"/a", "/a"},
		{"/a/", "/a"},
		{"/a/b", "/a/b"},
		{"/a/b/", "/a/b"},
	}
	for _, c := range cases {
		got := remove_trailing(c.in, '/')
		if got != c.want {
			t.Errorf("remove_trailing(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
