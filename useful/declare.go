package useful

// Create slice
func CreateSlice() []int {
	slice := []int{}
	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}
	return slice
}

// Create int map
func CreateMap() map[int]int {
	//var mymap map[int]int
	//or
	//mymap := make(map[int]int)
	//or
	mymap := map[int]int{}
	for v := 1; v <= 10; v++ {
		k := (v - 1)
		mymap[k] = v
	}
	return mymap
}
