package main

type Method func() error

// MethodKeys 方法类型，通过反射拿到key 值
type MethodKeys struct {
	Show Method `json:"show"`
}
