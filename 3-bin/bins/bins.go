package bins

import (
	"time"

	"github.com/google/uuid"
)

// Bin представляет собой структуру данных для хранения информации о бине
type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
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
	return []*Bin{NewBin(uuid.New().String(), true, "FirstBin"), NewBin(uuid.New().String(), false, "SecondBin")}
}
