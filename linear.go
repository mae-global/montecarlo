package montecarlo

func Linear(r []int,min,max int) int {

	list := make([]int,max - min)

	for i := 0; i < len(r); i++ {
		n := (r[i] % (max - min)) + min
		list[n - min] ++
	}

	largest := 0
	n := 0
	for key,value := range list {
		if value > largest {
			n = (key + min)
			largest = value
		}
	}
	
	return n	
}
