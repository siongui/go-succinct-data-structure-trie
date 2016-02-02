package bits

/**
  A Trie node, for use in building the encoding trie. This is not needed for
  the decoder.
*/
type TrieNode struct {
	letter   string
	final    bool
	children []*TrieNode
}

type Trie struct {
	previousWord string
	root         *TrieNode
	cache        []*TrieNode
	nodeCount    uint
}

func (t *Trie) Init() {
	t.previousWord = ""
	t.root = &TrieNode{
		letter: " ",
		final:  false,
	}
	t.cache = append(t.cache, t.root)
	t.nodeCount = 1
}

/**
  Returns the number of nodes in the trie
*/
func (t *Trie) GetNodeCount() uint {
	return t.nodeCount
}

/**
  Inserts a word into the trie. This function is fastest if the words are
  inserted in alphabetical order.
*/
func (t *Trie) Insert(word string) {

	commonPrefix := 0

	min := len(word)
	if min > len(t.previousWord) {
		min = len(t.previousWord)
	}

	for i := 0; i < min; i++ {
		if word[i] != t.previousWord[i] {
			break
		}
		commonPrefix += 1
	}

	t.cache = t.cache[:commonPrefix+1]
	node := t.cache[commonPrefix]

	for i := commonPrefix; i < len(word); i++ {
		// fix the bug if words not inserted in alphabetical order
		isLetterExist := false
		for _, cld := range node.children {
			if cld.letter == word[i:i+1] {
				t.cache = append(t.cache, cld)
				node = cld
				isLetterExist = true
				break
			}
		}
		if isLetterExist {
			continue
		}

		next := &TrieNode{
			letter: word[i : i+1],
			final:  false,
		}
		t.nodeCount++
		node.children = append(node.children, next)
		t.cache = append(t.cache, next)
		node = next
	}

	node.final = true
	t.previousWord = word
}

/**
  Apply a function to each node, traversing the trie in level order.
*/
func (t *Trie) Apply(fn func(*TrieNode)) {
	var level []*TrieNode
	level = append(level, t.root)
	for len(level) > 0 {
		node := level[0]
		level = level[1:]
		for i := 0; i < len(node.children); i++ {
			level = append(level, node.children[i])
		}
		fn(node)
	}
}

/**
  Encode the trie and all of its nodes. Returns a string representing the
  encoded data.
*/
func (t *Trie) Encode() string {
	// Write the unary encoding of the tree in level order.
	bits := BitWriter{}
	bits.Write(0x02, 2)
	t.Apply(func(node *TrieNode) {
		for i := 0; i < len(node.children); i++ {
			bits.Write(1, 1)
		}
		bits.Write(0, 1)
	})

	// Write the data for each node, using (dataBits) bits for one node.
	// 1 bit stores the "final" indicator. The other (dataBits-1) bits store
	// one of the characters of the alphabet.
	t.Apply(func(node *TrieNode) {
		value := mapCharToUint[node.letter]
		if node.final {
			value |= (1 << (dataBits - 1))
		}

		bits.Write(uint(value), dataBits)
	})

	return bits.GetData()
}
