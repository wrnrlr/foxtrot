package nbx

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/parser"
	"github.com/wrnrlr/foxtrot/cell"
	"github.com/wrnrlr/foxtrot/theme"
	"io"
	"io/ioutil"
	"os"
)

type notebookTag struct {
	XMLName xml.Name
	Version string    `xml:"version,attr"`
	Cells   []cellTag `xml:"cell"`
}

type cellTag struct {
	XMLName xml.Name
	Type    string `xml:"type,attr"`
	Content string `xml:",chardata"`
}

func ReadNBX(filename string) (cell.Cells, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Read(file)
}

func Read(r io.Reader) (cell.Cells, error) {
	var nb notebookTag
	nb.XMLName = xml.Name{Local: "notebook"}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("Failed to read file.")
		return nil, err
	}
	err = xml.Unmarshal(b, &nb)
	if err != nil && err != io.EOF {
		fmt.Println("Failed to parse file.")
		return nil, err
	}
	fmt.Printf("%v\n", nb)
	styles := theme.DefaultStyles()
	kernel := expreduce.NewEvalState()
	var cells cell.Cells
	for i, c := range nb.Cells {
		c2 := cell.NewCell(cell.ParseType(c.Type), string(i), styles)
		c2.SetText(c.Content)
		if c2.Type() == cell.Output {
			src := parser.ReplaceSyms(c.Content)
			buf := bytes.NewBufferString(src)
			expOut, err := parser.InterpBuf(buf, "nofile", kernel)
			if err != nil {
				fmt.Println("Failed to read output cell")
				return nil, err
			}
			c2.SetOut(expOut)
		}
		cells = append(cells, c2)
	}
	return cells, nil
}
