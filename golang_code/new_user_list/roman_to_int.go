package new_user_list

import (
	"fmt"
	"goland_code/type_def"
	//"go_puzz_code/utils"
)

func calculate(nextValue, num int) (result int) {
	result = 0
	if nextValue > num { // 当前值加减取决于下一个值
		result -= num
	} else {
		result += num
	}
	return result
}

func RomanToInt(s string) int {
	resultInt := 0
	nextValue := -1

	romanIntMap := make(map[byte]int)
	romanIntMap['I'] = 1
	romanIntMap['V'] = 5
	romanIntMap['X'] = 10
	romanIntMap['L'] = 50
	romanIntMap['C'] = 100
	romanIntMap['D'] = 500
	romanIntMap['M'] = 1000

	byteS := []byte(s)

	for i, tmpChar := range byteS {
		tmpNum := romanIntMap[tmpChar]
		nextValue = -1
		if i+1 < len(byteS) {
			nextValue = romanIntMap[byteS[i+1]]
		}

		resultInt += calculate(nextValue, tmpNum)
	}

	return resultInt
}

func InitRomanToIntFunc(funcMap map[int]type_def.LProjectImpl) {
	lP := type_def.NewLProjectImpl("罗马数字转整数", 0, runRomanToInt, 13)
	funcMap[lP.CodeNum] = lP
}

func runRomanToInt() {
	arguments := []string{"MCMXCIV"}
	fmt.Printf("args |\t %v \n", arguments)
	results := RomanToInt(arguments[0])
	fmt.Printf("resp |\t %v \n", results)
	//utils.PrintLine()
}
