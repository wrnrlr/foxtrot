package nbx

import (
	"github.com/stretchr/testify/assert"
	"github.com/wrnrlr/foxtrot/cell"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	r := strings.NewReader(`
	<notebook version="0.0.1">
		<cell type="Input">1+1</cell>
		<cell type="Output">2</cell>
	</notebook>`)
	cells, err := Read(r)
	assert.Nil(t, err)
	assert.NotNil(t, cells)
	assert.Equal(t, 2, len(cells))
	assert.Equal(t, cell.Input, cells[0].Type())
	assert.Equal(t, "1+1", cells[0].Text())
	assert.Equal(t, cell.Output, cells[1].Type())
}

func TestReadEmptyFile(t *testing.T) {
	r := strings.NewReader("")
	cells, err := Read(r)
	assert.Nil(t, err)
	assert.Nil(t, cells)
	assert.Equal(t, 0, len(cells))
}

func TestReadInvalidXML(t *testing.T) {
	r := strings.NewReader(`<notebook`)
	_, err := Read(r)
	assert.NotNil(t, err)
}
