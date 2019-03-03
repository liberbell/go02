package main

func main() {
  x := 2

  swith x {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3:
      fmt.Println("three")
    default:
      fmt.Printf("many")
    }
}
