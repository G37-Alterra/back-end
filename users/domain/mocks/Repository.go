// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	domain_users "ppob/users/domain"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ChangeStatusUsers provides a mock function with given fields: email
func (_m *Repository) ChangeStatusUsers(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeStatusVerif provides a mock function with given fields: email
func (_m *Repository) ChangeStatusVerif(email string) error {
	ret := _m.Called(email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckEmailPassword provides a mock function with given fields: email, password
func (_m *Repository) CheckEmailPassword(email string, password string) (domain_users.Users, error) {
	ret := _m.Called(email, password)

	var r0 domain_users.Users
	if rf, ok := ret.Get(0).(func(string, string) domain_users.Users); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(domain_users.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUser provides a mock function with given fields:
func (_m *Repository) GetAllUser() ([]domain_users.Users, error) {
	ret := _m.Called()

	var r0 []domain_users.Users
	if rf, ok := ret.Get(0).(func() []domain_users.Users); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain_users.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPhone provides a mock function with given fields: phone
func (_m *Repository) GetByPhone(phone string) (domain_users.Users, error) {
	ret := _m.Called(phone)

	var r0 domain_users.Users
	if rf, ok := ret.Get(0).(func(string) domain_users.Users); ok {
		r0 = rf(phone)
	} else {
		r0 = ret.Get(0).(domain_users.Users)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserAccount provides a mock function with given fields: phone
func (_m *Repository) GetUserAccount(phone string) (domain_users.Account, error) {
	ret := _m.Called(phone)

	var r0 domain_users.Account
	if rf, ok := ret.Get(0).(func(string) domain_users.Account); ok {
		r0 = rf(phone)
	} else {
		r0 = ret.Get(0).(domain_users.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: domain
func (_m *Repository) Store(domain domain_users.Users) (string, error) {
	ret := _m.Called(domain)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain_users.Users) string); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain_users.Users) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreAccount provides a mock function with given fields: domain
func (_m *Repository) StoreAccount(domain domain_users.Account) (domain_users.Account, error) {
	ret := _m.Called(domain)

	var r0 domain_users.Account
	if rf, ok := ret.Get(0).(func(domain_users.Account) domain_users.Account); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(domain_users.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain_users.Account) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreOtpUserVerif provides a mock function with given fields: code, email
func (_m *Repository) StoreOtpUserVerif(code string, email string) error {
	ret := _m.Called(code, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(code, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: phone, domain
func (_m *Repository) Update(phone string, domain domain_users.Users) error {
	ret := _m.Called(phone, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain_users.Users) error); ok {
		r0 = rf(phone, domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verif provides a mock function with given fields: code
func (_m *Repository) Verif(code string) (domain_users.UserVerif, error) {
	ret := _m.Called(code)

	var r0 domain_users.UserVerif
	if rf, ok := ret.Get(0).(func(string) domain_users.UserVerif); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Get(0).(domain_users.UserVerif)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t testing.TB) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
