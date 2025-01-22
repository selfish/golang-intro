package main

import(
"fmt"
)

func main(){
     fmt.Println("Hello, Gofmt world!");
     // Extra spaces, inconsistent indentation, etc.

    var someVar = "someVal"
    if (someVar!="") {
       fmt.Println("someVar is not empty");
     } else {
 fmt.Println("someVar is empty");
 }

  // Example slice to demonstrate '-s' simplification
  var numbers = []int{1,2,3,4,5}
   var newSlice = numbers[0: len(numbers)]
    fmt.Println("newSlice:", newSlice)
}
