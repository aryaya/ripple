package ledger

import (
	"fmt"
	"github.com/wangch/ripple/data"
	"github.com/wangch/ripple/storage"
	"sort"
)

type RadixAction byte

const (
	Addition RadixAction = 'A'
	Deletion RadixAction = 'D'
	Movement RadixAction = 'M'
)

// TODO Replace this with NodeId(Depth!)
type RadixOperation struct {
	*RadixNode
	Action RadixAction
}

func NewRadixOperation(action RadixAction, node data.Storer, depth uint8) *RadixOperation {
	return &RadixOperation{
		RadixNode: &RadixNode{
			Node:  node,
			Depth: depth,
		},
		Action: action,
	}
}

type RadixOperations []*RadixOperation

func (ro RadixOperations) Len() int      { return len(ro) }
func (ro RadixOperations) Swap(i, j int) { ro[i], ro[j] = ro[j], ro[i] }

type OpsByAction struct{ RadixOperations }

func (ops OpsByAction) Less(i, j int) bool {
	ro := ops.RadixOperations
	if ro[i].Action == ro[j].Action {
		return ro[i].Depth < ro[j].Depth
	}
	return ro[i].Action > ro[j].Action
}

type OpsByIndex struct{ RadixOperations }

func (ops OpsByIndex) Less(i, j int) bool {
	ro := ops.RadixOperations
	if ro[i].RadixNode.Node.GetHash() == ro[j].RadixNode.Node.GetHash() {
		return ro[i].Action > ro[j].Action
	}
	return ro[i].RadixNode.Node.GetHash().Compare(*ro[j].RadixNode.Node.GetHash()) < 0
}

func (ro *RadixOperations) Add(node data.Storer, action RadixAction, depth uint8) {
	*ro = append(*ro, NewRadixOperation(action, node, depth))
}

// Converts nodes that are both added and deleted into moved
func (ro RadixOperations) Fold() RadixOperations {
	var folded RadixOperations
	sort.Sort(OpsByAction{ro})
	for i := 0; i < len(ro)-1; i++ {
		if ro[i].Node.GetHash() == ro[i+1].Node.GetHash() {
			moved := *ro[i+1]
			moved.Action = Movement
			folded = append(folded, &moved)
			i++
		} else {
			folded = append(folded, ro[i])
		}
	}
	return folded
}

func Diff(left, right data.Hash256, db storage.DB) (RadixOperations, error) {
	var ops RadixOperations
	if err := diff(left, right, db, &ops, 0); err != nil {
		return nil, err
	}
	return ops, nil
}

func visitChildren(node data.Storer, db storage.DB, ops *RadixOperations, depth uint8, action RadixAction) error {
	inner, ok := node.(*data.InnerNode)
	if !ok {
		return nil
	}
	return inner.Each(func(pos int, h data.Hash256) error {
		child, err := db.Get(h)
		if err != nil {
			return nil
		}
		ops.Add(child, action, depth)
		return visitChildren(child, db, ops, depth+1, action)
	})
}

func diff(left, right data.Hash256, db storage.DB, ops *RadixOperations, depth uint8) error {
	var l, r data.Storer
	var err error
	switch {
	case left.IsZero() && right.IsZero():
		return nil
	case left.IsZero():
		r, err = db.Get(left)
		if err != nil {
			return err
		}
		ops.Add(r, Deletion, depth)
		return visitChildren(r, db, ops, depth+1, Deletion)
	case right.IsZero():
		l, err = db.Get(left)
		if err != nil {
			return err
		}
		ops.Add(l, Addition, depth)
		return visitChildren(l, db, ops, depth+1, Addition)
	}
	l, err = db.Get(left)
	if err != nil {
		return err
	}
	r, err = db.Get(right)
	if err != nil {
		return err
	}
	ops.Add(r, Deletion, depth)
	ops.Add(l, Addition, depth)
	leftInner, leftOk := l.(*data.InnerNode)
	rightInner, rightOk := r.(*data.InnerNode)
	switch {
	case !leftOk && !rightOk:
		return nil
	case !leftOk:
		return visitChildren(r, db, ops, depth+1, Deletion)
	case !rightOk:
		return visitChildren(l, db, ops, depth+1, Addition)
	default:
		for i := 0; i < 16; i++ {
			leftChild, rightChild := leftInner.Children[i], rightInner.Children[i]
			switch {
			case leftChild == rightChild:
				continue
			case leftChild.IsZero():
				child, err := db.Get(rightChild)
				if err != nil {
					return err
				}
				if err := visitChildren(child, db, ops, depth+1, Deletion); err != nil {
					return err
				}
			case rightChild.IsZero():
				child, err := db.Get(leftChild)
				if err != nil {
					return err
				}
				if err := visitChildren(child, db, ops, depth+1, Addition); err != nil {
					return err
				}
			default:
				if err := diff(leftChild, rightChild, db, ops, depth+1); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (ro RadixOperations) String() []string {
	s := make([]string, len(ro))
	for i := range ro {
		s[i] = ro[i].String()
	}
	return s
}

func (r RadixOperation) String() string {
	return fmt.Sprintf("%c,%s,%d,%s", r.Action, r.Node.GetType(), r.Depth, r.Node.NodeId())
}
