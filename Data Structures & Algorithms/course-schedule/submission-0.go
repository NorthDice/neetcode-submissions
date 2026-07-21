func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for _,pre := range prerequisites {
		course,prereq := pre[0],pre[1]
		adj[course] = append(adj[course], prereq)
	}

	states := make([]int,numCourses)

	var dfs func(course int) bool
	
	dfs = func(course int) bool {
		if states[course] == 1 {
			return false
		}
		if states[course] == 2 {
			return true
		}
		states[course] = 1

		for _, prereq := range adj[course] {
			if !dfs(prereq){
				return false
			}
		}

		states[course] = 2
		return true
	}

	for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return false
        }
    }	

	return true
	
}
