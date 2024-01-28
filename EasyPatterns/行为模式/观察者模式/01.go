package main

import "fmt"

/**
https://developer.aliyun.com/article/1267681

观察者模式概述：
观察者模式属于行为型设计模式，它用于在对象之间建立一对多的依赖关系。
当一个对象的状态发生改变时，它的所有依赖者都会收到通知并自动更新。观察者模式主要包含以下角色：
1. 主题（Subject）：负责维护一组观察者，并在状态改变时通知观察者。
2. 观察者（Observer）：定义了接收通知和更新的方法。
*/

// 观察者接口
type Observer interface {
	Update(subject Subject)
}

// 主题接口
type Subject interface {
	Register(observer Observer) // 注册
	Unregister(observer Observer)
	Notify() // 通知
}

// 具体主题
type NewsSubject struct {
	observers []Observer // 订阅者
	news      string
}

func (n *NewsSubject) Register(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NewsSubject) Unregister(observer Observer) {
	for i, o := range n.observers {
		if o == observer {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NewsSubject) Notify() {
	for _, observer := range n.observers {
		observer.Update(n)
	}
}

func (n *NewsSubject) SetNews(news string) {
	n.news = news
	n.Notify()
}

// 具体观察者
type Subscriber struct {
	name string
}

func (s *Subscriber) Update(subject Subject) {
	newsSubject := subject.(*NewsSubject)
	fmt.Printf("[%s] 收到新闻通知：%s\n", s.name, newsSubject.news)
}

func main() {
	subject := new(NewsSubject)

	subscriber1 := &Subscriber{
		name: "订阅者1",
	}

	subject.Register(subscriber1)

	subscriber2 := &Subscriber{
		name: "订阅者2",
	}

	subject.Register(subscriber2)

	subject.SetNews("新闻1发布了")
	subject.SetNews("新闻2发布了")

	//subject.Unregister(subscriber1)
	//subject.SetNews("新闻3发布了")
}
