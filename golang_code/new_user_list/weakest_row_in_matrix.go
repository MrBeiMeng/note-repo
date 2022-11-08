package new_user_list

import (
	"fmt"
	"goland_code/type_def"
	"goland_code/utils"
)

type Node struct {
	val   int
	index int
}

func kWeakestRows(mat [][]int, k int) []int {
	tmpMat := make([]Node, 0)

	for index, nums := range mat {
		tmpNum := 0
		for _, num := range nums {
			tmpNum += num
		}

		tmpMat = append(tmpMat, struct {
			val   int
			index int
		}{val: tmpNum, index: index})
	}

	// 排序
	selectSort(tmpMat, false)

	resultList := make([]int, 0)
	for _, item := range tmpMat[:k] {
		resultList = append(resultList, item.index)
	}
	return resultList
}

func selectSort(nums []Node, desc bool) {
	for i := 0; i < len(nums); i++ {
		selectedIndex := i
		for j := i + 1; j < len(nums); j++ {
			if desc {
				if nums[selectedIndex].val < nums[j].val {
					selectedIndex = j
				}
			} else {
				if nums[selectedIndex].val > nums[j].val {
					selectedIndex = j
				}
			}
		}

		nums[i], nums[selectedIndex] = nums[selectedIndex], nums[i]
	}
}

func InitWeakestRowInMatrixFunc(funcMap map[int]type_def.LProjectImpl) {
	lP := type_def.NewLProjectImpl("矩阵中战斗力最弱的 K 行", 0, runKWeakestRows, 1337)
	funcMap[lP.CodeNum] = lP
}

func runKWeakestRows() {
	arguments := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	fmt.Printf("args |\t %v \n", arguments)
	results := kWeakestRows(arguments, 3)
	fmt.Printf("resp |\t %v \n", results)
	utils.PrintLine()

	arguments2 := [][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}
	fmt.Printf("args |\t %v \n", arguments2)
	results2 := kWeakestRows(arguments2, 2)
	fmt.Printf("resp |\t %v \n", results2)
	utils.PrintLine()

	arguments3 := [][]int{{1, 0, 0, 0}, {1, 0, 0, 0}}
	fmt.Printf("args |\t %v \n", arguments3)
	results3 := kWeakestRows(arguments3, 2)
	fmt.Printf("resp |\t %v \n", results3)
	utils.PrintLine()

}
