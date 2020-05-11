package eventListener

import (
	o "github.com/GianlucaGuarini/go-observable"
)

type EventListenr struct {
	observer *o.Observable
}

func MakeEventListener() EventListenr {
	el := EventListenr{o.New()}
	return el
}

//Register actionがDispatchされた時に呼び出されるcallback関数を登録する
func (el EventListenr) Register(action string, callback interface{}) {
	el.observer.On(action, callback)
}

//UnregisterAll 登録されたactionとcallback関数を解除する
func (el EventListenr) UnregisterAll(action string, callback interface{}) {
	el.observer.Off("*")
}

//Dispatch actionに対応するcallbackを実行する
func (el EventListenr) Dispatch(action string, args ...interface{}) {
	go el.observer.Trigger(action, args...)
}
