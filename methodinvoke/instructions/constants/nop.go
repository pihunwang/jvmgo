package constants

import "jvmgo/methodinvoke/instructions/base"
import "jvmgo/methodinvoke/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
