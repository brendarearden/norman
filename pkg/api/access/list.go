package access

import (
	"fmt"

	builder2 "github.com/rancher/norman/pkg/parse/builder"

	"github.com/rancher/norman/pkg/types"
	"github.com/rancher/norman/pkg/types/convert"
)

func Create(context *types.APIOperation, typeName string, data map[string]interface{}, into interface{}) error {
	schema := context.Schemas.Schema(typeName)
	if schema == nil {
		return fmt.Errorf("failed to find schema " + typeName)
	}

	item, err := schema.Store.Create(context, schema, data)
	if err != nil {
		return err
	}

	b := builder2.NewBuilder(context)

	item, err = b.Construct(schema, item, builder2.List)
	if err != nil {
		return err
	}

	if into == nil {
		return nil
	}

	return convert.ToObj(item, into)
}

func ByID(context *types.APIOperation, typeName string, id string, into interface{}) error {
	schema := context.Schemas.Schema(typeName)
	if schema == nil {
		return fmt.Errorf("failed to find schema " + typeName)
	}

	item, err := schema.Store.ByID(context, schema, id)
	if err != nil {
		return err
	}

	b := builder2.NewBuilder(context)

	item, err = b.Construct(schema, item, builder2.List)
	if err != nil {
		return err
	}

	if into == nil {
		return nil
	}

	return convert.ToObj(item, into)
}

func List(context *types.APIOperation, typeName string, opts *types.QueryOptions, into interface{}) error {
	schema := context.Schemas.Schema(typeName)
	if schema == nil {
		return fmt.Errorf("failed to find schema " + typeName)
	}

	data, err := schema.Store.List(context, schema, opts)
	if err != nil {
		return err
	}

	b := builder2.NewBuilder(context)

	var newData []map[string]interface{}
	for _, item := range data {
		item, err = b.Construct(schema, item, builder2.List)
		if err != nil {
			return err
		}
		newData = append(newData, item)
	}

	return convert.ToObj(newData, into)
}