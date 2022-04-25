package main

func bubbleSort(s []int)[]int{
	for i := 0;i < len(s)-2 -i;i++ {
		if s[i] > s[i+1]{
			s[i],s[i+1] = s[i+1],s[i]
		}
	}
	return s
}
