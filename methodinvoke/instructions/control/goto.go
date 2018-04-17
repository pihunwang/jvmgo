package control

import "jvmgo/methodinvoke/instructions/base"
import "jvmgo/methodinvoke/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
