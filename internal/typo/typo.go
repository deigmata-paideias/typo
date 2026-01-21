package typo

import "fmt"

type ITypo interface {
	// Typo 输入错误的命令，返回正确的命令
	Typo(cmd string) string
}

type Typo struct {
}

func NewTypo() *Typo {
	return &Typo{}
}

func (t *Typo) Typo(cmd string) string {

	fmt.Println(cmd)
	return cmd
}
