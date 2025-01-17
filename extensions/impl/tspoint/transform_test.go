// Copyright 2023-2024 EMQ Technologies Co., Ltd.
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

package tspoint

import (
	"testing"

	"github.com/stretchr/testify/assert"

	mockContext "github.com/lf-edge/ekuiper/v2/pkg/mock/context"
)

func Test_parseTemplates(t *testing.T) {
	tests := []struct {
		name string
		conf WriteOptions
		err  string
	}{
		{
			name: "normal",
			conf: WriteOptions{
				Tags: map[string]string{
					"tag1": "value1",
				},
			},
		},
		{
			name: "normal with template",
			conf: WriteOptions{
				Tags: map[string]string{
					"tag1": "value1",
					"tag2": "{{.temperature}}",
					"tag3": "100",
				},
			},
		},
		{
			name: "error template",
			conf: WriteOptions{
				Tags: map[string]string{
					"tag1": "value1",
					"tag2": "{{abc .temperature}}",
					"tag3": "100",
				},
			},
			err: "Template Invalid: template: sink:1: function \"abc\" not defined",
		},
		{
			name: "prop template",
			conf: WriteOptions{
				Tags: map[string]string{
					"tag1": "value1",
					"tag2": "{{prop \"et\"}}",
					"tag3": "100",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := mockContext.NewMockContext("parseTemplate", "op")
			err := tt.conf.ValidateTagTemplates(ctx)
			if tt.err == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err.Error())
			}
		})
	}
}
