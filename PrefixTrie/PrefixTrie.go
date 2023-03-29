package PrefixTrie

import (
	"fmt"
	"strings"
	"sync"
)

type Trie struct {
	isWord   bool
	children [26]*Trie
}

// NewTrie initialize your data structure here.
func NewTrie() *Trie {
	return &Trie{}
}

// Add inserts a word into the trie.
func (this *Trie) Add(word string) {
	word = strings.ToLower(word)
	cur := this
	for i, c := range []rune(word) {
		n := c - 'a'

		if cur.children[n] == nil {
			cur.children[n] = NewTrie()
		}
		cur = cur.children[n]
		if i == len(word)-1 {
			cur.isWord = true
		}

	}
}

// IsIn returns if the word is in the trie.
func (this *Trie) IsIn(word string) bool {
	word = strings.ToLower(word)
	cur := this
	for _, c := range []rune(word) {
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return cur.isWord
}

// StartsWith returns if there is any word in the trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	prefix = strings.ToLower(prefix)
	for _, c := range []rune(prefix) {
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return true
}

// Search returns list of words starting with prefix
func (this *Trie) Search(prefix string, maxCountOfWords int) []string {
	cur := this
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}
	prefix = strings.ToLower(prefix)
	for _, c := range []rune(prefix) {
		n := c - 'a'
		if cur.children[n] == nil {
			fmt.Println("There is no words with this prefix")
			return []string{}
		}
		cur = cur.children[n]
	} // reach the prefix node

	result := make([]string, 0, maxCountOfWords)
	countOfWords := 0

	if cur.isWord {
		result = append(result, prefix)
		countOfWords++
	} // check if the prefix is word

	wg.Add(1)
	cur.depthSearch(&result, prefix, &countOfWords, &maxCountOfWords, wg, mtx)
	wg.Wait()

	fmt.Println("Count of founded words:", len(result))
	return result
}

// depthSearch is sub method. Search in trie in depth
func (this *Trie) depthSearch(result *[]string, modPrefix string, countOfWords *int, maxCountOfWords *int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	if *countOfWords > *maxCountOfWords-1 {
		fmt.Println("[INFO] the number of matching words is greater than the set maximum")
		return
	}
	isLeaf := true

	for i, node := range this.children {
		if node == nil {
			continue
		} // skip node if empty
		isLeaf = false
		prefix := modPrefix + string(rune(i)+'a')
		wg.Add(1)
		go node.depthSearch(result, prefix, countOfWords, maxCountOfWords, wg, mtx)
	}
	if isLeaf {
		//fmt.Println("[INFO] Prefix is", modPrefix)
		mtx.Lock()
		*result = append(*result, modPrefix)
		*countOfWords++
		mtx.Unlock()
	}
	wg.Done()
}

// example
func main() {
	trie := NewTrie()
	trie.Add("wqwe")
	trie.Add("dasd")
	trie.Add("wead")
	trie.Add("wqeasd")
	trie.Add("wqewqqwdzsdz")
	trie.Add("waqd")
	trie.Add("wqwdsc")
	trie.Add("wqqwed")
	trie.Add("asd")
	trie.Add("czxc")
	trie.Add("wq")
	//fmt.Println(trie.StartsWith("wq"))
	fmt.Println(strings.Join(trie.Search("wq", 10), " "))

}
