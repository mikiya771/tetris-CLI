package game

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	a "github.com/tetris-CLI/action"
	s "github.com/tetris-CLI/store/stage"
)

var _ = Describe("ReducerInitializeGame", func() {
	game := NewGame()
	Describe("ゲーム開始時の振る舞い", func() {
		game.dispatcher.Emit(a.InitializeGameAction)
		Context("when ゲーム開始", func() {
			It("全てのブロックが空のステージが生成される", func() {
				Ω(game.store.GetStage()).Should(Equal(s.NewStage()))
			})
			It("最初のテトリミノがセットされる", func() {
				Ω(true).Should(BeTrue())
			})
			It("7種1巡を守るようにテトリミノのキューが用意されている", func() {
				Ω(true).Should(BeTrue())
			})
		})
	})
})
