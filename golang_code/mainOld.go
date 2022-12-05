package main

//
//import (
//	"fmt"
//	new_user_list2 "goland_code/leetcode_code_place/new_user_list"
//	"goland_code/type_def"
//	"goland_code/utils"
//	"strconv"
//	"strings"
//)
//
//// 应该是反射获取整个文件夹下的所有程序，放到集合中，再去搞
//// 一个题目的 状态有 doing\done star:1,2,3 难度 是否显示
//
//func main() {
//	runFlag := true
//	keyProjMap := initMap()
//
//	println("please input something.")
//	for runFlag {
//		inputStr, err := utils.GetInput("")
//		if err != nil {
//			print(err.Error())
//			runFlag = false
//		}
//
//		analyseInput(inputStr, keyProjMap)
//	}
//}
//
//func analyseInput(inputStr string, tmap map[int]type_def.LProjectImpl) {
//	if strings.EqualFold(inputStr, "print") || strings.EqualFold(inputStr, "show") {
//		printProjList(tmap)
//		return
//	}
//
//	if strings.HasPrefix(inputStr, "print ") {
//		resultList := strings.Split(inputStr, "print ")
//		if len(resultList) < 2 {
//			println("错误的输入")
//			return
//		}
//		codeNum, _ := strconv.Atoi(resultList[1])
//		proj := tmap[codeNum]
//		proj.Print()
//		return
//	}
//	if strings.HasPrefix(inputStr, "show ") {
//		resultList := strings.Split(inputStr, "show ")
//		if len(resultList) < 2 {
//			println("错误的输入")
//			return
//		}
//		codeNum, _ := strconv.Atoi(resultList[1])
//		proj := tmap[codeNum]
//		proj.Print()
//		return
//	}
//
//	num, _ := strconv.Atoi(inputStr)
//	if proj, ok := tmap[num]; ok {
//		proj.GetInitFunc()()
//	}
//}
//
//func printProjList(keyProjMap map[int]type_def.LProjectImpl) {
//	count := 0
//	fmt.Print("\t")
//	for _, proj := range keyProjMap {
//		printProjEasy(proj)
//		count++
//		if count%10 == 0 {
//			count = 0
//			fmt.Println()
//			fmt.Print("\t")
//		}
//	}
//}
//
//func printProjEasy(proj type_def.LProjectImpl) {
//	switch proj.GetLLevel() {
//	case 0:
//		{
//			fmt.Printf("\033[1;32m%d %s\u001B[0m\t", proj.GetCodeNum(), proj.GetLName())
//		}
//		break
//	case 1:
//		{
//			fmt.Printf("\033[1;34m%d %s\u001B[0m\t", proj.GetCodeNum(), proj.GetLName())
//		}
//		break
//	case 2:
//		{
//			fmt.Printf("\033[1;35m%d %s\u001B[0m\t", proj.GetCodeNum(), proj.GetLName())
//		}
//		break
//	}
//}
//
//func initMap() map[int]type_def.LProjectImpl {
//	keyProjMap := make(map[int]type_def.LProjectImpl)
//
//	// 注册方法
//	new_user_list2.InitPalindromeLinkedListFunc(keyProjMap)
//	new_user_list2.InitRansomNoteFunc(keyProjMap)
//	new_user_list2.InitRomanToIntFunc(keyProjMap)
//	new_user_list2.InitWeakestRowInMatrixFunc(keyProjMap)
//
//	return keyProjMap
//}
