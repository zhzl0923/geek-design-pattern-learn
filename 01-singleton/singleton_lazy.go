package singleton

import "sync"

//懒汉式，在第一次获取实例时，初始化对象

var (
	lazySingleton *Signleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Signleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Signleton{}
		})
	}
	return lazySingleton
}
