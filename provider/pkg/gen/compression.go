// Copyright 2016-2022, Pulumi Corporation.
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

package gen

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

func CompressSchema(pkgSpec schema.PackageSpec) ([]byte, error) {
	compressedSchema := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedSchema)
	err := json.NewEncoder(compressedWriter).Encode(pkgSpec)
	if err != nil {
		return nil, errors.Wrap(err, "marshaling schema")
	}
	if err = compressedWriter.Close(); err != nil {
		return nil, err
	}
	return compressedSchema.Bytes(), nil
}

func DecompressSchema(compressedSchema []byte) ([]byte, error) {
	uncompressed, err := gzip.NewReader(bytes.NewReader(compressedSchema))
	if err != nil {
		return nil, errors.Wrap(err, "expand compressed schema")
	}
	uncompressedBuf := bytes.Buffer{}
	if _, err = io.CopyN(&uncompressedBuf, uncompressed, 1024); err != nil {
		return nil, err
	}
	if err = uncompressed.Close(); err != nil {
		return nil, errors.Wrap(err, "closing uncompress stream for schema")
	}
	return uncompressedBuf.Bytes(), err
}
