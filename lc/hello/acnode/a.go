package main

// Ac自动机
type AcNode struct {
	son   [26]*AcNode
	fail  *AcNode
	exist []int
}

func (root *AcNode) put(s string) {
	o := root
	for i := 0; i < len(s); i++ {
		pos := s[i] - 'a'
		if o.son[pos] == nil {
			o.son[pos] = &AcNode{}
		}
		o = o.son[pos]
	}
	o.exist = []int{len(s)}
}

func (root *AcNode) buildFail() {
	root.fail = root
	q := []*AcNode{}
	for i, son := range root.son[:] {
		if son == nil {
			root.son[i] = root
		} else {
			son.fail = root
			q = append(q, son)
		}
	}
	// BFS
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		f := o.fail
		for i, son := range o.son[:] {
			if son == nil {
				o.son[i] = f.son[i]
				continue
			}
			son.fail = f.son[i]
			son.exist = append(son.exist, son.fail.exist...)
			q = append(q, son)
		}
	}
}
