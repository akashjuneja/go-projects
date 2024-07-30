


In Go, variables are used to store values of various types. Here's a basic overview of how to define and use variables in Go:

### Defining Variables 
 
1. **Using the `var` keyword:** 

```go
var name string
var age int
var salary float64
```
 
2. **Declaring and initializing in one line:** 

```go
var name string = "Alice"
var age int = 30
var salary float64 = 50000.50
```
 
3. **Type inference:** 

```go
var name = "Alice"
var age = 30
var salary = 50000.50
```
 
4. **Short variable declaration:** 

```go
name := "Alice"
age := 30
salary := 50000.50
```

### Using Variables 

Once you have defined a variable, you can use it in your code. Here are some examples:
 
1. **Printing variables:** 

```go
package main

import "fmt"

func main() {
    var name string = "Alice"
    age := 30
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```
 
2. **Performing operations:** 

```go
package main

import "fmt"

func main() {
    var x int = 10
    var y int = 20
    sum := x + y
    fmt.Println("Sum:", sum)
}
```
 
3. **Updating variables:** 

```go
package main

import "fmt"

func main() {
    var count int = 1
    fmt.Println("Initial count:", count)
    count = 2
    fmt.Println("Updated count:", count)
}
```

### Example Program 

Here's a complete example that demonstrates variable declaration, initialization, and usage in a Go program:


```go
package main

import "fmt"

func main() {
    // Variable declarations
    var name string = "Alice"
    var age int = 30
    var salary float64 = 50000.50

    // Short variable declaration
    location := "New York"

    // Using variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Salary:", salary)
    fmt.Println("Location:", location)

    // Performing operations
    yearsOfExperience := 5
    newAge := age + yearsOfExperience
    fmt.Println("Age after", yearsOfExperience, "years:", newAge)

    // Updating variables
    salary += 10000
    fmt.Println("Updated salary:", salary)
}
```

This should give you a good starting point for working with variables in Go. Feel free to ask if you have more questions or need further examples!
