// Copyright 2024 EMQ Technologies Co., Ltd.
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

package modules

import "github.com/lf-edge/ekuiper/v2/pkg/message"

var Converters = map[string]message.ConverterProvider{}

// RegisterConverter registers a converter with the given name.
func RegisterConverter(name string, provider message.ConverterProvider) {
	Converters[name] = provider
}

func IsFormatSupported(format string) bool {
	_, ok := Converters[format]
	return ok
}