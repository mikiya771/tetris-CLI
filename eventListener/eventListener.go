package eventListener

import (
	o "github.com/GianlucaGuarini/go-observable"
	debug "github.com/tetris-CLI/debug"
)

type EventListenr struct {
	observer *o.Observable
}

func MakeEventListener() EventListenr {
	el := EventListenr{o.New()}
	return el
}

//Register actionがDispatchされた時に呼び出されるcallback関数を登録する
func (el EventListenr) On(action string, callback interface{}) {
	el.observer.On(action, callback)
}

//UnregisterAll 登録されたactionとcallback関数を解除する
func (el EventListenr) OffAll(action string, callback interface{}) {
	el.observer.Off("*")
}

//Dispatch actionに対応するcallbackを実行する
func (el EventListenr) Emit(action string, args ...interface{}) {
	debug.AddDebugLogs(action)
	go el.observer.Trigger(action, args...)
}
