package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	completed bool
	value     int
	children  map[rune]*TrieNode
}

func (t *Trie) insert(s string, val int) {
	curr_node := t.root
	for _, r := range s {
		value, exist := curr_node.children[r]
		if !exist {
			new_child := new(TrieNode)
			new_child.children = make(map[rune]*TrieNode)
			curr_node.children[r] = new_child
			curr_node = new_child
		} else {
			curr_node = value
		}
	}
	curr_node.completed = true
	curr_node.value = val
}

func (t *Trie) get(s string) (int, bool) {
	// fmt.Println(s)
	curr_node := t.root
	for _, r := range s {
		value, exist := curr_node.children[r]
		if !exist {
			return -1, false
		} else {
			if value.completed {
				return value.value, true
			}
			curr_node = value
		}
	}
	if curr_node.completed {
		return curr_node.value, true
	}
	return -1, false
}

func (t *Trie) getLeftNumber(s string) int {
	for idx, r := range s {
		if unicode.IsDigit(r) {
			return int(r - '0')
		} else {
			if value, exist := t.get(s[idx:min(idx+5, len(s))]); exist == true {
				return value
			}
		}
	}
	return 0
}

func (t *Trie) getRightNumber(s string) int {
	runes := []rune(s[0:])
	for idx := len(s) - 1; idx > -1; idx-- {
		if unicode.IsDigit(runes[idx]) {
			return int(runes[idx] - '0')
		} else {
			if value, exist := t.get(s[idx:min(idx+5, len(s))]); exist == true {
				return value
			}
		}
	}
	return 0
}

func main() {
	trie := new(Trie)
	trie.root = new(TrieNode)
	trie.root.children = make(map[rune]*TrieNode)
	number_list := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for idx, num := range number_list {
		trie.insert(num, idx+1)
	}
	number_list2 := []string{"oness", "twos", "three3", "four4", "five1", "sixd", "sevens", "eighty", "nineo"}
	for _, num := range number_list2 {
		val, _ := trie.get(num)
		fmt.Println(val)
	}

	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	total := 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		total += trie.getLeftNumber(line) * 10
		total += trie.getRightNumber(line)
	}
	fmt.Println(total)
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
