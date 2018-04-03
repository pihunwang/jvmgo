package rtda

type Frame struct {
	lower        *Frame	// 链表模拟栈帧
	localVars    LocalVars	// 保存局部变量指针
	operandStack *OperandStack	// 保存操作数栈指针
}

func newFrame(maxLocals, maxStack uint) *Frame {

}
