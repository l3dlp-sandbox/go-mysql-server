// Copyright 2020-2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plan

import (
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
)

// ApplyBindings replaces all `BindVar` expressions in the given sql.Node with
// their corresponding sql.Expression entries in the provided |bindings| map.
// If a binding for a |BindVar| expression is not found in the map, no error is
// returned and the |BindVar| expression is left in place. There is no check on
// whether all entries in |bindings| are used at least once throughout the |n|.
func ApplyBindings(n sql.Node, bindings map[string]sql.Expression) (sql.Node, error) {
	fixBindings := func(expr sql.Expression) (sql.Expression, error) {
		switch e := expr.(type) {
		case *expression.BindVar:
			val, found := bindings[e.Name]
			if found {
				return val, nil
			}
		case *expression.GetField:
			t, ok := e.Type().(sql.DeferredType)
			if !ok {
				return expr, nil
			}
			val, found := bindings[t.Name()]
			if !found {
				return expr, nil
			}
			return expression.NewGetFieldWithTable(e.Index(), val.Type(), e.Table(), e.Name(), val.IsNullable()), nil
		}
		return expr, nil
	}

	return TransformUpWithOpaque(n, func(node sql.Node) (sql.Node, error) {
		switch n := node.(type) {
		case *IndexedJoin:
			cond, err := expression.TransformUp(n.Cond, fixBindings)
			if err != nil {
				return nil, err
			}
			return NewIndexedJoin(n.left, n.right, n.joinType, cond, n.scopeLen), nil
		}
		return TransformExpressionsUp(node, fixBindings)
	})
}