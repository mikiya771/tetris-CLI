package action

//ActionType Actionの型を示す
type ActionType int

const (
	//RotateLeftAction Tetriminoを左方向 (反時計回り) に回転するAction
	RotateLeftAction ActionType = iota + 1
	//RotateRightAction Tetriminoを右方向 (時計回り) に回転するAction
	RotateRightAction
	//MoveLeftAction Tetriminoを左方向に1ミノ分移動するAction
	MoveLeftAction
	//MoveRightAction Tetriminoを右方向に1ミノ分移動するAction
	MoveRightAction
	//SoftDropAction Tetriminoを下方向に1ミノ分移動するAction
	SoftDropAction
	//HardDropAction Tetriminoを可能な限り下方向に移動し，固定するAction
	HardDropAction
)
