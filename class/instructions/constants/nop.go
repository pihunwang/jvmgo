package constants

import (
	"jvmgo/class/instructions/base"
	"jvmgo/class/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
