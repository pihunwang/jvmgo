package control

import (
	"jvmgo/class/instructions/base"
	"jvmgo/class/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}