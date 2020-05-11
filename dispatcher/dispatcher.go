package dispatcher

import (
	el "github.com/tetris-CLI/eventListener"
)

var Dispatcher el.EventListenr = el.MakeEventListener()
