// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package extensiondefinition

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Define Object Type
// ----------------------------------------------------------------------

/*
ExtensionDefinition - This type implements the STIX 2 Extension Definition object and
defines all of the properties and methods needed to create and work with this
object. All of the methods not defined local to this type are inherited from the
individual properties.
*/
type ExtensionDefinition struct {
	objects.CommonObjectProperties
	objects.NameProperty
	objects.DescriptionProperty
	Schema              string   `json:"schema,omitempty" bson:"schema,omitempty"`
	Version             string   `json:"version,omitempty" bson:"version,omitempty"`
	ExtensionTypes      []string `json:"extension_types,omitempty" bson:"extension_types,omitempty"`
	ExtensionProperties []string `json:"extension_properties,omitempty" bson:"extension_properties,omitempty"`
}

/*
GetPropertyList - This method will return a list of all of the properties that
are unique to this object. This is used by the custom UnmarshalJSON for this
object. It is defined here in this file to make it easy to keep in sync.
*/
func (o *ExtensionDefinition) GetPropertyList() []string {
	return []string{"name", "description", "schema", "version", "extension_types", "extension_properties"}
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/*
New - This function will create a new STIX Identity object and return it as a
pointer.
*/
func New() *ExtensionDefinition {
	var obj ExtensionDefinition
	obj.InitSDO("extension-definition")
	return &obj
}
