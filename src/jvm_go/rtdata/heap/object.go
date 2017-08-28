package heap

//方法区，运行时数据区的一块逻辑区域，由多个线程共享
//方法区主要存放从class文件中获取的类的信息
//类变量也存放在方法区中

type Object struct {
	class *Class
	data  interface{} //为了支持数组，把fields换成data
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
	return class.isAssignableFrom(self.class)
}
