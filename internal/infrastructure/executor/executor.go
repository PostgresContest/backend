package executor

import (
	dbPublic "backend/internal/infrastructure/db/public"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	pgxdec "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Executor struct {
	connection *pgxpool.Pool
	typeMap    *pgtype.Map
}

func NewProvider(connection *dbPublic.Connection) *Executor {
	typeMap := pgtype.NewMap()
	pgxdec.Register(typeMap)

	return &Executor{
		connection: connection.Pool,
		typeMap:    typeMap,
	}
}

type FieldDescription struct {
	Name     string
	Datatype string
}

type Row []string

type Result struct {
	Query            string
	Rows             []Row
	FieldDescription []FieldDescription
	QueryHash        string
	ResultHash       string
}

func (e *Executor) normalizeDatatype(datatype pgconn.FieldDescription) string {
	result := ""
	t, ok := e.typeMap.TypeForOID(datatype.DataTypeOID)
	if !ok {
		return result
	}

	result = t.Name
	if datatype.DataTypeSize > -1 {
		result += fmt.Sprintf("(%d)", datatype.DataTypeSize)
	}

	return result
}

func generateQueryHash(query string) string {
	sum := sha256.Sum256([]byte(query))

	return fmt.Sprintf("%x", sum)
}

func generateResultHash(fd []FieldDescription, rows []Row) string {
	buf := strings.Builder{}
	for i, description := range fd {
		buf.WriteString(fmt.Sprintf("%d %s %s", i, description.Name, description.Datatype))
	}
	for r, row := range rows {
		for c, col := range row {
			buf.WriteString(fmt.Sprintf("%d %d %s", r, c, col))
		}
	}

	sum := sha256.Sum256([]byte(buf.String()))

	return fmt.Sprintf("%x", sum)
}

func (e *Executor) makeResponse(query string, rawValues [][][]uint8, fieldDesc []pgconn.FieldDescription) *Result {
	resp := new(Result)
	resp.Query = query
	resp.Rows = make([]Row, len(rawValues))
	for r, value := range rawValues {
		resp.Rows[r] = make(Row, len(value))
		for c, uint8s := range value {
			resp.Rows[r][c] = string(uint8s)
		}
	}

	resp.FieldDescription = make([]FieldDescription, len(fieldDesc))

	for i, description := range fieldDesc {
		resp.FieldDescription[i].Name = description.Name
		resp.FieldDescription[i].Datatype = e.normalizeDatatype(description)
	}

	resp.QueryHash = generateQueryHash(query)
	resp.ResultHash = generateResultHash(resp.FieldDescription, resp.Rows)

	return resp
}

var (
	banWords                 = []string{"SAVEPOINT", "BEGIN", "COMMIT", "RELEASE"}
	ErrQueryContainsBanWords = errors.New("query contains ban words")
)

func checkBanWords(query string) bool {
	query = strings.ToUpper(query)
	for _, word := range banWords {
		if strings.Contains(query, word) {
			return false
		}
	}
	return true
}

func (e *Executor) Execute(ctx context.Context, query string) (*Result, error) {
	if ok := checkBanWords(query); !ok {
		return nil, ErrQueryContainsBanWords
	}

	tr, err := e.connection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tr.Rollback(ctx)

	rows, err := tr.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values [][][]uint8
	for rows.Next() {
		rawValues := rows.RawValues()
		value := make([][]byte, len(rawValues))
		for i, rValue := range rawValues {
			value[i] = make([]byte, len(rValue))
			copy(value[i], rValue)
		}

		values = append(values, value)
	}

	response := e.makeResponse(query, values, rows.FieldDescriptions())

	return response, nil
}
