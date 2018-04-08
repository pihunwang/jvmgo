package constants

import (
	"jvmgo/order/instructions/base"
	"jvmgo/order/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
