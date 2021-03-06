package action

//action for store event
const (
	//InitializeGameAction ゲームを初期化するAction
	InitializeGameAction = "Initialize_game"
	//StartTimerAction 不操作時の自動落下用のTimerを開始する
	StartTimerAction = "Start_timer"
	//StopTimerAction 不操作時の自動落下用のTimerを停止する
	StopTimerAction = "Stop_timer"
	//ResetTimerAction 不操作時の自動落下用のTimerをリセットする
	ResetTimerAction = "Reset_timer"
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
	//FixTetriminoToStageAction TetriminoをStageに固定するAction
	FixTetriminoToStageAction = "Fix_tetrimino_to_stage"
	//RefreshStageAction StageのLineを確認し、埋まっているなら削除するAction
	RefreshStageAction = "Refresh_stage"
	//ExitGameAction ゲームを終了するAction
	ExitGameAction = "Exit_game"
)

// action for view event
const (
	//UpdateViewAction Storeを元にゲーム画面を更新するAction
	UpdateViewAction = "Update_view"
)
