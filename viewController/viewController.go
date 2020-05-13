package viewController

import o "github.com/GianlucaGuarini/go-observable"

//今後viewが増えたとしても，event管理の責務をここに一括する
var ViewEventManager *o.Observable = o.New()
