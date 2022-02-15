package notifier

import "fmt"

type INotiManager interface {
	AddSubscriber(s ISubscriber)
	Publish(e IEvent)
}

type NotiManager struct {
	subscribers map[string][]ISubscriber
}

func NewNotiManager() *NotiManager {
	return &NotiManager{map[string][]ISubscriber{}}
}

func (nm *NotiManager) AddSubscriber(s ISubscriber) {
	nm.subscribers[s.Token()] = append(nm.subscribers[s.Token()], s)
	fmt.Println(nm.subscribers)
}

func (nm *NotiManager) Publish(e IEvent) {
	sublist, ok := nm.subscribers[e.Token()]
	if !ok {
		return
	}

	tail := len(sublist) - 1
	idx := 0
	for idx <= tail {
		sublist[idx].Handle(e)
		if sublist[idx].Type() == SubtypeOnce {
			// sublist = append(sublist[:i], sublist[i+1:]...)
			sublist[idx], sublist[tail] = sublist[tail], sublist[idx]
			tail--
		} else {
			idx++
		}
	}

	sublist = sublist[:idx]
	if len(sublist) == 0 {
		delete(nm.subscribers, e.Token())
	}
	fmt.Println(nm.subscribers)
}
