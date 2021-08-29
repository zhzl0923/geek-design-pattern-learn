package singleton

//饿汉式，类在加载时，已经初始化好了
type Signleton struct {
}

var singleton *Signleton

func init() {
	singleton = &Signleton{}
}

func GetInstance() *Signleton {
	return singleton
}
