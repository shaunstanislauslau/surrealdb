// Copyright © 2016 Abcum Ltd
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

package fncs

import (
	"context"

	"github.com/abcum/surreal/sql"
	"github.com/abcum/surreal/util/data"
)

func get(ctx context.Context, args ...interface{}) (interface{}, error) {
	doc := data.Consume(args[0])
	switch fld := args[1].(type) {
	case *sql.Value:
		return doc.Get(fld.ID).Data(), nil
	case string:
		return doc.Get(fld).Data(), nil
	}
	return nil, nil
}