package libltdl

import (
	"path/filepath"
	"runtime"
)

// defaultModuleExtension returns the typical shared library extension for the
// current platform.
func defaultModuleExtension() string {
	switch runtime.GOOS {
	case "windows":
		return ".dll"
	case "darwin":
		return ".dylib"
	default:
		return ".so"
	}
}

// DlOpenSearch attempts to open filename from the search path using plugin.
// It tries the filename as given, then in each search directory. When advise is
// non-nil, TryExt allows automatic appending of the system extension when the
// file is not found. IsResident marks the handle as resident.
func DlOpenSearch(filename string, advise *DlAdvise) *DlHandle {
	if filename == "" {
		return nil
	}
	// Try as given first.
	if h := DlOpen(filename); h != nil {
		if advise != nil && advise.IsResident {
			h.IsResident = true
		}
		return h
	}
	// Try appending system extension if requested.
	if advise != nil && advise.TryExt {
		ext := defaultModuleExtension()
		if ext != "" && filepath.Ext(filename) != ext {
			if h := DlOpen(filename + ext); h != nil {
				if advise.IsResident {
					h.IsResident = true
				}
				return h
			}
		}
	}
	mu.RLock()
	dirs := append([]string(nil), userSearchPath...)
	mu.RUnlock()
	for _, dir := range dirs {
		if dir == "" {
			continue
		}
		full := filepath.Join(dir, filename)
		if h := DlOpen(full); h != nil {
			if advise != nil && advise.IsResident {
				h.IsResident = true
			}
			return h
		}
		if advise != nil && advise.TryExt {
			ext := defaultModuleExtension()
			if ext != "" && filepath.Ext(full) != ext {
				if h := DlOpen(full + ext); h != nil {
					if advise.IsResident {
						h.IsResident = true
					}
					return h
				}
			}
		}
	}
	return nil
}

// DlOpenExt is a convenience wrapper for DlOpenSearch with TryExt enabled.
func DlOpenExt(filename string) *DlHandle {
	adv := DlAdvise{TryExt: true}
	return DlOpenSearch(filename, &adv)
}
