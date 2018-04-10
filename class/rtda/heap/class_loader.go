package heap

import (
	"fmt"
	"jvmgo/class/classfile"
	"jvmgo/class/classpath"
)

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class //loaded Classes
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {

}

func (self *ClassLoader) LoadClass(name string) *Class {

}