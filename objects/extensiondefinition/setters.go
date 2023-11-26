// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package extensiondefinition

import (
	"github.com/freetaxii/libstix2/objects"
)

// ----------------------------------------------------------------------
// Public Methods - Identity
// ----------------------------------------------------------------------

/*
SetSchema - This method takes in a string value representing an extension schema.

STIX description:
The normative definition of the extension, either as a URL or as plain text explaining the definition.
A URL SHOULD point to a JSON schema or a location that contains information about the schema.
*/
func (o *ExtensionDefinition) SetSchema(s string) error {
	o.Schema = s
	return nil
}

/*
SetVersion - This method takes in a string value representing an extension version.

STIX description:
The version of this extension. Producers of STIX extensions are encouraged to follow standard
semantic versioning procedures where the version number follows the pattern, MAJOR.MINOR.PATCH.
This will allow consumers to distinguish between the three different levels of compatibility
typically identified by such versioning strings.
*/
func (o *ExtensionDefinition) SetVersion(v string) error {
	o.Version = v
	return nil
}

/*
AddExtensionTypes - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an extension type and adds it
to the extension_types property.
*/
func (o *ExtensionDefinition) AddExtensionTypes(values interface{}) error {
	return objects.AddValuesToList(&o.ExtensionTypes, values)
}

/*
AddExtensionProperties - This method takes in a string value, a comma separated list of
string values, or a slice of string values that represents an extension type and adds it
to the extension_properties property.
*/
func (o *ExtensionDefinition) AddExtensionProperties(values interface{}) error {
	return objects.AddValuesToList(&o.ExtensionProperties, values)
}
