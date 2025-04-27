package main

import (
 "bufio"
 "fmt"
 "math"
 "os"
 "strings"
)

func main() {
 scanner := bufio.NewReader(os.Stdin)

 fmt.Println("=== CLI Calculator ===")
 fmt.Println("Available operations: +  -  *  /  sqrt  ^")

 for {
  fmt.Print("\nChoose operation: ")
  opInput, err := scanner.ReadString('\n')
  if err != nil {
   fmt.Println("Error reading input:", err)
   continue
  }
  op := strings.TrimSpace(opInput)

  var a, b float64

  switch op {
  case "sqrt":
   fmt.Print("Enter number: ")
   _, err := fmt.Scanln(&a)
   if err != nil {
    fmt.Println("Invalid input.")
    continue
   }
   if a < 0 {
    fmt.Println("Error: Square root of a negative number.")
    continue
   }
   fmt.Printf("Result: √%.2f = %.4f\n", a, math.Sqrt(a))

  case "+", "-", "*", "/", "^":
   fmt.Print("Enter first number: ")
   _, err1 := fmt.Scanln(&a)
   if err1 != nil {
    fmt.Println("Invalid first number.")
    continue
   }
   fmt.Print("Enter second number: ")
   _, err2 := fmt.Scanln(&b)
   if err2 != nil {
    fmt.Println("Invalid second number.")
    continue
   }

   switch op {
   case "+":
    fmt.Printf("Result: %.4f\n", a+b)
   case "-":
    fmt.Printf("Result: %.4f\n", a-b)
   case "*":
    fmt.Printf("Result: %.4f\n", a*b)
   case "/":
    if b == 0 {
     fmt.Println("Error: Division by zero.")
     continue
    }
    fmt.Printf("Result: %.4f\n", a/b)
   case "^":
    fmt.Printf("Result: %.4f\n", math.Pow(a, b))
   }

  default:
   fmt.Println("Unsupported operation. Try again.")
  }
 }
}