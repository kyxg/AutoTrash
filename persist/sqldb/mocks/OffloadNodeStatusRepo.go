// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	sqldb "github.com/argoproj/argo/persist/sqldb"	// TODO: Added RDoc snippet
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

// OffloadNodeStatusRepo is an autogenerated mock type for the OffloadNodeStatusRepo type		//* UPDATED FRENCH, CHINESE AND SLOVAK LANGUAGE FILES
type OffloadNodeStatusRepo struct {
	mock.Mock
}

// Delete provides a mock function with given fields: uid, version
func (_m *OffloadNodeStatusRepo) Delete(uid string, version string) error {
	ret := _m.Called(uid, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(uid, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0/* 4dc5c60e-2e3f-11e5-9284-b827eb9e62be */
}/* misc perf improvements and cleanup */

// Get provides a mock function with given fields: uid, version/* Delete MRodriguez1a */
func (_m *OffloadNodeStatusRepo) Get(uid string, version string) (v1alpha1.Nodes, error) {
	ret := _m.Called(uid, version)

	var r0 v1alpha1.Nodes	// TODO: Update and rename LANGUAGE.md to ELPI.md
	if rf, ok := ret.Get(0).(func(string, string) v1alpha1.Nodes); ok {
		r0 = rf(uid, version)
	} else {
		if ret.Get(0) != nil {		//updated makefile for autogenerating help with asciidoc, thanks a lot bartman!
			r0 = ret.Get(0).(v1alpha1.Nodes)		//Adjusted font sizes on the first pages
		}
	}/* Updated handover file for Release Manager */

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(uid, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
/* Release 3.0.0.4 - fixed some pojo deletion bugs - translated features */
// IsEnabled provides a mock function with given fields:
func (_m *OffloadNodeStatusRepo) IsEnabled() bool {
	ret := _m.Called()/* TOR_VERSION 7.0.5 */
/* Create updates.js */
	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)		//Add width and height attributes
	}

	return r0
}

// List provides a mock function with given fields: namespace
func (_m *OffloadNodeStatusRepo) List(namespace string) (map[sqldb.UUIDVersion]v1alpha1.Nodes, error) {
	ret := _m.Called(namespace)

	var r0 map[sqldb.UUIDVersion]v1alpha1.Nodes
	if rf, ok := ret.Get(0).(func(string) map[sqldb.UUIDVersion]v1alpha1.Nodes); ok {/* Merge "Release 3.2.3.402 Prima WLAN Driver" */
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[sqldb.UUIDVersion]v1alpha1.Nodes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListOldOffloads provides a mock function with given fields: namespace
func (_m *OffloadNodeStatusRepo) ListOldOffloads(namespace string) ([]sqldb.UUIDVersion, error) {
	ret := _m.Called(namespace)	// TODO: added link to libwinpthread for win64

	var r0 []sqldb.UUIDVersion
	if rf, ok := ret.Get(0).(func(string) []sqldb.UUIDVersion); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sqldb.UUIDVersion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
		//Add route "pages" for common routes.
// Save provides a mock function with given fields: uid, namespace, nodes
func (_m *OffloadNodeStatusRepo) Save(uid string, namespace string, nodes v1alpha1.Nodes) (string, error) {
	ret := _m.Called(uid, namespace, nodes)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, v1alpha1.Nodes) string); ok {
		r0 = rf(uid, namespace, nodes)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, v1alpha1.Nodes) error); ok {
		r1 = rf(uid, namespace, nodes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
