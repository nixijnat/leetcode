package main

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"), true)   // 返回 true
	fmt.Println(trie.Search("app"), false)    // 返回 false
	fmt.Println(trie.StartsWith("app"), true) // 返回 true
	trie.Insert("app")
	fmt.Println(trie.Search("app"), true) // 返回 true
}
