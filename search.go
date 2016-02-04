package bits

/**
 * Given a word, returns array of words, prefix of which is word
 */
func (f *FrozenTrie) GetSuggestedWords(word string, limit int) []string {
	var result []string

	node := f.GetRoot()

	// find the node corresponding to the last char of input
	for _, runeValue := range word {
		var child FrozenTrieNode
		var j uint = 0
		for ; j < node.GetChildCount(); j++ {
			child = node.GetChild(j)
			if child.letter == string(runeValue) {
				break
			}
		}

		// not found, return.
		if j == node.GetChildCount() {
			return result
		}

		node = child
	}

	// The node corresponding to the last letter of word is found.
	// Use this node as root. traversing the trie in level order.
	return f.traverseSubTrie(node, word, limit)
}

func (f *FrozenTrie) traverseSubTrie(node FrozenTrieNode, prefix string, limit int) []string {
	var result []string

	var level []FrozenTrieNode
	level = append(level, node)
	var prefixLevel []string
	prefixLevel = append(prefixLevel, prefix)

	for len(level) > 0 {
		nodeNow := level[0]
		level = level[1:]
		prefixNow := prefixLevel[0]
		prefixLevel = prefixLevel[1:]

		// if the prefix is a legal word.
		if nodeNow.final {
			result = append(result, prefixNow)
			if len(result) > limit {
				return result
			}
		}

		var i uint = 0
		for ; i < nodeNow.GetChildCount(); i++ {
			child := nodeNow.GetChild(i)
			level = append(level, child)
			prefixLevel = append(prefixLevel, prefixNow+child.letter)
		}
	}

	return result
}
