package nbx

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/corywalker/expreduce/expreduce"
	"github.com/corywalker/expreduce/expreduce/atoms"
	"github.com/corywalker/expreduce/expreduce/parser"
	"github.com/corywalker/expreduce/pkg/expreduceapi"
	"github.com/wrnrlr/foxtrot/cell"
	"github.com/wrnrlr/foxtrot/theme"
	"io"
	"io/ioutil"
	"os"
	"strconv"
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

type cll struct {
	Rules    []string
	Children string
}

func ReadNBX(filename string) (cell.Cells, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Read(file)
}

func ReadCell(filename string) (cell.Cells, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readCells(file, filename)
}

func readCells(r io.Reader, filename string) (cell.Cells, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		fmt.Println("Failed to read file.")
		return nil, err
	}
	content := buf.String()
	src := parser.ReplaceSyms(content)
	buf = bytes.NewBufferString(src)
	kernel := expreduce.NewEvalState()
	ex, err := parser.InterpBuf(buf, filename, kernel)
	if err != nil {
		fmt.Println("Failed to parse file.")
		return nil, err
	}
	//ex = kernel.Eval(ex)
	cells, _ := parseCells(ex)
	return cells, nil
}

func parseCells(ex expreduceapi.Ex) (cells cell.Cells, err error) {
	// parse is Expression
	e, ok := ex.(*atoms.Expression)
	if !ok {
		return nil, errors.New("ex not an Expression")
	}
	// Check is List
	if e.HeadStr() != "System`List" && e.HeadStr() != "Global`Notebook" {
		return nil, errors.New("ex not an List")
	}
	// Loop over arguments
	for _, part := range e.Parts[1:] {
		cell, err := parseCell(part)
		if err == nil {
			cells = append(cells, cell)
		}
	}
	// Create Cell
	// Add to cells
	fmt.Println(ex)
	return cells, err
}

func parseCell(ex expreduceapi.Ex) (cell.Cell, error) {
	// parse is Expression
	e, ok := ex.(*atoms.Expression)
	if !ok {
		return nil, errors.New("ex not an Expression")
	}
	t, ok := ParseCellType(e)
	if !ok {
		return nil, errors.New("ex not an Expression")
	}
	// Set Text/Content
	content, ok := ParseString(e.GetPart(1))
	if !ok {
		return nil, errors.New("ex not an Expression")
	}
	txt, err := UnquoteString(content)
	th := theme.DefaultStyles()
	label := ParseCellLabel(t)
	c := cell.NewCell(t, label, th)
	c.SetText(txt)
	return c, err
}

func ParseCellLabel(t cell.Type) string {
	switch t {
	case cell.Input:
		return "In[] :="
	case cell.Output:
		return "Out[] ="
	default:
		return ""
	}
}

func ParseCellType(ex *atoms.Expression) (cell.Type, bool) {
	switch ex.HeadStr() {
	case "Global`Header1":
		return cell.H1, true
	case "Global`Header2":
		return cell.H2, true
	case "Global`Header3":
		return cell.H3, true
	case "Global`Header4":
		return cell.H4, true
	case "Global`Header5":
		return cell.H5, true
	case "Global`Header6":
		return cell.H6, true
	case "Global`Paragraph":
		return cell.Paragraph, true
	case "Global`Input":
		return cell.Input, true
	case "Global`Output":
		return cell.Output, true
	case "Global`Code":
		return cell.Code, true
	default:
		return cell.Empty, false
	}
}

func UnquoteString(a *atoms.String) (string, error) {
	s := a.StringForm(atoms.DefaultStringParams())
	s, err := strconv.Unquote(s)
	return s, err
}

func ParseString(ex expreduceapi.Ex) (*atoms.String, bool) {
	s, ok := ex.(*atoms.String)
	return s, ok
}

func ParseExpression(ex expreduceapi.Ex) (*atoms.Expression, bool) {
	e, ok := ex.(*atoms.Expression)
	return e, ok
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
