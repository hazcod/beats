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

package netflow

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "netflow", asset.ModuleFieldsPri, AssetNetflow); err != nil {
		panic(err)
	}
}

// AssetNetflow returns asset data.
// This is the base64 encoded gzipped contents of module/netflow.
func AssetNetflow() string {
	return "eJw8jjFOw0AQRfs9xbtAcoAtqFCkFKAUINGazBiPWHas3Ymt3B4Z4fTv/f8OfOs9UzXG4uvhx+VWNEFYFM28apyKrwlE+7XZHOY185QAXv5gRm80vaotVr92g6EK58vp/ME2vAHepOOLNt6fL0feJuVxB+LaqR4MIoymRTqfevcqrNMQxKR7JVbnWzA3X0y0HxP/Qk6/AQAA//9CcUYh"
}
