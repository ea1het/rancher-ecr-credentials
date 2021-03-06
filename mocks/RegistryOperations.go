package mocks

import client "github.com/rancher/go-rancher/client"
import mock "github.com/stretchr/testify/mock"

// RegistryOperations is an autogenerated mock type for the RegistryOperations type
type RegistryOperations struct {
	mock.Mock
}

// ActionActivate provides a mock function with given fields: _a0
func (_m *RegistryOperations) ActionActivate(_a0 *client.Registry) (*client.StoragePool, error) {
	ret := _m.Called(_a0)

	var r0 *client.StoragePool
	if rf, ok := ret.Get(0).(func(*client.Registry) *client.StoragePool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.StoragePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionCreate provides a mock function with given fields: _a0
func (_m *RegistryOperations) ActionCreate(_a0 *client.Registry) (*client.StoragePool, error) {
	ret := _m.Called(_a0)

	var r0 *client.StoragePool
	if rf, ok := ret.Get(0).(func(*client.Registry) *client.StoragePool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.StoragePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionDeactivate provides a mock function with given fields: _a0
func (_m *RegistryOperations) ActionDeactivate(_a0 *client.Registry) (*client.StoragePool, error) {
	ret := _m.Called(_a0)

	var r0 *client.StoragePool
	if rf, ok := ret.Get(0).(func(*client.Registry) *client.StoragePool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.StoragePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionPurge provides a mock function with given fields: _a0
func (_m *RegistryOperations) ActionPurge(_a0 *client.Registry) (*client.StoragePool, error) {
	ret := _m.Called(_a0)

	var r0 *client.StoragePool
	if rf, ok := ret.Get(0).(func(*client.Registry) *client.StoragePool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.StoragePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionRemove provides a mock function with given fields: _a0
func (_m *RegistryOperations) ActionRemove(_a0 *client.Registry) (*client.StoragePool, error) {
	ret := _m.Called(_a0)

	var r0 *client.StoragePool
	if rf, ok := ret.Get(0).(func(*client.Registry) *client.StoragePool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.StoragePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionRestore provides a mock function with given fields: _a0
func (_m *RegistryOperations) ActionRestore(_a0 *client.Registry) (*client.StoragePool, error) {
	ret := _m.Called(_a0)

	var r0 *client.StoragePool
	if rf, ok := ret.Get(0).(func(*client.Registry) *client.StoragePool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.StoragePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionUpdate provides a mock function with given fields: _a0
func (_m *RegistryOperations) ActionUpdate(_a0 *client.Registry) (*client.StoragePool, error) {
	ret := _m.Called(_a0)

	var r0 *client.StoragePool
	if rf, ok := ret.Get(0).(func(*client.Registry) *client.StoragePool); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.StoragePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ById provides a mock function with given fields: id
func (_m *RegistryOperations) ById(id string) (*client.Registry, error) {
	ret := _m.Called(id)

	var r0 *client.Registry
	if rf, ok := ret.Get(0).(func(string) *client.Registry); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: opts
func (_m *RegistryOperations) Create(opts *client.Registry) (*client.Registry, error) {
	ret := _m.Called(opts)

	var r0 *client.Registry
	if rf, ok := ret.Get(0).(func(*client.Registry) *client.Registry); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: container
func (_m *RegistryOperations) Delete(container *client.Registry) error {
	ret := _m.Called(container)

	var r0 error
	if rf, ok := ret.Get(0).(func(*client.Registry) error); ok {
		r0 = rf(container)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: opts
func (_m *RegistryOperations) List(opts *client.ListOpts) (*client.RegistryCollection, error) {
	ret := _m.Called(opts)

	var r0 *client.RegistryCollection
	if rf, ok := ret.Get(0).(func(*client.ListOpts) *client.RegistryCollection); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.RegistryCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.ListOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: existing, updates
func (_m *RegistryOperations) Update(existing *client.Registry, updates interface{}) (*client.Registry, error) {
	ret := _m.Called(existing, updates)

	var r0 *client.Registry
	if rf, ok := ret.Get(0).(func(*client.Registry, interface{}) *client.Registry); ok {
		r0 = rf(existing, updates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Registry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.Registry, interface{}) error); ok {
		r1 = rf(existing, updates)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
