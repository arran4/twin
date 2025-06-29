package libltdl

// LoadAddOnVtable emulates the BeOS loader using the plugin API.
var LoadAddOnVtable = &Dlvtable{
	Name:        "lt_load_add_on",
	ModuleOpen:  dlopenOpen,
	ModuleClose: dlopenClose,
	FindSym:     dlopenSym,
	Priority:    DlloaderAppend,
}
