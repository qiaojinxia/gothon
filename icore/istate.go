package icore

type IStack interface {
	Load_VALUE(interface{}) //往栈里丢对象
	PRINT_ANSWER()          //打印栈的数据
	ADD_TWO_VALUES()        //弹出栈顶的值相加

}
