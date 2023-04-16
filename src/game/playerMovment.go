package game

import (
	"gopherLand2/src/game/entity"
)

// Each game tick, move player
func (g *Game) ComputeTick(playerId int) {
	p := g.Players[playerId]

	if p.MovesLeft && !p.MovesRight {
		g.moveHorizonally(playerId, -entity.DEFAULT_SPEED)
	}
	if !p.MovesLeft && p.MovesRight {
		g.moveHorizonally(playerId, entity.DEFAULT_SPEED)
	}

	if !p.TouchesGround {
		g.moveVertically(playerId, p.VerticalVelocity)
		g.Players[playerId].VerticalVelocity += entity.GRAVITY
	}
}

// Handle player's control
func (g *Game) RunPlayer(playerId int) {
	var action string

	for {
		action = <-g.Channel

		switch action {
		case "left":
			g.Players[playerId].MovesLeft = true
		case "right":
			g.Players[playerId].MovesRight = true
		case "up":
			if g.Players[playerId].TouchesGround {
				g.Players[playerId].TouchesGround = false
				g.Players[playerId].VerticalVelocity -= g.Players[playerId].JumpSpeed
			}
		case "released_left":
			g.Players[playerId].MovesLeft = false
		case "released_right":
			g.Players[playerId].MovesRight = false
		}
	}
}

// Moves player towards direction if possible
func (g *Game) moveVertically(playerId int, y float64) {
	p := g.Players[playerId]

	if y > 0.0 { // Player goes down
		if !g.Collide(p.Pos.X+p.EatBox[0], p.Pos.Y+p.EatBox[3]+y) &&
			!g.Collide(p.Pos.X+p.EatBox[2], p.Pos.Y+p.EatBox[3]+y) {
			g.Players[playerId].Pos.Y += y
		} else {
			g.Players[playerId].TouchesGround = true
			g.Players[playerId].VerticalVelocity = 0.0
		}
	} else if y < 0.0 { // Player goes up
		if !g.Collide(p.Pos.X+p.EatBox[0], p.Pos.Y+p.EatBox[1]+y) &&
			!g.Collide(p.Pos.X+p.EatBox[2], p.Pos.Y+p.EatBox[1]+y) {
			g.Players[playerId].Pos.Y += y
		} else {
			g.Players[playerId].VerticalVelocity = 0.0
		}
	}
}

// Moves left or right and check collisions
func (g *Game) moveHorizonally(playerId int, x float64) {
	p := g.Players[playerId]

	if x > 0.0 { // Player goes right
		if !g.Collide(p.Pos.X+p.EatBox[2]+x, p.Pos.Y+p.EatBox[1]) &&
			!g.Collide(p.Pos.X+p.EatBox[2]+x, p.Pos.Y+1.0) &&
			!g.Collide(p.Pos.X+p.EatBox[2]+x, p.Pos.Y+p.EatBox[3]) {
			g.Players[playerId].Pos.X += x
			g.checkNothingUnder(playerId)
		}
	} else if x < 0.0 { // Player goes left
		if !g.Collide(p.Pos.X+p.EatBox[0]+x, p.Pos.Y+p.EatBox[1]) &&
			!g.Collide(p.Pos.X+p.EatBox[0]+x, p.Pos.Y+1.0) &&
			!g.Collide(p.Pos.X+p.EatBox[0]+x, p.Pos.Y+p.EatBox[3]) {
			g.Players[playerId].Pos.X += x
			g.checkNothingUnder(playerId)
		}
	}
}

// Check if a specific point in the map is in collision with a solid block
func (g *Game) Collide(x, y float64) bool {
	xPos := int(x)
	yPos := int(y)
	if yPos < len(g.GameMap.Blocks) && yPos >= 0 {
		if xPos < len(g.GameMap.Blocks[yPos]) && xPos >= 0 {
			if g.Ressources.Elements[g.GameMap.Blocks[yPos][xPos]].Solid {
				return true
			}
		}
	}
	return false
}

// Checks if something a solid block is under the player, if not, players is not touching the ground anymore
func (g *Game) checkNothingUnder(playerId int) {
	p := g.Players[playerId]

	if p.TouchesGround {
		if !g.Collide(p.Pos.X+p.EatBox[0], p.Pos.Y+p.EatBox[3]+0.5) &&
			!g.Collide(p.Pos.X+p.EatBox[2], p.Pos.Y+p.EatBox[3]+0.5) {
			g.Players[playerId].TouchesGround = false
		}
	}
}
