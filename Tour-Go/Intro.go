package main

import "fmt"

func main3()  {
   var a,b int
   a = 42
   b = 7
   fmt.Println("a vale:", a)
   fmt.Printf("a as binary: %b \n", a)
   fmt.Printf("b as hexadecimal: %#X \n\n", b)

   // print these values as both binary & hexadecimal
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
   fmt.Printf("Value \t Binary\t Hexadecimal\n")
   fmt.Printf("----- \t ------\t -----------\n")
   fmt.Printf("%v \t %b \t %#X\n", a,a,a)
   fmt.Printf("%v \t %b \t %#X\n", b,b,b)
   fmt.Printf("%v \t %b \t %#X\n", c,c,c)
   fmt.Printf("%v \t %b \t %#X\n", d,d,d)
   fmt.Printf("%v \t %b \t %#X\n", e,e,e)
   fmt.Printf("%v \t %b \t %#X\n", f,f,f)
   
   y := 42
   z := 42.0
   fmt.Printf("%v of type %T \n", y, y)
   fmt.Printf("%v of type %T \n", z, z)

   var m float32 = 43.742
   fmt.Printf("%v of type %T \n", m, m)

   /*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it 
		// in a variable that is declared to hold a VALUE of float64
		z = m 
		fmt.Printf("%v of type %T \n", z, z)
   */

   // this does work
   z = float64(m)
   fmt.Printf("%v of type %T \n", z, z)  
	
}