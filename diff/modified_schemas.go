package diff

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// ModifiedSchemas is map of schema names to their respective diffs
type ModifiedSchemas map[string]*SchemaDiff

func (modifiedSchemas ModifiedSchemas) addSchemaDiff(schema1 string, schemaRef1, schemaRef2 *openapi3.SchemaRef) {

	if diff := getSchemaDiff(schemaRef1, schemaRef2); !diff.empty() {
		modifiedSchemas[schema1] = &diff
	}
}