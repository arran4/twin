package libltdl

import "sort"

// SList is a simple singly linked list.
type SList struct {
	Next     *SList
	Userdata interface{}
}

type SListCallback func(item *SList, userdata interface{}) interface{}
type SListCompare func(item1, item2 *SList, userdata interface{}) int

// SlistDelete calls deleteFct on each element and returns nil.
func SlistDelete(head *SList, deleteFct func(*SList)) *SList {
	for head != nil {
		next := head.Next
		deleteFct(head)
		head = next
	}
	return nil
}

// SlistRemove removes the first matching element according to find.
// It returns the value returned by find when the match occurs.
func SlistRemove(phead **SList, find SListCallback, matchdata interface{}) interface{} {
	if phead == nil || *phead == nil {
		return nil
	}
	if res := find(*phead, matchdata); res != nil {
		*phead = (*phead).Next
		return res
	}
	for cur := *phead; cur.Next != nil; cur = cur.Next {
		if res := find(cur.Next, matchdata); res != nil {
			cur.Next = cur.Next.Next
			return res
		}
	}
	return nil
}

// SlistFind searches for a match using find and returns its value.
func SlistFind(slist *SList, find SListCallback, matchdata interface{}) interface{} {
	for slist != nil {
		if res := find(slist, matchdata); res != nil {
			return res
		}
		slist = slist.Next
	}
	return nil
}

// SlistConcat concatenates head and tail and returns the resulting list.
func SlistConcat(head, tail *SList) *SList {
	if head == nil {
		return tail
	}
	last := head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = tail
	return head
}

// SlistCons prepends item to slist.
func SlistCons(item, slist *SList) *SList {
	if item == nil {
		return slist
	}
	item.Next = slist
	return item
}

// SlistTail returns the list excluding its first element.
func SlistTail(slist *SList) *SList {
	if slist == nil {
		return nil
	}
	return slist.Next
}

// SlistNth returns the nth element of slist starting from 1.
func SlistNth(slist *SList, n int) *SList {
	for n > 1 && slist != nil {
		slist = slist.Next
		n--
	}
	return slist
}

// SlistLength returns the number of items in slist.
func SlistLength(slist *SList) int {
	n := 0
	for slist != nil {
		n++
		slist = slist.Next
	}
	return n
}

// SlistReverse destructively reverses slist.
func SlistReverse(slist *SList) *SList {
	var result *SList
	for slist != nil {
		next := slist.Next
		slist.Next = result
		result = slist
		slist = next
	}
	return result
}

// SlistForeach calls foreach for each element until it returns non-nil.
func SlistForeach(slist *SList, foreach SListCallback, userdata interface{}) interface{} {
	for slist != nil {
		next := slist.Next
		if res := foreach(slist, userdata); res != nil {
			return res
		}
		slist = next
	}
	return nil
}

// SlistSort sorts slist using compare.
func SlistSort(slist *SList, compare SListCompare, userdata interface{}) *SList {
	var nodes []*SList
	for node := slist; node != nil; node = node.Next {
		node.Next = nil
		nodes = append(nodes, node)
	}
	sort.Slice(nodes, func(i, j int) bool {
		return compare(nodes[i], nodes[j], userdata) <= 0
	})
	var head *SList
	for i := len(nodes) - 1; i >= 0; i-- {
		nodes[i].Next = head
		head = nodes[i]
	}
	return head
}

// SlistBox creates a new list element containing userdata.
func SlistBox(userdata interface{}) *SList {
	return &SList{Userdata: userdata}
}

// SlistUnbox returns the userdata stored in item.
func SlistUnbox(item *SList) interface{} {
	if item == nil {
		return nil
	}
	return item.Userdata
}
