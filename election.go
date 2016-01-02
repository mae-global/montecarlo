package montecarlo


func Election(candiates int,r []int) []int {

	order := make(map[int]int,0)
	for i := 0; i < candiates; i++ {
		order[i] = 0
	}

	for i := 0; i < len(r); i++ {
		order[r[i] % candiates] ++
	}

	out := make([]int,0)

	for {
		if len(order) <= 0 {
			break
		}

		largest := 0
		for key,value := range order {
			if order[largest] < value {
				largest = key
			}
		}
		
		out = append(out,largest + 1)
		delete(order,largest)
	}

	return out
}

func Elimination(candiates,rounds int,r []int) []int {

	/* TODO: add elimination code here */	
	return Election(candiates - rounds,r)
}

