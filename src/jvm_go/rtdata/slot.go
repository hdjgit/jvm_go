package rtdata

/*
局部变量表是按索引访问的
每个地方可以容纳一个int或者引用值
1）使用int，但是int存地址，当成引用时候，无法垃圾回收 P72
2）interface 可读性太差
3）结构体
 */

type Slot struct {
	num int32 //存放整数
	ref *Object //存放指针
}
