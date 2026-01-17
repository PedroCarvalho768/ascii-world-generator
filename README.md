# ASCII World Generator

A random ASCII world generator built with Go and raylib-go. This project generates procedural terrain worlds displayed using ASCII characters.

## Features

- Procedural world generation using height map algorithm
- Multiple terrain types: Water (~), Sand (.), Grass ("), Forest (♣), Mountains (^), and Snow (*)
- Interactive controls to regenerate worlds
- Built with Go and raylib-go for graphical rendering

## Prerequisites

### System Dependencies

Before building this project, you need to install raylib dependencies:

**Ubuntu/Debian:**
```bash
sudo apt-get install -y libasound2-dev libx11-dev libxrandr-dev libxi-dev libgl1-mesa-dev libglu1-mesa-dev libxcursor-dev libxinerama-dev libwayland-dev libxkbcommon-dev
```

**Fedora:**
```bash
sudo dnf install -y alsa-lib-devel mesa-libGL-devel libXrandr-devel libXi-devel libXcursor-devel libXinerama-devel libatomic
```

**macOS:**
```bash
# No additional dependencies needed on macOS
```

**Windows:**
```bash
# Install Go from https://golang.org/dl/
# Install TDM-GCC or MinGW-w64
```

### Go

You need Go 1.21 or later installed. Download from [golang.org](https://golang.org/dl/).

## Building

```bash
# Clone the repository
git clone https://github.com/PedroCarvalho768/ascii-world-generator.git
cd ascii-world-generator

# Download dependencies
go mod download

# Build the project
go build -o ascii-world-generator .
```

## Running

```bash
./ascii-world-generator
```

## Controls

- **SPACE**: Regenerate the current world
- **R**: Create a new random world
- **ESC**: Exit the application

## World Generation

The world is generated using a simple height map algorithm:
1. Create a random noise grid
2. Smooth the grid using neighborhood averaging
3. Map height values to terrain types:
   - Very Low: Water (~)
   - Low: Sand (.)
   - Medium: Grass (")
   - Medium-High: Forest (♣)
   - High: Mountains (^)
   - Very High: Snow (*)

## Project Structure

```
.
├── main.go         # Main application with world generation and rendering
├── go.mod          # Go module definition
├── go.sum          # Dependency checksums
└── README.md       # This file
```

## Learning Resources

This project demonstrates:
- Go programming basics
- Random world generation algorithms
- ASCII rendering techniques
- Game loop architecture
- Using CGo and external libraries (raylib)

## License

This is a learning project created to explore Go and ASCII game development.