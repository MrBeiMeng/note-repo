package type_def

// 一个 leetcode 的方法，应该包含了一些基本方法
// 生成参数 - 有几个随机参数
// 运行，通过参数去调对应的方法
// 查看结果 > 最好有一个log文件
// 获取难度
// 获取喜欢程度 star 123
// 获取状态 status doing/done
// 是否显示
// 获取题号
// 获取题目
// 获取描述
// 获取参数限制
// 获取示例

type LeetCodeProject interface {
	Run()                       // 运行 返回的str 作为log
	RunWithoutArgs()            // 运行 但不需要参数
	GetLevel() string           // 获取难度等级
	GetStar() string            // 根据程度返回图标
	GetStatus() string          // 带颜色的哦
	WillShow() bool             // 是否显示
	GetCodeNum() int            // 获取题号
	GetCodeTitle() string       // 获取题目
	GetDescription() string     // 获取描述
	GetArgsDescription() string // 获取参数限制
	GetExamples() string        // 获取示例
	GetTags() []string
	GetUrl() string

	//GetArgs() interface{}       // 生成参数
}
