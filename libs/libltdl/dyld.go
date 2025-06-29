package libltdl

// DyldVtable provides a minimal macOS loader using the plugin API.
var DyldVtable = &Dlvtable{
	Name:        "lt_dyld",
	SymPrefix:   "_",
	ModuleOpen:  dlopenOpen,
	ModuleClose: dlopenClose,
	FindSym:     dlopenSym,
	Priority:    DlloaderAppend,
}
