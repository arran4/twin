package libltdl

import (
	"os"
	"path/filepath"
	"strings"
)

// ArgzCreateSep splits str on delim and returns the elements.
func ArgzCreateSep(str string, delim rune) []string {
	fields := strings.FieldsFunc(str, func(r rune) bool { return r == delim })
	return fields
}

// ArgzCreate returns a copy of argv as a new argz slice.
func ArgzCreate(argv []string) []string {
	if len(argv) == 0 {
		return nil
	}
	out := make([]string, len(argv))
	copy(out, argv)
	return out
}

// ArgzCount returns the number of elements in argz.
func ArgzCount(argz []string) int {
	return len(argz)
}

// ArgzAdd appends str to argz if it is non-empty.
func ArgzAdd(argz *[]string, str string) {
	if str == "" {
		return
	}
	*argz = append(*argz, str)
}

// ArgzAddSep splits str on delim and appends the fields to argz.
func ArgzAddSep(argz *[]string, str string, delim rune) {
	if str == "" {
		return
	}
	parts := strings.FieldsFunc(str, func(r rune) bool { return r == delim })
	for _, p := range parts {
		if p != "" {
			*argz = append(*argz, p)
		}
	}
}

// ArgzDelete removes the first occurrence of entry from argz.
func ArgzDelete(argz *[]string, entry string) {
	for i, v := range *argz {
		if v == entry {
			*argz = append((*argz)[:i], (*argz)[i+1:]...)
			return
		}
	}
}

// ArgzAppend appends buf as a single element to argz.
func ArgzAppend(argz *[]string, buf string) {
	if buf != "" {
		*argz = append(*argz, buf)
	}
}

// ArgzInsert inserts entry before the first instance equal to before.
// If before is empty or not found, entry is appended.
func ArgzInsert(argz *[]string, before, entry string) {
	if before == "" {
		*argz = append(*argz, entry)
		return
	}
	for i, v := range *argz {
		if v == before {
			*argz = append((*argz)[:i], append([]string{entry}, (*argz)[i:]...)...)
			return
		}
	}
	*argz = append(*argz, entry)
}

// ArgzNext returns the element following entry or the first element when entry is empty.
func ArgzNext(argz []string, entry string) string {
	if entry == "" {
		if len(argz) > 0 {
			return argz[0]
		}
		return ""
	}
	for i, v := range argz {
		if v == entry {
			if i+1 < len(argz) {
				return argz[i+1]
			}
			break
		}
	}
	return ""
}

// ArgzStringify joins argz using sep as separator.
func ArgzStringify(argz []string, sep rune) string {
	return strings.Join(argz, string(sep))
}

// ArgzInsertInOrder inserts entry into argz maintaining lexicographic order with no duplicates.
func ArgzInsertInOrder(argz *[]string, entry string) {
	if entry == "" {
		return
	}
	// check duplicates and find insertion index
	idx := len(*argz)
	for i, v := range *argz {
		if v == entry {
			return
		}
		if entry < v && idx == len(*argz) {
			idx = i
		}
	}
	*argz = append((*argz)[:idx], append([]string{entry}, (*argz)[idx:]...)...)
}

// ArgzInsertDir adds non-hidden files in dirname (without extensions) to argz in order.
// It returns 0 on success and 1 on error.
func ArgzInsertDir(argz *[]string, dirname string) int {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return 1
	}
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		base := strings.TrimSuffix(name, filepath.Ext(name))
		base = strings.TrimRightFunc(base, func(r rune) bool {
			return r == '.' || ('0' <= r && r <= '9')
		})
		if base == "" {
			continue
		}
		ArgzInsertInOrder(argz, filepath.Join(dirname, base))
	}
	return 0
}
