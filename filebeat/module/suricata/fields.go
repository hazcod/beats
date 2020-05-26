// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package suricata

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "suricata", asset.ModuleFieldsPri, AssetSuricata); err != nil {
		panic(err)
	}
}

// AssetSuricata returns asset data.
// This is the base64 encoded gzipped contents of module/suricata.
func AssetSuricata() string {
	return "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"
}
