package libltdl

import (
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

// DlHandle represents a loaded module.
type DlHandle struct {
	Name       string
	Module     *plugin.Plugin
	IsResident bool
}

var userSearchPath []string

// DlAddSearchDir appends dir to the module search path.
func DlAddSearchDir(dir string) int {
	if dir == "" {
		return 0
	}
	userSearchPath = append(userSearchPath, dir)
	return 0
}

// DlInsertSearchDir inserts dir before the path element "before".
func DlInsertSearchDir(before, dir string) int {
	if dir == "" {
		return 0
	}
	idx := -1
	for i, d := range userSearchPath {
		if d == before {
			idx = i
			break
		}
	}
	if idx < 0 {
		userSearchPath = append(userSearchPath, dir)
	} else {
		userSearchPath = append(userSearchPath[:idx], append([]string{dir}, userSearchPath[idx:]...)...)
	}
	return 0
}

// DlSetSearchPath sets the module search path to the colon-separated list path.
func DlSetSearchPath(path string) int {
	if path == "" {
		userSearchPath = nil
		return 0
	}
	userSearchPath = strings.Split(path, ":")
	return 0
}

// DlGetSearchPath returns the current module search path as a colon-separated string.
func DlGetSearchPath() string {
	return strings.Join(userSearchPath, ":")
}

// DlOpen loads the shared object filename using the Go plugin loader.
func DlOpen(filename string) *DlHandle {
	p, err := plugin.Open(filename)
	if err != nil {
		LtSetLastError(err.Error())
		return nil
	}
	return &DlHandle{Name: filename, Module: p}
}

// DlClose unloads a module. The Go plugin package does not support unloading.
func DlClose(handle *DlHandle) int {
	if handle == nil {
		return 1
	}
	return 0
}

// DlSym looks up symbolname in handle.
func DlSym(handle *DlHandle, symbolname string) interface{} {
	if handle == nil || handle.Module == nil {
		LtSetLastError("invalid handle")
		return nil
	}
	sym, err := handle.Module.Lookup(symbolname)
	if err != nil {
		LtSetLastError(err.Error())
		return nil
	}
	return sym
}

// DlMakeResident marks handle as resident so DlClose is a no-op.
func DlMakeResident(handle *DlHandle) int {
	if handle == nil {
		return 1
	}
	handle.IsResident = true
	return 0
}

// DlIsResident reports whether handle is resident. Returns -1 if handle is nil.
func DlIsResident(handle *DlHandle) int {
	if handle == nil {
		LtSetLastError("invalid handle")
		return -1
	}
	if handle.IsResident {
		return 1
	}
	return 0
}

// DlAdvise contains loader advice options.
type DlAdvise struct {
	TryExt         bool
	IsResident     bool
	IsSymGlobal    bool
	IsSymLocal     bool
	TryPreloadOnly bool
}

// DladviseInit creates a new DlAdvise object.
func DladviseInit() *DlAdvise { return &DlAdvise{} }

// DladviseDestroy frees the advise object.
func DladviseDestroy(a **DlAdvise) int {
	if a != nil {
		*a = nil
	}
	return 0
}

// DladviseExt enables extension searching.
func DladviseExt(a *DlAdvise) int {
	if a != nil {
		a.TryExt = true
	}
	return 0
}

// DladviseResident marks modules opened with this advice as resident.
func DladviseResident(a *DlAdvise) int {
	if a != nil {
		a.IsResident = true
	}
	return 0
}

// DladviseLocal marks modules opened with this advice as local.
func DladviseLocal(a *DlAdvise) int {
	if a != nil {
		a.IsSymLocal = true
	}
	return 0
}

// DladviseGlobal marks modules opened with this advice as global.
func DladviseGlobal(a *DlAdvise) int {
	if a != nil {
		a.IsSymGlobal = true
	}
	return 0
}

// DladvisePreload restricts loading to preloaded modules.
func DladvisePreload(a *DlAdvise) int {
	if a != nil {
		a.TryPreloadOnly = true
	}
	return 0
}

// DlOpenAdvise opens filename using the given advice.
func DlOpenAdvise(filename string, advise *DlAdvise) *DlHandle {
	return DlOpenSearch(filename, advise)
}

// DlForeachFile calls fn for each module found in searchPath.
// If searchPath is empty, the current user search path is used.
// The callback receives the path to a candidate module without extension.
// If fn returns non-zero, the iteration stops and that value is returned.
func DlForeachFile(searchPath string, fn func(string, interface{}) int, data interface{}) int {
	var dirs []string
	if searchPath != "" {
		dirs = strings.Split(searchPath, ":")
	} else {
		dirs = userSearchPath
	}
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		seen := make(map[string]struct{})
		for _, e := range entries {
			name := e.Name()
			base := strings.TrimSuffix(name, filepath.Ext(name))
			if _, ok := seen[base]; ok {
				continue
			}
			seen[base] = struct{}{}
			candidate := filepath.Join(dir, base)
			if fn != nil {
				if rc := fn(candidate, data); rc != 0 {
					return rc
				}
			}
		}
	}
	return 0
}
