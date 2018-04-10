package heap

import (
	"jvmgo/class/classfile"
)

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {

}
