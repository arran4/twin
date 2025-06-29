package libltdl

// LoadlibraryVtable provides a Windows loader using the plugin API as a simplification.
var LoadlibraryVtable = &Dlvtable{
	Name:        "lt_loadlibrary",
	ModuleOpen:  dlopenOpen,
	ModuleClose: dlopenClose,
	FindSym:     dlopenSym,
	Priority:    DlloaderAppend,
}
