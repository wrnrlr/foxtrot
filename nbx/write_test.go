package nbx

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/wrnrlr/foxtrot/cell"
	"testing"
)

func TestWrite(t *testing.T) {
	buffer := new(bytes.Buffer)
	var cells cell.Cells
	c := cell.NewCell(cell.Input, "In[0]:=", nil)
	cells = append(cells, c)
	err := Write(buffer, cells)
	assert.Nil(t, err)
	cells, err = Read(buffer)
	assert.Equal(t, 1, len(cells))
}
