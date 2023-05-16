# gopherLand2

So I'll try to build a 2D game, exploration, construction, fights, lvl up, multi players.

## Doc

> `gopherLand2.exe`

*Run the game in solo mode as Gopher.*

> `gopherLand2.exe solo gina`

*Run the game in solo mode as Gina.*

> `gopherLand2.exe multiplayer localhost 4565 Pazu george`

*Run the game in multiplayer mode defining host, port, player nickname (Pazu) and playing as George*

> `gopherLand2.exe serer`

*Run a TCP server on localhost trough port* **12387**.

## Characters

- **George** - *George is based. He maintains a beautiful moustache.*
- **Gina** - *Gina is pretty and use a lot of makeup. Though she is completely stupid.*
- **Gavin** - *Gavin thinks the world is easy. He is also a good student.*
- **Gwen** - *Gwen is a dark depressive girl. She has a very bad memory.*
- **Galien** - *Fkzjvqo mxbdpylkw nruaties hgc. Bibble wibble flib.*
- **Gopher** - *Gopher knows the truth about this world. He will never tell.*

## To do

| State | To do         | Description                                                                             |
| ----- | ------------- | --------------------------------------------------------------------------------------- |
| Nok   | Mob/NPC       | Player can interact with NPC and mobs                                                   |
| Nok   | Items         | Items just exists in the game, can pop, can be stored                                   |
| Ok    | Multiplatform | The game works for Win32, Win64, and most used Linux archs, server can run on raspberry |
| Ok    | Player        | A player is animated and can move through a simple key set                              |
| Ok    | Multiplayer   | Possible to play multiplayer through TCP server                                         |
| Ok    | Map           | Read a map in a file format and displays it correctly                                   |
| Ok    | Window        | Open resizable window                                                                   |

## Bugs to fix

| State | To fix                 | Description                                                                       |
| ----- | ---------------------- | --------------------------------------------------------------------------------- |
| Bug   | Map race conditions    | Golang maps does not support concurrent access                                    |
| Fixed | Pass movement, not pos | For TCP multiplayer, do not provide player position, but rather player keyPressed |