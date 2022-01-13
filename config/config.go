//  Copyright (c) 2015 Couchbase, Inc.
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

package config

import (
	// token maps
	_ "github.com/bpizzi/bleve/analysis/tokenmap"

	// fragment formatters
	_ "github.com/bpizzi/bleve/search/highlight/format/ansi"
	_ "github.com/bpizzi/bleve/search/highlight/format/html"

	// fragmenters
	_ "github.com/bpizzi/bleve/search/highlight/fragmenter/simple"

	// highlighters
	_ "github.com/bpizzi/bleve/search/highlight/highlighter/ansi"
	_ "github.com/bpizzi/bleve/search/highlight/highlighter/html"
	_ "github.com/bpizzi/bleve/search/highlight/highlighter/simple"

	// char filters
	_ "github.com/bpizzi/bleve/analysis/char/asciifolding"
	_ "github.com/bpizzi/bleve/analysis/char/html"
	_ "github.com/bpizzi/bleve/analysis/char/regexp"
	_ "github.com/bpizzi/bleve/analysis/char/zerowidthnonjoiner"

	// analyzers
	_ "github.com/bpizzi/bleve/analysis/analyzer/custom"
	_ "github.com/bpizzi/bleve/analysis/analyzer/keyword"
	_ "github.com/bpizzi/bleve/analysis/analyzer/simple"
	_ "github.com/bpizzi/bleve/analysis/analyzer/standard"
	_ "github.com/bpizzi/bleve/analysis/analyzer/web"

	// token filters
	_ "github.com/bpizzi/bleve/analysis/token/apostrophe"
	_ "github.com/bpizzi/bleve/analysis/token/camelcase"
	_ "github.com/bpizzi/bleve/analysis/token/compound"
	_ "github.com/bpizzi/bleve/analysis/token/edgengram"
	_ "github.com/bpizzi/bleve/analysis/token/elision"
	_ "github.com/bpizzi/bleve/analysis/token/keyword"
	_ "github.com/bpizzi/bleve/analysis/token/length"
	_ "github.com/bpizzi/bleve/analysis/token/lowercase"
	_ "github.com/bpizzi/bleve/analysis/token/ngram"
	_ "github.com/bpizzi/bleve/analysis/token/reverse"
	_ "github.com/bpizzi/bleve/analysis/token/shingle"
	_ "github.com/bpizzi/bleve/analysis/token/stop"
	_ "github.com/bpizzi/bleve/analysis/token/truncate"
	_ "github.com/bpizzi/bleve/analysis/token/unicodenorm"
	_ "github.com/bpizzi/bleve/analysis/token/unique"

	// tokenizers
	_ "github.com/bpizzi/bleve/analysis/tokenizer/exception"
	_ "github.com/bpizzi/bleve/analysis/tokenizer/regexp"
	_ "github.com/bpizzi/bleve/analysis/tokenizer/single"
	_ "github.com/bpizzi/bleve/analysis/tokenizer/unicode"
	_ "github.com/bpizzi/bleve/analysis/tokenizer/web"
	_ "github.com/bpizzi/bleve/analysis/tokenizer/whitespace"

	// date time parsers
	_ "github.com/bpizzi/bleve/analysis/datetime/flexible"
	_ "github.com/bpizzi/bleve/analysis/datetime/optional"

	// languages
	_ "github.com/bpizzi/bleve/analysis/lang/ar"
	_ "github.com/bpizzi/bleve/analysis/lang/bg"
	_ "github.com/bpizzi/bleve/analysis/lang/ca"
	_ "github.com/bpizzi/bleve/analysis/lang/cjk"
	_ "github.com/bpizzi/bleve/analysis/lang/ckb"
	_ "github.com/bpizzi/bleve/analysis/lang/cs"
	_ "github.com/bpizzi/bleve/analysis/lang/da"
	_ "github.com/bpizzi/bleve/analysis/lang/de"
	_ "github.com/bpizzi/bleve/analysis/lang/el"
	_ "github.com/bpizzi/bleve/analysis/lang/en"
	_ "github.com/bpizzi/bleve/analysis/lang/es"
	_ "github.com/bpizzi/bleve/analysis/lang/eu"
	_ "github.com/bpizzi/bleve/analysis/lang/fa"
	_ "github.com/bpizzi/bleve/analysis/lang/fi"
	_ "github.com/bpizzi/bleve/analysis/lang/fr"
	_ "github.com/bpizzi/bleve/analysis/lang/ga"
	_ "github.com/bpizzi/bleve/analysis/lang/gl"
	_ "github.com/bpizzi/bleve/analysis/lang/hi"
	_ "github.com/bpizzi/bleve/analysis/lang/hu"
	_ "github.com/bpizzi/bleve/analysis/lang/hy"
	_ "github.com/bpizzi/bleve/analysis/lang/id"
	_ "github.com/bpizzi/bleve/analysis/lang/in"
	_ "github.com/bpizzi/bleve/analysis/lang/it"
	_ "github.com/bpizzi/bleve/analysis/lang/nl"
	_ "github.com/bpizzi/bleve/analysis/lang/no"
	_ "github.com/bpizzi/bleve/analysis/lang/pt"
	_ "github.com/bpizzi/bleve/analysis/lang/ro"
	_ "github.com/bpizzi/bleve/analysis/lang/ru"
	_ "github.com/bpizzi/bleve/analysis/lang/sv"
	_ "github.com/bpizzi/bleve/analysis/lang/tr"

	// kv stores
	_ "github.com/bpizzi/bleve/index/store/boltdb"
	_ "github.com/bpizzi/bleve/index/store/goleveldb"
	_ "github.com/bpizzi/bleve/index/store/gtreap"
	_ "github.com/bpizzi/bleve/index/store/moss"

	// index types
	_ "github.com/bpizzi/bleve/index/upsidedown"
)
