// package main

// import "fmt"

// func main() {
// 	fullName := "Anna Soriano"
// 	age := 29
// 	fmt.Println(fullName, age)
// }
// package main

// import (
//     "fmt"
//     "strconv"
// )

// func main() {
//     i, _ := strconv.Atoi("-42")
//     s := strconv.Itoa(-42)
//     fmt.Println(i, s)
// }

// package main

// import (
//     "fmt"
//     "os"
//     "strconv"
// )

// func main() {
//     number1, _ := strconv.Atoi(os.Args[1])
//     number2, _ := strconv.Atoi(os.Args[2])
//     fmt.Println("Sum:", number1+number2)
// }
// package main

// import (
//     "fmt"
//     "os"
//     "strconv"
// )

// func main() {
//     sum, mul := calc(os.Args[1], os.Args[2])
//     fmt.Println("Sum:", sum)
//     fmt.Println("Mul:", mul)
// }

// func calc(number1 string, number2 string) (sum int, mul int) {
//     int1, _ := strconv.Atoi(number1)
//     int2, _ := strconv.Atoi(number2)
//     sum = int1 + int2
//     mul = int1 * int2
//     return
// }

// package main

// import "fmt"



// func main() {
// 	var x int = 10
// 	ptr := &x
// 	fmt.Println(*ptr)

// 	*ptr = 20
// 	fmt.Println(x)
// }
package main

import "fmt"

// Función que recibe un puntero a un int y modifica su valor
func changeValue(num *int) {
    *num = 42  // Modifica el valor apuntado por num
}

func main() {
    x := 10
    fmt.Println("Antes de la función:", x)
    changeValue(&x)  // Pasa la dirección de x a la función
    fmt.Println("Después de la función:", x)
}

