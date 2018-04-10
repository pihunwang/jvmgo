package heap

import (
	"jvmgo/class/classfile"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *classfile.ConstantPool
	fields            []*Field  // 字段表
	methods           []*Method // 方法表
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint      // 实例变量占据的空间大小
	staticSlotCount   uint      // 类变量占据的空间大小
	staticVars        *Slots    // 静态变量
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}