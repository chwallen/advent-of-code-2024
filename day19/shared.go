package day19

import (
	"strings"
)

type node struct {
	white   *node
	blue    *node
	black   *node
	red     *node
	green   *node
	pattern string
}

func createTree(input string) *node {
	root := &node{}
	patterns := strings.Split(input, ", ")
	for _, pattern := range patterns {
		currentNode := root
		for i := range pattern {
			currentNode = currentNode.getOrCreateChild(pattern, i)
		}
	}
	return root
}

func (n *node) getChild(char uint8) *node {
	return *n.getChildField(char)
}

func (n *node) getPattern() string {
	return n.pattern
}

func (n *node) getChildField(char uint8) **node {
	switch char {
	case 'w':
		return &n.white
	case 'u':
		return &n.blue
	case 'b':
		return &n.black
	case 'r':
		return &n.red
	case 'g':
		return &n.green
	default:
		panic("invalid char " + string(char))
	}
}

func (n *node) getOrCreateChild(pattern string, patternIndex int) *node {
	char := pattern[patternIndex]
	childField := n.getChildField(char)
	child := *childField
	if child == nil {
		child = &node{}
		*childField = child
	}
	if len(pattern)-1 == patternIndex {
		child.pattern = pattern
	}
	return child
}
