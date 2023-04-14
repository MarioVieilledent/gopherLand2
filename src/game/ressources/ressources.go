package ressources

import (
	"encoding/json"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Ressources are the description of all single ressources
type Ressources struct {
	Blocks        map[string]Block `json:"blocks"`
	RessourcePack *ebiten.Image
}

func New(size int) Ressources {
	r := Ressources{
		Blocks:        map[string]Block{},
		RessourcePack: loadRessourcePack(),
	}

	r.loadRessources(size)

	return r
}

func loadRessourcePack() *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("data/ressources/defaultRessourcePack.png")
	if err != nil {
		log.Println(err)
	}
	return img
}

func (r *Ressources) loadRessources(size int) {
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
				r.Blocks[k].X*size,
				r.Blocks[k].Y*size,
				r.Blocks[k].X*size+size,
				r.Blocks[k].X*size+size)).(*ebiten.Image)
			r.Blocks[k] = block
		}
	}
}
