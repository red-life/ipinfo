// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	maxminddb "github.com/oschwald/maxminddb-golang"
	mock "github.com/stretchr/testify/mock"

	net "net"

	ports "github.com/red-life/ipinfo/internal/ports"
)

// IMaxMind is an autogenerated mock type for the IMaxMind type
type IMaxMind struct {
	mock.Mock
}

type IMaxMind_Expecter struct {
	mock *mock.Mock
}

func (_m *IMaxMind) EXPECT() *IMaxMind_Expecter {
	return &IMaxMind_Expecter{mock: &_m.Mock}
}

// GetASN provides a mock function with given fields: ip
func (_m *IMaxMind) GetASN(ip net.IP) (ports.ASN, error) {
	ret := _m.Called(ip)

	var r0 ports.ASN
	var r1 error
	if rf, ok := ret.Get(0).(func(net.IP) (ports.ASN, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(net.IP) ports.ASN); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(ports.ASN)
	}

	if rf, ok := ret.Get(1).(func(net.IP) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IMaxMind_GetASN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetASN'
type IMaxMind_GetASN_Call struct {
	*mock.Call
}

// GetASN is a helper method to define mock.On call
//   - ip net.IP
func (_e *IMaxMind_Expecter) GetASN(ip interface{}) *IMaxMind_GetASN_Call {
	return &IMaxMind_GetASN_Call{Call: _e.mock.On("GetASN", ip)}
}

func (_c *IMaxMind_GetASN_Call) Run(run func(ip net.IP)) *IMaxMind_GetASN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.IP))
	})
	return _c
}

func (_c *IMaxMind_GetASN_Call) Return(_a0 ports.ASN, _a1 error) *IMaxMind_GetASN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IMaxMind_GetASN_Call) RunAndReturn(run func(net.IP) (ports.ASN, error)) *IMaxMind_GetASN_Call {
	_c.Call.Return(run)
	return _c
}

// GetCity provides a mock function with given fields: ip
func (_m *IMaxMind) GetCity(ip net.IP) (ports.City, error) {
	ret := _m.Called(ip)

	var r0 ports.City
	var r1 error
	if rf, ok := ret.Get(0).(func(net.IP) (ports.City, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(net.IP) ports.City); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(ports.City)
	}

	if rf, ok := ret.Get(1).(func(net.IP) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IMaxMind_GetCity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCity'
type IMaxMind_GetCity_Call struct {
	*mock.Call
}

// GetCity is a helper method to define mock.On call
//   - ip net.IP
func (_e *IMaxMind_Expecter) GetCity(ip interface{}) *IMaxMind_GetCity_Call {
	return &IMaxMind_GetCity_Call{Call: _e.mock.On("GetCity", ip)}
}

func (_c *IMaxMind_GetCity_Call) Run(run func(ip net.IP)) *IMaxMind_GetCity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.IP))
	})
	return _c
}

func (_c *IMaxMind_GetCity_Call) Return(_a0 ports.City, _a1 error) *IMaxMind_GetCity_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IMaxMind_GetCity_Call) RunAndReturn(run func(net.IP) (ports.City, error)) *IMaxMind_GetCity_Call {
	_c.Call.Return(run)
	return _c
}

// GetContinent provides a mock function with given fields: ip
func (_m *IMaxMind) GetContinent(ip net.IP) (ports.Continent, error) {
	ret := _m.Called(ip)

	var r0 ports.Continent
	var r1 error
	if rf, ok := ret.Get(0).(func(net.IP) (ports.Continent, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(net.IP) ports.Continent); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(ports.Continent)
	}

	if rf, ok := ret.Get(1).(func(net.IP) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IMaxMind_GetContinent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContinent'
type IMaxMind_GetContinent_Call struct {
	*mock.Call
}

// GetContinent is a helper method to define mock.On call
//   - ip net.IP
func (_e *IMaxMind_Expecter) GetContinent(ip interface{}) *IMaxMind_GetContinent_Call {
	return &IMaxMind_GetContinent_Call{Call: _e.mock.On("GetContinent", ip)}
}

func (_c *IMaxMind_GetContinent_Call) Run(run func(ip net.IP)) *IMaxMind_GetContinent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.IP))
	})
	return _c
}

func (_c *IMaxMind_GetContinent_Call) Return(_a0 ports.Continent, _a1 error) *IMaxMind_GetContinent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IMaxMind_GetContinent_Call) RunAndReturn(run func(net.IP) (ports.Continent, error)) *IMaxMind_GetContinent_Call {
	_c.Call.Return(run)
	return _c
}

// GetCountry provides a mock function with given fields: ip
func (_m *IMaxMind) GetCountry(ip net.IP) (ports.Country, error) {
	ret := _m.Called(ip)

	var r0 ports.Country
	var r1 error
	if rf, ok := ret.Get(0).(func(net.IP) (ports.Country, error)); ok {
		return rf(ip)
	}
	if rf, ok := ret.Get(0).(func(net.IP) ports.Country); ok {
		r0 = rf(ip)
	} else {
		r0 = ret.Get(0).(ports.Country)
	}

	if rf, ok := ret.Get(1).(func(net.IP) error); ok {
		r1 = rf(ip)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IMaxMind_GetCountry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCountry'
type IMaxMind_GetCountry_Call struct {
	*mock.Call
}

// GetCountry is a helper method to define mock.On call
//   - ip net.IP
func (_e *IMaxMind_Expecter) GetCountry(ip interface{}) *IMaxMind_GetCountry_Call {
	return &IMaxMind_GetCountry_Call{Call: _e.mock.On("GetCountry", ip)}
}

func (_c *IMaxMind_GetCountry_Call) Run(run func(ip net.IP)) *IMaxMind_GetCountry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.IP))
	})
	return _c
}

func (_c *IMaxMind_GetCountry_Call) Return(_a0 ports.Country, _a1 error) *IMaxMind_GetCountry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IMaxMind_GetCountry_Call) RunAndReturn(run func(net.IP) (ports.Country, error)) *IMaxMind_GetCountry_Call {
	_c.Call.Return(run)
	return _c
}

// Load provides a mock function with given fields: countryReader, cityReader, asnReader
func (_m *IMaxMind) Load(countryReader *maxminddb.Reader, cityReader *maxminddb.Reader, asnReader *maxminddb.Reader) {
	_m.Called(countryReader, cityReader, asnReader)
}

// IMaxMind_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type IMaxMind_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
//   - countryReader *maxminddb.Reader
//   - cityReader *maxminddb.Reader
//   - asnReader *maxminddb.Reader
func (_e *IMaxMind_Expecter) Load(countryReader interface{}, cityReader interface{}, asnReader interface{}) *IMaxMind_Load_Call {
	return &IMaxMind_Load_Call{Call: _e.mock.On("Load", countryReader, cityReader, asnReader)}
}

func (_c *IMaxMind_Load_Call) Run(run func(countryReader *maxminddb.Reader, cityReader *maxminddb.Reader, asnReader *maxminddb.Reader)) *IMaxMind_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*maxminddb.Reader), args[1].(*maxminddb.Reader), args[2].(*maxminddb.Reader))
	})
	return _c
}

func (_c *IMaxMind_Load_Call) Return() *IMaxMind_Load_Call {
	_c.Call.Return()
	return _c
}

func (_c *IMaxMind_Load_Call) RunAndReturn(run func(*maxminddb.Reader, *maxminddb.Reader, *maxminddb.Reader)) *IMaxMind_Load_Call {
	_c.Call.Return(run)
	return _c
}

// NewIMaxMind creates a new instance of IMaxMind. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMaxMind(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMaxMind {
	mock := &IMaxMind{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}