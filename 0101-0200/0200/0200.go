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

func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	set := newSet(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {
				continue
			}
			grid[i][j] = 'v' // 标记已经走过，避免重复的合并判断
			i1 := i*n + j
			if i > 0 && grid[i-1][j] == '1' {
				set.union(i1, (i-1)*n+j)
			}
			if i < m-1 && grid[i+1][j] == '1' {
				set.union(i1, (i+1)*n+j)
			}
			if j > 0 && grid[i][j-1] == '1' {
				set.union(i1, i*n+j-1)
			}
			if j < n-1 && grid[i][j+1] == '1' {
				set.union(i1, i*n+j+1)
			}
		}
	}
	return set.count
}

type disjoinSet struct {
	parents []int
	count   int // 计数，以免后续统计根节点
}

func newSet(grid [][]byte) *disjoinSet {
	m, n := len(grid), len(grid[0])
	set := &disjoinSet{
		parents: make([]int, 0, m*n),
	}
	flag := 0
	// 初始化
	// 合法点的父节点标记为自己
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				set.parents = append(set.parents, flag)
				set.count++
			} else {
				set.parents = append(set.parents, -1)
			}
			flag++
		}
	}
	return set
}

func (d *disjoinSet) find(i int) int {
	if i != d.parents[i] { // 查询根节点的过程中就 压缩路径
		d.parents[i] = d.find(d.parents[i])
	}
	return d.parents[i]
}

func (d *disjoinSet) union(v1, v2 int) {
	i1 := d.find(v1)
	i2 := d.find(v2)
	if i1 == i2 { // 同一颗树就不合并
		return
	}
	d.parents[i2] = i1 // 合并并查集
	/*  这里没必要压缩路径，造成时间复杂度极高
	for i := range d.parents {
		if d.parents[i] == i2 {
			d.parents[i] = i1
		}
	}
	*/
	d.count-- // 消失一个岛
}
