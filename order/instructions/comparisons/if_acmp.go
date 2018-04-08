package comparisons

import (
	"jvmgo/order/instructions/base"
	"jvmgo/order/rtda"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}