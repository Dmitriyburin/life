package life

import (
	"math/rand"
	"time"
)

type World struct {
	Height int // Высота сетки
	Width  int // Ширина сетки
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}
func (w *World) next(x, y int) bool {
	n := w.neighbors(x, y)       // получим количество живых соседей
	alive := w.Cells[y][x]       // текущее состояние клетки
	if n < 4 && n > 1 && alive { // если соседей двое или трое, а клетка жива
		return true // то следующее состояние — жива
	}
	if n == 3 && !alive { // если клетка мертва, но у неё трое соседей
		return true // клетка оживает
	}

	return false // в любых других случаях — клетка мертва
}
func (w *World) neighbors(x, y int) int {
	n := 0
	tor := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	for _, v := range tor {
		xx, yy := x+v[0], y+v[1]
		if xx < 0 || xx >= w.Height || yy < 0 || yy >= w.Width {
			continue
		}
		if w.Cells[yy][xx] {
			n++
		}
	}

	return n
}

func NextState(oldWorld, newWorld *World) {
	for i := 0; i < oldWorld.Height; i++ {
		for j := 0; j < oldWorld.Width; j++ {
			newWorld.Cells[i][j] = oldWorld.next(j, i)
		}
	}
}

// RandInit заполняет поля на указанное число процентов
func (w *World) RandInit(percentage int) {
	numAlive := percentage * w.Height * w.Width / 100
	w.fillAlive(numAlive)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < w.Height*w.Width; i++ {
		randRowLeft := r.Intn(w.Width)
		randColLeft := r.Intn(w.Height)
		randRowRight := r.Intn(w.Width)
		randColRight := r.Intn(w.Height)

		w.Cells[randRowLeft][randColLeft] = w.Cells[randRowRight][randColRight]
	}
}

func (w *World) fillAlive(num int) {
	aliveCount := 0
	for j, row := range w.Cells {
		for k := range row {
			w.Cells[j][k] = true
			aliveCount++
			if aliveCount == num {

				return
			}
		}
	}
}
