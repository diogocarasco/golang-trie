package main

import (
	"fmt"
	"strings"
)

const AlphabetSize = 26

// Node structure containing 26 positions for each letter of the alphabet and a flag for the end of the word.
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie structure containing a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie initializes the trie structure and return its pointer.
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert receives a word and add it to the trie structure
func (t *Trie) Insert(word string) {

	word = strings.ToLower(word)
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {

		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}

		currentNode = currentNode.children[charIndex]

	}

	currentNode.isEnd = true

}

// Search for a given word in the trie structure
func (t *Trie) Search(word string) bool {

	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {

		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]

	}

	if currentNode.isEnd {
		return true
	}

	return false

}

func main() {

	myTrie := InitTrie()

	myTrie.Insert("summer")
	myTrie.Insert("winter")
	found := myTrie.Search("summer")
	notFound := myTrie.Search("sum")
	fmt.Println(found)
	fmt.Println(notFound)

}
