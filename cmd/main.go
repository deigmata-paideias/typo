package main

import (
  "fmt"
  "os"

  "github.com/deigmata-paideias/typo/cmd/root"
)

func main() {

  if err := root.GetRootCommand().Execute(); err != nil {
    _, _ = fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
