package leetcode

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 *
 * https://leetcode.com/problems/implement-trie-prefix-tree/description/
 *
 * algorithms
 * Medium (51.64%)
 * Likes:    4134
 * Dislikes: 65
 * Total Accepted:    384.6K
 * Total Submissions: 743.8K
 * Testcase Example:  '["Trie","insert","search","search","startsWith","insert","search"]\n' +
  '[[],["apple"],["apple"],["app"],["app"],["app"],["app"]]'
 *
 * Implement a trie with insert, search, and startsWith methods.
 *
 * Example:
 *
 *
 * Trie trie = new Trie();
 *
 * trie.insert("apple");
 * trie.search("apple");   // returns true
 * trie.search("app");     // returns false
 * trie.startsWith("app"); // returns true
 * trie.insert("app");
 * trie.search("app");     // returns true
 *
 *
 * Note:
 *
 *
 * You may assume that all inputs are consist of lowercase letters a-z.
 * All inputs are guaranteed to be non-empty strings.
 *
 *
*/

// @lc code=start
type Trie struct {
	Next   map[rune]*Trie
	IsLeaf bool
}

func NewTrie() *Trie {
	return &Trie{
		Next: make(map[rune]*Trie),
	}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Next: make(map[rune]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	tail := this
	for _, char := range word {
		if _, ok := tail.Next[char]; !ok {
			tail.Next[char] = NewTrie()
		}
		tail = tail.Next[char]
	}
	tail.IsLeaf = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	tail := this
	for _, char := range word {
		if tail = tail.Next[char]; tail == nil {
			return false
		}
	}
	return tail.IsLeaf
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	tail := this
	for _, char := range prefix {
		if tail = tail.Next[char]; tail == nil {
			return false
		}
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
// @lc code=end
