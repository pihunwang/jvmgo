package math

import (
	"jvmgo/order/instructions/base"
	"jvmgo/order/rtda"
	"math"
)

// int左位移
type ISHL struct {
	base.NoOperandsInstruction
}

// int算术右位移(有符号右移)
type ISHR struct {
	base.NoOperandsInstruction
}

// int逻辑右位移(无符号右移)
type IUSHR struct {
	base.NoOperandsInstruction
}

// long左位移
type LSHL struct {
	base.NoOperandsInstruction
}

// long算术右位移
type LSHR struct {
	base.NoOperandsInstruction
}

// long逻辑右位移
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushInt(result)
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushLong(result)
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x1f
	result := int64(uint64(v1) >> s)
	stack.PushInt(result)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}