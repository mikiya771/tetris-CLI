package viewController

import el "github.com/tetris-CLI/eventListener"

//今後viewが増えたとしても，event管理の責務をここに一括する
var ViewEventManager el.EventListenr = el.MakeEventListener()
