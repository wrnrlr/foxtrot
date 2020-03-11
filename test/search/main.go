package main

import (
	"github.com/blevesearch/bleve"
	"github.com/corywalker/expreduce/expreduce"
	"path"
)

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)

	es := expreduce.NewEvalState()
	es.GetDefinedMap()

	defSets := expreduce.GetAllDefinitions()
	for _, defSet := range defSets {
		defSet.Defs[0].Tests
		categoryFn := fmt.Sprintf("builtin/%s/index.md", defSet.Name)
		writeCategoryIndex(path.Join(*docs_location, categoryFn), defSet)
		categoryDef := fmt.Sprintf(
			"    - '%s': '%s'\n",
			defSet.Name,
			categoryFn,
		)
		f.WriteString(categoryDef)

		for _, def := range defSet.Defs {
			if def.OmitDocumentation {
				continue
			}
			def.AnnotateWithDynamic(es)
			symbolFn := fmt.Sprintf(
				"builtin/%s/%s.md",
				defSet.Name,
				defNameFile(def.Name),
			)
			writeSymbol(path.Join(*docs_location, symbolFn), defSet, def, es)
			symbolDef := fmt.Sprintf(
				"    - '%s ': '%s'\n",
				defNamePrint(def.Name),
				symbolFn,
			)
			f.WriteString(symbolDef)
		}
	}

	// index some data
	err = index.Index(identifier, your_data)

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
}
