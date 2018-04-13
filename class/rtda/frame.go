package rtda

import (
	"jvmgo/class/rtda/heap"
	"fmt"
)

type Frame struct {
	lower        *Frame        // 链表模拟栈帧
	localVars    LocalVars     // 保存局部变量指针
	operandStack *OperandStack // 保存操作数栈指针
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread: thread,
		method :method,
		localVars: newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	fmt.Printf("OperandStack , top = %v \n",self.operandStack.size)
	for _,slot := range self.operandStack.slots{
		fmt.Printf("slot = %v \n",slot)
	}
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}