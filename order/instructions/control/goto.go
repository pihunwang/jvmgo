package control

import (
	"jvmgo/order/instructions/base"
	"jvmgo/order/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}