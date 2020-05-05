package action

const (
	//InitializeGameAction ゲームを初期化するAction
	InitializeGameAction = "Initialize_game"
	//SetNewTetriminoAction 新しいTetriminoをセットするAction
	SetNewTetriminoAction = "Set_new_tetrimino"
	//RotateTetriminoToLeftAction Tetriminoを左方向 (反時計回り) に回転するAction
	RotateTetriminoToLeftAction = "Rotate_tetrimino_to_left"
	//RotateTetriminoToRightAction Tetriminoを右方向 (時計回り) に回転するAction
	RotateTetriminoToRightAction = "Rotate_tetrimino_to_right"
	//MoveTetriminoToLeftAction Tetriminoを左方向に1移動するAction
	MoveTetriminoToLeftAction = "Move_tetrimino_to_left"
	//MoveTetriminoToRightAction Tetriminoを右方向に1移動するAction
	MoveTetriminoToRightAction = "Move_tetrimino_to_right"
	//SoftDropTetriminoAction Tetriminoを下方向に1移動するAction
	SoftDropTetriminoAction = "Soft_drop_tetrimino"
	//HardDropTetriminoAction Tetriminoを可能な限り下方向に移動し，固定するAction
	HardDropTetriminoAction = "Hard_drop_tetrimino"
	//UpdateTetriminoAction TetriminoをStageに固定するAction
	UpdateTetriminoAction = "Update_tetrimino"
	//FixTetriminoToStageAction TetriminoをStageに固定するAction
	FixTetriminoToStageAction = "Fix_tetrimino_to_stage"
	//RefreshStageAction StageのLineを確信し、埋まっているなら削除するAction
	RefreshStageAction = "Refresh_stage"
)
