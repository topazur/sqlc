// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: group_concat.sql

package aggregate_functions

import (
	"context"
	"database/sql"
)

const groupConcat = `-- name: GroupConcat :many
SELECT student_name, GROUP_CONCAT(test_score)
FROM student
GROUP BY student_name
`

type GroupConcatRow struct {
	StudentName sql.NullString
	GroupConcat sql.NullString
}

func (q *Queries) GroupConcat(ctx context.Context) ([]GroupConcatRow, error) {
	rows, err := q.db.QueryContext(ctx, groupConcat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroupConcatRow
	for rows.Next() {
		var i GroupConcatRow
		if err := rows.Scan(&i.StudentName, &i.GroupConcat); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const groupConcatOrderBy = `-- name: GroupConcatOrderBy :many
SELECT student_name,
    GROUP_CONCAT(DISTINCT test_score ORDER BY test_score DESC SEPARATOR ' ')
FROM student
GROUP BY student_name
`

type GroupConcatOrderByRow struct {
	StudentName sql.NullString
	GroupConcat sql.NullString
}

func (q *Queries) GroupConcatOrderBy(ctx context.Context) ([]GroupConcatOrderByRow, error) {
	rows, err := q.db.QueryContext(ctx, groupConcatOrderBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroupConcatOrderByRow
	for rows.Next() {
		var i GroupConcatOrderByRow
		if err := rows.Scan(&i.StudentName, &i.GroupConcat); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
