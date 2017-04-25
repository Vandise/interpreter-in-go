package main

import (
  "fmt"
  "os"
  "os/user"
  "lang/repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
	panic(err)
  }
  fmt.Printf("Hello %s, this is lang v 0.0.0\n", user.Username)
  repl.Start(os.Stdin, os.Stdout)
}
