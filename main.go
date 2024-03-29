package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)
  if eludedGuards >= 50 {
    fmt.Print("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.\n")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
  fmt.Println("isHeistOn is currently: ", isHeistOn)
  openedVault := rand.Intn(100)
  if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault can't be opened schmucko!")
  }
  leftSafely := rand.Intn(5)
  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("You tripped the alarm nimwit")
      case 1:
        isHeistOn = false
        fmt.Println("Turns out vaults don't open from the inside dimwit.")
      case 2:
        isHeistOn = false
        fmt.Println("Jeez you stink at this.")
      case 3:
        isHeistOn = false
        fmt.Println("You forgot to add gas to your getaway car duh!")
      default:
        isHeistOn = true
        fmt.Println("Let's go go go go go go go go!!!!!")
    }
    if isHeistOn {
      amtStolen := 10000 + rand.Intn(1000000)
      fmt.Println("$", amtStolen, "not bad!!!")
    }
  }
}