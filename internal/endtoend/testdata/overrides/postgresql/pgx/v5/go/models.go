// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package override

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kyleconroy/sqlc-testdata/pkg"
	"github.com/lib/pq"
)

type Foo struct {
	Other   string
	Total   int64
	Tags    pgtype.FlatArray[string]
	ByteSeq []byte
	Retyped pkg.CustomType
	Langs   pq.StringArray
}
