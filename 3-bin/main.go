package main

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id string, private bool, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func newBinList() []*Bin {
	return []*Bin{newBin("bin1", true, "FirstBin"), newBin("bin2", false, "SecondBin")}
}

func main() {
	bin := newBin("12345", true, "MyFirstBin")
	println("Bin ID:", bin.id)
	println("Bin Private:", bin.private)
	println("Bin Created At:", bin.createdAt.String())
	println("Bin Name:", bin.name)
}
