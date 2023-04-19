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
	Elements        map[string]Element `json:"blocks"`
	RessourcePack   *ebiten.Image
	BackgroundImage *ebiten.Image
}

func New(size int) Ressources {
	r := Ressources{
		Elements:        map[string]Element{},
		RessourcePack:   loadImage("data/ressources/defaultRessourcePack.png"),
		BackgroundImage: loadImage("data/ressources/defaultBackground.png"),
	}

	r.loadElements(size)

	return r
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Println(err)
	}
	return img
}

func (r *Ressources) loadElements(size int) {
	data, err := os.ReadFile("data/tables/ressources.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &r.Elements)
	if err != nil {
		panic(err)
	}

	for k, elem := range r.Elements {
		elem.Img = r.RessourcePack.SubImage(image.Rect(
			size*elem.X,
			size*elem.Y,
			size*elem.X+size*elem.W,
			size*elem.Y+size*elem.H)).(*ebiten.Image)

		r.Elements[k] = elem
	}
}
