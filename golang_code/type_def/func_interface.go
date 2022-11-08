package type_def

import "fmt"

type LProject interface {
	GetLName() string
	GetLLevel() int // 0 easy | 1 mid | 2 hard
	GetInitFunc() RunFunc
	GetCodeNum() int
	Print()
}

type LProjectImpl struct {
	Name     string
	Level    int
	InitFunc RunFunc
	CodeNum  int
}

func NewLProjectImpl(name string, level int, initFunc RunFunc, codeNum int) LProjectImpl {
	return LProjectImpl{Name: name, Level: level, InitFunc: initFunc, CodeNum: codeNum}
}

func (l *LProjectImpl) GetLName() string {
	return l.Name
}

func (l *LProjectImpl) GetLLevel() int {
	return l.Level
}

func (l *LProjectImpl) GetInitFunc() RunFunc {
	return l.InitFunc
}

func (l *LProjectImpl) GetCodeNum() int {
	return l.CodeNum
}

func (l *LProjectImpl) Print() {
	fmt.Println("\tNo.\tdifficulty\ttitle")
	difficultyStr := ""
	switch l.GetLLevel() {
	case 0:
		{
			difficultyStr = fmt.Sprintf("\033[1;32m%s\u001B[0m\t", "easy")
		}
		break
	case 1:
		{
			difficultyStr = fmt.Sprintf("\033[1;34m%s\u001B[0m\t", "Medium")
		}
		break
	case 2:
		{
			difficultyStr = fmt.Sprintf("\033[1;35m%s\u001B[0m\t", "hard")
		}
		break
	}
	fmt.Printf("\t%d\t%s\t%s", l.GetCodeNum(), difficultyStr, l.GetLName())
}
