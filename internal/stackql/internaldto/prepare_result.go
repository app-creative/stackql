package internaldto

import (
	"github.com/stackql/go-openapistackql/openapistackql"
)

type PrepareResultSetDTO struct {
	OutputBody    map[string]interface{}
	Msg           *BackendMessages
	RawRows       map[int]map[int]interface{}
	RowMap        map[string]map[string]interface{}
	ColumnOrder   []string
	ColumnSchemas []*openapistackql.Schema
	RowSort       func(map[string]map[string]interface{}) []string
	Err           error
}

func NewPrepareResultSetDTO(
	body map[string]interface{},
	rowMap map[string]map[string]interface{},
	columnOrder []string,
	rowSort func(map[string]map[string]interface{}) []string,
	err error,
	msg *BackendMessages,
) PrepareResultSetDTO {
	return PrepareResultSetDTO{
		OutputBody:  body,
		RowMap:      rowMap,
		ColumnOrder: columnOrder,
		RowSort:     rowSort,
		Err:         err,
		Msg:         msg,
		RawRows:     map[int]map[int]interface{}{},
	}
}

func NewPrepareResultSetPlusRawDTO(
	body map[string]interface{},
	rowMap map[string]map[string]interface{},
	columnOrder []string,
	rowSort func(map[string]map[string]interface{}) []string,
	err error,
	msg *BackendMessages,
	rawRows map[int]map[int]interface{},
) PrepareResultSetDTO {
	return PrepareResultSetDTO{
		OutputBody:  body,
		RowMap:      rowMap,
		ColumnOrder: columnOrder,
		RowSort:     rowSort,
		Err:         err,
		Msg:         msg,
		RawRows:     rawRows,
	}
}

func NewPrepareResultSetPlusRawAndTypesDTO(
	body map[string]interface{},
	rowMap map[string]map[string]interface{},
	columnOrder []string,
	columnSchemas []*openapistackql.Schema,
	rowSort func(map[string]map[string]interface{}) []string,
	err error,
	msg *BackendMessages,
	rawRows map[int]map[int]interface{},
) PrepareResultSetDTO {
	return PrepareResultSetDTO{
		OutputBody:    body,
		RowMap:        rowMap,
		ColumnOrder:   columnOrder,
		ColumnSchemas: columnSchemas,
		RowSort:       rowSort,
		Err:           err,
		Msg:           msg,
		RawRows:       rawRows,
	}
}