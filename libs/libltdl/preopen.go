package libltdl

// PreloadSym represents a preloaded symbol.
type PreloadSym struct {
	Name    string
	Address interface{}
}

var (
	defaultPreloaded []PreloadSym
	preloadedLists   [][]PreloadSym
)

// DlPreloadDefault records preloaded symbols to be used by default.
func DlPreloadDefault(preloaded []PreloadSym) int {
	mu.Lock()
	defaultPreloaded = preloaded
	mu.Unlock()
	return 0
}

// DlPreload adds preloaded symbols. Passing nil resets to the defaults.
func DlPreload(preloaded []PreloadSym) int {
	if preloaded == nil {
		mu.Lock()
		preloadedLists = nil
		if defaultPreloaded != nil {
			preloadedLists = append(preloadedLists, defaultPreloaded)
		}
		mu.Unlock()
		return 0
	}
	mu.Lock()
	preloadedLists = append(preloadedLists, preloaded)
	mu.Unlock()
	return 0
}

// DlPreloadOpen invokes cb for each module from originator.
func DlPreloadOpen(originator string, cb func(*DlHandle) int) int {
	mu.RLock()
	lists := append([][]PreloadSym(nil), preloadedLists...)
	mu.RUnlock()
	var errors int
	for _, list := range lists {
		if len(list) == 0 {
			continue
		}
		if originator != "" && list[0].Name != originator {
			continue
		}
		for i := 1; i < len(list); i++ {
			if list[i].Address == nil && list[i].Name != "" {
				h := DlOpen(list[i].Name)
				if h == nil {
					errors++
				} else if cb != nil {
					errors += cb(h)
				}
			}
		}
	}
	if errors > 0 {
		return errors
	}
	return 0
}
