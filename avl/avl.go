package avl

// Node represents a node in an AVL tree.
type Node struct {
	Left, Right, Parent *Node
	Key                 uint
	Height              uint8
}

// Compare is used to compare two nodes.
type Compare func(a, b *Node) int

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return int(n.Height)
}

func avlInsert(node, parent, old *Node, root **Node) {
	if parent != nil {
		if parent.Left == old {
			parent.Left = node
		} else {
			parent.Right = node
		}
	} else if root != nil {
		*root = node
	}
	if node != nil {
		node.Parent = parent
	}
}

// AVLFind searches for key starting from base using cmp.
func AVLFind(key *Node, base *Node, cmp Compare) *Node {
	kkey := key.Key
	for base != nil {
		if kkey < base.Key {
			base = base.Left
		} else if kkey > base.Key {
			base = base.Right
		} else {
			c := cmp(key, base)
			if c < 0 {
				base = base.Left
			} else if c > 0 {
				base = base.Right
			} else {
				break
			}
		}
	}
	return base
}

// AVLRebalance rebalances the tree starting from P.
func AVLRebalance(P *Node, cmp Compare, root **Node) {
	for P != nil {
		L := P.Left
		HL := height(L)
		R := P.Right
		HR := height(R)

		if HL > HR+1 {
			LL := L.Left
			LR := L.Right
			HLR := height(LR)
			PP := P.Parent
			if height(LL) >= HLR {
				if P.Left = LR; LR != nil {
					LR.Parent = P
				}
				L.Right = P
				avlInsert(L, PP, P, root)
				P.Parent = L
				P.Height = uint8(1 + HLR)
				L.Height = 1 + P.Height
				P = PP
			} else {
				if P.Left = LR.Right; LR.Right != nil {
					LR.Right.Parent = P
				}
				LR.Right = P
				avlInsert(LR, PP, P, root)
				P.Parent = LR
				if L.Right = LR.Left; LR.Left != nil {
					LR.Left.Parent = L
				}
				LR.Left = L
				L.Parent = LR
				L.Height = uint8(HLR)
				P.Height = uint8(HLR)
				LR.Height = uint8(HL)
				P = PP
			}
		} else if HL+1 < HR {
			RR := R.Right
			RL := R.Left
			HRL := height(RL)
			PP := P.Parent
			if height(RR) >= HRL {
				if P.Right = RL; RL != nil {
					RL.Parent = P
				}
				R.Left = P
				avlInsert(R, PP, P, root)
				P.Parent = R
				P.Height = uint8(1 + HRL)
				R.Height = 1 + P.Height
				P = PP
			} else {
				if P.Right = RL.Left; RL.Left != nil {
					RL.Left.Parent = P
				}
				RL.Left = P
				avlInsert(RL, PP, P, root)
				P.Parent = RL
				if R.Left = RL.Right; RL.Right != nil {
					RL.Right.Parent = R
				}
				RL.Right = R
				R.Parent = RL
				R.Height = uint8(HRL)
				P.Height = uint8(HRL)
				RL.Height = uint8(HR)
				P = PP
			}
		} else {
			h := HL
			if HR > HL {
				h = HR
			}
			h++
			if int(P.Height) != h {
				P.Height = uint8(h)
				P = P.Parent
			} else {
				break
			}
		}
	}
}

// AVLInsert inserts node L into the tree rooted at P.
func AVLInsert(L *Node, P *Node, cmp Compare, root **Node) {
	Lkey := L.Key
	C := P
	var Ckey uint

	for C != nil {
		P = C
		Ckey = C.Key
		if Lkey < Ckey || (Lkey == Ckey && cmp(L, P) < 0) {
			C = P.Left
		} else {
			C = P.Right
		}
	}

	L.Left = nil
	L.Right = nil
	L.Parent = P
	L.Height = 1
	if P != nil {
		if Lkey < Ckey || (Lkey == Ckey && cmp(L, P) < 0) {
			P.Left = L
		} else {
			P.Right = L
		}
		AVLRebalance(P, cmp, root)
	} else if root != nil {
		*root = L
	}
}

// AVLRemove removes node P from the tree.
func AVLRemove(P *Node, cmp Compare, root **Node) {
	L := P.Left
	R := P.Right
	LC := L
	RC := R
	var Y *Node

	if L != nil && R != nil {
		for LC != nil && RC != nil {
			L = LC
			LC = L.Right
			R = RC
			RC = R.Left
		}
		if LC == nil {
			LC = L.Left
			if Y = L.Parent; Y == P {
				Y.Left = LC
			} else {
				Y.Right = LC
			}
			if LC != nil {
				LC.Parent = Y
			}
			avlInsert(L, P.Parent, P, root)
			if L.Right = P.Right; L.Right != nil {
				L.Right.Parent = L
			}
			if L.Left = P.Left; L.Left != nil {
				L.Left.Parent = L
			}
			HL := height(L.Left)
			HR := height(L.Right)
			if HL > HR {
				L.Height = uint8(HL + 1)
			} else {
				L.Height = uint8(HR + 1)
			}
			Y = L.Parent
		} else {
			RC = R.Right
			if Y = L.Parent; Y == P {
				Y.Right = RC
			} else {
				Y.Left = RC
			}
			if RC != nil {
				RC.Parent = Y
			}
			avlInsert(R, P.Parent, P, root)
			if R.Left = P.Left; R.Left != nil {
				R.Left.Parent = R
			}
			if R.Right = P.Right; R.Right != nil {
				R.Right.Parent = R
			}
			HR := height(R.Right)
			HL := height(R.Left)
			if HR > HL {
				R.Height = uint8(HR + 1)
			} else {
				R.Height = uint8(HL + 1)
			}
			Y = R.Parent
		}
	} else {
		if R != nil {
			avlInsert(R, P.Parent, P, root)
			Y = P.Parent
		} else {
			avlInsert(L, P.Parent, P, root)
			Y = P.Parent
		}
	}
	AVLRebalance(Y, cmp, root)
	P.Parent = nil
	P.Left = nil
	P.Right = nil
}
