// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"compress/gzip"
	"encoding/base64"
	"io"
	"strings"
)

func NewDefaultGenesisReader() (io.Reader, error) {
	return gzip.NewReader(base64.NewDecoder(base64.StdEncoding, strings.NewReader(defaultGenesisBlock)))
}

const defaultGenesisBlock = "H4sIAAAJbogA/6yRz0oDMRDG74Iv4KnsuYf8T8azBw++xCQzsYHdbelGWJG+u9HWFf9BFb9DQvhmfky+ebq8WDV143ZM3F2vOjGLDzKqW59qahl4qjjsTnWLscM9j/UWp813hN9rAfNc93iDFReuW7x7nO7KUOrRkjqExaKSc0kPfX08muYzdyjz5v+nTdsyRpxeYyTDGhBVlgYcO7SAXtjkUyYWrGLOkJON7+Fi329T63xbiJgJITutg/McwBsbjQNjUztkiDp6lXJEQfaladVF7PG0Qv/zqIf1gsf2BmJJgVGRwygNOYPaBkisMRMoSEbCF7w8hy7RCEuWvQBtGlz77FhIQp9lat+wUWsmRX+jn53umfTucES363D1DAAA//8BAAD//waatjUTAwAA"
