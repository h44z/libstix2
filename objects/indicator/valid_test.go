// Copyright 2015-2022 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import "testing"

// ----------------------------------------------------------------------
// Tests
// These tests will not use the setters as some setters will have their
// own logic and verification steps in them.
// ----------------------------------------------------------------------

// ----------------------------------------------------------------------
// Test the public Valid method - Make sure we hit each level and make sure
// required property checks are working when they are left blank.
// ----------------------------------------------------------------------

/*
TestValid1 - Make sure we get a value of false when Indicator Type is blank.
*/
func TestValid1(t *testing.T) {
	i := New()
	want := false

	if got, err, _ := i.Valid(false); got != want {
		t.Error("Fail Indicator Type Check 0")
		t.Log(err)
	}
}

/*
TestValid2 - Make sure we get a value of false when Pattern is blank.
*/
func TestValid2(t *testing.T) {
	i := New()
	want := false
	// Set the Indicator Type value so we can move to next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")

	if got, err, _ := i.Valid(false); got != want {
		t.Error("Fail Pattern Check 0")
		t.Log(err)
	}
}

/*
TestValid3 - Make sure we get a value of false when Pattern Type is blank.
*/
func TestValid3(t *testing.T) {
	i := New()
	want := false
	// Set the Indicator Type and Pattern value so we can move to the next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"

	if got, err, _ := i.Valid(false); got != want {
		t.Error("Fail Pattern Type Check 0")
		t.Log(err)
	}
}

/*
TestValid4 - Make sure we get a value of false when Valid From is blank.
*/
func TestValid4(t *testing.T) {
	i := New()
	want := false
	// Set the Indicator Type, Pattern, and Pattern Type value so we can move to
	// the next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"
	i.PatternType = "stix"

	if got, err, _ := i.Valid(false); got != want {
		t.Error("Fail Valid From Check 0")
		t.Log(err)
	}
}

/*
TestValid5 - Make sure we get a value of false when Valid Until is invalid.
*/
func TestValid5(t *testing.T) {
	i := New()
	want := false

	// Set the Indicator Type, Pattern, and Pattern Type value so we can move to.
	// the next test
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"
	i.PatternType = "stix"

	// We need the next test for Valid Until to fail so lets set a bad time.
	i.ValidUntil = "2019-0924T20:49:12.123456Z"

	if got, err, _ := i.Valid(false); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
	}
}

/*
TestValid6 - Make sure we get a value of false when Valid Until is before Valid
From.
*/
func TestValid6(t *testing.T) {
	i := New()
	want := false

	// Set the Indicator Type, Pattern, and Pattern Type value so we can move to
	// the next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"
	i.PatternType = "stix"

	// We need the next test for Valid Until to fail so lets set a bad time.
	i.ValidFrom = "2019-09-24T20:49:13.123456Z"
	i.ValidUntil = "2019-09-24T20:49:12.123456Z"

	if got, err, _ := i.Valid(false); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
	}
}

/*
TestValid7 - Make sure we get a value of true when everything is filled out
correctly.
*/
func TestValid7(t *testing.T) {
	i := New()
	want := false

	// Set the Indicator Type, Pattern, and Pattern Type value so we can move to
	// the next test.
	i.IndicatorTypes = append(i.IndicatorTypes, "TestValue")
	i.Pattern = "TestPattern"
	i.PatternType = "stix"

	// Set the timestamps correctly
	i.ValidFrom = "2019-09-24T20:49:12.123456Z"
	i.ValidUntil = "2019-09-24T20:49:13.123456Z"

	if got, err, details := i.Valid(false); got != want {
		t.Error("Fail Valid Until Check 0")
		t.Log(err)
		t.Log(details)
	}
}
