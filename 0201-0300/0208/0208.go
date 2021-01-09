package main

type Trie struct {
	end  bool           // 标记是不是单词结尾
	vals map[byte]*Trie // map 可以快速查询，儿子也是字典树
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		vals: make(map[byte]*Trie, 4),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var l = this.vals
	var n *Trie
	for i := range word {
		n = l[word[i]] // 是否已经存在节点
		if n == nil {  // 不存在则创建
			tmp := Constructor()
			n = &tmp
			l[word[i]] = n // 创建节点
		}
		l = n.vals // 儿子
	}
	n.end = true // 标记为单词结束
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var l = this.vals
	var n *Trie
	for i := range word {
		n = l[word[i]]
		if n == nil {
			return false
		}
		l = n.vals
	}
	return n.end // 单词结束才是存在单词
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	l := this.vals
	for i := range prefix {
		n := l[prefix[i]]
		if n == nil {
			return false
		}
		l = n.vals
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
