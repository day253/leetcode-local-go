package external

type sortProblems struct {
	question [][]int
	answer   []int
}

func problems() []sortProblems {
	return []sortProblems{
		{
			question: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			answer: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}
}
