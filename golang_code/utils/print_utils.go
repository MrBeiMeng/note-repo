package utils

import (
	"fmt"
	"goland_code/type_def"
	"strings"
)

func PrintLeetCodeProject(l type_def.LeetCodeProject) {
	str := fmt.Sprintf("\t标题： %s\t标签： %s\t链接： %s\n\t等级： %s\t评价：%s\t状态：%s\n\t描述： %s\n\t示例： %s\n\t参数： %s",
		l.GetCodeTitle(), strings.Join(l.GetTags(), "|"), l.GetUrl(), l.GetLevel(), l.GetStar(), l.GetStatus(), l.GetDescription(), l.GetExamples(), l.GetArgsDescription())

	println(str)
}

func PrintLeetCodeProjectEasy(l type_def.LeetCodeProject) {
	str := fmt.Sprintf("\t|%d\t|%s\t|%s\t|[%s]\t|%s",
		l.GetCodeNum(), l.GetCodeTitle(), l.GetLevel(), strings.Join(l.GetTags(), "、"), l.GetUrl())

	println(str)
}
