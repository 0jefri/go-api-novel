package constant

const (
	NOVEL_INSERT = "INSERT INTO novel(id,judul,penerbit,tahun_terbit,penulis)VALUES($1, $2, $3, $4, $5)"
	LIST_NOVEL   = "SELECT * FROM novel;"
	GET_NOVEL_ID = "SELECT * FROM novel WHERE id = $1;"
	NOVEL_UPDATE = "UPDATE novel SET judul=$1, penerbit=$2, tahun_terbit=$3, penulis=$4 WHERE id=$5"
	NOVEL_DELETE = "DELETE FROM novel WHERE id=$1"
)
