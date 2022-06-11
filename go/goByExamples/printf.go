package main

import "fmt"

func main(){
	var int32 int32 = 2147483647
	fmt.Printf("%T: %d (maximum value)\n", int32, int32)

	var int64 int64 = 9223372036854775807
	fmt.Printf("%T: %d (maximum value)\n", int64, int64)

	// var int31 int32 = 0
	// fmt.Printf("maximum value can be storted in int32: %d\n", int32)
	
	var float1 float32 = 3.14151617181920212223242526 
	fmt.Printf("%T: %f (using %%f)\n", float1, float1)

	var float2 float64 = 3.14151617181920212223242526 
	fmt.Printf("%T: %f (using %%f)\n", float2, float2)

	var float32 float32 = 3.14151617181920212223242526 
	fmt.Println("float32: ", float32)

	var float64 float64 = 3.14151617181920212223242526 
	fmt.Println("float64: ", float64)

	fmt.Printf("%T: %.2f (2 precision)\n", float2, float2)
	fmt.Printf("%T: %.3f (3 precision)\n", float2, float2)
	fmt.Printf("%T: %.4f (4 precision)\n", float2, float2)
	fmt.Printf("%T: %.5f (5 precision)\n", float2, float2)
	fmt.Printf("%T: %.6f (6 precision)\n", float2, float2)

	var myBool bool = true
	fmt.Println("boolean: ", myBool)

	var myName string = "sushil sanjay bhile"
	fmt.Println("string:", myName)	

	var binary int = 25
	fmt.Printf("binary of 25: %b\n", binary)

	var char string = "s"
	fmt.Println("character:", char)

	var hexNum int = 15
	fmt.Printf("hex of %d: %x\n", hexNum, hexNum)

}