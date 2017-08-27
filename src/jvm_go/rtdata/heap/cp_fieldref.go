package heap

import "jvm_go/fileparser"

type FieldRef struct {
	MemberRef
	field *Field
}

//字段解析
func newFieldRef(cp *ConstantPool, refInfo *fileparser.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}
