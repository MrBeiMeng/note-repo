package main

import (
	"bufio"
	"errors"
	"fmt"
	"goland_code/new_user_list"
	"goland_code/type_def"
	"os"
	"strconv"
	"strings"
)

func main() {
	runFlag := true
	keyProjMap := initMap()

	println("please input something.")
	for runFlag {
		inputStr, err := getInput("")
		if err != nil {
			print(err.Error())
			runFlag = false
		}

		analyseInput(inputStr, keyProjMap)
	}
}

func analyseInput(inputStr string, tmap map[int]type_def.LProjectImpl) {
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
		proj.Print()
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
		proj.Print()
		return
	}

	num, _ := strconv.Atoi(inputStr)
	if proj, ok := tmap[num]; ok {
		proj.GetInitFunc()()
	}
}

func printProjList(keyProjMap map[int]type_def.LProjectImpl) {
	count := 0
	fmt.Print("\t")
	for _, proj := range keyProjMap {
		printProjEasy(proj)
		count++
		if count%10 == 0 {
			count = 0
			fmt.Println()
			fmt.Print("\t")
		}
	}
}

func printProjEasy(proj type_def.LProjectImpl) {
	switch proj.GetLLevel() {
	case 0:
		{
			fmt.Printf("\033[1;32m%d %s\u001B[0m\t", proj.GetCodeNum(), proj.GetLName())
		}
		break
	case 1:
		{
			fmt.Printf("\033[1;34m%d %s\u001B[0m\t", proj.GetCodeNum(), proj.GetLName())
		}
		break
	case 2:
		{
			fmt.Printf("\033[1;35m%d %s\u001B[0m\t", proj.GetCodeNum(), proj.GetLName())
		}
		break
	}
}

func initMap() map[int]type_def.LProjectImpl {
	keyProjMap := make(map[int]type_def.LProjectImpl)

	// 注册方法
	new_user_list.InitPalindromeLinkedListFunc(keyProjMap)
	new_user_list.InitRansomNoteFunc(keyProjMap)
	new_user_list.InitRomanToIntFunc(keyProjMap)
	new_user_list.InitWeakestRowInMatrixFunc(keyProjMap)

	return keyProjMap
}

func getInput(message string) (string, error) {
	println(message)
	print("bl >> ")
	inputStr := ""
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	inputStr = string(line)

	if strings.EqualFold(inputStr, "quit") || strings.EqualFold(inputStr, "stop") {
		return "", errors.New("用户结束进程")
	}

	return inputStr, nil
}
