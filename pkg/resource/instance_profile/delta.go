// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package instance_profile

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	compareTags(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Path, b.ko.Spec.Path) {
		delta.Add("Spec.Path", a.ko.Spec.Path, b.ko.Spec.Path)
	} else if a.ko.Spec.Path != nil && b.ko.Spec.Path != nil {
		if *a.ko.Spec.Path != *b.ko.Spec.Path {
			delta.Add("Spec.Path", a.ko.Spec.Path, b.ko.Spec.Path)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Role, b.ko.Spec.Role) {
		delta.Add("Spec.Role", a.ko.Spec.Role, b.ko.Spec.Role)
	} else if a.ko.Spec.Role != nil && b.ko.Spec.Role != nil {
		if *a.ko.Spec.Role != *b.ko.Spec.Role {
			delta.Add("Spec.Role", a.ko.Spec.Role, b.ko.Spec.Role)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.RoleRef, b.ko.Spec.RoleRef) {
		delta.Add("Spec.RoleRef", a.ko.Spec.RoleRef, b.ko.Spec.RoleRef)
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}
