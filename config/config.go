package config

import "time"

//InvisibleStageHeight 判定用の描画されないStageの縦の長さ
const InvisibleStageHeight int = 2

//VisibleStageHeight 描画するStageの縦の長さ
const VisibleStageHeight int = 20

//StageHeight Stageの縦の長さ
const StageHeight int = InvisibleStageHeight + VisibleStageHeight

//StageWidth Stageの横の長さ
const StageWidth int = 10

//AutoDropIntervalTime 不操作時の自動落下の間隔
const AutoDropIntervalTime time.Duration = 1 * time.Second

//Debug デバッグ用ログを表示するかどうか
const Debug bool = true
