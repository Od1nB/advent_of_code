package comparison

func Intersection[C comparable](x, y []C) []C {
	seen := map[C]bool{}
	set := []C{}
	for _, v := range x {
		seen[v] = true
	}

	for _, v := range y {
		if seen[v] {
			set = append(set, v)
		}
	}
	return set
}
