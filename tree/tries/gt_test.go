package tries

import "testing"

func TestAddRoute(t *testing.T) {
	e := &engine{trees: make([]methodTree, 0)}
	e.addRoute("GET", "node", nil)
	e.addRoute("GET", "node/add", nil)
	e.addRoute("GET", "user", nil)
}
