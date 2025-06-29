package libltdl

import "os"

// Dirent represents a directory entry.
type Dirent struct {
	Name string
}

// Dir provides simple directory iteration.
type Dir struct {
	entries []os.DirEntry
	index   int
}

// Opendir reads directory entries from path.
func Opendir(path string) (*Dir, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return &Dir{entries: entries}, nil
}

// Readdir returns the next directory entry or nil when exhausted.
func (d *Dir) Readdir() *Dirent {
	if d == nil || d.index >= len(d.entries) {
		return nil
	}
	entry := d.entries[d.index]
	d.index++
	return &Dirent{Name: entry.Name()}
}

// Closedir closes the directory. It is a no-op in Go.
func (d *Dir) Closedir() {}
