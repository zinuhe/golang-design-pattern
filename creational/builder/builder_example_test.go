// go test
// go test -test.run=TestBuilder

package builder_test

import (
  "fmt"
  "testing"
  b "github.com/zinuhe/golang-design-pattern/creational/builder"
)

func TestBuilder(t *testing.T) {
  normalBuilder := b.GetBuilder("normal")
  iglooBuilder := b.GetBuilder("igloo")

  director := b.NewDirector(normalBuilder)
  normalHouse := director.BuildHouse()

  fmt.Println("Normal House Door Type  :", normalHouse.DoorType)
  fmt.Println("Normal House Window Type:", normalHouse.WindowType)
  fmt.Println("Normal House Num Floor  :", normalHouse.Floor)

  director.SetBuilder(iglooBuilder)
  iglooHouse := director.BuildHouse()

  fmt.Println("\nIgloo House Door Type  :", iglooHouse.DoorType)
  fmt.Println("Igloo House Window Type:", iglooHouse.WindowType)
  fmt.Println("Igloo House Num Floor  :", iglooHouse.Floor)
}
