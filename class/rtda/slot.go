package rtda

import (
	"jvmgo/class/rtda/heap"
)

type Slot struct {
	num int32
	ref *heap.Object
}
