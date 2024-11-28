package sources

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/turbot/tailpipe-plugin-chaos/config"
	"github.com/turbot/tailpipe-plugin-chaos/rows"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

const StructColumnsSourceIdentifier = "chaos_struct_columns"

// init function should register the source
func init() {
	row_source.RegisterRowSource[*StructColumnsSource]()
}

// StructColumnsSource source is responsible for generating some rows of data
type StructColumnsSource struct {
	// row_source.RowSourceImpl[*StructColumnsSourceConfig]
	row_source.RowSourceImpl[*StructColumnsSourceConfig, *config.ChaosConnection]
}

func (s *StructColumnsSource) Identifier() string {
	return StructColumnsSourceIdentifier
}

func (s *StructColumnsSource) Collect(ctx context.Context) error {
	sourceName := StructColumnsSourceIdentifier
	sourceEnrichmentFields := &enrichment.CommonFields{
		TpSourceName: &sourceName,
		TpSourceType: StructColumnsSourceIdentifier,
	}

	for i := 1; i <= s.Config.RowCount; i++ {
		rowData := s.populateRowData(i)
		row := &types.RowData{Data: rowData, Metadata: sourceEnrichmentFields}
		if err := s.OnRow(ctx, row, nil); err != nil {
			return fmt.Errorf("error processing row: %w", err)
		}
	}

	return nil
}

func (s *StructColumnsSource) populateRowData(i int) *rows.StructColumns {
	var rowData rows.StructColumns
	rowData.Identifier = fmt.Sprintf("row-%d", i)
	rowData.Timestamp = time.Now()
	rowData.SimpleStruct = s.randomSimpleStruct()
	rowData.ArrayStruct = s.randomStructWithArray()
	rowData.NestedStruct = s.randomNestedStruct(rowData.SimpleStruct)
	return &rowData
}

func (s *StructColumnsSource) randomSimpleStruct() *rows.SimpleStruct {
	if rand.Intn(2) == 0 {
		return nil
	}

	return &rows.SimpleStruct{
		Id:   rand.Int63(),
		Name: fmt.Sprintf("name-%d", rand.Intn(100)),
	}
}

func (s *StructColumnsSource) randomStructWithArray() *rows.StructWithArray {
	if rand.Intn(2) == 0 {
		return nil
	}

	var numbers []int
	for i := 0; i < rand.Intn(10); i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	return &rows.StructWithArray{
		Id:      rand.Int63(),
		Numbers: numbers,
	}
}

func (s *StructColumnsSource) randomNestedStruct(sub *rows.SimpleStruct) *rows.NestedStruct {
	if rand.Intn(2) == 0 {
		return nil
	}

	return &rows.NestedStruct{
		Id:   rand.Int63(),
		Name: fmt.Sprintf("name-%d", rand.Intn(100)),
		Sub:  sub,
	}
}
