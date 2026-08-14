# 1. Pointers:

- Go has pointers. A pointer holds the memory address of a value.
- The type *T is a pointer to a value of type T. Its zero value is nil.

```bash
var p *int
```

- The & operator generates a pointer to its operand.

```bash
i := 42
p = &i
```
```bash
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```
# 2. Structs:

- A struct is a collection of fields.
- Struct fields are accessed using a dot.
```bash
package main
import "fmt"
type Vertex struct {
	X int
	Y int
}
func main() {
    v := Vertex{1, 2}
	fmt.Println(v)
    # It prints {1 2}
    fmt.Println(v.X)
    # It prints 1
}

```

# 3. Pointers to structs:
- Struct fields can be accessed through a struct pointer.

- To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

```bash
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

# 4. Struct Literals:
- A struct literal denotes a newly allocated struct value by listing the values of its fields.

- You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
- The special prefix & returns a pointer to the struct value.
```bash
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

# 5. Arrays:

- The type [n]T is an array of n values of type T.

```bash
primes := [6]int{2, 3, 5, 7, 11, 13}
```
```bash
var a [10]int
# declares a variable a as an array of ten integers
```
- An array's length is part of its type, so arrays cannot be resized.


# 6. Slices:

- An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
- The type []T is a slice with elements of type T.
- A slice is formed by specifying two indices, a low and high bound, separated by a colon: `a[low : high]`. This selects a half-open range which includes the first element, but excludes the last one.
```bash
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
# It prints [3 5 7]
```

- A slice does not store any data, it just describes a section of an underlying array.
- Changing the elements of a slice modifies the corresponding elements of its underlying array.

# 7. Slice literals:

- A slice literal is like an array literal without the length.
- This is an array literal:
```bash
a := [3]bool{true, true, false}
```
- And this creates the same array as above, then builds a slice that references it:
```bash
a := []bool{true, true, false}
# You can modify inplace, for example: a[1] = false
```
# 8. Slice defaults:
- When slicing, you may omit the high or low bounds to use their defaults instead.
- The default is zero for the low bound and the length of the slice for the high bound.
```bash
var a [10]int

# those slieces are equivalent
a[0:10]
a[:10]
a[0:]
a[:]
```

# 9. Slice length and capacity:
- The length of a slice is the number of elements it contains.
- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
- The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
- You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
```bash
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	# Slice the slice to give it zero length.
	s = s[:7]
	printSlice(s)

	# Extend its length.
	s = s[:4]
	printSlice(s)

	# Drop its first two values.
	s = s[2:]
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
# len=6 cap=6 [2 3 5 7 11 13]
# len=0 cap=6 []
# len=4 cap=6 [2 3 5 7]
# len=2 cap=4 [5 7]

```

- ****A slice that has a length and capacity of zero is called a nil slice. A nil slice has no underlying array.**

# 10. Creating a slcie with make:

- Slices can be created with the built-in make function; this is how you create dynamically-sized arrays (dynamic here means create at runtime).
- The make function allocates a zeroed array and returns a slice that refers to that array:
```bash
a := make([]int, 5)  // len(a)=5
```
- To specify a capacity, pass a third argument to make:
```bash
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

# 11. Slices of slices:
Slices can contain any type, including other slices.
```bash
board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
```

# 11. Appending to a slice:
```
func append(s []T, vs ...T) []T
```
- The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
- If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array

```bash
func main() {
	var s []int
	printSlice(s)
	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)
	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)
	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

# 12. Range:
- The range form of the for loop iterates over a slice or map.
- When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
```bash
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
# 2**0 = 1
# 2**1 = 2
# 2**2 = 4
# 2**3 = 8
# 2**4 = 16
# 2**5 = 32
# 2**6 = 64
# 2**7 = 128
```
- We can skip the index or value by assigning to _.
```
for i, _ := range pow
for _, value := range pow
```
- **If we only want the index, we can omit the second variable entirely**
