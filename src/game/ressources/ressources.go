package ressources

import (
	"encoding/json"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ressources struct {
	Size          int              `json:"size"`
	Blocks        map[string]Block `json:"blocks"`
	RessourcePack *ebiten.Image
}

type Block struct {
	Name string `json:"name"`
	Code int    `json:"code"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Img  *ebiten.Image
}

func New() Ressources {
	r := Ressources{
		Size:          32,
		Blocks:        map[string]Block{},
		RessourcePack: loadRessourcePack(),
	}

	r.loadRessources()

	return r
}

func loadRessourcePack() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("data/ressources/defaultRessourcePack.png")
	if err != nil {
		log.Println(err)
	}
	return img
}

func (r *Ressources) loadRessources() {
	data, err := os.ReadFile("data/tables/ressources.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &r.Blocks)
	if err != nil {
		panic(err)
	}

	for k := range r.Blocks {
		block, ok := r.Blocks[k]
		if ok {
			block.Img = r.RessourcePack.SubImage(image.Rect(
				r.Blocks[k].X*r.Size,
				r.Blocks[k].Y*r.Size,
				r.Blocks[k].X*r.Size+r.Size,
				r.Blocks[k].X*r.Size+r.Size)).(*ebiten.Image)
			r.Blocks[k] = block
		}
	}
}
