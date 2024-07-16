package main

import "fmt"

type data struct {
	key   string
	value int
}

func get(key string, slc []data) (int, []data) {
	for k, v := range slc {
		if v.key == key {
			slcnew := make([]data, 0)
			slcnew = append(slcnew, slc[k])
			slcnew = append(slcnew, slc[:k]...)
			slcnew = append(slcnew, slc[k+1:]...)
			return v.value, slcnew
		}
	}

	return -1, slc
}

func put(key string, value int, slc []data, capacity int) []data {
	userdata := data{
		key:   key,
		value: value,
	}

	if len(slc) != capacity {
		slc = append(slc, userdata)
		return slc
	} else {
		slcnew := make([]data, 0)
		slcnew = append(slcnew, slc[1:]...)
		slcnew = append(slcnew, userdata)
		return slcnew
	}
}

func main() {
	var capacity = 4
	slc := make([]data, 0)

	slc = put("a", 1, slc, capacity)
	slc = put("b", 2, slc, capacity)
	slc = put("c", 3, slc, capacity)
	slc = put("d", 4, slc, capacity)

	slc = put("g", 6, slc, capacity) //b,c,d,g
	slc = put("h", 7, slc, capacity) //c,d,g,h
	slc = put("i", 8, slc, capacity) //d,g,h,i

	fmt.Println(get("h", slc)) // d,g,i,h
	fmt.Println(get("a", slc))

	slc = put("j", 10, slc, capacity)
	slc = put("k", 11, slc, capacity)

	fmt.Println(get("h", slc))
	fmt.Println(get("a", slc))
	fmt.Println(get("g", slc))

	// fmt.Println(slc)
	// slc = put("e", 5, slc, capacity)
	// fmt.Println(slc)

	// fmt.Print("Key: c\t")
	// fmt.Println(get("c", slc))

	// fmt.Print("Key: k\t")
	// fmt.Println(get("k", slc))
}
