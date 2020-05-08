package dispatcher

import (
	o "github.com/GianlucaGuarini/go-observable"
)

var observable *o.Observable = o.New()

//Register actionがDispatchされた時に呼び出されるcallback関数を登録する
func Register(action string, callback interface{}) {
	observable.On(action, callback)
}

//UnregisterAll 登録されたactionとcallback関数を解除する
func UnregisterAll(action string, callback interface{}) {
	observable.On(action, callback)
}

//Dispatch actionに対応するcallbackを実行する
func Dispatch(action string, args ...interface{}) {
	go observable.Trigger(action, args...)
}
