package cell

//Cell Lineを構成する正方形
type Cell struct {
	//Cellが埋まっているか
	IsFilled bool
}

//NewCell Cellインスタンスを初期化して返す
func NewCell(isFilled bool) Cell {
	return Cell{
		IsFilled: isFilled,
	}
}
