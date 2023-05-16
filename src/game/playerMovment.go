package game

import (
	"gopherLand2/src/game/entity"
	"gopherLand2/src/localInstance/input"
	"math"
)

// Each game tick, move player
func (g *Game) ManagePlayerMovement(nickname string) {
	if g.Players[nickname].MovesLeft && !g.Players[nickname].MovesRight {
		g.moveHorizonally(nickname, -entity.DEFAULT_SPEED)
	}
	if !g.Players[nickname].MovesLeft && g.Players[nickname].MovesRight {
		g.moveHorizonally(nickname, entity.DEFAULT_SPEED)
	}
	if g.Players[nickname].IsJumping {
		if g.Players[nickname].TouchesGround {
			if p, ok := g.Players[nickname]; ok {
				p.TouchesGround = false
				p.VerticalVelocity -= g.Players[nickname].JumpSpeed
				g.Players[nickname] = p
			}
		}
	}

	if !g.Players[nickname].TouchesGround {
		g.moveVertically(nickname, g.Players[nickname].VerticalVelocity)
		if p, ok := g.Players[nickname]; ok {
			p.VerticalVelocity += entity.GRAVITY
			g.Players[nickname] = p
		}
	}
}

// Handle player's control
func (g *Game) RunAllPlayers() {
	var keyPressed input.KeyPressed

	for {
		keyPressed = <-g.PlayerInputChannel

		if keyPressed.Nickname == g.Nickname {
			if g.PlayerPosChannel != nil {
				g.PlayerPosChannel <- entity.PlayerInfo{
					Nickname:   g.Players[keyPressed.Nickname].Nickname,
					Character:  g.Players[keyPressed.Nickname].Character,
					Pos:        g.Players[keyPressed.Nickname].Pos,
					KeyPressed: keyPressed,
				}
			}
		}

		if p, ok := g.Players[keyPressed.Nickname]; ok {
			if keyPressed.Left {
				p.MovesLeft = true
			} else {
				p.MovesLeft = false
			}
			if keyPressed.Right {
				p.MovesRight = true
			} else {
				p.MovesRight = false
			}
			if keyPressed.Up {
				p.IsJumping = true
			} else {
				p.IsJumping = false
			}

			g.Players[keyPressed.Nickname] = p
		}
	}
}

// Moves player towards direction if possible
func (g *Game) moveVertically(nickname string, y float64) {
	if p, ok := g.Players[nickname]; ok {
		if y > 0.0 { // Player goes down
			if !g.Collide(nickname, p.Pos.X+p.EatBox[0], p.Pos.Y+p.EatBox[3]+y) &&
				!g.Collide(nickname, p.Pos.X+p.EatBox[2], p.Pos.Y+p.EatBox[3]+y) {
				p.Pos.Y += y
			} else {
				// Move down as much as possible
				p.Pos.Y = math.Round(p.Pos.Y)
				// Land in ground
				p.TouchesGround = true
				p.VerticalVelocity = 0.0
			}
		} else if y < 0.0 { // Player goes up
			if !g.Collide(nickname, p.Pos.X+p.EatBox[0], p.Pos.Y+p.EatBox[1]+y) &&
				!g.Collide(nickname, p.Pos.X+p.EatBox[2], p.Pos.Y+p.EatBox[1]+y) {
				p.Pos.Y += y
			} else {
				// Hit ceiling
				p.VerticalVelocity = 0.0
			}
		}
		g.Players[nickname] = p
	}
}

// Moves left or right and check collisions
func (g *Game) moveHorizonally(nickname string, x float64) {
	if p, ok := g.Players[nickname]; ok {
		if x > 0.0 { // Player goes right
			if !g.Collide(nickname, p.Pos.X+p.EatBox[2]+x, p.Pos.Y+p.EatBox[1]) &&
				!g.Collide(nickname, p.Pos.X+p.EatBox[2]+x, p.Pos.Y+1.0) &&
				!g.Collide(nickname, p.Pos.X+p.EatBox[2]+x, p.Pos.Y+p.EatBox[3]) {
				// Move right and check if floor under
				p.Pos.X += x
				g.checkTouchesGround(nickname)
			}
		} else if x < 0.0 { // Player goes left
			if !g.Collide(nickname, p.Pos.X+p.EatBox[0]+x, p.Pos.Y+p.EatBox[1]) &&
				!g.Collide(nickname, p.Pos.X+p.EatBox[0]+x, p.Pos.Y+1.0) &&
				!g.Collide(nickname, p.Pos.X+p.EatBox[0]+x, p.Pos.Y+p.EatBox[3]) {
				// Move left and check if floor under
				p.Pos.X += x
				g.checkTouchesGround(nickname)
			}
		}
		g.Players[nickname] = p
	}
}

// Check if a specific point in the map is in collision with a solid block
func (g *Game) Collide(nickname string, x, y float64) bool {
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
func (g *Game) checkTouchesGround(nickname string) {
	if g.Players[nickname].TouchesGround {
		if !g.Collide(nickname, g.Players[nickname].Pos.X+g.Players[nickname].EatBox[0], g.Players[nickname].Pos.Y+g.Players[nickname].EatBox[3]+0.5) &&
			!g.Collide(nickname, g.Players[nickname].Pos.X+g.Players[nickname].EatBox[2], g.Players[nickname].Pos.Y+g.Players[nickname].EatBox[3]+0.5) {
			if p, ok := g.Players[nickname]; ok {
				p.TouchesGround = false
				g.Players[nickname] = p
			}
		}
	}
}
