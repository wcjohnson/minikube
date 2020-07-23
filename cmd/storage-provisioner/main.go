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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"k8s.io/minikube/pkg/storage"
)

var defaultPvDir = "/tmp/hostpath-provisioner"

func main() {
	// Glog requires that /tmp exists.
	if err := os.MkdirAll("/tmp", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating tmpdir: %v\n", err)
		os.Exit(1)
	}

	var pvDir string
	flag.StringVar(&pvDir, "pvdir", defaultPvDir, "Path where persistent volumes should be stored.")
	flag.Parse()

	if err := storage.StartStorageProvisioner(pvDir); err != nil {
		glog.Exit(err)
	}

}
