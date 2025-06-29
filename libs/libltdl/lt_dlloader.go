package libltdl

import (
	"fmt"
	"os"
	"strings"
)

// DlloaderPriority corresponds to lt_dlloader_priority values.
const (
	DlloaderPrepend = iota
	DlloaderAppend
)

// Dlvtable describes a module loader.
type Dlvtable struct {
	Name         string
	SymPrefix    string
	ModuleOpen   func(data interface{}, filename string, advise interface{}) interface{}
	ModuleClose  func(data interface{}, module interface{}) int
	FindSym      func(data interface{}, module interface{}, symbolname string) interface{}
	DlloaderInit func(data interface{}) int
	DlloaderExit func(data interface{}) int
	DlloaderData interface{}
	Priority     int
}

var loaders []*Dlvtable

// DlloaderAdd registers a loader according to its priority.
func DlloaderAdd(vtable *Dlvtable) int {
	if vtable == nil || vtable.ModuleOpen == nil || vtable.ModuleClose == nil || vtable.FindSym == nil {
		return 1
	}
	if vtable.Priority != DlloaderPrepend && vtable.Priority != DlloaderAppend {
		return 1
	}
	if vtable.Priority == DlloaderPrepend {
		loaders = append([]*Dlvtable{vtable}, loaders...)
	} else {
		loaders = append(loaders, vtable)
	}
	return 0
}

// DlloaderNext returns the index of the loader following loader.
// Pass -1 to retrieve the first loader.
func DlloaderNext(loader int) int {
	if loader < 0 {
		if len(loaders) > 0 {
			return 0
		}
		return -1
	}
	if loader+1 < len(loaders) {
		return loader + 1
	}
	return -1
}

// DlloaderGet returns the vtable for loader index.
func DlloaderGet(loader int) *Dlvtable {
	if loader < 0 || loader >= len(loaders) {
		return nil
	}
	return loaders[loader]
}

// DlloaderRemove removes and returns the loader with matching name.
func DlloaderRemove(name string) *Dlvtable {
	for i, l := range loaders {
		if l != nil && l.Name == name {
			loaders = append(loaders[:i], loaders[i+1:]...)
			return l
		}
	}
	return nil
}

// DlloaderFind returns the first loader with matching name.
func DlloaderFind(name string) *Dlvtable {
	for _, l := range loaders {
		if l != nil && l.Name == name {
			return l
		}
	}
	return nil
}

// DlloaderDump prints the list of loader names to stderr.
func DlloaderDump() {
	names := make([]string, 0, len(loaders))
	for _, l := range loaders {
		if l == nil {
			continue
		}
		name := l.Name
		if name == "" {
			name = "(null)"
		}
		names = append(names, name)
	}
	if len(names) == 0 {
		fmt.Fprintln(os.Stderr, "loaders: (empty)")
	} else {
		fmt.Fprintln(os.Stderr, "loaders:", strings.Join(names, ", "))
	}
}
