package game

import (
	"gopherLand2/src/game/entity"
	"gopherLand2/src/localInstance/input"
	"math"
)

// Each game tick, move player
func (g *Game) ComputeTick() {
	if g.Player.MovesLeft && !g.Player.MovesRight {
		g.moveHorizonally(-entity.DEFAULT_SPEED)
	}
	if !g.Player.MovesLeft && g.Player.MovesRight {
		g.moveHorizonally(entity.DEFAULT_SPEED)
	}
	if g.Player.IsJumping {
		if g.Player.TouchesGround {
			g.Player.TouchesGround = false
			g.Player.VerticalVelocity -= g.Player.JumpSpeed
		}
	}

	if !g.Player.TouchesGround {
		g.moveVertically(g.Player.VerticalVelocity)
		g.Player.VerticalVelocity += entity.GRAVITY
	}
}

// Handle player's control
func (g *Game) RunPlayer() {
	var keyPressed input.KeyPressed

	for {
		keyPressed = <-g.PlayerInputChannel

		if keyPressed.Left {
			g.Player.MovesLeft = true
		} else {
			g.Player.MovesLeft = false
		}
		if keyPressed.Right {
			g.Player.MovesRight = true
		} else {
			g.Player.MovesRight = false
		}
		if keyPressed.Up {
			g.Player.IsJumping = true
		} else {
			g.Player.IsJumping = false
		}
	}
}

// Moves player towards direction if possible
func (g *Game) moveVertically(y float64) {
	if y > 0.0 { // Player goes down
		if !g.Collide(g.Player.Pos.X+g.Player.EatBox[0], g.Player.Pos.Y+g.Player.EatBox[3]+y) &&
			!g.Collide(g.Player.Pos.X+g.Player.EatBox[2], g.Player.Pos.Y+g.Player.EatBox[3]+y) {
			g.Player.Pos.Y += y
		} else {
			// Move down as much as possible
			g.Player.Pos.Y = math.Round(g.Player.Pos.Y)
			// Land in ground
			g.Player.TouchesGround = true
			g.Player.VerticalVelocity = 0.0
		}
	} else if y < 0.0 { // Player goes up
		if !g.Collide(g.Player.Pos.X+g.Player.EatBox[0], g.Player.Pos.Y+g.Player.EatBox[1]+y) &&
			!g.Collide(g.Player.Pos.X+g.Player.EatBox[2], g.Player.Pos.Y+g.Player.EatBox[1]+y) {
			g.Player.Pos.Y += y
		} else {
			// Hit ceiling
			g.Player.VerticalVelocity = 0.0
		}
	}

	if g.PlayerPosChannel != nil {
		g.PlayerPosChannel <- entity.PlayerInfo{
			Nickname: g.Player.Nickname,
			Pos:      g.Player.Pos,
		}
	}
}

// Moves left or right and check collisions
func (g *Game) moveHorizonally(x float64) {
	if x > 0.0 { // Player goes right
		if !g.Collide(g.Player.Pos.X+g.Player.EatBox[2]+x, g.Player.Pos.Y+g.Player.EatBox[1]) &&
			!g.Collide(g.Player.Pos.X+g.Player.EatBox[2]+x, g.Player.Pos.Y+1.0) &&
			!g.Collide(g.Player.Pos.X+g.Player.EatBox[2]+x, g.Player.Pos.Y+g.Player.EatBox[3]) {
			// Move right and check if floor under
			g.Player.Pos.X += x
			g.checkTouchesGround()
		}
	} else if x < 0.0 { // Player goes left
		if !g.Collide(g.Player.Pos.X+g.Player.EatBox[0]+x, g.Player.Pos.Y+g.Player.EatBox[1]) &&
			!g.Collide(g.Player.Pos.X+g.Player.EatBox[0]+x, g.Player.Pos.Y+1.0) &&
			!g.Collide(g.Player.Pos.X+g.Player.EatBox[0]+x, g.Player.Pos.Y+g.Player.EatBox[3]) {
			// Move left and check if floor under
			g.Player.Pos.X += x
			g.checkTouchesGround()
		}
	}

	if g.PlayerPosChannel != nil {
		g.PlayerPosChannel <- entity.PlayerInfo{
			Nickname: g.Player.Nickname,
			Pos:      g.Player.Pos,
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
func (g *Game) checkTouchesGround() {
	if g.Player.TouchesGround {
		if !g.Collide(g.Player.Pos.X+g.Player.EatBox[0], g.Player.Pos.Y+g.Player.EatBox[3]+0.5) &&
			!g.Collide(g.Player.Pos.X+g.Player.EatBox[2], g.Player.Pos.Y+g.Player.EatBox[3]+0.5) {
			g.Player.TouchesGround = false
		}
	}
}
