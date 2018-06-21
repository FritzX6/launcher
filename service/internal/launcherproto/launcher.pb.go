// Code generated by protoc-gen-go. DO NOT EDIT.
// source: launcher.proto

/*
Package kolide_agent is a generated protocol buffer package.

It is generated from these files:
	launcher.proto

It has these top-level messages:
	AgentApiRequest
	AgentApiResponse
	EnrollmentRequest
	EnrollmentDetails
	EnrollmentResponse
	ConfigResponse
	LogCollection
	QueryCollection
	ResultCollection
	HealthCheckResponse
*/
package kolide_agent

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AGENT is added as a new log type, for adding new
// logging capabilities from kolide/agent
type LogCollection_LogType int32

const (
	LogCollection_RESULT LogCollection_LogType = 0
	LogCollection_STATUS LogCollection_LogType = 1
	LogCollection_AGENT  LogCollection_LogType = 2
)

var LogCollection_LogType_name = map[int32]string{
	0: "RESULT",
	1: "STATUS",
	2: "AGENT",
}
var LogCollection_LogType_value = map[string]int32{
	"RESULT": 0,
	"STATUS": 1,
	"AGENT":  2,
}

func (x LogCollection_LogType) String() string {
	return proto.EnumName(LogCollection_LogType_name, int32(x))
}
func (LogCollection_LogType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{9, 0}
}

type AgentApiRequest struct {
	NodeKey string `protobuf:"bytes,1,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
}

func (m *AgentApiRequest) Reset()                    { *m = AgentApiRequest{} }
func (m *AgentApiRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentApiRequest) ProtoMessage()               {}
func (*AgentApiRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AgentApiRequest) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

type AgentApiResponse struct {
	Message     string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	ErrorCode   string `protobuf:"bytes,2,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	NodeInvalid bool   `protobuf:"varint,3,opt,name=node_invalid,json=nodeInvalid" json:"node_invalid,omitempty"`
}

func (m *AgentApiResponse) Reset()                    { *m = AgentApiResponse{} }
func (m *AgentApiResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentApiResponse) ProtoMessage()               {}
func (*AgentApiResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AgentApiResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AgentApiResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *AgentApiResponse) GetNodeInvalid() bool {
	if m != nil {
		return m.NodeInvalid
	}
	return false
}

type EnrollmentRequest struct {
	EnrollSecret      string             `protobuf:"bytes,1,opt,name=enroll_secret,json=enrollSecret" json:"enroll_secret,omitempty"`
	HostIdentifier    string             `protobuf:"bytes,2,opt,name=host_identifier,json=hostIdentifier" json:"host_identifier,omitempty"`
	EnrollmentDetails *EnrollmentDetails `protobuf:"bytes,3,opt,name=enrollment_details,json=enrollmentDetails" json:"enrollment_details,omitempty"`
}

func (m *EnrollmentRequest) Reset()                    { *m = EnrollmentRequest{} }
func (m *EnrollmentRequest) String() string            { return proto.CompactTextString(m) }
func (*EnrollmentRequest) ProtoMessage()               {}
func (*EnrollmentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EnrollmentRequest) GetEnrollSecret() string {
	if m != nil {
		return m.EnrollSecret
	}
	return ""
}

func (m *EnrollmentRequest) GetHostIdentifier() string {
	if m != nil {
		return m.HostIdentifier
	}
	return ""
}

func (m *EnrollmentRequest) GetEnrollmentDetails() *EnrollmentDetails {
	if m != nil {
		return m.EnrollmentDetails
	}
	return nil
}

type EnrollmentDetails struct {
	OsVersion       string `protobuf:"bytes,1,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	OsBuild         string `protobuf:"bytes,2,opt,name=os_build,json=osBuild" json:"os_build,omitempty"`
	OsPlatform      string `protobuf:"bytes,3,opt,name=os_platform,json=osPlatform" json:"os_platform,omitempty"`
	Hostname        string `protobuf:"bytes,4,opt,name=hostname" json:"hostname,omitempty"`
	HardwareVendor  string `protobuf:"bytes,5,opt,name=hardware_vendor,json=hardwareVendor" json:"hardware_vendor,omitempty"`
	HardwareModel   string `protobuf:"bytes,6,opt,name=hardware_model,json=hardwareModel" json:"hardware_model,omitempty"`
	HardwareSerial  string `protobuf:"bytes,7,opt,name=hardware_serial,json=hardwareSerial" json:"hardware_serial,omitempty"`
	OsqueryVersion  string `protobuf:"bytes,8,opt,name=osquery_version,json=osqueryVersion" json:"osquery_version,omitempty"`
	LauncherVersion string `protobuf:"bytes,9,opt,name=launcher_version,json=launcherVersion" json:"launcher_version,omitempty"`
	OsName          string `protobuf:"bytes,10,opt,name=os_name,json=osName" json:"os_name,omitempty"`
	OsPlatformLike  string `protobuf:"bytes,11,opt,name=os_platform_like,json=osPlatformLike" json:"os_platform_like,omitempty"`
}

func (m *EnrollmentDetails) Reset()                    { *m = EnrollmentDetails{} }
func (m *EnrollmentDetails) String() string            { return proto.CompactTextString(m) }
func (*EnrollmentDetails) ProtoMessage()               {}
func (*EnrollmentDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EnrollmentDetails) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *EnrollmentDetails) GetOsBuild() string {
	if m != nil {
		return m.OsBuild
	}
	return ""
}

func (m *EnrollmentDetails) GetOsPlatform() string {
	if m != nil {
		return m.OsPlatform
	}
	return ""
}

func (m *EnrollmentDetails) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *EnrollmentDetails) GetHardwareVendor() string {
	if m != nil {
		return m.HardwareVendor
	}
	return ""
}

func (m *EnrollmentDetails) GetHardwareModel() string {
	if m != nil {
		return m.HardwareModel
	}
	return ""
}

func (m *EnrollmentDetails) GetHardwareSerial() string {
	if m != nil {
		return m.HardwareSerial
	}
	return ""
}

func (m *EnrollmentDetails) GetOsqueryVersion() string {
	if m != nil {
		return m.OsqueryVersion
	}
	return ""
}

func (m *EnrollmentDetails) GetLauncherVersion() string {
	if m != nil {
		return m.LauncherVersion
	}
	return ""
}

func (m *EnrollmentDetails) GetOsName() string {
	if m != nil {
		return m.OsName
	}
	return ""
}

func (m *EnrollmentDetails) GetOsPlatformLike() string {
	if m != nil {
		return m.OsPlatformLike
	}
	return ""
}

type EnrollmentResponse struct {
	NodeKey     string `protobuf:"bytes,1,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
	NodeInvalid bool   `protobuf:"varint,2,opt,name=node_invalid,json=nodeInvalid" json:"node_invalid,omitempty"`
	ErrorCode   string `protobuf:"bytes,3,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *EnrollmentResponse) Reset()                    { *m = EnrollmentResponse{} }
func (m *EnrollmentResponse) String() string            { return proto.CompactTextString(m) }
func (*EnrollmentResponse) ProtoMessage()               {}
func (*EnrollmentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EnrollmentResponse) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

func (m *EnrollmentResponse) GetNodeInvalid() bool {
	if m != nil {
		return m.NodeInvalid
	}
	return false
}

func (m *EnrollmentResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

// kolide/cloud will be generating well-structured json already, so forward
// it along instead of de/re/de-serializing it as a protobuf too
// this might make sense to convert to full proto later
type ConfigResponse struct {
	ConfigJsonBlob string `protobuf:"bytes,1,opt,name=config_json_blob,json=configJsonBlob" json:"config_json_blob,omitempty"`
	NodeInvalid    bool   `protobuf:"varint,2,opt,name=node_invalid,json=nodeInvalid" json:"node_invalid,omitempty"`
	ErrorCode      string `protobuf:"bytes,3,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *ConfigResponse) Reset()                    { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()               {}
func (*ConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ConfigResponse) GetConfigJsonBlob() string {
	if m != nil {
		return m.ConfigJsonBlob
	}
	return ""
}

func (m *ConfigResponse) GetNodeInvalid() bool {
	if m != nil {
		return m.NodeInvalid
	}
	return false
}

func (m *ConfigResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

type LogCollection struct {
	NodeKey   string                `protobuf:"bytes,1,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
	LogType   LogCollection_LogType `protobuf:"varint,2,opt,name=log_type,json=logType,enum=kolide.agent.LogCollection_LogType" json:"log_type,omitempty"`
	Logs      []*LogCollection_Log  `protobuf:"bytes,3,rep,name=logs" json:"logs,omitempty"`
	ErrorCode string                `protobuf:"bytes,4,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *LogCollection) Reset()                    { *m = LogCollection{} }
func (m *LogCollection) String() string            { return proto.CompactTextString(m) }
func (*LogCollection) ProtoMessage()               {}
func (*LogCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LogCollection) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

func (m *LogCollection) GetLogType() LogCollection_LogType {
	if m != nil {
		return m.LogType
	}
	return LogCollection_RESULT
}

func (m *LogCollection) GetLogs() []*LogCollection_Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *LogCollection) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

type LogCollection_Log struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *LogCollection_Log) Reset()                    { *m = LogCollection_Log{} }
func (m *LogCollection_Log) String() string            { return proto.CompactTextString(m) }
func (*LogCollection_Log) ProtoMessage()               {}
func (*LogCollection_Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *LogCollection_Log) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// a query collection contains many queries
type QueryCollection struct {
	Queries     []*QueryCollection_Query `protobuf:"bytes,1,rep,name=queries" json:"queries,omitempty"`
	NodeInvalid bool                     `protobuf:"varint,2,opt,name=node_invalid,json=nodeInvalid" json:"node_invalid,omitempty"`
	ErrorCode   string                   `protobuf:"bytes,3,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *QueryCollection) Reset()                    { *m = QueryCollection{} }
func (m *QueryCollection) String() string            { return proto.CompactTextString(m) }
func (*QueryCollection) ProtoMessage()               {}
func (*QueryCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueryCollection) GetQueries() []*QueryCollection_Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *QueryCollection) GetNodeInvalid() bool {
	if m != nil {
		return m.NodeInvalid
	}
	return false
}

func (m *QueryCollection) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

type QueryCollection_Query struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Query string `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
}

func (m *QueryCollection_Query) Reset()                    { *m = QueryCollection_Query{} }
func (m *QueryCollection_Query) String() string            { return proto.CompactTextString(m) }
func (*QueryCollection_Query) ProtoMessage()               {}
func (*QueryCollection_Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *QueryCollection_Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryCollection_Query) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

// a result collection contains many results
type ResultCollection struct {
	NodeKey   string                     `protobuf:"bytes,1,opt,name=node_key,json=nodeKey" json:"node_key,omitempty"`
	Results   []*ResultCollection_Result `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	ErrorCode string                     `protobuf:"bytes,3,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
}

func (m *ResultCollection) Reset()                    { *m = ResultCollection{} }
func (m *ResultCollection) String() string            { return proto.CompactTextString(m) }
func (*ResultCollection) ProtoMessage()               {}
func (*ResultCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ResultCollection) GetNodeKey() string {
	if m != nil {
		return m.NodeKey
	}
	return ""
}

func (m *ResultCollection) GetResults() []*ResultCollection_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ResultCollection) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

// status is moved here instead of appearing as a map of id[status]
// on the ResultCollection, as it does in the osq docs
type ResultCollection_Result struct {
	Id     string                               `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Rows   []*ResultCollection_Result_ResultRow `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	Status int32                                `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
}

func (m *ResultCollection_Result) Reset()                    { *m = ResultCollection_Result{} }
func (m *ResultCollection_Result) String() string            { return proto.CompactTextString(m) }
func (*ResultCollection_Result) ProtoMessage()               {}
func (*ResultCollection_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

func (m *ResultCollection_Result) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResultCollection_Result) GetRows() []*ResultCollection_Result_ResultRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *ResultCollection_Result) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ResultCollection_Result_ResultRow struct {
	Columns []*ResultCollection_Result_ResultRow_Column `protobuf:"bytes,1,rep,name=columns" json:"columns,omitempty"`
}

func (m *ResultCollection_Result_ResultRow) Reset()         { *m = ResultCollection_Result_ResultRow{} }
func (m *ResultCollection_Result_ResultRow) String() string { return proto.CompactTextString(m) }
func (*ResultCollection_Result_ResultRow) ProtoMessage()    {}
func (*ResultCollection_Result_ResultRow) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0, 0}
}

func (m *ResultCollection_Result_ResultRow) GetColumns() []*ResultCollection_Result_ResultRow_Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

type ResultCollection_Result_ResultRow_Column struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ResultCollection_Result_ResultRow_Column) Reset() {
	*m = ResultCollection_Result_ResultRow_Column{}
}
func (m *ResultCollection_Result_ResultRow_Column) String() string { return proto.CompactTextString(m) }
func (*ResultCollection_Result_ResultRow_Column) ProtoMessage()    {}
func (*ResultCollection_Result_ResultRow_Column) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0, 0, 0}
}

func (m *ResultCollection_Result_ResultRow_Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResultCollection_Result_ResultRow_Column) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=kolide.agent.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*AgentApiRequest)(nil), "kolide.agent.AgentApiRequest")
	proto.RegisterType((*AgentApiResponse)(nil), "kolide.agent.AgentApiResponse")
	proto.RegisterType((*EnrollmentRequest)(nil), "kolide.agent.EnrollmentRequest")
	proto.RegisterType((*EnrollmentDetails)(nil), "kolide.agent.EnrollmentDetails")
	proto.RegisterType((*EnrollmentResponse)(nil), "kolide.agent.EnrollmentResponse")
	proto.RegisterType((*ConfigResponse)(nil), "kolide.agent.ConfigResponse")
	proto.RegisterType((*LogCollection)(nil), "kolide.agent.LogCollection")
	proto.RegisterType((*LogCollection_Log)(nil), "kolide.agent.LogCollection.Log")
	proto.RegisterType((*QueryCollection)(nil), "kolide.agent.QueryCollection")
	proto.RegisterType((*QueryCollection_Query)(nil), "kolide.agent.QueryCollection.Query")
	proto.RegisterType((*ResultCollection)(nil), "kolide.agent.ResultCollection")
	proto.RegisterType((*ResultCollection_Result)(nil), "kolide.agent.ResultCollection.Result")
	proto.RegisterType((*ResultCollection_Result_ResultRow)(nil), "kolide.agent.ResultCollection.Result.ResultRow")
	proto.RegisterType((*ResultCollection_Result_ResultRow_Column)(nil), "kolide.agent.ResultCollection.Result.ResultRow.Column")
	proto.RegisterType((*HealthCheckResponse)(nil), "kolide.agent.HealthCheckResponse")
	proto.RegisterEnum("kolide.agent.LogCollection_LogType", LogCollection_LogType_name, LogCollection_LogType_value)
	proto.RegisterEnum("kolide.agent.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Api service

type ApiClient interface {
	// Attempt to enroll a host with kolide/cloud
	RequestEnrollment(ctx context.Context, in *EnrollmentRequest, opts ...grpc.CallOption) (*EnrollmentResponse, error)
	// request an updated configuration from kolide/cloud
	// a generic request object is sent
	RequestConfig(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	// request/pull Dist queries from kolide/cloud
	// a generic request object is sent
	RequestQueries(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*QueryCollection, error)
	// publish logs from osqueryd to kolide/cloud
	// a generic response object is returned
	PublishLogs(ctx context.Context, in *LogCollection, opts ...grpc.CallOption) (*AgentApiResponse, error)
	// publish results from Dist queries to kolide/cloud
	// a generic response object is returned
	PublishResults(ctx context.Context, in *ResultCollection, opts ...grpc.CallOption) (*AgentApiResponse, error)
	// check the health of the GRPC server
	// a value indicating healthiness is returned
	// if you don't hear from this endpoint assume the worst
	CheckHealth(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) RequestEnrollment(ctx context.Context, in *EnrollmentRequest, opts ...grpc.CallOption) (*EnrollmentResponse, error) {
	out := new(EnrollmentResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/RequestEnrollment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RequestConfig(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/RequestConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RequestQueries(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*QueryCollection, error) {
	out := new(QueryCollection)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/RequestQueries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PublishLogs(ctx context.Context, in *LogCollection, opts ...grpc.CallOption) (*AgentApiResponse, error) {
	out := new(AgentApiResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/PublishLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) PublishResults(ctx context.Context, in *ResultCollection, opts ...grpc.CallOption) (*AgentApiResponse, error) {
	out := new(AgentApiResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/PublishResults", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CheckHealth(ctx context.Context, in *AgentApiRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/kolide.agent.Api/CheckHealth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	// Attempt to enroll a host with kolide/cloud
	RequestEnrollment(context.Context, *EnrollmentRequest) (*EnrollmentResponse, error)
	// request an updated configuration from kolide/cloud
	// a generic request object is sent
	RequestConfig(context.Context, *AgentApiRequest) (*ConfigResponse, error)
	// request/pull Dist queries from kolide/cloud
	// a generic request object is sent
	RequestQueries(context.Context, *AgentApiRequest) (*QueryCollection, error)
	// publish logs from osqueryd to kolide/cloud
	// a generic response object is returned
	PublishLogs(context.Context, *LogCollection) (*AgentApiResponse, error)
	// publish results from Dist queries to kolide/cloud
	// a generic response object is returned
	PublishResults(context.Context, *ResultCollection) (*AgentApiResponse, error)
	// check the health of the GRPC server
	// a value indicating healthiness is returned
	// if you don't hear from this endpoint assume the worst
	CheckHealth(context.Context, *AgentApiRequest) (*HealthCheckResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_RequestEnrollment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnrollmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RequestEnrollment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/RequestEnrollment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RequestEnrollment(ctx, req.(*EnrollmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RequestConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RequestConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/RequestConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RequestConfig(ctx, req.(*AgentApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RequestQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RequestQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/RequestQueries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RequestQueries(ctx, req.(*AgentApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PublishLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PublishLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/PublishLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PublishLogs(ctx, req.(*LogCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_PublishResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResultCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).PublishResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/PublishResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).PublishResults(ctx, req.(*ResultCollection))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kolide.agent.Api/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CheckHealth(ctx, req.(*AgentApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kolide.agent.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestEnrollment",
			Handler:    _Api_RequestEnrollment_Handler,
		},
		{
			MethodName: "RequestConfig",
			Handler:    _Api_RequestConfig_Handler,
		},
		{
			MethodName: "RequestQueries",
			Handler:    _Api_RequestQueries_Handler,
		},
		{
			MethodName: "PublishLogs",
			Handler:    _Api_PublishLogs_Handler,
		},
		{
			MethodName: "PublishResults",
			Handler:    _Api_PublishResults_Handler,
		},
		{
			MethodName: "CheckHealth",
			Handler:    _Api_CheckHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "launcher.proto",
}

func init() { proto.RegisterFile("launcher.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 997 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x8f, 0x93, 0x26, 0x6e, 0x5e, 0xda, 0xc4, 0x1d, 0x56, 0xe0, 0x0d, 0x74, 0xb7, 0xeb, 0xd5,
	0x8a, 0x22, 0x2d, 0x41, 0xea, 0x4a, 0x1c, 0x90, 0x00, 0x65, 0x43, 0x55, 0xca, 0x86, 0x6c, 0xeb,
	0xa4, 0x85, 0x9b, 0xe5, 0xc4, 0xaf, 0xc9, 0x90, 0x89, 0x27, 0xf5, 0x38, 0xa9, 0x22, 0x71, 0x47,
	0x1c, 0x39, 0xb2, 0xdf, 0x82, 0x8f, 0xc0, 0x97, 0xe2, 0x8c, 0x3c, 0x1e, 0xbb, 0x89, 0x43, 0xda,
	0x22, 0xb8, 0xf9, 0xfd, 0xfc, 0x7b, 0x7f, 0xf3, 0x7b, 0xcf, 0x81, 0x2a, 0x73, 0x67, 0xfe, 0x60,
	0x84, 0x41, 0x63, 0x1a, 0xf0, 0x90, 0x93, 0x9d, 0x31, 0x67, 0xd4, 0xc3, 0x86, 0x3b, 0x44, 0x3f,
	0xb4, 0x5e, 0x42, 0xad, 0x19, 0x3d, 0x34, 0xa7, 0xd4, 0xc6, 0xeb, 0x19, 0x8a, 0x90, 0x3c, 0x86,
	0x6d, 0x9f, 0x7b, 0xe8, 0x8c, 0x71, 0x61, 0x6a, 0x07, 0xda, 0x61, 0xd9, 0xd6, 0x23, 0xfb, 0x0d,
	0x2e, 0x2c, 0x1f, 0x8c, 0x5b, 0xb6, 0x98, 0x72, 0x5f, 0x20, 0x31, 0x41, 0x9f, 0xa0, 0x10, 0xee,
	0x10, 0x13, 0xb6, 0x32, 0xc9, 0x3e, 0x00, 0x06, 0x01, 0x0f, 0x9c, 0x01, 0xf7, 0xd0, 0xcc, 0xcb,
	0x97, 0x65, 0x89, 0xb4, 0xb8, 0x87, 0xe4, 0x19, 0xec, 0xc8, 0x3c, 0xd4, 0x9f, 0xbb, 0x8c, 0x7a,
	0x66, 0xe1, 0x40, 0x3b, 0xdc, 0xb6, 0x2b, 0x11, 0x76, 0x1a, 0x43, 0xd6, 0x1f, 0x1a, 0xec, 0x1d,
	0xfb, 0x01, 0x67, 0x6c, 0x82, 0x7e, 0x98, 0x14, 0xf8, 0x1c, 0x76, 0x51, 0x82, 0x8e, 0xc0, 0x41,
	0x80, 0xa1, 0xca, 0xbb, 0x13, 0x83, 0x5d, 0x89, 0x91, 0x8f, 0xa1, 0x36, 0xe2, 0x22, 0x74, 0xa8,
	0x87, 0x7e, 0x48, 0xaf, 0x28, 0x06, 0xaa, 0x82, 0x6a, 0x04, 0x9f, 0xa6, 0x28, 0xe9, 0x00, 0xc1,
	0x34, 0x85, 0xe3, 0x61, 0xe8, 0x52, 0x26, 0x64, 0x31, 0x95, 0xa3, 0xa7, 0x8d, 0xe5, 0x61, 0x35,
	0x6e, 0x4b, 0xf9, 0x26, 0xa6, 0xd9, 0x7b, 0x98, 0x85, 0xac, 0x77, 0x85, 0xe5, 0x9a, 0x15, 0x1a,
	0xcd, 0x82, 0x0b, 0x67, 0x8e, 0x81, 0xa0, 0xdc, 0x57, 0x05, 0x97, 0xb9, 0xb8, 0x8c, 0x81, 0x68,
	0xe6, 0x5c, 0x38, 0xfd, 0x19, 0x65, 0x9e, 0x2a, 0x53, 0xe7, 0xe2, 0x75, 0x64, 0x92, 0xa7, 0x50,
	0xe1, 0xc2, 0x99, 0x32, 0x37, 0xbc, 0xe2, 0xc1, 0x44, 0x16, 0x56, 0xb6, 0x81, 0x8b, 0x33, 0x85,
	0x90, 0x3a, 0x6c, 0x47, 0x2d, 0xf9, 0xee, 0x04, 0xcd, 0x2d, 0xf9, 0x36, 0xb5, 0xe5, 0x14, 0xdc,
	0xc0, 0xbb, 0x71, 0x03, 0x74, 0xe6, 0xe8, 0x7b, 0x3c, 0x30, 0x8b, 0x6a, 0x0a, 0x0a, 0xbe, 0x94,
	0x28, 0x79, 0x01, 0x29, 0xe2, 0x4c, 0xb8, 0x87, 0xcc, 0x2c, 0x49, 0xde, 0x6e, 0x82, 0x7e, 0x1f,
	0x81, 0x2b, 0xf1, 0x04, 0x06, 0xd4, 0x65, 0xa6, 0xbe, 0x1a, 0xaf, 0x2b, 0xd1, 0x88, 0xc8, 0xc5,
	0xf5, 0x0c, 0x83, 0x45, 0xda, 0xf4, 0x76, 0x4c, 0x54, 0x70, 0xd2, 0xf9, 0x27, 0x60, 0x24, 0x02,
	0x4d, 0x99, 0x65, 0xc9, 0xac, 0x25, 0x78, 0x42, 0xfd, 0x00, 0x74, 0x2e, 0x1c, 0xd9, 0x27, 0x48,
	0x46, 0x89, 0x8b, 0x4e, 0xd4, 0xe5, 0x21, 0x18, 0x4b, 0x23, 0x72, 0x18, 0x1d, 0xa3, 0x59, 0x49,
	0xb2, 0x25, 0x73, 0x6a, 0xd3, 0x31, 0x5a, 0xd7, 0x40, 0x96, 0xf5, 0xa4, 0x24, 0xbc, 0x59, 0xf1,
	0x6b, 0x22, 0xcd, 0xaf, 0x89, 0x34, 0x23, 0xf3, 0x42, 0x46, 0xe6, 0xd6, 0xcf, 0x50, 0x6d, 0x71,
	0xff, 0x8a, 0x0e, 0xd3, 0x74, 0x87, 0x60, 0x0c, 0x24, 0xe2, 0xfc, 0x24, 0xb8, 0xef, 0xf4, 0x19,
	0xef, 0xab, 0xb4, 0xd5, 0x18, 0xff, 0x4e, 0x70, 0xff, 0x35, 0xe3, 0xfd, 0xff, 0x21, 0xfb, 0x2f,
	0x79, 0xd8, 0x6d, 0xf3, 0x61, 0x8b, 0x33, 0x86, 0x83, 0x50, 0x49, 0x6d, 0x53, 0xb3, 0x5f, 0xc1,
	0x36, 0xe3, 0x43, 0x27, 0x5c, 0x4c, 0xe3, 0x75, 0xad, 0x1e, 0x3d, 0x5f, 0x5d, 0x80, 0x95, 0x48,
	0x91, 0xd5, 0x5b, 0x4c, 0xd1, 0xd6, 0x59, 0xfc, 0x40, 0x5e, 0xc1, 0x16, 0xe3, 0xc3, 0x68, 0x79,
	0x0a, 0xeb, 0xcb, 0xb3, 0xe6, 0x6b, 0x4b, 0x72, 0xa6, 0x81, 0xad, 0x4c, 0x03, 0xf5, 0xc7, 0x50,
	0x68, 0xf3, 0x21, 0x21, 0xb0, 0xe5, 0xb9, 0xa1, 0xab, 0x2a, 0x96, 0xcf, 0xd6, 0x4b, 0xd0, 0x55,
	0x09, 0x04, 0xa0, 0x64, 0x1f, 0x77, 0x2f, 0xda, 0x3d, 0x23, 0x17, 0x3d, 0x77, 0x7b, 0xcd, 0xde,
	0x45, 0xd7, 0xd0, 0x48, 0x19, 0x8a, 0xcd, 0x93, 0xe3, 0x4e, 0xcf, 0xc8, 0x5b, 0x7f, 0x6a, 0x50,
	0x3b, 0x8f, 0x94, 0xb7, 0x34, 0x8b, 0x2f, 0x41, 0x8f, 0xc4, 0x48, 0x51, 0x98, 0x9a, 0xac, 0x39,
	0xd3, 0x6f, 0x86, 0x1f, 0xdb, 0x76, 0xe2, 0xf3, 0xdf, 0x7f, 0x9e, 0xfa, 0xa7, 0x50, 0x94, 0x31,
	0x49, 0x15, 0xf2, 0xd4, 0x53, 0xdd, 0xe5, 0xa9, 0x47, 0x1e, 0x41, 0x51, 0xae, 0x89, 0xba, 0x06,
	0xb1, 0x61, 0xfd, 0x5e, 0x00, 0xc3, 0x46, 0x31, 0x63, 0xe1, 0xc3, 0x7e, 0xd0, 0xaf, 0x41, 0x0f,
	0x24, 0x5d, 0x98, 0x79, 0xd9, 0xdf, 0x8b, 0xd5, 0xfe, 0xb2, 0xb1, 0x14, 0x60, 0x27, 0x5e, 0xf7,
	0x95, 0xff, 0x6b, 0x1e, 0x4a, 0xb1, 0xcb, 0x5a, 0x03, 0x2d, 0xd8, 0x0a, 0xf8, 0x4d, 0x92, 0xf7,
	0xb3, 0x07, 0xe5, 0x4d, 0xd2, 0xf3, 0x1b, 0x5b, 0x3a, 0x93, 0xf7, 0xa1, 0x24, 0x42, 0x37, 0x9c,
	0xc5, 0xf7, 0xb8, 0x68, 0x2b, 0xab, 0xfe, 0x9b, 0x06, 0xe5, 0x94, 0x4b, 0xce, 0x40, 0x1f, 0x70,
	0x36, 0x9b, 0xf8, 0xc9, 0xaf, 0xf8, 0xf9, 0xbf, 0xcc, 0xd6, 0x68, 0x49, 0x77, 0x3b, 0x09, 0x53,
	0x3f, 0x82, 0x52, 0x0c, 0x45, 0xba, 0x93, 0x07, 0x47, 0xe9, 0x4e, 0x1e, 0xd5, 0x47, 0x50, 0x9c,
	0xbb, 0x6c, 0x96, 0x7c, 0xd2, 0x62, 0xc3, 0x7a, 0xa7, 0xc1, 0x7b, 0xdf, 0xa2, 0xcb, 0xc2, 0x51,
	0x6b, 0x84, 0x83, 0x71, 0xba, 0xed, 0x27, 0x69, 0x0f, 0x9a, 0x5c, 0xa9, 0xcc, 0x28, 0xfe, 0xc1,
	0xa5, 0xd1, 0xc5, 0x60, 0x4e, 0xfd, 0x61, 0x57, 0xba, 0x25, 0x4d, 0x5b, 0x5f, 0xc0, 0xee, 0xca,
	0x0b, 0x52, 0x01, 0xfd, 0xa2, 0xf3, 0xa6, 0xf3, 0xf6, 0x87, 0x8e, 0x91, 0x8b, 0x8c, 0xee, 0xb1,
	0x7d, 0x79, 0xda, 0x39, 0x31, 0x34, 0x52, 0x83, 0x4a, 0xe7, 0x6d, 0xcf, 0x49, 0x80, 0xfc, 0xd1,
	0x5f, 0x05, 0x28, 0x34, 0xa7, 0x94, 0xfc, 0x08, 0x7b, 0xea, 0x2b, 0x7a, 0x7b, 0x06, 0xc9, 0xc6,
	0xaf, 0x9c, 0xa2, 0xd6, 0x0f, 0x36, 0x13, 0xe2, 0x8a, 0xad, 0x1c, 0xe9, 0xc0, 0xae, 0xa2, 0xc7,
	0xd7, 0x8e, 0xec, 0xaf, 0x3a, 0x65, 0xfe, 0x65, 0xd4, 0x3f, 0x5a, 0x7d, 0xbd, 0x7a, 0x22, 0xad,
	0x1c, 0x39, 0x83, 0xaa, 0xa2, 0x9e, 0xab, 0x6d, 0xbb, 0x27, 0xe0, 0xfe, 0x9d, 0xab, 0x6b, 0xe5,
	0x48, 0x1b, 0x2a, 0x67, 0xb3, 0x3e, 0xa3, 0x62, 0xd4, 0x8e, 0xee, 0xce, 0x87, 0x77, 0x9c, 0xa7,
	0xfa, 0x93, 0x4d, 0xb9, 0xd2, 0xfa, 0x6c, 0xa8, 0xaa, 0x68, 0xb6, 0xda, 0x95, 0x27, 0x77, 0xab,
	0xee, 0x01, 0x31, 0xcf, 0xa1, 0x22, 0x85, 0x10, 0x6b, 0xe2, 0xbe, 0x86, 0x9f, 0xdd, 0x2b, 0x24,
	0x2b, 0xd7, 0x2f, 0xc9, 0x3f, 0x7d, 0xaf, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x60, 0xf5,
	0x82, 0x06, 0x0a, 0x00, 0x00,
}
