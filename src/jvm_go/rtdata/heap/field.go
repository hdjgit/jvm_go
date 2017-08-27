package heap

import "jvm_go/fileparser"

type Field struct {
	ClassMember
	slotId uint
}

func newFields(class *Class, cfFields []*fileparser.MemberInfo) []*Field {
	fileds := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fileds[i] = &Field{}
		fileds[i].class = class
		fileds[i].copyMemberInfo(cfField)
	}
	return fileds
}
