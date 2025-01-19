package graph

func ComplexityConfig() ComplexityRoot {
	var c ComplexityRoot

	c.Repository.Issues = func(childComplexity int, after *string, before *string, first *int32, last *int32) int {
		var cnt int
		switch {
		case first != nil && last != nil:
			if *first > *last {
				cnt += int(*first)
			} else {
				cnt += int(*last)
			}
		case first != nil && last == nil:
			cnt = int(*first)
		case first == nil && last != nil:
			cnt = int(*last)
		default:
			cnt = 1
		}
		return cnt * childComplexity
	}		
	return c
}