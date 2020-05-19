package emitter

import (
	o "github.com/GianlucaGuarini/go-observable"
	debug "github.com/tetris-CLI/debug"
)

//Emitter messageに対応したhandlerを呼び出す
type Emitter struct {
	observable *o.Observable
}

//NewEmitter Emitterをを初期化して返す
func NewEmitter() Emitter {
	return Emitter{
		observable: o.New(),
	}
}

//On messageがEmitされた時に呼び出されるhandlerを登録する
func (emitter Emitter) On(message string, handler interface{}) {
	emitter.observable.On(message, handler)
}

//Off messageに対応するhandlerの登録を解除する
func (emitter Emitter) Off(message string) {
	emitter.observable.Off(message)
}

//Emit messageに対応するhandlerを呼び出す
func (emitter Emitter) Emit(message string) {
	debug.AddDebugLogs(message)
	go emitter.observable.Trigger(message)
}
