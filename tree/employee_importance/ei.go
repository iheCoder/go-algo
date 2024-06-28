package employee_importance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	m := make(map[int]*Employee)
	var target *Employee
	for _, employee := range employees {
		if employee.Id == id {
			target = employee
		}
		m[employee.Id] = employee
	}
	if target == nil {
		return 0
	}

	var dfs func(e *Employee) int
	dfs = func(e *Employee) int {
		imp := e.Importance
		for _, sub := range e.Subordinates {
			imp += dfs(m[sub])
		}
		return imp
	}
	return dfs(target)
}
