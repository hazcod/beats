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

package googlecloud

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "googlecloud", asset.ModuleFieldsPri, AssetGooglecloud); err != nil {
		panic(err)
	}
}

// AssetGooglecloud returns asset data.
// This is the base64 encoded gzipped contents of module/googlecloud.
func AssetGooglecloud() string {
	return "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"
}
