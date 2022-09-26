// Package kubernetes
// Copyright © 2022 Ajay K <ajaykemparaj@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	storagev1 "k8s.io/api/storage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"log"
)

func CheckStorageClasses(clusterConfig model.Drifter, client kubernetes.Interface, ctx context.Context) bool {
	hasDrifts := false
	if len(clusterConfig.Kubernetes.Storage.StorageClasses) > 0 {
		scList, err := client.StorageV1().StorageClasses().List(ctx, v1.ListOptions{})
		if err != nil {
			log.Fatal("Unable to get storage classes ", err)
		}
		scInstalledMap := make(map[string]storagev1.StorageClass)
		for _, sc := range scList.Items {
			scInstalledMap[sc.Name] = sc
		}

		for _, expectSc := range clusterConfig.Kubernetes.Storage.StorageClasses {
			if _, ok := scInstalledMap[expectSc]; ok {
			} else {
				log.Printf("Missing storage class: %s\n", expectSc)
				hasDrifts = true
			}
		}
	}
	return hasDrifts
}
