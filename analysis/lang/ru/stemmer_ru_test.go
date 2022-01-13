//  Copyright (c) 2018 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ru

import (
	"reflect"
	"testing"

	"github.com/bpizzi/bleve/analysis"
	"github.com/bpizzi/bleve/registry"
)

func TestSnowballRussianStemmer(t *testing.T) {
	tests := []struct {
		input  analysis.TokenStream
		output analysis.TokenStream
	}{
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("актеров"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("актер"),
				},
			},
		},
		{
			input: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("километров"),
				},
			},
			output: analysis.TokenStream{
				&analysis.Token{
					Term: []byte("километр"),
				},
			},
		},
	}

	cache := registry.NewCache()
	filter, err := cache.TokenFilterNamed(SnowballStemmerName)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		actual := filter.Filter(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected %s, got %s", test.output[0].Term, actual[0].Term)
		}
	}
}
