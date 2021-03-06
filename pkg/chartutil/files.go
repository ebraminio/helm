/*
Copyright 2016 The Kubernetes Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package chartutil

import (
	"github.com/golang/protobuf/ptypes/any"
)

// Files is a map of files in a chart that can be accessed from a template.
type Files map[string][]byte

// NewFiles creates a new Files from chart files.
// Given an []*any.Any (the format for files in a chart.Chart), extract a map of files.
func NewFiles(from []*any.Any) Files {
	files := map[string][]byte{}
	for _, f := range from {
		files[f.TypeUrl] = f.Value
	}
	return files
}

// Get a file by path.
//
// This is intended to be accessed from within a template, so a missed key returns
// an empty []byte.
func (f Files) Get(name string) []byte {
	v, ok := f[name]
	if !ok {
		return []byte{}
	}
	return v
}

// GetString returns a string representation of the given file.
//
// This is a convenience for the otherwise cumbersome template logic
// for '{{.Files.Get "foo" | printf "%s"}}'.
func (f Files) GetString(name string) string {
	return string(f.Get(name))
}
