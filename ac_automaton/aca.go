package ac_automaton

type Node struct {
	children map[rune]*Node // 子节点
	fail     *Node          // 失败指针
	output   []string       // 输出模式串列表
}

func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		output:   []string{},
	}
}

type ACAutomaton struct {
	root *Node
}

func NewACAutomaton() *ACAutomaton {
	return &ACAutomaton{
		root: NewNode(),
	}
}

// Insert 插入模式串
func (ac *ACAutomaton) Insert(pattern string) {
	// 从根节点开始插入模式串
	current := ac.root

	// 逐字符插入，如果当前字符不存在，则创建新节点
	for _, ch := range pattern {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = NewNode()
		}
		current = current.children[ch]
	}
	current.output = append(current.output, pattern)
}

// Build 构建失败指针
func (ac *ACAutomaton) Build() {
	queue := []*Node{ac.root}

	// 根节点的子节点 fail 指向 root
	for _, child := range ac.root.children {
		child.fail = ac.root
		queue = append(queue, child)
	}

	// BFS 构建失败指针
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for ch, child := range current.children {
			// 失败指针初始指向当前节点的 fail
			failNode := current.fail

			// 沿着失败指针寻找匹配的子节点
			for failNode != nil && failNode.children[ch] == nil {
				failNode = failNode.fail
			}

			// 如果找到匹配的子节点，则设置失败指针
			if failNode == nil {
				child.fail = ac.root
			} else {
				child.fail = failNode.children[ch]
				// 将失败指针指向的节点的输出模式串添加到当前节点的输出模式串列表中
				child.output = append(child.output, child.fail.output...)
			}

			queue = append(queue, child)
		}
	}
}
