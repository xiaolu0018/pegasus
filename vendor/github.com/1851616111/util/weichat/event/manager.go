package event

import (
	"errors"
	"time"
	"fmt"
)

var ErrEventNotFound error = errors.New("event not found")
var ErrEventExist error = errors.New("event already exist")
var ErrEventActionNotFount error = errors.New("event action not found")
var ErrEventCBNotFount error = errors.New("event action not found")
func NewEventManager() *EventManager {
	return &EventManager{
		eventActionNum:    0,
		eventToActionM:    make(map[string]*Action),
		eventCallBackNum : 0,
		eventToCallBackM:  make(map[string]func(*Event) error),
	}
}

//reference https://mp.weixin.qq.com/wiki/2/5baf56ce4947d35003b86a9805634b1e.html
type EventManager struct {
	eventActionNum  int
	eventToActionM  map[string]*Action

	eventCallBackNum int
	eventToCallBackM map[string] func(*Event) error
}

func (m *EventManager) RegistAction(tp string, act *Action) error {
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
	m.eventActionNum++

	return nil
}

func (m *EventManager) RegistEventCallBack(tp string, cb func(*Event) error ) error {
	if tp != E_News && tp != E_Subscribe && tp != E_UnSubscribe {
		return ErrEventNotFound
	}

	if cb == nil {
		return ErrEventCBNotFount
	}

	if _, exist := m.eventToCallBackM[tp]; exist {
		return ErrEventExist
	}

	m.eventToCallBackM[tp] = cb
	m.eventCallBackNum ++
	return nil
}

func (m *EventManager) Handle(e *Event) *Action {
	if m.eventActionNum == 0 {
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

func (m *EventManager) CallBack(e *Event) error {
	fmt.Println("------------------------01")
	if m.eventCallBackNum == 0 {
		return nil
	}
	fmt.Println("------------------------02")
	cb, exist := m.eventToCallBackM[string(e.E)]
	if !exist {
		return nil
	}
	fmt.Println("------------------------03")
	return cb(e)
}