package psql

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	stmt := Select(
		// top
		42,

		// distinct
		false,

		// columns
		[]string{
			AS(Identifier(`id`), Identifier(`ID`)),
			AS(Identifier(`fn`), Identifier(`FirstName`)),
			AS(Identifier(`ln`), Identifier(`LastName`)),
		},

		// table
		TableName(
			Identifier(`test`),   //db
			Identifier(`public`), //sch
			Identifier(`people`), //tb
		),

		// where
		AND(
			Where(
				map[string]string{
					Identifier("fn"): Text("Jac%"),
					Identifier("ln"): Text("?choa"),
				},
				LIKE,
				false,
				true,
			),
			Where(
				map[string]string{
					Identifier("bd"): Date(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
				},
				GE,
				false,
				true,
			),
		),

		// group by
		[]string{
			Identifier(`id`),
			Identifier(`fn`),
			Identifier(`ln`),
		},

		// order by
		[]string{
			Identifier(`ln`),
			Identifier(`fn`),
		},

		// terminate `;`
		true,
	)
	fmt.Println(stmt)
}
