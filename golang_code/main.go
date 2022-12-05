package main

import (
	"fmt"
	"goland_code/leetcode_code_place/new_gay_list"
	"goland_code/type_def"
	"goland_code/utils"
	"strconv"
	"strings"
	"time"
)

// 应该是反射获取整个文件夹下的所有程序，放到集合中，再去搞
// 一个题目的 状态有 doing\done star:1,2,3 难度 是否显示

func main() {
	runFlag := true
	keyProjMap := initMap()

	println("please input something.")
	for runFlag {
		inputStr, err := utils.GetInput("", 0)
		if err != nil {
			print(err.Error())
			runFlag = false
		}

		analyseInput(inputStr, keyProjMap)
	}
}

func analyseInput(inputStr string, tmap map[int]type_def.LeetCodeProject) {
	if strings.EqualFold(inputStr, "print") || strings.EqualFold(inputStr, "show") {
		printProjList(tmap)
		return
	}

	if strings.HasPrefix(inputStr, "print ") {
		resultList := strings.Split(inputStr, "print ")
		if len(resultList) < 2 {
			println("错误的输入")
			return
		}
		codeNum, _ := strconv.Atoi(resultList[1])
		proj := tmap[codeNum]
		utils.PrintLeetCodeProject(proj)
		return
	}
	if strings.HasPrefix(inputStr, "show ") {
		resultList := strings.Split(inputStr, "show ")
		if len(resultList) < 2 {
			println("错误的输入")
			return
		}
		codeNum, _ := strconv.Atoi(resultList[1])
		proj := tmap[codeNum]
		utils.PrintLeetCodeProject(proj)
		return
	}

	if strings.HasPrefix(inputStr, "run ") {
		resultList := strings.Split(inputStr, "run ")
		if len(resultList) < 2 {
			println("错误的输入")
			return
		}
		num, _ := strconv.Atoi(resultList[1])
		if proj, ok := tmap[num]; ok {
			//proj.GetInitFunc()()
			proj.RunWithoutArgs()
		}
		return
	}

	if strings.HasPrefix(inputStr, "test ") {
		resultList := strings.Split(inputStr, "test ")
		if len(resultList) < 2 {
			println("错误的输入")
			return
		}
		num, _ := strconv.Atoi(resultList[1])
		if proj, ok := tmap[num]; ok {
			preTime := time.Now()
			proj.Run()
			fmt.Printf("\t方法用时 [%s]\n", time.Now().Sub(preTime))
		}
		return
	}

	num, _ := strconv.Atoi(inputStr)
	if proj, ok := tmap[num]; ok {
		//proj.GetInitFunc()()
		proj.RunWithoutArgs()
	}
}

func printProjList(keyProjMap map[int]type_def.LeetCodeProject) {
	count := 0
	fmt.Print("\t")
	for _, proj := range keyProjMap {
		utils.PrintLeetCodeProjectEasy(proj)
		//utils.PrintLine()
		count++
		if count%10 == 0 {
			count = 0
			fmt.Println()
			fmt.Print("\t")
		}
	}
}

func initMap() map[int]type_def.LeetCodeProject {
	keyProjMap := make(map[int]type_def.LeetCodeProject)

	var pl new_gay_list.PalindromeLinkedList
	keyProjMap[pl.GetCodeNum()] = pl

	return keyProjMap
}
