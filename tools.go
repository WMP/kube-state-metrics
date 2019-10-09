// +build tools

/*
Copyright 2019 The Kubernetes Authors All rights reserved.

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

package tools

import (
    _ "github.com/campoy/embedmd"
    _ "golang.org/x/tools/cmd/benchcmp"
    _ "github.com/google/go-jsonnet/cmd/jsonnet"
    _ "github.com/jsonnet-bundler/jsonnet-bundler/cmd/jb"
)
