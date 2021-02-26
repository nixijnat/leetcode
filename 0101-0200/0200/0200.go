package main

func numIslands_breath(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	queue := make([][2]int, 0, m)
	// 封装入队函数
	appendQueue := func(i, j int) {
		queue = append(queue, [2]int{i, j})
		// NOTE 入队标记为 visit 避免后续重复入队甚至造成死循环
		grid[i][j] = 'v'
	}
	res := 0
	// 遍历所有节点
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 非陆地则忽略
			if grid[i][j] != '1' {
				continue
			}
			res++ // 找到陆地就是岛屿
			// 广度遍历，找到岛屿所有陆地
			// 并标记为 visit
			appendQueue(i, j)
			for len(queue) != 0 {
				v := queue[0]
				queue = queue[1:]
				if v[0] > 0 && grid[v[0]-1][v[1]] == '1' {
					appendQueue(v[0]-1, v[1])
				}
				if v[0] < m-1 && grid[v[0]+1][v[1]] == '1' {
					appendQueue(v[0]+1, v[1])
				}
				if v[1] > 0 && grid[v[0]][v[1]-1] == '1' {
					appendQueue(v[0], v[1]-1)
				}
				if v[1] < n-1 && grid[v[0]][v[1]+1] == '1' {
					appendQueue(v[0], v[1]+1)
				}
			}
		}
	}
	return res
}

type disjointSet struct {
	parent []int
	rank   []int
	count  int // 记录岛屿数量
}

func newSet(nums [][]byte) *disjointSet {
	s := &disjointSet{
		parent: make([]int, len(nums)*len(nums[0])),
		rank:   make([]int, len(nums)*len(nums[0])),
	}
	flag := 0
	for i := range nums {
		for j := range nums[i] {
			if nums[i][j] == '0' {
				s.parent[flag] = -1
			} else {
				s.parent[flag] = flag
				s.count++
			}
			s.rank[flag] = 1
			flag++
		}
	}
	return s
}

func (s *disjointSet) find(i int) int {
	if s.parent[i] != i {
		// 路径压缩
		s.parent[i] = s.find(s.parent[i])
	}
	return s.parent[i]
}

func (s *disjointSet) union(i, j int) {
	p1 := s.find(i)
	p2 := s.find(j)
	if p1 == p2 {
		return
	}
	// rank 小的靠向 rank 高的
	if s.rank[p1] < s.rank[p2] {
		p1, p2 = p2, p1
	}
	s.parent[p2] = p1
	// rank 一样，将被靠的秩加一
	if s.rank[p1] == s.rank[p2] {
		s.rank[p1]++
	}
	s.count-- // 岛屿减少一个
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	s := newSet(grid)
	for i := range grid {
		for j, v := range grid[i] {
			if v == '0' {
				continue
			}
			grid[i][j] = '0' // 避免后续重复判断
			index := i*n + j
			if i-1 >= 0 && grid[i-1][j] == '1' {
				s.union(index, index-n)
			}
			if i+1 < m && grid[i+1][j] == '1' {
				s.union(index, index+n)
			}
			if j-1 >= 0 && grid[i][j-1] == '1' {
				s.union(index, index-1)
			}
			if j+1 < n && grid[i][j+1] == '1' {
				s.union(index, index+1)
			}
		}
	}
	return s.count
}
