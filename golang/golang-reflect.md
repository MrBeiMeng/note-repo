## Golang 反射解析 



**反射简介**

> 反射可以在运行期间，获取对象的类型[TypeOf]，以及获取对象的值[ValueOf]



### Type 接口源码分析

> Type 接口是 go类型的一种描述，不是所有方法都能调用，要根据具体类型来判断。
>
> `go doc reflect Type `

```go
type Type interface {
  
  Align()	int // 内存字节对齐
  
  FieldAlign()	int // 结构体内存字节对齐
  
  Method(int)	Method // 结构体实现的方法，公共方法首字母小写会找不到
  
  MethodByName(string)	(Method,bool) // 根据方法名获取方法
  
  NumMethod()	int	// 获取方法的数量
  
  Name()	string // 类型名称 
  
  PkgPath()	string // 需要存储的字节数
  
  Size() uintptr // 需要存储的字节数
  
  String() string // 类型的字符串描述
  
  Kind()	Kind // Kind returns the specific kind of this type | 返回特定类型
  
  Implements(u Type) bool // 判断是否实现了某个接口
  
  AssignableTo(u Type) bool // 判断该值是否可以赋值给u
  
  ConvertibleTo(u Type) bool // 是否可以转换为u类型
  
  Bits() int // 类型的字节长度
  
  ChanDir() ChanDir // 通道类型的方向
  
  IsVariadic() bool // 是否是可变参数
  
  Elem() Type // 返回类型的元素类型
  
  Field(i int) StructField // 返回结构体的第几个字段
  
  FieldByIndex(index []int) StructField
  
  FildByName(name string)(StructField,bool)
  
  FieldByNameFunc(match func(string))
  
  In(i int) Type // 函数类型的第几个参数
  
  Key() Type // 返回 map 的key,前提是函数类型对象的参数是map
  
  Len() int // 返回数组的长度
  
  NumField() int // 返回结构体的字段数
  
  NumIn() int // 返回函数类型的输入参数数量
  
  NumOut() int // 函数类型的返回值数量
  
  Out(i int) Type // 函数类型的第几个返回值
}
```



### Value 源码分析

> go doc reflect Value



**value.go 中的函数**

```go
type abc interface{
  Append(s Value,x ...Value) Value // 添加到切片
  
  AppendSlice(s,t Value)Value // 添加一个到切片
  
  Indirect(v Value) Value // 返回v的指针值
  
  MakeChan(typ Type,buffer int) Value // 创建一个通道，返回 Value 类型
  
  MakeFunc(typ Type,fn func(args []Value)(result []Value)) Value // 创建一个函数
  
  MakeMap(typ Type) Value // 创建一个map
  
  MakeMapWithSize(typ Type,n int) Value // 创建一个map
  
  MakeSlice(typ Type,len, cap int) Value // 创建一个切片
  
  New(typ Type) Value // 创建一个类型
  
  NewAt(typ Type,p unsafe.Pointer) Value// 在某个指针地址处创建一个类型
  
  ValueOf(i interface{}) Value // 获取值类型
  
  Zero(typ Type) value // 零值描述
}
```

**测试代码** 
[golang-reflectDemo](../code_place/GoCode/reflect_demo.go)


