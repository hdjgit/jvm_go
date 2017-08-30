package heap

//方法区，运行时数据区的一块逻辑区域，由多个线程共享
//方法区主要存放从class文件中获取的类的信息
//类变量也存放在方法区中

type Object struct {
	class *Class
	data  interface{} //为了支持数组，把fields换成data
	//记录Object结构体实例的额外信息
	extra interface{}
}

//如何知道静态变量和实例变量需要多少空间，以及哪个字段对应Slots中的哪个位置呢？

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(self.class)
}
func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}
func (self *Object) Data() interface{} {
	return self.data
}
func (self *Object) GetIntVar(name, descriptor string) int32 {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetInt(field.slotId)
}