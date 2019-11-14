package foxtrot

// https://www.geeksforgeeks.org/gap-buffer-data-structure/
// https://darrenburns.net/posts/piece-table/
// https://www.cs.unm.edu/~crowley/papers/sds.pdf
// https://web.archive.org/web/20160308183811/http://1017.songtrellisopml.com/whatsbeenwroughtusingpiecetables

type PieceTable struct {
}

func (pt PieceTable) Empty() {}

func (pt PieceTable) Insert() {}

func (pt PieceTable) Delete() {}

func (pt PieceTable) ItemAt() {}

type Piece struct {
	start, length int
}
