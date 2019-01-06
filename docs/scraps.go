
// getSelectedFields reads the requested fields in Resolve.
//
// Invoke it like this:
// 		getSelectedFields([]string{"companies"}, resolveParams)
// Returns -> []string{"id", "name"}
//
// In case you have a "path" you want to select from, e.g.
// 		query { a { b { x, y, z }}}
//
// Then you'd call it like this:
// 		getSelectedFields([]string{"a", "b"}, resolveParams)
// Returns []string{"x", "y", "z"}
//
// Source: https://github.com/graphql-go/graphql/issues/125

// func getSelectedFields(selectionPath []string, params graphql.ResolveParams) []string {
// 	fields := params.Info.FieldASTs
// 	for _, propName := range selectionPath {
// 		found := false
// 		for _, field := range fields {
// 			if field.Name.Value == propName {
// 				selections := field.SelectionSet.Selections

// 				fields = make([]*ast.Field, 0)
// 				for _, selection := range selections {
// 					fields = append(fields, selection.(*ast.Field))
// 				}
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			return []string{}
// 		}
// 	}
// 	var collect []string
// 	for _, field := range fields {
// 		collect = append(collect, field.Name.Value)
// 	}
// 	return collect
// }
