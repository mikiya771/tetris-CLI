package debug

import "github.com/tetris-CLI/config"

var debugLogs []string

//GetDebugLogs デバッグ用ログのSliceを返す
func GetDebugLogs() []string {
	return debugLogs
}

//AddDebugLogs デバッグ用のログを追加する
func AddDebugLogs(log string) {
	if len(debugLogs) < config.StageHeight {
		debugLogs = append(debugLogs, log)
	} else {
		debugLogs = append(debugLogs, log)[len(debugLogs)-config.StageHeight+1:]
	}
}
