<p align="center"><img width=50% src="https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/31c5ef3a-c8fd-4d20-b19c-bace7f78f285/dapghvb-3ddef599-4b47-4fb0-8b4f-f3175e4bbf70.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcLzMxYzVlZjNhLWM4ZmQtNGQyMC1iMTljLWJhY2U3Zjc4ZjI4NVwvZGFwZ2h2Yi0zZGRlZjU5OS00YjQ3LTRmYjAtOGI0Zi1mMzE3NWU0YmJmNzAucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.S2WHekX4e7XK9lOjU7v1rKgd2OrjOVND_fN3dsD7cGg"></p>

## Table of Content 📚
- [Overview](#overview)
- [Description](#description)
- [Contributors](#contributors)

## Overview
This pacman is a multithreaded version of the arcade video game [Pacman](https://en.wikipedia.org/wiki/Pac-Man). This version is a
Computer vs Human game. Each enemy is independent and the number of enemies is configurable. 

## Description
- The game's maze layout is static.
- The `pacman` gamer is controlled by the user.
- Enemies are autonomous entities that will move a random way.
- Enemies and pacman respect the layout limits and walls.
- Enemies number can be configured on game's start.
- Each enemy's behaviour is implemented as a separated thread.
- Enemies and pacman threads use the same map or game layout data structure resource.
- Display obtained pacman's scores.
- Pacman loses a life when an enemy touches it.
- Pacman loses the game when it ran out of lifes.
- Pacman wins the game when it has taken all coins in the map.

## Contributors
<table>
  <tr>
    <td align="center"><a href="https://github.com/SYM1000"><img src="https://avatars.githubusercontent.com/u/20364366?v=4" width="100px;" alt=""/><br /><sub><b>Santiago Yeomans</b></sub></a><br /></td>
    <td align="center"><a href="https://github.com/UlisesBojorquez"><img src="https://avatars.githubusercontent.com/u/35876113?v=4" width="100px;" alt=""/><br /><sub><b>Ulises Bojorquez</b></sub></a><br /></td>
  </tr>
