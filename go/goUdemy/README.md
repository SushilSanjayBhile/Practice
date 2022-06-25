# Interesting Things/Features/Facts of golang

Fact 1) Slices are resizable, we can append elements to slice or we can remove elements from slice

Fact 2) When we append element to slice and the capacity is full then go will create new slice and will assign address of the new slice to the variable

Fact 3) If the capacity of slice is 4 and you add 4 elements and print address of the slice (by &slice_name), and then if you add few elements in the slice and again print the address of slice, it will show the same address. To check if in real address of slice is changed, you need to print address of first element. like: &slice_name[0]

Fact 4) Try not to import package as static import. Like:

        import (
            . "fmt"
        )

        This might cause confusion, when we have multiple static imports, then it is difficult to identify in which package the function's definition is.