package new_user_list

func kWeakestRows(mat [][]int, k int) []int {
	tmpMat := make([]int, 0)

	for _, nums := range mat {
		tmpNum := 0
		for _, num := range nums {
			tmpNum += num
		}
		tmpMat = append(tmpMat, tmpNum)
	}

	// 排序
	selectSort(tmpMat, false)

	return tmpMat[:k+1]
}

func selectSort(nums []int, desc bool) {
	for i := 0; i < len(nums); i++ {
		selectedIndex := i
		for j := i + 1; j < len(nums); j++ {
			if desc {
				if nums[selectedIndex] <= nums[j] {
					selectedIndex = j
				}
			} else {
				if nums[selectedIndex] >= nums[j] {
					selectedIndex = j
				}
			}
		}

		nums[i], nums[selectedIndex] = nums[selectedIndex], nums[i]
	}
}

//func RunCanConstruct() {
//	arguments := []string{"abc", "aabbcc"}
//	fmt.Printf("args |\t %v \n", arguments)
//	results := canConstruct(arguments[0], arguments[1])
//	fmt.Printf("resp |\t %v \n", results)
//	utils.PrintLine()
//}
