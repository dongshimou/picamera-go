package model

// 订阅
type Subscriber interface {

}

// 观察
type Observer interface {
	Start()error
	Stop()error
	Shoot()error
	AddSub(Subscriber)
}

