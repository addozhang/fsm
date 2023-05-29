// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/flomesh-io/fsm/pkg/catalog (interfaces: MeshCataloger)

// Package catalog is a generated GoMock package.
package catalog

import (
	reflect "reflect"

	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/policy/v1alpha1"
	endpoint "github.com/flomesh-io/fsm/pkg/endpoint"
	identity "github.com/flomesh-io/fsm/pkg/identity"
	k8s "github.com/flomesh-io/fsm/pkg/k8s"
	service "github.com/flomesh-io/fsm/pkg/service"
	trafficpolicy "github.com/flomesh-io/fsm/pkg/trafficpolicy"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockMeshCataloger is a mock of MeshCataloger interface.
type MockMeshCataloger struct {
	ctrl     *gomock.Controller
	recorder *MockMeshCatalogerMockRecorder
}

// MockMeshCatalogerMockRecorder is the mock recorder for MockMeshCataloger.
type MockMeshCatalogerMockRecorder struct {
	mock *MockMeshCataloger
}

// NewMockMeshCataloger creates a new mock instance.
func NewMockMeshCataloger(ctrl *gomock.Controller) *MockMeshCataloger {
	mock := &MockMeshCataloger{ctrl: ctrl}
	mock.recorder = &MockMeshCatalogerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshCataloger) EXPECT() *MockMeshCatalogerMockRecorder {
	return m.recorder
}

// GetAccessControlTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetAccessControlTrafficPolicy(arg0 service.MeshService) (*trafficpolicy.AccessControlTrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessControlTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.AccessControlTrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessControlTrafficPolicy indicates an expected call of GetAccessControlTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetAccessControlTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessControlTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetAccessControlTrafficPolicy), arg0)
}

// GetEgressGatewayPolicy mocks base method.
func (m *MockMeshCataloger) GetEgressGatewayPolicy() (*trafficpolicy.EgressGatewayPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEgressGatewayPolicy")
	ret0, _ := ret[0].(*trafficpolicy.EgressGatewayPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEgressGatewayPolicy indicates an expected call of GetEgressGatewayPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetEgressGatewayPolicy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEgressGatewayPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetEgressGatewayPolicy))
}

// GetEgressSourceSecret mocks base method.
func (m *MockMeshCataloger) GetEgressSourceSecret(arg0 v1.SecretReference) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEgressSourceSecret", arg0)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEgressSourceSecret indicates an expected call of GetEgressSourceSecret.
func (mr *MockMeshCatalogerMockRecorder) GetEgressSourceSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEgressSourceSecret", reflect.TypeOf((*MockMeshCataloger)(nil).GetEgressSourceSecret), arg0)
}

// GetEgressTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetEgressTrafficPolicy(arg0 identity.ServiceIdentity) (*trafficpolicy.EgressTrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEgressTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.EgressTrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEgressTrafficPolicy indicates an expected call of GetEgressTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetEgressTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEgressTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetEgressTrafficPolicy), arg0)
}

// GetExportTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetExportTrafficPolicy(arg0 service.MeshService) (*trafficpolicy.ServiceExportTrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExportTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.ServiceExportTrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExportTrafficPolicy indicates an expected call of GetExportTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetExportTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExportTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetExportTrafficPolicy), arg0)
}

// GetInboundMeshTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetInboundMeshTrafficPolicy(arg0 identity.ServiceIdentity, arg1 []service.MeshService) *trafficpolicy.InboundMeshTrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInboundMeshTrafficPolicy", arg0, arg1)
	ret0, _ := ret[0].(*trafficpolicy.InboundMeshTrafficPolicy)
	return ret0
}

// GetInboundMeshTrafficPolicy indicates an expected call of GetInboundMeshTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetInboundMeshTrafficPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInboundMeshTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetInboundMeshTrafficPolicy), arg0, arg1)
}

// GetIngressTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetIngressTrafficPolicy(arg0 service.MeshService) (*trafficpolicy.IngressTrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.IngressTrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngressTrafficPolicy indicates an expected call of GetIngressTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetIngressTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetIngressTrafficPolicy), arg0)
}

// GetKubeController mocks base method.
func (m *MockMeshCataloger) GetKubeController() k8s.Controller {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubeController")
	ret0, _ := ret[0].(k8s.Controller)
	return ret0
}

// GetKubeController indicates an expected call of GetKubeController.
func (mr *MockMeshCatalogerMockRecorder) GetKubeController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubeController", reflect.TypeOf((*MockMeshCataloger)(nil).GetKubeController))
}

// GetOutboundMeshTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetOutboundMeshTrafficPolicy(arg0 identity.ServiceIdentity) *trafficpolicy.OutboundMeshTrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutboundMeshTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.OutboundMeshTrafficPolicy)
	return ret0
}

// GetOutboundMeshTrafficPolicy indicates an expected call of GetOutboundMeshTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetOutboundMeshTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutboundMeshTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetOutboundMeshTrafficPolicy), arg0)
}

// GetPluginChains mocks base method.
func (m *MockMeshCataloger) GetPluginChains() []*trafficpolicy.PluginChain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginChains")
	ret0, _ := ret[0].([]*trafficpolicy.PluginChain)
	return ret0
}

// GetPluginChains indicates an expected call of GetPluginChains.
func (mr *MockMeshCatalogerMockRecorder) GetPluginChains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginChains", reflect.TypeOf((*MockMeshCataloger)(nil).GetPluginChains))
}

// GetPluginConfigs mocks base method.
func (m *MockMeshCataloger) GetPluginConfigs() []*trafficpolicy.PluginConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPluginConfigs")
	ret0, _ := ret[0].([]*trafficpolicy.PluginConfig)
	return ret0
}

// GetPluginConfigs indicates an expected call of GetPluginConfigs.
func (mr *MockMeshCatalogerMockRecorder) GetPluginConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginConfigs", reflect.TypeOf((*MockMeshCataloger)(nil).GetPluginConfigs))
}

// GetPlugins mocks base method.
func (m *MockMeshCataloger) GetPlugins() []*trafficpolicy.Plugin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlugins")
	ret0, _ := ret[0].([]*trafficpolicy.Plugin)
	return ret0
}

// GetPlugins indicates an expected call of GetPlugins.
func (mr *MockMeshCatalogerMockRecorder) GetPlugins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlugins", reflect.TypeOf((*MockMeshCataloger)(nil).GetPlugins))
}

// GetRetryPolicy mocks base method.
func (m *MockMeshCataloger) GetRetryPolicy(arg0 identity.ServiceIdentity, arg1 service.MeshService) *v1alpha1.RetryPolicySpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRetryPolicy", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.RetryPolicySpec)
	return ret0
}

// GetRetryPolicy indicates an expected call of GetRetryPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetRetryPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRetryPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetRetryPolicy), arg0, arg1)
}

// ListAllowedUpstreamEndpointsForService mocks base method.
func (m *MockMeshCataloger) ListAllowedUpstreamEndpointsForService(arg0 identity.ServiceIdentity, arg1 service.MeshService) []endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedUpstreamEndpointsForService", arg0, arg1)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	return ret0
}

// ListAllowedUpstreamEndpointsForService indicates an expected call of ListAllowedUpstreamEndpointsForService.
func (mr *MockMeshCatalogerMockRecorder) ListAllowedUpstreamEndpointsForService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedUpstreamEndpointsForService", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedUpstreamEndpointsForService), arg0, arg1)
}

// ListInboundServiceIdentities mocks base method.
func (m *MockMeshCataloger) ListInboundServiceIdentities(arg0 identity.ServiceIdentity) []identity.ServiceIdentity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInboundServiceIdentities", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	return ret0
}

// ListInboundServiceIdentities indicates an expected call of ListInboundServiceIdentities.
func (mr *MockMeshCatalogerMockRecorder) ListInboundServiceIdentities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInboundServiceIdentities", reflect.TypeOf((*MockMeshCataloger)(nil).ListInboundServiceIdentities), arg0)
}

// ListInboundTrafficTargetsWithRoutes mocks base method.
func (m *MockMeshCataloger) ListInboundTrafficTargetsWithRoutes(arg0 identity.ServiceIdentity) ([]trafficpolicy.TrafficTargetWithRoutes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInboundTrafficTargetsWithRoutes", arg0)
	ret0, _ := ret[0].([]trafficpolicy.TrafficTargetWithRoutes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInboundTrafficTargetsWithRoutes indicates an expected call of ListInboundTrafficTargetsWithRoutes.
func (mr *MockMeshCatalogerMockRecorder) ListInboundTrafficTargetsWithRoutes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInboundTrafficTargetsWithRoutes", reflect.TypeOf((*MockMeshCataloger)(nil).ListInboundTrafficTargetsWithRoutes), arg0)
}

// ListOutboundServiceIdentities mocks base method.
func (m *MockMeshCataloger) ListOutboundServiceIdentities(arg0 identity.ServiceIdentity) []identity.ServiceIdentity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutboundServiceIdentities", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	return ret0
}

// ListOutboundServiceIdentities indicates an expected call of ListOutboundServiceIdentities.
func (mr *MockMeshCatalogerMockRecorder) ListOutboundServiceIdentities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutboundServiceIdentities", reflect.TypeOf((*MockMeshCataloger)(nil).ListOutboundServiceIdentities), arg0)
}

// ListOutboundServicesForIdentity mocks base method.
func (m *MockMeshCataloger) ListOutboundServicesForIdentity(arg0 identity.ServiceIdentity) []service.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutboundServicesForIdentity", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	return ret0
}

// ListOutboundServicesForIdentity indicates an expected call of ListOutboundServicesForIdentity.
func (mr *MockMeshCatalogerMockRecorder) ListOutboundServicesForIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutboundServicesForIdentity", reflect.TypeOf((*MockMeshCataloger)(nil).ListOutboundServicesForIdentity), arg0)
}

// ListServiceIdentitiesForService mocks base method.
func (m *MockMeshCataloger) ListServiceIdentitiesForService(arg0 service.MeshService) []identity.ServiceIdentity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServiceIdentitiesForService", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	return ret0
}

// ListServiceIdentitiesForService indicates an expected call of ListServiceIdentitiesForService.
func (mr *MockMeshCatalogerMockRecorder) ListServiceIdentitiesForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceIdentitiesForService", reflect.TypeOf((*MockMeshCataloger)(nil).ListServiceIdentitiesForService), arg0)
}
