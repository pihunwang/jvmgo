package comparisons

import (
	"jvmgo/order/instructions/base"
	"jvmgo/order/rtda"
)

// x == 0
type IFEQ struct {
	base.BranchInstruction
}

// x != 0
type IFNE struct {
	base.BranchInstruction
}

// x < 0
type IFLT struct {
	base.BranchInstruction
}

// x <= 0
type IFLE struct {
	base.BranchInstruction
}

// x > 0
type IFGT struct {
	base.BranchInstruction
}

// x >= 0
type IFGE struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}