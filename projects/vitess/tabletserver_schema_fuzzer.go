// Copyright 2022 ADA Logics Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package schema

import (
	"fmt"
	"testing"

	"vitess.io/vitess/go/sqltypes"
	"vitess.io/vitess/go/mysql/fakesqldb"
)


func FuzzLoadTable(f *testing.F) {
	f.Fuzz(func(t *testing.T, comment, query, tableType string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
			}
		}()
		db := fakesqldb.New(t)
		defer db.Close()
		db.AddQuery(query, &sqltypes.Result{})
		_, _ = newTestLoadTable(tableType, comment, db)
	})
	
}