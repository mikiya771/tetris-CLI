package mino

//Mino Tetriminoを構成する正方形
type Mino struct {
	//水平方向の座標値
	X int
	//重力方向の座標値
	Y int
}

//NewMino Minoインスタンスを初期化して返す
func NewMino(x, y int) Mino {
	return Mino{
		X: x,
		Y: y,
	}
}

//Clone Minoインスタンスを複製して返す
func (mino *Mino) Clone() Mino {
	return Mino{
		X: mino.X,
		Y: mino.Y,
	}
}
