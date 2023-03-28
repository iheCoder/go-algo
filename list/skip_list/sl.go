package skip_list

type skipList struct {
	head  *node
	count int64
}

type node struct {
	o        *obj
	score    float64
	backward *node
	nl       *nodeLevel
}

type nodeLevel struct {
	forward *node
	span    uint
}

type obj struct {
}
