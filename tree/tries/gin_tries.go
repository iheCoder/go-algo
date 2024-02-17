package tries

type engine struct {
	trees methodsTrees
}

type methodsTrees []methodTree

type methodTree struct {
	method string
	root   *node
}

type Handler func()

type node struct {
	// 当前节点path
	path string
	// 子节点首字符组成的string
	indices string
	// 按照优先级排序的子节点
	children []*node
	// 优先级
	priority uint32
	// 根节点到该节点的字符串
	fullPath string
	// 处理器
	handlers []Handler
}

func (n *node) addRoute(path string, handlers []Handler) {
	fullPath := path
	n.priority++

	// 若为空节点，则直接插入子节点
	if len(n.path) == 0 && len(n.children) == 0 {
		n.path = path
		n.fullPath = fullPath
		n.handlers = handlers
		return
	}

	// 根据路径插入到节点
	n.addChild(path, fullPath, 0, handlers)
}

func (n *node) addChild(path, fullPath string, parentFullPathIndex int, handlers []Handler) {
	// 找到新增路径与节点路径的最长公共前缀
	i := longestCommonPrefix(path, n.path)

	// 节点的path不是最长公共前缀，分裂原子节点和新增节点
	if i < len(n.path) {
		child := &node{
			path:     n.path[i:],
			indices:  n.indices,
			children: n.children,
			priority: n.priority - 1,
			fullPath: n.fullPath,
		}
		n.children = []*node{child}
		n.indices = string(n.path[i])
		n.path = path[:i]
		n.fullPath = fullPath[:parentFullPathIndex+i]
	}

	// 在新增路径存在公共前缀时，添加含有该路径的子节点
	if i < len(path) {
		path = path[i:]
		c := path[0]

		// 若该路径的非公共前缀首字符在节点children已经存在，那么就将该节点添加到子节点中
		for j, max := 0, len(n.indices); j < max; j++ {
			if c == n.indices[j] {
				parentFullPathIndex += len(n.path)
				j = n.incrementChildPriority(j)
				n.children[j].addChild(path, fullPath, parentFullPathIndex, handlers)
				return
			}
		}

		// 若该路径的非公共前缀首字符在节点children不存在，就直接添加为节点的子节点
		n.indices += string(c)
		child := &node{fullPath: fullPath, path: path, handlers: handlers}
		n.children = append(n.children, child)
		n.incrementChildPriority(len(n.indices) - 1)
		return
	}

	// 路径与节点路径完全匹配时，直接覆盖原全路径和处理器
	n.handlers = handlers
	n.fullPath = fullPath
	return
}

// 增加子节点的优先级
func (n *node) incrementChildPriority(pos int) int {
	// 增加该子节点的优先级
	cs := n.children
	cs[pos].priority++
	prio := cs[pos].priority

	// 重新根据优先级排序
	newPos := pos
	for ; newPos > 0 && cs[newPos-1].priority < prio; newPos-- {
		cs[newPos-1], cs[newPos] = cs[newPos], cs[newPos-1]
	}

	// 重新组合indices
	if newPos != pos {
		n.indices = n.indices[:newPos] + n.indices[pos:pos+1] + n.indices[newPos:pos] + n.indices[pos+1:]
	}

	return newPos
}

type nodeValue struct {
	handlers []Handler
	tsr      bool
}

func (n *node) getValue(path string) *nodeValue {
	prefix := n.path

	// 若路径不完全匹配当前节点路径
	if len(path) > len(prefix) {
		// 若路径存在该前缀
		if path[:len(prefix)] == prefix {
			path = path[len(prefix):]

			// 尝试找到匹配该路径的子节点
			idxc := path[0]
			for i := 0; i < len(n.indices); i++ {
				if idxc == n.indices[i] {
					n.children[i].getValue(path)
				}
			}

			// 若无法找到匹配子节点，那么就是当前节点了
			return &nodeValue{
				tsr: path == "/" && n.handlers != nil,
			}
		}
	}

	// 路径与节点路径完全匹配
	if path == prefix {
		// 若节点处理器存在，则返回
		if n.handlers != nil {
			return &nodeValue{
				handlers: n.handlers,
			}
		}

		for i := 0; i < len(n.indices); i++ {
			if n.indices[i] == '/' {
				child := n.children[i]
				return &nodeValue{
					tsr: len(child.path) == 1 && child.handlers != nil,
				}
			}
		}
	}

	return nil
}

func (e *engine) addRoute(method, path string, handlers []Handler) {
	// 1. 根据方法获取对应的前缀树。若root为nil，则初始化
	root := e.trees.get(method)
	if root == nil {
		root = &node{
			fullPath: "/",
		}
		e.trees = append(e.trees, methodTree{
			method: method,
			root:   root,
		})
	}

	// 2. 添加路径到前缀树
	root.addRoute(path, handlers)
}

// 遍历树，若存在对应方法所在的树就返回
func (t methodsTrees) get(method string) *node {
	for _, tree := range t {
		if tree.method == method {
			return tree.root
		}
	}

	return nil
}

func longestCommonPrefix(a, b string) int {
	i := 0
	max := minInt(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
