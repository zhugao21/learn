package designpatterns

import "sync"

// Singleton 单例模式：一个类只允许创建一个实例，（数据只需要保存一份）
type Singleton struct {
}

// 饿汉模式：使用之前已经初始化，使用时直接获取
var singleton Singleton

func init() {
	singleton = Singleton{}
}

func GetInstance() Singleton {
	return singleton
}

// 懒汉模式: 第一次使用时，才进行初始化
var lazySingleton Singleton
var once sync.Once

func GetLazyInstance() Singleton {
	once.Do(func() {
		lazySingleton = Singleton{}
	})
	return lazySingleton
}
