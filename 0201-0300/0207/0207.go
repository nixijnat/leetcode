package main

type node struct {
	ele   []int
	visit bool
}

// 2079
func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return false
	}
	if len(prerequisites) <= 1 {
		return true
	}
	mt := make(map[int]*node, 8)
	for _, v := range prerequisites {
		l := mt[v[0]]
		if l == nil {
			l = &node{make([]int, 0, 2), false}
			mt[v[0]] = l
		}
		l.ele = append(l.ele, v[1])
		l = mt[v[1]]
		if l == nil {
			l = &node{make([]int, 0, 2), false}
			mt[v[1]] = l
		}
	}
	for k, v := range mt {
		if len(v) > 0 {
			if !find(mt, )
		}
	}
	return find(mt, prerequisites[0][0])
}
func find(prerequisites map[int]*node, val int) bool {
	n := prerequisites[val]
	if n == nil {
		return true
	}
	if n.visit {
		return false
	}
	n.visit = true
	for _, v := range n.ele {
		if !find(prerequisites, v) {
			return false
		}
	}
	n.visit = false
	return true
}
