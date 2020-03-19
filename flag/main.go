package main

import (
  "flag"
  "fmt"
)

func main() {

  var ip = flag.String("kubeconfig", "/asd/asd/sad/sad", "default value")

  flag.Parse()

  fmt.Println(*ip)
}
