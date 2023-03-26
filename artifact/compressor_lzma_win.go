// Copyright 2021 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package artifact

import (
	"io"

	"github.com/ulikunitz/xz"
)

type CompressorLzmaWin struct {
}

func NewCompressorLzmaWin() Compressor {
	return &CompressorLzmaWin{}
}

func (c *CompressorLzmaWin) GetFileExtension() string {
	return ".xz"
}

func (c *CompressorLzmaWin) NewReader(r io.Reader) (io.ReadCloser, error) {
	rc, err := xz.NewReader(r)
	if err != nil {
		return nil, err
	}
	return io.NopCloser(rc), nil
}

func (c *CompressorLzmaWin) NewWriter(w io.Writer) (io.WriteCloser, error) {
	return xz.NewWriter(w)
}

func init() {
	RegisterCompressor("lzma", &CompressorLzmaWin{})
}
