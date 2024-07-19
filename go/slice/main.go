package main
 
import "fmt"
 
func main() {
 
    //sliceLengthAndCapacity()
    // sliceMemory()
    // mapMemorySpace()
    // sliceCapacityIncrease()
    // printSliceAddress()
    sliceMemoryAllocation()
}

func sliceMemoryAllocation() {
    x := []int{0,1,2}
    y := x
    fmt.Println("x:", x, "y:", y)
    y[0] = 10
    fmt.Println("x:", x, "y:", y)

    z := []int{}
    z = append(z, x...)
    fmt.Println("x:", x, "z:", z)
    z[0] = 100
    fmt.Println("x:", x, "z:", z)
}

func printSliceAddress() {
    x := []int{}
    fmt.Printf("length: %d\t capacity: %d\taddress: %p\t data: %v\n", len(x), cap(x), &x, x)
}

func sliceCapacityIncrease() {
    x := []int{1,2,3,4}
    fmt.Printf("length: %d\t capacity: %d\taddress: %p\t data: %v\n", len(x), cap(x), &x, x)
    x = append(x, 3)
    fmt.Printf("length: %d\t capacity: %d\taddress: %p\t data: %v\n", len(x), cap(x), &x, x)
    x = append(x, 3,3,4,5)
    fmt.Printf("length: %d\t capacity: %d\taddress: %p\t data: %v\n", len(x), cap(x), &x, x)
}

func mapMemorySpace() {
    var map1 = make(map[string]int, 0)
    map1["sushil"] = 30
    map2 := map1
    fmt.Println(map2)
    map2["sushil"] = 31
    fmt.Println(map1)
}

func sliceLengthAndCapacity() {

    // Creating an array
    arr := [12]string{"This","is", "a","Go","interview","question","This","is", "a","Go","interview","question"}
 
    // Print array
    fmt.Println("Original Array:", arr)
 
    // Create a slice
    slicedArr := arr[2:4]
 
    // Display slice
    fmt.Println("Sliced Array:", slicedArr)
 
    // Length of slice calculated using len()
    fmt.Printf("Length of the slice: %d\n", len(slicedArr))
 
    // Capacity of slice calculated using cap()
    fmt.Printf("Capacity of the slice: %d\n\n\n", cap(slicedArr))

}
func sliceMemory() {
    slice1 := []int{1, 2}
    slice3 := slice1
    slice2 := []int{3, 4}
    copy(slice1, slice2)
    fmt.Println(slice1, slice2, slice3)
    fmt.Printf("%p, %p, %p", &slice1, &slice2, &slice3)
}
