package rtda

import "jvmgo/methodinvoke/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
