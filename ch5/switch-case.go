package main
import "fmt"

func main(){

    fmt.Println("Enter number 0-9: ")
    var input  int64
    fmt.Scanf("%d", &input)

    switch input{
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    default: fmt.Println("Unknow Number")
    }
}
