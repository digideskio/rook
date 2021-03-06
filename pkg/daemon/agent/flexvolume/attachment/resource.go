/*
Copyright 2017 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Some of the code below came from https://github.com/coreos/etcd-operator
which also has the apache 2.0 license.
*/

// Package attachment to manage Kubernetes storage attach events.
package attachment

import (
	"reflect"

	opkit "github.com/rook/operator-kit"
	rookalpha "github.com/rook/rook/pkg/apis/rook.io/v1alpha1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

const (
	CustomResourceName       = "volumeattachment"
	CustomResourceNamePlural = "volumeattachments"
)

// VolumeAttachmentResource represents the VolumeAttachment custom resource object
var VolumeAttachmentResource = opkit.CustomResource{
	Name:    CustomResourceName,
	Plural:  CustomResourceNamePlural,
	Group:   rookalpha.CustomResourceGroup,
	Version: rookalpha.Version,
	Scope:   apiextensionsv1beta1.NamespaceScoped,
	Kind:    reflect.TypeOf(rookalpha.VolumeAttachment{}).Name(),
}
