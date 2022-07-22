// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloud/metadata_interface.go

// Package cloud is a generated GoMock package.
package cloud

import (
	reflect "reflect"

	arn "github.com/aws/aws-sdk-go/aws/arn"
	ec2metadata "github.com/aws/aws-sdk-go/aws/ec2metadata"
	gomock "github.com/golang/mock/gomock"
)

// MockMetadataService is a mock of MetadataService interface.
type MockMetadataService struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataServiceMockRecorder
}

// MockMetadataServiceMockRecorder is the mock recorder for MockMetadataService.
type MockMetadataServiceMockRecorder struct {
	mock *MockMetadataService
}

// NewMockMetadataService creates a new mock instance.
func NewMockMetadataService(ctrl *gomock.Controller) *MockMetadataService {
	mock := &MockMetadataService{ctrl: ctrl}
	mock.recorder = &MockMetadataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataService) EXPECT() *MockMetadataServiceMockRecorder {
	return m.recorder
}

// GetAvailabilityZone mocks base method.
func (m *MockMetadataService) GetAvailabilityZone() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailabilityZone")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAvailabilityZone indicates an expected call of GetAvailabilityZone.
func (mr *MockMetadataServiceMockRecorder) GetAvailabilityZone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailabilityZone", reflect.TypeOf((*MockMetadataService)(nil).GetAvailabilityZone))
}

// GetInstanceID mocks base method.
func (m *MockMetadataService) GetInstanceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInstanceID indicates an expected call of GetInstanceID.
func (mr *MockMetadataServiceMockRecorder) GetInstanceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceID", reflect.TypeOf((*MockMetadataService)(nil).GetInstanceID))
}

// GetInstanceType mocks base method.
func (m *MockMetadataService) GetInstanceType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInstanceType indicates an expected call of GetInstanceType.
func (mr *MockMetadataServiceMockRecorder) GetInstanceType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceType", reflect.TypeOf((*MockMetadataService)(nil).GetInstanceType))
}

// GetOutpostArn mocks base method.
func (m *MockMetadataService) GetOutpostArn() arn.ARN {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutpostArn")
	ret0, _ := ret[0].(arn.ARN)
	return ret0
}

// GetOutpostArn indicates an expected call of GetOutpostArn.
func (mr *MockMetadataServiceMockRecorder) GetOutpostArn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutpostArn", reflect.TypeOf((*MockMetadataService)(nil).GetOutpostArn))
}

// GetRegion mocks base method.
func (m *MockMetadataService) GetRegion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRegion indicates an expected call of GetRegion.
func (mr *MockMetadataServiceMockRecorder) GetRegion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegion", reflect.TypeOf((*MockMetadataService)(nil).GetRegion))
}

// MockEC2Metadata is a mock of EC2Metadata interface.
type MockEC2Metadata struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MetadataMockRecorder
}

// MockEC2MetadataMockRecorder is the mock recorder for MockEC2Metadata.
type MockEC2MetadataMockRecorder struct {
	mock *MockEC2Metadata
}

// NewMockEC2Metadata creates a new mock instance.
func NewMockEC2Metadata(ctrl *gomock.Controller) *MockEC2Metadata {
	mock := &MockEC2Metadata{ctrl: ctrl}
	mock.recorder = &MockEC2MetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2Metadata) EXPECT() *MockEC2MetadataMockRecorder {
	return m.recorder
}

// Available mocks base method.
func (m *MockEC2Metadata) Available() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Available")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Available indicates an expected call of Available.
func (mr *MockEC2MetadataMockRecorder) Available() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Available", reflect.TypeOf((*MockEC2Metadata)(nil).Available))
}

// GetInstanceIdentityDocument mocks base method.
func (m *MockEC2Metadata) GetInstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceIdentityDocument")
	ret0, _ := ret[0].(ec2metadata.EC2InstanceIdentityDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceIdentityDocument indicates an expected call of GetInstanceIdentityDocument.
func (mr *MockEC2MetadataMockRecorder) GetInstanceIdentityDocument() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceIdentityDocument", reflect.TypeOf((*MockEC2Metadata)(nil).GetInstanceIdentityDocument))
}

// GetMetadata mocks base method.
func (m *MockEC2Metadata) GetMetadata(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockEC2MetadataMockRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockEC2Metadata)(nil).GetMetadata), arg0)
}
