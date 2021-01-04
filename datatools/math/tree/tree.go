package tree

type Things []interface{}

type Node struct {
	tag      string
	text     string
	parent   *Node
	children []*Node
}

func NewTree(tag string, text string) *Node {
	return &Node{tag: tag, text: text, parent: nil}
}
