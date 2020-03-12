package cell

type Cells []Cell

func (cells Cells) Add(c Cell) Cells {
	return append(cells, c)
}

func (cells Cells) Get(i int) Cell {
	return cells[i]
}

func (cells Cells) Length() int {
	return len(cells)
}

func (cells Cells) Insert(i int, c Cell) Cells {
	cells = append(cells, c)
	copy(cells[i+1:], cells[i:])
	cells[i] = c
	return cells
}

func (cells Cells) Delete(i int) Cells {
	if i < len(cells)-1 {
		copy(cells[i:], cells[i+1:])
	}
	cells[len(cells)-1] = nil // or the zero value of T
	cells = cells[:len(cells)-1]
	return cells
}
