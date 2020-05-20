package game

import (
	"github.com/nsf/termbox-go"
	a "github.com/tetris-CLI/action"
	e "github.com/tetris-CLI/emitter"
	r "github.com/tetris-CLI/reducer"
	s "github.com/tetris-CLI/store"
	v "github.com/tetris-CLI/view"
)

//Game テトリスのゲームを表現する構造体
type Game struct {
	dispatcher e.Emitter
	reducer    r.Reducer
	store      s.Store
	view       v.View
}

//NewGame Game構造体を初期化して返す
func NewGame() Game {
	dispatcher := e.NewEmitter()
	store := s.NewStore()
	reducer := r.NewReducer(&store)
	reducer.Register(dispatcher)
	view := v.NewView()
	view.Watch(&store)
	return Game{
		dispatcher: dispatcher,
		reducer:    reducer,
		store:      store,
		view:       view,
	}
}

//Run Gameが終了するまで、キー入力を受け取りを描画を行う
func (game Game) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer game.destruct()

	game.dispatcher.Emit(a.InitializeGameAction)
	game.pollKeyEvent()
}

func (game Game) destruct() {
	game.dispatcher.Off("*")
	game.store.UpdateNotifier.Off("*")
	termbox.Close()
}

func (game Game) pollKeyEvent() {
	game.dispatcher.On(a.ExitGameAction, func() {
		termbox.Interrupt()
	})

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				game.dispatcher.Emit(a.ExitGameAction)
			case termbox.KeyArrowLeft:
				game.dispatcher.Emit(a.MoveTetriminoToLeftAction)
			case termbox.KeyArrowRight:
				game.dispatcher.Emit(a.MoveTetriminoToRightAction)
			case termbox.KeyArrowDown:
				game.dispatcher.Emit(a.SoftDropTetriminoAction)
			case termbox.KeyArrowUp:
				game.dispatcher.Emit(a.RotateTetriminoToRightAction)
			case termbox.KeySpace:
				game.dispatcher.Emit(a.HardDropTetriminoAction)
			default:
			}
		case termbox.EventInterrupt:
			return
		default:
		}
	}
}
