package event

import (
	"errors"
	"fmt"
	"time"
)

var ErrEventNotFound error = errors.New("event not found")
var ErrEventExist error = errors.New("event already exist")
var ErrEventActionNotFount error = errors.New("event action not found")

func NewEventManager() *EventManager {
	return &EventManager{
		eventNum:       0,
		eventToActionM: make(map[string]*Action),
	}
}

//reference https://mp.weixin.qq.com/wiki/2/5baf56ce4947d35003b86a9805634b1e.html
type EventManager struct {
	eventNum       int
	eventToActionM map[string]*Action
}

func (m *EventManager) Registe(tp string, act *Action) error {
	if tp != E_News && tp != E_Subscribe && tp != E_UnSubscribe {
		return ErrEventNotFound
	}

	if act == nil {
		return ErrEventActionNotFount
	}

	if _, exist := m.eventToActionM[tp]; exist {
		return ErrEventExist
	}

	m.eventToActionM[tp] = act
	m.eventNum++

	return nil
}

func (m *EventManager) Handle(e *Event) *Action {
	if m.eventNum == 0 {
		return nil
	}

	act, exist := m.eventToActionM[string(e.E)]
	if !exist {
		return nil
	}

	retAct := Action{
		Common:       act.Common,
		ArticleCount: act.ArticleCount,
		Items:        act.Items,
	}

	retAct.CreateTime = time.Now().Unix()
	retAct.From, act.To = e.To, e.From

	return &retAct
}
