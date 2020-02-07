package life

import "math/rand"

type gameField struct {
	field [][]bool
}

func (g *gameField) set(x int, y int, b bool) {
	g.field[y][x] = b
}

func (g *gameField) init(width int, height int) {
	g.field = make([][]bool, height)
	for i := range g.field {
		g.field[i] = make([]bool, width)
	}
}

func (g *gameField) next(cellCount int, x int, y int) bool {
	aliveCount := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && g.Alive(cellCount, x+i, y+j) {
				aliveCount++
			}
		}
	}
	return aliveCount == 3 || aliveCount == 2 && g.Alive(cellCount, x, y)
}

func (g *gameField) Alive(cellCount int, x int, y int) bool {
	// x += cellCount
	// x %= cellCount
	// y += cellCount
	// y %= cellCount

	if x < 0 || y < 0 {
		return false
	} else if x >= cellCount || y >= cellCount {
		return false
	}

	return g.field[y][x]
}

// Life has cell count and two gameFields (current, next).
type Life struct {
	CellCount     int
	Current, Next gameField
}

// Step is simulate next generation and copy from next gameField to current gameField.
func (g *Life) Step() {
	for y := 0; y < g.CellCount; y++ {
		for x := 0; x < g.CellCount; x++ {
			g.Next.set(x, y, g.Current.next(g.CellCount, x, y))
		}
	}

	for i := 0; i < len(g.Next.field); i++ {
		copy(g.Current.field[i], g.Next.field[i])
	}
}

// Setup for Current & Next gameField
func (g *Life) Setup(cellCount int) {
	g.CellCount = cellCount

	current := gameField{}
	next := gameField{}

	current.init(cellCount, cellCount)
	next.init(cellCount, cellCount)

	g.Current = current
	g.Next = next
}

// InitCurrentField set Life to Current Field
func (g *Life) InitCurrentField() {
	// Random
	for i := 0; i < (g.CellCount * g.CellCount / 4); i++ {
		g.Current.set(rand.Intn(g.CellCount), rand.Intn(g.CellCount), true)
	}

	// Glider
	// x := int(g.CellCount / 2)
	// y := int(g.CellCount / 2)
	// g.Current.set(x+1, y+4, true)
	// g.Current.set(x+2, y+2, true)
	// g.Current.set(x+2, y+4, true)
	// g.Current.set(x+3, y+3, true)
	// g.Current.set(x+3, y+4, true)

	// Thunderbird
	// x := int(g.CellCount / 2)
	// y := int(g.CellCount / 2)
	// g.Current.set(x-1, y+1, true)
	// g.Current.set(x, y+1, true)
	// g.Current.set(x+1, y+1, true)
	// g.Current.set(x, y+3, true)
	// g.Current.set(x, y+4, true)
	// g.Current.set(x, y+5, true)

	// Cross
	// for x := 2; x < g.CellCount-1; x++ {
	// 	y := math.Floor(float64(g.CellCount * x / g.CellCount))
	// 	if y > 1 {
	// 		g.Current.set(x, int(y), true)
	// 		g.Current.set(x, g.CellCount-int(y), true)
	// 	}
	// }
}

func (g *Life) SetCurrentField(x, y int, set bool) {
	g.Current.set(x, y, set)
}
