package lc

type LazyTag struct {
	toAdd int
}

func (l *LazyTag) Add(o *LazyTag) *LazyTag {
	l.toAdd += o.toAdd

	return l
}

func (l *LazyTag) HasTag() bool {
	return l.toAdd != 0
}

func (l *LazyTag) Clear() {
	l.toAdd = 0
}

type SegmentTreeNode struct {
	min int
	max int
	tag *LazyTag
}

func NewSegmentTreeNode() *SegmentTreeNode {
	return &SegmentTreeNode{
		min: 0,
		max: 0,
		tag: &LazyTag{},
	}
}

type SegmentTree struct {
	n    int
	tree []*SegmentTreeNode
}

func NewSegmentTree(data []int) *SegmentTree {
	n := len(data)
	tree := make([]*SegmentTreeNode, n*4+1)

	for i := range tree {
		tree[i] = NewSegmentTreeNode()
	}

	seg := &SegmentTree{
		n:    n,
		tree: tree,
	}

	seg.build(data, 1, n, 1)
	return seg
}

func (s *SegmentTree) Add(l, r, val int) {
	tag := &LazyTag{toAdd: val}
	s.update(l, r, tag, 1, s.n, 1)
}

func (s *SegmentTree) FindLast(start, val int) int {
	if start > s.n {
		return -1
	}

	return s.find(start, s.n, val, 1, s.n, 1)
}

func (s *SegmentTree) applyTag(i int, tag *LazyTag) {
	s.tree[i].min += tag.toAdd
	s.tree[i].max += tag.toAdd
	s.tree[i].tag.Add(tag)
}

func (s *SegmentTree) pushdown(i int) {
	if s.tree[i].tag.HasTag() {
		tag := &LazyTag{toAdd: s.tree[i].tag.toAdd}
		s.applyTag(i<<1, tag)
		s.applyTag((i<<1)|1, tag)
		s.tree[i].tag.Clear()
	}
}

func (s *SegmentTree) pushup(i int) {
	left := s.tree[i<<1]
	right := s.tree[(i<<1)|1]
	s.tree[i].min = min(left.min, right.min)
	s.tree[i].max = max(left.max, right.max)
}

func (s *SegmentTree) build(data []int, l, r, i int) {
	if l == r {
		s.tree[i].min = data[l-1]
		s.tree[i].max = data[l-1]
		return
	}

	mid := l + ((r - l) >> 1)
	s.build(data, l, mid, i<<1)
	s.build(data, mid+1, r, (i<<1)|1)
	s.pushup(i)
}

func (s *SegmentTree) update(targetL, targetR int, tag *LazyTag, l, r, i int) {
	if targetL <= l && r <= targetR {
		s.applyTag(i, tag)
		return
	}

	s.pushdown(i)
	mid := l + ((r - l) >> 1)
	if targetL <= mid {
		s.update(targetL, targetR, tag, l, mid, i<<1)
	}
	if targetR > mid {
		s.update(targetL, targetR, tag, mid+1, r, (i<<1)|1)
	}

	s.pushup(i)
}

func (s *SegmentTree) find(targetL, targetR, val, l, r, i int) int {
	if s.tree[i].min > val || s.tree[i].max < val {
		return -1
	}

	if l == r {
		return l
	}

	s.pushdown(i)
	mid := l + ((r - l) >> 1)

	if targetR >= mid+1 {
		res := s.find(targetL, targetR, val, mid+1, r, (i<<1)|1)
		if res != -1 {
			return res
		}
	}

	if l <= targetR && mid >= targetL {
		return s.find(targetL, targetR, val, l, mid, i<<1)
	}

	return -1
}

func longestBalanced(nums []int) int {
	occurrences := make(map[int][]int)

	sgn := func(x int) int {
		if x%2 == 0 {
			return 1
		}
		return -1
	}

	length := 0
	prefixSum := make([]int, len(nums))
	prefixSum[0] = sgn(nums[0])
	occurrences[nums[0]] = append(occurrences[nums[0]], 1)

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1]
		occ := occurrences[nums[i]]
		if len(occ) == 0 {
			prefixSum[i] += sgn(nums[i])
		}
		occurrences[nums[i]] = append(occ, i+1)
	}

	seg := NewSegmentTree(prefixSum)
	for i := 0; i < len(nums); i++ {
		length = max(length, seg.FindLast(i+length, 0)-i)
		nextPos := len(nums) + 1
		occurrences[nums[i]] = occurrences[nums[i]][1:]
		if len(occurrences[nums[i]]) > 0 {
			nextPos = occurrences[nums[i]][0]
		}

		seg.Add(i+1, nextPos-1, -sgn(nums[i]))
	}

	return length
}
