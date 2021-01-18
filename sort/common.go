package sort

type sortProblems struct {
	question []int
	answer   []int
}

func problems() []sortProblems {
	return []sortProblems{
		{
			question: []int{3, 4, 5, 2, 1, 7, 8, -1, -3},
			answer:   []int{-3, -1, 1, 2, 3, 4, 5, 7, 8},
		},
		{
			question: []int{3, 4, 5, 2, 1},
			answer:   []int{1, 2, 3, 4, 5},
		},
	}
}
