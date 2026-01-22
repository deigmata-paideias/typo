package scanner

import (
  "fmt"

  "github.com/deigmata-paideias/typo/internal/scanner"
  "github.com/deigmata-paideias/typo/internal/scanner/custom"
  "github.com/deigmata-paideias/typo/internal/types"
)

func RunScanner(t types.CommandType) error {

  switch t {
  case types.Alias:
    if err := execAliasScanner(); err != nil {
      return err
    }
  case types.Man:
    if err := execManScanner(); err != nil {
      return err
    }
  default:
    fmt.Println("not support")
  }

  return nil
}

func execAliasScanner() error {
  aliasScanner := scanner.NewAliasScanner()
  output, err := aliasScanner.Scan()
  if err != nil {
    return err
  }
  fmt.Println(output)
  // add custom
  gitAliasScanner := custom.NewGitAliasScanner()
  gitOutput, err := gitAliasScanner.Scan()
  if err != nil {
    return err
  }
  fmt.Println(gitOutput)

  return nil
}

func execManScanner() error {
  manScanner := scanner.NewManScanner()
  output, err := manScanner.Scan()
  if err != nil {
    return err
  }
  fmt.Println(output)

  return nil
}
