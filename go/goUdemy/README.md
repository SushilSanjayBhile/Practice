# Interesting Things/Features/Facts of golang

Fact 1) Slices are resizable, we can append elements to slice or we can remove elements from slice

Fact 2) When we append element to slice and the capacity is full then go will create new slice and will assign address of the new slice to the variable

Fact 3) If the capacity of slice is 4 and you add 4 elements and print address of the slice (by &slice_name), and then if you add few elements in the slice and again print the address of slice, it will show the same address. To check if in real address of slice is changed, you need to print address of first element. like: &slice_name[0]

Fact 4) Try not to import package as static import. Like:

        import (
            . "fmt"
        )

        This might cause confusion, when we have multiple static imports, then it is difficult to identify in which package the function's definition is.

        ex. Look inside dontDo/staticImport

Fact 5) In return of a function, we don't have to write name of variables we want to return, if we have mentioned the names in function's return definition.

        ex. Look inside interesting/return

Fact 6) In parameter of function, if multiple variables have same datatype, we can declare that datatype only once and all the variables left to that datatype will have that datatype.

        ex. Look inside interesting/functionparams

Fact 7) We can declare multiple variables in same line.
        Also, The interesting thing is those 2 or more variable need not be of same datatype.

        ex. Look inside interesting/multipleVariables
