package references

import (
	"jvmgo/class/instructions/base"
	"jvmgo/class/rtda"
	"jvmgo/class/rtda/heap"
)

type CHECK_CAST struct {
	base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	// null可以转换成任何类型
	if ref == nil{
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class){
		panic("java.lang.ClassCastException")
	}
}
