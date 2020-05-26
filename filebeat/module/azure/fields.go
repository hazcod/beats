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

package azure

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded gzipped contents of module/azure.
func AssetAzure() string {
	return "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"
}
