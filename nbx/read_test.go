package nbx

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	r := strings.NewReader(`
	<notebook version="0.0.1">
		<cell type="Input">1+1</cell>
	</notebook>`)
	cells, err := Read(r)
	assert.Nil(t, err)
	assert.NotNil(t, cells)
	assert.Equal(t, 1, len(cells))
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
