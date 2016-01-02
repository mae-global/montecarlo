package montecarlo


func RedWhiteBlackBox(r []int,max,mid int) (int,int,int) {

	blacks := 0
	whites := 0

	for i := 0; i < len(r); i++ {
		if (r[i] % max) > mid {
			blacks ++
		} else if (r[i] % max) == mid {
			whites ++
		}
	}

	return len(r) - (blacks + whites),whites,blacks
}

func RedBlack(r []int,max,mid int) float64 {

	reds,whites,blacks := RedWhiteBlackBox(r,max,mid)
	if reds > (blacks + whites) {
		return 1.0
	}
	return 0.0
}	

func WhiteBlack(r []int,max,mid int) float64 {

	reds,whites,blacks := RedWhiteBlackBox(r,max,mid)
	if (reds + whites) > blacks {
		return 1.0
	}
	return 0.0
}


