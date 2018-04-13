package rtda

import (
	"jvmgo/class/rtda/heap"
)

type Slot struct {
	num int32
	ref *heap.Object
}

func (self *Slot) Num() int32{
	return self.num
}

func (self *Slot) Ref() *heap.Object{
	return self.ref
}