# 1. Packages:

- Structure of a Go program: A module is the whole project, it can contain multiple packages, and each package can contain multiple source files.

```
Example project structure:

myproject/
├── go.mod          ← defines the module
├── main.go         ← package main (entry point)
├── greet/
│   └── greet.go   ← package greet
└── math/
    └── math.go    ← package math
```

- In a file, the package name is the same as the last element for the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

# 2. Exported Names:
- In Go, a name is exported if it begins with a capital letter. Any "unexported" names are not accessible from outside the package.

```bash
#For example, if we use math.pi -> error, it must me math.Pi

package main
import (
    "fmt"
    "math"
)
func main() {
    fmt.Println(math.Pi) // Correct usage
    # fmt.Println(math.pi) // This would cause an error
}
```

# 3. Functions:
- A function can take zero or more arguments.
- Notice that the type comes after the variable name.

```bash
package main
import "fmt"
func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(42, 13))
}
```
- When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

```bash
func add(x, y int) int{
    ...
}
```
- A function can return any number of results.

```bash
func swap(x, y string) (string, string) {
	return y, x
}
```

- Go's return values may be named. If so, they are treated as variables defined at the top of the function. These names should be used to document the meaning of the return values.

```bash
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
# A return statement without arguments returns the named return values. This is known as a "naked" return.
```

# 4. Variables:

- The var statement declares a list of variables; as in function argument lists, the type is last.
- A var statement can be at package or function level. We see both in this example.
```bash
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

```
- A var declaration can include initializers, one per variable. If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```bash
var i, j int = 1, 2
```

- Short variable declaration: Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

```bash
func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
- Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

- Variables declared without an explicit initial value are given their zero value. (0 for numeric types, false for the boolean type, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.)

# 5. Type Conversions:

- The expression T(v) converts the value v to the type T.
```bash
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```
- **Unlike in C, in Go assignment between items of different type requires an explicit conversion. In the above example, if we remove var f float64 = math.Sqrt(x*x + y*y), we will get an error: cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt.**

# 6. Type Inference:

When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type
```bash
var i int
j := i // j is an int
```