package bins

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func NewBinList() []*Bin {
	return []*Bin{NewBin("bin1", true, "FirstBin"), NewBin("bin2", false, "SecondBin")}
}
