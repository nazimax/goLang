package main

import("fmt"
        "math/rand")
func pointers(a int){
  fmt.Println("memory address of a is : ",&a)
  fmt.Println("the content of memory address of a is : ",*(&a))
}
func random(){
  fmt.Println("A number from 1-100 ",rand.Intn(100))
}
func add(x,y float64) float64{

  return x+y

}
func main()  {
//random()
a := 5
pointers(a)
}
