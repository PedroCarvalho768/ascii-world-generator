package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000
	screenHeight = 800
	worldWidth   = 80
	worldHeight  = 50
	cellSize     = 16
)

// TerrainType represents different types of terrain
type TerrainType int

const (
	Water TerrainType = iota
	Sand
	Grass
	Forest
	Mountain
	Snow
)

// Tile represents a single tile in the world
type Tile struct {
	Terrain TerrainType
	Char    rune
	Color   rl.Color
}

// World represents the game world
type World struct {
	Tiles  [][]Tile
	Width  int
	Height int
}

// NewWorld creates a new random world
func NewWorld(width, height int) *World {
	world := &World{
		Width:  width,
		Height: height,
		Tiles:  make([][]Tile, height),
	}

	for i := range world.Tiles {
		world.Tiles[i] = make([]Tile, width)
	}

	world.Generate()
	return world
}

// Generate creates a random world using simple noise-like algorithm
func (w *World) Generate() {
	// Create height map using simple diamond-square-like algorithm
	heightMap := make([][]float64, w.Height)
	for i := range heightMap {
		heightMap[i] = make([]float64, w.Width)
	}

	// Fill with random values and smooth
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			heightMap[y][x] = rand.Float64()
		}
	}

	// Smooth the height map
	for iter := 0; iter < 3; iter++ {
		newMap := make([][]float64, w.Height)
		for i := range newMap {
			newMap[i] = make([]float64, w.Width)
		}

		for y := 0; y < w.Height; y++ {
			for x := 0; x < w.Width; x++ {
				sum := 0.0
				count := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						ny, nx := y+dy, x+dx
						if ny >= 0 && ny < w.Height && nx >= 0 && nx < w.Width {
							sum += heightMap[ny][nx]
							count++
						}
					}
				}
				newMap[y][x] = sum / float64(count)
			}
		}
		heightMap = newMap
	}

	// Convert height map to terrain
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			height := heightMap[y][x]
			w.Tiles[y][x] = w.heightToTile(height)
		}
	}
}

// heightToTile converts a height value to a terrain tile
func (w *World) heightToTile(height float64) Tile {
	var tile Tile

	switch {
	case height < 0.25:
		tile.Terrain = Water
		tile.Char = '~'
		tile.Color = rl.NewColor(65, 105, 225, 255) // Royal Blue
	case height < 0.35:
		tile.Terrain = Sand
		tile.Char = '.'
		tile.Color = rl.NewColor(238, 214, 175, 255) // Tan
	case height < 0.55:
		tile.Terrain = Grass
		tile.Char = '"'
		tile.Color = rl.NewColor(34, 139, 34, 255) // Forest Green
	case height < 0.70:
		tile.Terrain = Forest
		tile.Char = '♣'
		tile.Color = rl.NewColor(0, 100, 0, 255) // Dark Green
	case height < 0.85:
		tile.Terrain = Mountain
		tile.Char = '^'
		tile.Color = rl.NewColor(139, 137, 137, 255) // Gray
	default:
		tile.Terrain = Snow
		tile.Char = '*'
		tile.Color = rl.NewColor(255, 250, 250, 255) // Snow White
	}

	return tile
}

// Render draws the world to the screen
func (w *World) Render() {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			tile := w.Tiles[y][x]
			posX := int32(x * cellSize)
			posY := int32(y * cellSize)

			// Draw character
			rl.DrawText(string(tile.Char), posX, posY, cellSize, tile.Color)
		}
	}
}

func main() {
	// Initialize window
	rl.InitWindow(screenWidth, screenHeight, "ASCII World Generator")
	rl.SetTargetFPS(60)

	// Create initial world (using auto-seeded global random source in Go 1.20+)
	world := NewWorld(worldWidth, worldHeight)

	// Main game loop
	for !rl.WindowShouldClose() {
		// Update
		if rl.IsKeyPressed(rl.KeySpace) {
			world.Generate()
		}

		if rl.IsKeyPressed(rl.KeyR) {
			world = NewWorld(worldWidth, worldHeight)
		}

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		world.Render()

		// Draw instructions
		instructionY := int32(screenHeight - 60)
		rl.DrawText("SPACE: Regenerate World", 10, instructionY, 20, rl.White)
		rl.DrawText("R: New Random World", 10, instructionY+25, 20, rl.White)
		rl.DrawText("ESC: Exit", 10, instructionY+50, 20, rl.White)

		rl.EndDrawing()
	}

	// Close window
	rl.CloseWindow()
}
