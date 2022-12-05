package new_user_list

import (
	"fmt"
	"goland_code/type_def"
	"goland_code/utils"
)

// 判断 magazine 字符串是否包含所有 ransomNote 中出现的字符串 | 包括个数
func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	// 思路1 用map统计 magazine 中出现的字母个数, ransomNote 出现则次数减一
	letterPool := make(map[byte]int)
	for _, char := range []byte(magazine) {
		letterPool[char] += 1
	}

	for _, char := range []byte(ransomNote) {
		if value, ok := letterPool[char]; ok && value > 0 {
			letterPool[char] = value - 1
		} else {
			return false
		}
	}

	return true
}

func canConstruct2(ransomNote, magazine string) bool { // 官方题解：思路很好
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

func InitRansomNoteFunc(funcMap map[int]type_def.LProjectImpl) {
	impl := type_def.NewLProjectImpl("赎金信", 0, runCanConstruct, 383)
	funcMap[impl.CodeNum] = impl
}

func runCanConstruct() {
	arguments := []string{"abc", "aabbcc"}
	fmt.Printf("args |\t %v \n", arguments)
	results := canConstruct(arguments[0], arguments[1])
	fmt.Printf("resp |\t %v \n", results)
	utils.PrintLine()
}
