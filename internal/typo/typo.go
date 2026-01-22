package typo

import (
  "fmt"

  "github.com/deigmata-paideias/typo/internal/repository"
)

type ITypo interface {
  // Typo 输入错误的命令，返回正确的命令
  Typo(cmd string) string
}

type LocalTypo struct {
  repository.IRepository
}

func NewLocalTypo(repo repository.IRepository) ITypo {

  return &LocalTypo{
    repo,
  }
}

func (t *LocalTypo) Typo(cmd string) string {

  fmt.Println(cmd)
  return cmd
}

type LlmTypo struct {
}

func NewLlmTypo() ITypo {
  return &LlmTypo{}
}

func (t *LlmTypo) Typo(cmd string) string {

  // todo: call llm api to get correct command
  fmt.Println(cmd)
  return cmd
}
