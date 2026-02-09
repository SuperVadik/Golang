package bins

import (
	"time"
)

// Bin представляет собой структуру данных для хранения информации о бине
type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

// NewBin создает и возвращает новый бин с заданными параметрами
func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

// NewBinList создает и возвращает список бинов для демонстрации
func NewBinList() []*Bin {
	return []*Bin{NewBin("bin1", true, "FirstBin"), NewBin("bin2", false, "SecondBin")}
}
