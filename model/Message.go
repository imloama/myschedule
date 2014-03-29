package model

/**
 * 返回前台的信息，最终要生成Json格式的字符串
 */
type Message struct {
	Code int
	Msg  string
	Data interface{}
}
