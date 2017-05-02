package main

import (
	"log"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

type myScene struct{}

// Gopher gopherくんのstruct定義
type Gopher struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// Type uniquely defines your game type
func (*myScene) Type() string { return "myGame" }

// Preload is called before loading any assets from the disk,
// to allow you to register / queue them
func (*myScene) Preload() {
	engo.Files.Load("textures/gopher.png")
}

// Setup メインループが開始する前に実行される関数
func (*myScene) Setup(world *ecs.World) {

	// worldに対してRenderSystemを追加
	world.AddSystem(&common.RenderSystem{})

	// gopherくんの初期化

	// gopherくんのSpaceComponentの初期化。位置と大きさを設定する。
	gopher := Gopher{BasicEntity: ecs.NewBasic()}
	gopher.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{10, 10},
		Width:    303,
		Height:   641,
	}

	// gopherくんのRenderComponentにpreLoadしていた画像を設定する
	texture, err := common.LoadedSprite("textures/gopher.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	gopher.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}

	// WorldのRenderSystemにgopherを登録
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&gopher.BasicEntity, &gopher.RenderComponent, &gopher.SpaceComponent)
		}
	}

}

func main() {
	opts := engo.RunOptions{
		Title:  "Go lang のゲーム",
		Width:  400,
		Height: 400,
	}
	engo.Run(opts, &myScene{})
}
