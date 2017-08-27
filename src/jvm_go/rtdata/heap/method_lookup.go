package heap

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	//找类里面的方法，找不到找父类的方法
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	//找接口里的方法
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}

		//找不到找接口的接口的方法
		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil
}
