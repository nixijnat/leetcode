package main

import "fmt"

func calcEquation_1(equations [][]string, values []float64, queries [][]string) []float64 {
	s := newSet(equations, values)
	res := make([]float64, 0, 4)
	for _, v := range queries {
		p1, w1 := s.find(v[0])
		p2, w2 := s.find(v[1])
		if p1 == "" || p2 == "" || p1 != p2 {
			res = append(res, -1)
			continue
		}
		res = append(res, w2/w1) // NOTE 注意不要写反了
	}
	return res
}

type dSet struct {
	attr   map[string]*node // 存储秩和权值
	parent map[string]string
}
type node struct {
	rank   int     // 秩
	weight float64 // 到父节点的权值
}

func (d *dSet) print() {
	for i, v := range d.attr {
		fmt.Println(i, *v)
	}
	fmt.Println(d.parent)
}

func newSet(equations [][]string, values []float64) *dSet {
	s := &dSet{
		attr:   make(map[string]*node),
		parent: make(map[string]string),
	}
	// 初始化结点
	for _, v := range equations {
		if _, ok := s.attr[v[0]]; !ok {
			s.attr[v[0]] = &node{1, 1}
			s.parent[v[0]] = v[0]
		}
		if _, ok := s.attr[v[1]]; !ok {
			s.attr[v[1]] = &node{1, 1}
			s.parent[v[1]] = v[1]
		}
	}
	// 并查 构造最终树
	for i, v := range equations {
		s.union(v[0], v[1], values[i])
	}
	return s
}

// 返回根节点，以及 v 到根节点的权值
func (s *dSet) find(v string) (string, float64) {
	attr := s.attr[v]
	if attr == nil { // 不存在的的节点
		return "", 0
	}
	if v != s.parent[v] {
		// 压缩路径 以及 更新权限
		p, w := s.find(s.parent[v])
		s.parent[v] = p
		attr.weight *= w
	}
	return s.parent[v], attr.weight
}

// 带权合并
func (s *dSet) union(v1, v2 string, weight float64) {
	p1, w1 := s.find(v1)
	p2, w2 := s.find(v2)
	if p1 == p2 {
		return
	}
	attr1, attr2 := s.attr[p1], s.attr[p2]
	if attr1.rank > attr2.rank { // 比较秩
		p1, p2 = p2, p1
		w1, w2 = w2, w1
		weight = 1 / weight
		attr1, attr2 = attr2, attr1
	}
	s.parent[p1] = p2
	attr1.weight = w2 / w1 / weight // NOTE 更新旧的根节点的权值
	if attr1.rank == attr2.rank {
		attr2.rank++
	}
}

//

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i, v := range equations {
		ps, ok := graph[v[0]]
		if !ok {
			ps = make(map[string]float64)
			graph[v[0]] = ps
		}
		ps[v[1]] = values[i]
		ps, ok = graph[v[1]]
		if !ok {
			ps = make(map[string]float64)
			graph[v[1]] = ps
		}
		ps[v[0]] = 1.0 / values[i]
	}
	type point struct {
		to     string
		weight float64
	}
	// 广度优先搜索
	bfs := func(start, end string) float64 {
		if start == end {
			return 1.0
		}
		// 广度优先搜索的队列
		queue := make([]string, 0, 8)
		queue = append(queue, start)
		// record 一方面记录已访问 一方面记录到这个节点的数值
		record := make(map[string]float64, len(graph))
		record[start] = 1 // 初始结点为1
		for len(queue) != 0 {
			front := queue[0]
			queue = queue[1:]
			res := record[front]
			for k, v := range graph[front] {
				if k == end { // 到达终点
					return res * v
				}
				// 已访问忽略
				if _, ok := record[k]; ok {
					continue
				}
				// 标注已访问 并记录数值
				record[k] = res * v
				queue = append(queue, k)
			}
		}
		// 没有找到
		return -1.0
	}
	// 寻找结果，搜索图
	res := make([]float64, 0, len(queries))
	for _, v := range queries {
		res = append(res, bfs(v[0], v[1]))
	}
	return res
}
