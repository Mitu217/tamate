// Package cloudtasks provides access to the Cloud Tasks API.
//
// See https://cloud.google.com/cloud-tasks/
//
// Usage example:
//
//   import "google.golang.org/api/cloudtasks/v2beta2"
//   ...
//   cloudtasksService, err := cloudtasks.New(oauthHttpClient)
package cloudtasks // import "google.golang.org/api/cloudtasks/v2beta2"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "cloudtasks:v2beta2"
const apiName = "cloudtasks"
const apiVersion = "v2beta2"
const basePath = "https://cloudtasks.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Locations = NewProjectsLocationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Locations *ProjectsLocationsService
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	rs.Queues = NewProjectsLocationsQueuesService(s)
	return rs
}

type ProjectsLocationsService struct {
	s *Service

	Queues *ProjectsLocationsQueuesService
}

func NewProjectsLocationsQueuesService(s *Service) *ProjectsLocationsQueuesService {
	rs := &ProjectsLocationsQueuesService{s: s}
	rs.Tasks = NewProjectsLocationsQueuesTasksService(s)
	return rs
}

type ProjectsLocationsQueuesService struct {
	s *Service

	Tasks *ProjectsLocationsQueuesTasksService
}

func NewProjectsLocationsQueuesTasksService(s *Service) *ProjectsLocationsQueuesTasksService {
	rs := &ProjectsLocationsQueuesTasksService{s: s}
	return rs
}

type ProjectsLocationsQueuesTasksService struct {
	s *Service
}

// AcknowledgeTaskRequest: Request message for acknowledging a task
// using
// AcknowledgeTask.
type AcknowledgeTaskRequest struct {
	// ScheduleTime: Required.
	//
	// The task's current schedule time, available in the
	// schedule_time returned by
	// LeaseTasks response or
	// RenewLease response. This restriction is
	// to ensure that your worker currently holds the lease.
	ScheduleTime string `json:"scheduleTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ScheduleTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ScheduleTime") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AcknowledgeTaskRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AcknowledgeTaskRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AppEngineHttpRequest: App Engine HTTP request.
//
// The message defines the HTTP request that is sent to an App Engine
// app when
// the task is dispatched.
//
// This proto can only be used for tasks in a queue which
// has
// app_engine_http_target set.
//
// Using AppEngineHttpRequest
// requires
// [`appengine.applications.get`](/appengine/docs/admin-api/acce
// ss-control)
// Google IAM permission for the project
// and the following
// scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
//
// The task will be delivered to the App Engine app which belongs to the
// same
// project as the queue. For more information, see
// [How Requests are
// Routed](/appengine/docs/standard/python/how-requests-are-routed)
// and how routing is affected by
// [dispatch files](/appengine/docs/python/config/dispatchref).
//
// The AppEngineRouting used to construct the URL that the task
// is
// delivered to can be set at the queue-level or task-level:
//
// * If set,
//    app_engine_routing_override
//    is used for all tasks in the queue, no matter what the setting
//    is for the
//    task-level app_engine_routing.
//
//
// The `url` that the task will be sent to is:
//
// * `url =` host `+`
//   relative_url
//
// The task attempt has succeeded if the app's request handler
// returns
// an HTTP response code in the range [`200` - `299`]. `503`
// is
// considered an App Engine system error instead of an
// application
// error. Requests returning error `503` will be retried regardless
// of
// retry configuration and not counted against retry counts.
// Any other response code or a failure to receive a response before
// the
// deadline is a failed attempt.
type AppEngineHttpRequest struct {
	// AppEngineRouting: Task-level setting for App Engine routing.
	//
	// If set,
	// app_engine_routing_override
	// is used for all tasks in the queue, no matter what the setting is for
	// the
	// task-level app_engine_routing.
	AppEngineRouting *AppEngineRouting `json:"appEngineRouting,omitempty"`

	// Headers: HTTP request headers.
	//
	// This map contains the header field names and values.
	// Headers can be set when the
	// [task is
	// created](google.cloud.tasks.v2beta2.CloudTasks.CreateTask).
	// Repeated headers are not supported but a header value can contain
	// commas.
	//
	// Cloud Tasks sets some headers to default values:
	//
	// * `User-Agent`: By default, this header is
	//   "AppEngine-Google; (+http://code.google.com/appengine)".
	//   This header can be modified, but Cloud Tasks will append
	//   "AppEngine-Google; (+http://code.google.com/appengine)" to the
	//   modified `User-Agent`.
	//
	// If the task has a payload, Cloud
	// Tasks sets the following headers:
	//
	// * `Content-Type`: By default, the `Content-Type` header is set to
	//   "application/octet-stream". The default can be overridden by
	// explicitly
	//   setting `Content-Type` to a particular media type when the
	//   [task is
	// created](google.cloud.tasks.v2beta2.CloudTasks.CreateTask).
	//   For example, `Content-Type` can be set to "application/json".
	// * `Content-Length`: This is computed by Cloud Tasks. This value is
	//   output only.   It cannot be changed.
	//
	// The headers below cannot be set or overridden:
	//
	// * `Host`
	// * `X-Google-*`
	// * `X-AppEngine-*`
	//
	// In addition, Cloud Tasks sets some headers when the task is
	// dispatched,
	// such as headers containing information about the task; see
	// [request
	// headers](/appengine/docs/python/taskqueue/push/creating-handlers#readi
	// ng_request_headers).
	// These headers are set only when the task is dispatched, so they are
	// not
	// visible when the task is returned in a Cloud Tasks
	// response.
	//
	// Although there is no specific limit for the maximum number of headers
	// or
	// the size, there is a limit on the maximum size of the Task. For
	// more
	// information, see the CreateTask documentation.
	Headers map[string]string `json:"headers,omitempty"`

	// HttpMethod: The HTTP method to use for the request. The default is
	// POST.
	//
	// The app's request handler for the task's target URL must be able to
	// handle
	// HTTP requests with this http_method, otherwise the task attempt will
	// fail
	// with error code 405 (Method Not Allowed). See
	// [Writing a push task request
	// handler](/appengine/docs/java/taskqueue/push/creating-handlers#writing
	// _a_push_task_request_handler)
	// and the documentation for the request handlers in the language your
	// app is
	// written in e.g.
	// [Python Request
	// Handler](/appengine/docs/python/tools/webapp/requesthandlerclass).
	//
	// Possible values:
	//   "HTTP_METHOD_UNSPECIFIED" - HTTP method unspecified
	//   "POST" - HTTP Post
	//   "GET" - HTTP Get
	//   "HEAD" - HTTP Head
	//   "PUT" - HTTP Put
	//   "DELETE" - HTTP Delete
	HttpMethod string `json:"httpMethod,omitempty"`

	// Payload: Payload.
	//
	// The payload will be sent as the HTTP message body. A message
	// body, and thus a payload, is allowed only if the HTTP method is
	// POST or PUT. It is an error to set a data payload on a task with
	// an incompatible HttpMethod.
	Payload string `json:"payload,omitempty"`

	// RelativeUrl: The relative URL.
	//
	// The relative URL must begin with "/" and must be a valid HTTP
	// relative URL.
	// It can contain a path and query string arguments.
	// If the relative URL is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083
	// characters.
	RelativeUrl string `json:"relativeUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppEngineRouting") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppEngineRouting") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AppEngineHttpRequest) MarshalJSON() ([]byte, error) {
	type NoMethod AppEngineHttpRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AppEngineHttpTarget: App Engine HTTP target.
//
// The task will be delivered to the App Engine application
// hostname
// specified by its AppEngineHttpTarget and AppEngineHttpRequest.
// The documentation for AppEngineHttpRequest explains how the
// task's host URL is constructed.
//
// Using AppEngineHttpTarget
// requires
// [`appengine.applications.get`](/appengine/docs/admin-api/acce
// ss-control)
// Google IAM permission for the project
// and the following
// scope:
//
// `https://www.googleapis.com/auth/cloud-platform`
type AppEngineHttpTarget struct {
	// AppEngineRoutingOverride: Overrides for the
	// task-level app_engine_routing.
	//
	// If set, `app_engine_routing_override` is used for all tasks in
	// the queue, no matter what the setting is for the
	// task-level app_engine_routing.
	AppEngineRoutingOverride *AppEngineRouting `json:"appEngineRoutingOverride,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AppEngineRoutingOverride") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppEngineRoutingOverride")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *AppEngineHttpTarget) MarshalJSON() ([]byte, error) {
	type NoMethod AppEngineHttpTarget
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AppEngineRouting: App Engine Routing.
//
// For more information about services, versions, and instances see
// [An Overview of App
// Engine](/appengine/docs/python/an-overview-of-app-engine),
// [Microservi
// ces Architecture on Google App
// Engine](/appengine/docs/python/microservices-on-app-engine),
// [App Engine Standard request
// routing](/appengine/docs/standard/python/how-requests-are-routed),
// and
//  [App Engine Flex request
// routing](/appengine/docs/flexible/python/how-requests-are-routed).
type AppEngineRouting struct {
	// Host: Output only. The host that the task is sent to.
	//
	// For more information, see
	// [How Requests are
	// Routed](/appengine/docs/standard/python/how-requests-are-routed).
	//
	// The
	//  host is constructed as:
	//
	//
	// * `host = [application_domain_name]`</br>
	//   `| [service] + '.' + [application_domain_name]`</br>
	//   `| [version] + '.' + [application_domain_name]`</br>
	//   `| [version_dot_service]+ '.' + [application_domain_name]`</br>
	//   `| [instance] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_service] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version] + '.' + [application_domain_name]`</br>
	//   `| [instance_dot_version_dot_service] + '.' +
	// [application_domain_name]`
	//
	// * `application_domain_name` = The domain name of the app, for
	//   example <app-id>.appspot.com, which is associated with the
	//   queue's project ID. Some tasks which were created using the App
	// Engine
	//   SDK use a custom domain name.
	//
	// * `service =` service
	//
	// * `version =` version
	//
	// * `version_dot_service =`
	//   version `+ '.' +`
	//   service
	//
	// * `instance =` instance
	//
	// * `instance_dot_service =`
	//   instance `+ '.' +`
	//   service
	//
	// * `instance_dot_version =`
	//   instance `+ '.' +`
	//   version
	//
	// * `instance_dot_version_dot_service =`
	//   instance `+ '.' +`
	//   version `+ '.' +`
	//   service
	//
	// If service is empty, then the task will be sent
	// to the service which is the default service when the task is
	// attempted.
	//
	// If version is empty, then the task will be sent
	// to the version which is the default version when the task is
	// attempted.
	//
	// If instance is empty, then the task
	// will be sent to an instance which is available when the task
	// is
	// attempted.
	//
	// When service is "default",
	// version is "default", and
	// instance is empty,
	// host is shortened to just the
	// `application_domain_name`.
	//
	// If service,
	// version, or
	// instance is invalid, then the task
	// will be sent to the default version of the default service when
	// the task is attempted.
	Host string `json:"host,omitempty"`

	// Instance: App instance.
	//
	// By default, the task is sent to an instance which is available
	// when
	// the task is attempted.
	//
	// Requests can only be sent to a specific instance if
	// [manual scaling is used in App Engine
	// Standard](/appengine/docs/python/an-overview-of-app-engine?hl=en_US#sc
	// aling_types_and_instance_classes).
	// App Engine Flex does not support instances. For more information,
	// see
	// [App Engine Standard request
	// routing](/appengine/docs/standard/python/how-requests-are-routed)
	// and [App Engine Flex request
	// routing](/appengine/docs/flexible/python/how-requests-are-routed).
	Instance string `json:"instance,omitempty"`

	// Service: App service.
	//
	// By default, the task is sent to the service which is the
	// default
	// service when the task is attempted ("default").
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, host is not parsable
	// into service,
	// version, and
	// instance. For example, some tasks
	// which were created using the App Engine SDK use a custom domain
	// name; custom domains are not parsed by Cloud Tasks. If
	// host is not parsable, then
	// service,
	// version, and
	// instance are the empty string.
	Service string `json:"service,omitempty"`

	// Version: App version.
	//
	// By default, the task is sent to the version which is the
	// default
	// version when the task is attempted ("default").
	//
	// For some queues or tasks which were created using the App Engine
	// Task Queue API, host is not parsable
	// into service,
	// version, and
	// instance. For example, some tasks
	// which were created using the App Engine SDK use a custom domain
	// name; custom domains are not parsed by Cloud Tasks. If
	// host is not parsable, then
	// service,
	// version, and
	// instance are the empty string.
	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Host") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Host") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AppEngineRouting) MarshalJSON() ([]byte, error) {
	type NoMethod AppEngineRouting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AttemptStatus: The status of a task attempt.
type AttemptStatus struct {
	// DispatchTime: Output only. The time that this attempt was
	// dispatched.
	//
	// `dispatch_time` will be truncated to the nearest microsecond.
	DispatchTime string `json:"dispatchTime,omitempty"`

	// ResponseStatus: Output only. The response from the target for this
	// attempt.
	//
	// If the task has not been attempted or the task is currently
	// running
	// then the response status is unset.
	ResponseStatus *Status `json:"responseStatus,omitempty"`

	// ResponseTime: Output only. The time that this attempt response was
	// received.
	//
	// `response_time` will be truncated to the nearest microsecond.
	ResponseTime string `json:"responseTime,omitempty"`

	// ScheduleTime: Output only. The time that this attempt was
	// scheduled.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime string `json:"scheduleTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DispatchTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DispatchTime") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AttemptStatus) MarshalJSON() ([]byte, error) {
	type NoMethod AttemptStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Binding: Associates `members` with a `role`.
type Binding struct {
	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents
	// anyone
	//    who is authenticated with a Google account or a service
	// account.
	//
	// * `user:{emailid}`: An email address that represents a specific
	// Google
	//    account. For example, `alice@gmail.com` or `joe@example.com`.
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a
	// service
	//    account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google
	// group.
	//    For example, `admins@example.com`.
	//
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all
	// the
	//    users of that domain. For example, `google.com` or
	// `example.com`.
	//
	//
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or
	// `roles/owner`.
	// Required
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Members") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Members") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Binding) MarshalJSON() ([]byte, error) {
	type NoMethod Binding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CancelLeaseRequest: Request message for canceling a lease
// using
// CancelLease.
type CancelLeaseRequest struct {
	// ResponseView: The response_view specifies which subset of the Task
	// will be
	// returned.
	//
	// By default response_view is BASIC; not all
	// information is retrieved by default because some data, such
	// as
	// payloads, might be desirable to return only when needed because
	// of its large size or because of the sensitivity of data that
	// it
	// contains.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on
	// the
	// Task resource.
	//
	// Possible values:
	//   "VIEW_UNSPECIFIED" - Unspecified. Defaults to BASIC.
	//   "BASIC" - The basic view omits fields which can be large or can
	// contain
	// sensitive data.
	//
	// This view does not include the
	// (payload in AppEngineHttpRequest and
	// payload in PullMessage). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	//   "FULL" - All information is returned.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google
	// IAM](https://cloud.google.com/iam/)
	// permission on the Queue resource.
	ResponseView string `json:"responseView,omitempty"`

	// ScheduleTime: Required.
	//
	// The task's current schedule time, available in the
	// schedule_time returned by
	// LeaseTasks response or
	// RenewLease response. This restriction is
	// to ensure that your worker currently holds the lease.
	ScheduleTime string `json:"scheduleTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResponseView") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResponseView") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CancelLeaseRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CancelLeaseRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateTaskRequest: Request message for CreateTask.
type CreateTaskRequest struct {
	// ResponseView: The response_view specifies which subset of the Task
	// will be
	// returned.
	//
	// By default response_view is BASIC; not all
	// information is retrieved by default because some data, such
	// as
	// payloads, might be desirable to return only when needed because
	// of its large size or because of the sensitivity of data that
	// it
	// contains.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on
	// the
	// Task resource.
	//
	// Possible values:
	//   "VIEW_UNSPECIFIED" - Unspecified. Defaults to BASIC.
	//   "BASIC" - The basic view omits fields which can be large or can
	// contain
	// sensitive data.
	//
	// This view does not include the
	// (payload in AppEngineHttpRequest and
	// payload in PullMessage). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	//   "FULL" - All information is returned.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google
	// IAM](https://cloud.google.com/iam/)
	// permission on the Queue resource.
	ResponseView string `json:"responseView,omitempty"`

	// Task: Required.
	//
	// The task to add.
	//
	// Task names have the following
	// format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tas
	// ks/TASK_ID`.
	// The user can optionally specify a task name. If a
	// name is not specified then the system will generate a random
	// unique task id, which will be set in the task returned in
	// the
	// response.
	//
	// If schedule_time is not set or is in the
	// past then Cloud Tasks will set it to the current time.
	//
	// Task De-duplication:
	//
	// Explicitly specifying a task ID enables task de-duplication.  If
	// a task's ID is identical to that of an existing task or a task
	// that was deleted or completed recently then the call will fail
	// with ALREADY_EXISTS.
	// If the task's queue was created using Cloud Tasks, then another task
	// with
	// the same name can't be created for ~1hour after the original task
	// was
	// deleted or completed. If the task's queue was created using
	// queue.yaml or
	// queue.xml, then another task with the same name can't be created
	// for ~9days after the original task was deleted or completed.
	//
	// Because there is an extra lookup cost to identify duplicate
	// task
	// names, these CreateTask calls have significantly
	// increased latency. Using hashed strings for the task id or for
	// the prefix of the task id is recommended. Choosing task ids that
	// are sequential or have sequential prefixes, for example using
	// a
	// timestamp, causes an increase in latency and error rates in all
	// task commands. The infrastructure relies on an approximately
	// uniform distribution of task ids to store and serve
	// tasks
	// efficiently.
	Task *Task `json:"task,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResponseView") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResponseView") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateTaskRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateTaskRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// GetIamPolicyRequest: Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
}

// LeaseTasksRequest: Request message for leasing tasks using
// LeaseTasks.
type LeaseTasksRequest struct {
	// Filter: `filter` can be used to specify a subset of tasks to
	// lease.
	//
	// When `filter` is set to `tag=<my-tag>` then the
	// response will contain only tasks whose
	// tag is equal to `<my-tag>`. `<my-tag>` must be
	// less than 500 characters.
	//
	// When `filter` is set to `tag_function=oldest_tag()`, only tasks which
	// have
	// the same tag as the task with the oldest schedule_time will be
	// returned.
	//
	// Grammar Syntax:
	//
	// * `filter = "tag=" tag | "tag_function=" function`
	//
	// * `tag = string`
	//
	// * `function = "oldest_tag()"
	//
	// The `oldest_tag()` function returns tasks which have the same tag as
	// the
	// oldest task (ordered by schedule time).
	//
	// SDK compatibility: Although the SDK allows tags to be either
	// string
	// or
	// [bytes](/appengine/docs/standard/java/javadoc/com/google/appengine/
	// api/taskqueue/TaskOptions.html#tag-byte:A-),
	// only UTF-8 encoded tags can be used in Cloud Tasks. Tag which
	// aren't UTF-8 encoded can't be used in the
	// filter and the task's
	// tag will be displayed as empty in Cloud Tasks.
	Filter string `json:"filter,omitempty"`

	// LeaseDuration:
	// After the worker has successfully finished the work associated
	// with the task, the worker must call via
	// AcknowledgeTask before the
	// schedule_time. Otherwise the task will be
	// returned to a later LeaseTasks call so
	// that another worker can retry it.
	//
	// The maximum lease duration is 1 week.
	// `lease_duration` will be truncated to the nearest second.
	LeaseDuration string `json:"leaseDuration,omitempty"`

	// MaxTasks: The maximum number of tasks to lease. The maximum that can
	// be
	// requested is 1000.
	MaxTasks int64 `json:"maxTasks,omitempty"`

	// ResponseView: The response_view specifies which subset of the Task
	// will be
	// returned.
	//
	// By default response_view is BASIC; not all
	// information is retrieved by default because some data, such
	// as
	// payloads, might be desirable to return only when needed because
	// of its large size or because of the sensitivity of data that
	// it
	// contains.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on
	// the
	// Task resource.
	//
	// Possible values:
	//   "VIEW_UNSPECIFIED" - Unspecified. Defaults to BASIC.
	//   "BASIC" - The basic view omits fields which can be large or can
	// contain
	// sensitive data.
	//
	// This view does not include the
	// (payload in AppEngineHttpRequest and
	// payload in PullMessage). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	//   "FULL" - All information is returned.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google
	// IAM](https://cloud.google.com/iam/)
	// permission on the Queue resource.
	ResponseView string `json:"responseView,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Filter") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LeaseTasksRequest) MarshalJSON() ([]byte, error) {
	type NoMethod LeaseTasksRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LeaseTasksResponse: Response message for leasing tasks using
// LeaseTasks.
type LeaseTasksResponse struct {
	// Tasks: The leased tasks.
	Tasks []*Task `json:"tasks,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Tasks") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Tasks") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LeaseTasksResponse) MarshalJSON() ([]byte, error) {
	type NoMethod LeaseTasksResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListLocationsResponse: The response message for
// Locations.ListLocations.
type ListLocationsResponse struct {
	// Locations: A list of locations that matches the specified filter in
	// the request.
	Locations []*Location `json:"locations,omitempty"`

	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Locations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Locations") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListLocationsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListLocationsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListQueuesResponse: Response message for ListQueues.
type ListQueuesResponse struct {
	// NextPageToken: A token to retrieve next page of results.
	//
	// To return the next page of results, call
	// ListQueues with this value as the
	// page_token.
	//
	// If the next_page_token is empty, there are no more results.
	//
	// The page token is valid for only 2 hours.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Queues: The list of queues.
	Queues []*Queue `json:"queues,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListQueuesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListQueuesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListTasksResponse: Response message for listing tasks using
// ListTasks.
type ListTasksResponse struct {
	// NextPageToken: A token to retrieve next page of results.
	//
	// To return the next page of results, call
	// ListTasks with this value as the
	// page_token.
	//
	// If the next_page_token is empty, there are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Tasks: The list of tasks.
	Tasks []*Task `json:"tasks,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListTasksResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListTasksResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Location: A resource that represents Google Cloud Platform location.
type Location struct {
	// Labels: Cross-service attributes for the location. For example
	//
	//     {"cloud.googleapis.com/region": "us-east1"}
	Labels map[string]string `json:"labels,omitempty"`

	// LocationId: The canonical id for this location. For example:
	// "us-east1".
	LocationId string `json:"locationId,omitempty"`

	// Metadata: Service-specific metadata. For example the available
	// capacity at the given
	// location.
	Metadata googleapi.RawMessage `json:"metadata,omitempty"`

	// Name: Resource name for the location, which may vary between
	// implementations.
	// For example: "projects/example-project/locations/us-east1"
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Labels") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Location) MarshalJSON() ([]byte, error) {
	type NoMethod Location
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PauseQueueRequest: Request message for PauseQueue.
type PauseQueueRequest struct {
}

// Policy: Defines an Identity and Access Management (IAM) policy. It is
// used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `Binding` binds a list
// of
// `members` to a `role`, where the members can be user accounts, Google
// groups,
// Google domains, and service accounts. A `role` is a named list of
// permissions
// defined by IAM.
//
// **Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//
// "serviceAccount:my-other-app@appspot.gserviceaccount.com",
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam/docs).
type Policy struct {
	// Bindings: Associates a list of `members` to a `role`.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `json:"bindings,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a policy from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform policy updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `getIamPolicy`,
	// and
	// systems are expected to put that etag in the request to
	// `setIamPolicy` to
	// ensure that their change will be applied to the same version of the
	// policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the
	// existing
	// policy is overwritten blindly.
	Etag string `json:"etag,omitempty"`

	// Version: Deprecated.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Bindings") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bindings") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Policy) MarshalJSON() ([]byte, error) {
	type NoMethod Policy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PullMessage: The pull message contains data that can be used by the
// caller of
// LeaseTasks to process the task.
//
// This proto can only be used for tasks in a queue which
// has
// pull_target set.
type PullMessage struct {
	// Payload: A data payload consumed by the worker to execute the task.
	Payload string `json:"payload,omitempty"`

	// Tag: The task's tag.
	//
	// Tags allow similar tasks to be processed in a batch. If you
	// label
	// tasks with a tag, your worker can
	// lease tasks with the same tag using
	// filter. For example, if you want to
	// aggregate the events associated with a specific user once a day,
	// you could tag tasks with the user ID.
	//
	// The task's tag can only be set when the
	// task is created.
	//
	// The tag must be less than 500 characters.
	//
	// SDK compatibility: Although the SDK allows tags to be either
	// string or
	// [bytes](/appengine/docs/standard/java/javadoc/com/google/appengine/api
	// /taskqueue/TaskOptions.html#tag-byte:A-),
	// only UTF-8 encoded tags can be used in Cloud Tasks. If a tag isn't
	// UTF-8
	// encoded, the tag will be empty when the task is returned by Cloud
	// Tasks.
	Tag string `json:"tag,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Payload") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Payload") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *PullMessage) MarshalJSON() ([]byte, error) {
	type NoMethod PullMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PullTarget: Pull target.
type PullTarget struct {
}

// PurgeQueueRequest: Request message for PurgeQueue.
type PurgeQueueRequest struct {
}

// Queue: A queue is a container of related tasks. Queues are configured
// to manage
// how those tasks are dispatched. Configurable properties include rate
// limits,
// retry options, target types, and others.
type Queue struct {
	// AppEngineHttpTarget: App Engine HTTP target.
	//
	// An App Engine queue is a queue that has an AppEngineHttpTarget.
	AppEngineHttpTarget *AppEngineHttpTarget `json:"appEngineHttpTarget,omitempty"`

	// Name: The queue name.
	//
	// The queue name must have the following
	// format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`
	//
	// *
	//  `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	// projects](/resource-manager/docs/creating-managing-projects#identifyin
	// g_projects)
	// * `LOCATION_ID` is the canonical ID for the queue's location.
	//    The list of available locations can be obtained by calling
	//    ListLocations.
	//    For more information, see
	// https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100
	// characters.
	//
	// Caller-specified and required in CreateQueue,
	// after which it becomes output only.
	Name string `json:"name,omitempty"`

	// PullTarget: Pull target.
	//
	// A pull queue is a queue that has a PullTarget.
	PullTarget *PullTarget `json:"pullTarget,omitempty"`

	// PurgeTime: Output only. The last time this queue was purged.
	//
	// All tasks that were created before this time
	// were purged.
	//
	// A queue can be purged using PurgeQueue, the
	// [App Engine Task Queue SDK, or the Cloud
	// Console](/appengine/docs/standard/python/taskqueue/push/deleting-tasks
	// -and-queues#purging_all_tasks_from_a_queue).
	//
	// Purge time will be truncated to the nearest microsecond. Purge
	// time will be unset if the queue has never been purged.
	PurgeTime string `json:"purgeTime,omitempty"`

	// RateLimits: Rate limits for task dispatches.
	//
	// rate_limits and
	// retry_config are related because they both
	// control task attempts however they control how tasks are
	// attempted in different ways:
	//
	// * rate_limits controls the total rate of
	//   dispatches from a queue (i.e. all traffic dispatched from the
	//   queue, regardless of whether the dispatch is from a first
	//   attempt or a retry).
	// * retry_config controls what happens to
	//   particular a task after its first attempt fails. That is,
	//   retry_config controls task retries (the
	//   second attempt, third attempt, etc).
	RateLimits *RateLimits `json:"rateLimits,omitempty"`

	// RetryConfig: Settings that determine the retry behavior.
	//
	// * For tasks created using Cloud Tasks: the queue-level retry
	// settings
	//   apply to all tasks in the queue that were created using Cloud
	// Tasks.
	//   Retry settings cannot be set on individual tasks.
	// * For tasks created using the App Engine SDK: the queue-level retry
	//   settings apply to all tasks in the queue which do not have retry
	// settings
	//   explicitly set on the task and were created by the App Engine SDK.
	// See
	//   [App Engine
	// documentation](/appengine/docs/standard/python/taskqueue/push/retrying
	// -tasks).
	RetryConfig *RetryConfig `json:"retryConfig,omitempty"`

	// State: Output only. The state of the queue.
	//
	// `state` can only be changed by called
	// PauseQueue,
	// ResumeQueue, or
	// uploading
	// [queue.yaml/xml](/appengine/docs/python/config/queueref).
	// Up
	// dateQueue cannot be used to change `state`.
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - Unspecified state.
	//   "RUNNING" - The queue is running. Tasks can be dispatched.
	//
	// If the queue was created using Cloud Tasks and the queue has
	// had no activity (method calls or task dispatches) for 30 days,
	// the queue may take a few minutes to re-activate. Some method
	// calls may return NOT_FOUND and
	// tasks may not be dispatched for a few minutes until the queue
	// has been re-activated.
	//   "PAUSED" - Tasks are paused by the user. If the queue is paused
	// then Cloud
	// Tasks will stop delivering tasks from it, but more tasks can
	// still be added to it by the user. When a pull queue is paused,
	// all LeaseTasks calls will return a
	// FAILED_PRECONDITION.
	//   "DISABLED" - The queue is disabled.
	//
	// A queue becomes `DISABLED`
	// when
	// [queue.yaml](/appengine/docs/python/config/queueref)
	// or
	// [queue.xml](appengine/docs/standard/java/config/queueref) is
	// uploaded
	// which does not contain the queue. You cannot directly disable a
	// queue.
	//
	// When a queue is disabled, tasks can still be added to a queue
	// but the tasks are not dispatched and
	// LeaseTasks calls return a
	// `FAILED_PRECONDITION` error.
	//
	// To permanently delete this queue and all of its tasks,
	// call
	// DeleteQueue.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AppEngineHttpTarget")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppEngineHttpTarget") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Queue) MarshalJSON() ([]byte, error) {
	type NoMethod Queue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RateLimits: Rate limits.
//
// This message determines the maximum rate that tasks can be dispatched
// by a
// queue, regardless of whether the dispatch is a first task attempt or
// a retry.
type RateLimits struct {
	// MaxBurstSize: Output only. The max burst size.
	//
	// Max burst size limits how fast tasks in queue are processed when
	// many tasks are in the queue and the rate is high. This field
	// allows the queue to have a high rate so processing starts
	// shortly
	// after a task is enqueued, but still limits resource usage when
	// many tasks are enqueued in a short period of time.
	//
	// The [token bucket](https://wikipedia.org/wiki/Token_Bucket)
	// algorithm is used to control the rate of task dispatches. Each
	// queue has a token bucket that holds tokens, up to the
	// maximum
	// specified by `max_burst_size`. Each time a task is dispatched,
	// a
	// token is removed from the bucket. Tasks will be dispatched until
	// the queue's bucket runs out of tokens. The bucket will
	// be
	// continuously refilled with new tokens based
	// on
	// max_tasks_dispatched_per_second.
	//
	// Cloud Tasks will pick the value of `max_burst_size` based on
	// the
	// value of
	// max_tasks_dispatched_per_second.
	//
	// For App Engine queues that were created or updated
	// using
	// `queue.yaml/xml`, `max_burst_size` is equal
	// to
	// [bucket_size](/appengine/docs/standard/python/config/queueref#bucke
	// t_size).
	// Since `max_burst_size` is output only, if
	// UpdateQueue is called on a queue
	// created by `queue.yaml/xml`, `max_burst_size` will be reset based
	// on the value of
	// max_tasks_dispatched_per_second,
	// regardless of whether
	// max_tasks_dispatched_per_second
	// is updated.
	//
	MaxBurstSize int64 `json:"maxBurstSize,omitempty"`

	// MaxConcurrentTasks: The maximum number of concurrent tasks that Cloud
	// Tasks allows
	// to be dispatched for this queue. After this threshold has
	// been
	// reached, Cloud Tasks stops dispatching tasks until the number
	// of
	// concurrent requests decreases.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the
	// default.
	//
	//
	// The maximum allowed value is 5,000. -1 indicates no limit.
	//
	// This field is output only for
	// [pull queues](google.cloud.tasks.v2beta2.PullTarget).
	//
	//
	// This field has the same meaning as
	// [max_concurrent_requests in
	// queue.yaml/xml](/appengine/docs/standard/python/config/queueref#max_co
	// ncurrent_requests).
	MaxConcurrentTasks int64 `json:"maxConcurrentTasks,omitempty"`

	// MaxTasksDispatchedPerSecond: The maximum rate at which tasks are
	// dispatched from this queue.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the
	// default.
	//
	// * For App Engine queues, the maximum allowed value is 500.
	// * This field is output only   for [pull
	// queues](google.cloud.tasks.v2beta2.PullTarget). In
	//   addition to the `max_tasks_dispatched_per_second` limit, a
	//   maximum of 10 QPS of LeaseTasks
	//   requests are allowed per pull queue.
	//
	//
	// This field has the same meaning as
	// [rate in
	// queue.yaml/xml](/appengine/docs/standard/python/config/queueref#rate).
	MaxTasksDispatchedPerSecond float64 `json:"maxTasksDispatchedPerSecond,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaxBurstSize") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MaxBurstSize") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RateLimits) MarshalJSON() ([]byte, error) {
	type NoMethod RateLimits
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *RateLimits) UnmarshalJSON(data []byte) error {
	type NoMethod RateLimits
	var s1 struct {
		MaxTasksDispatchedPerSecond gensupport.JSONFloat64 `json:"maxTasksDispatchedPerSecond"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.MaxTasksDispatchedPerSecond = float64(s1.MaxTasksDispatchedPerSecond)
	return nil
}

// RenewLeaseRequest: Request message for renewing a lease
// using
// RenewLease.
type RenewLeaseRequest struct {
	// LeaseDuration: Required.
	//
	// The desired new lease duration, starting from now.
	//
	//
	// The maximum lease duration is 1 week.
	// `lease_duration` will be truncated to the nearest second.
	LeaseDuration string `json:"leaseDuration,omitempty"`

	// ResponseView: The response_view specifies which subset of the Task
	// will be
	// returned.
	//
	// By default response_view is BASIC; not all
	// information is retrieved by default because some data, such
	// as
	// payloads, might be desirable to return only when needed because
	// of its large size or because of the sensitivity of data that
	// it
	// contains.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on
	// the
	// Task resource.
	//
	// Possible values:
	//   "VIEW_UNSPECIFIED" - Unspecified. Defaults to BASIC.
	//   "BASIC" - The basic view omits fields which can be large or can
	// contain
	// sensitive data.
	//
	// This view does not include the
	// (payload in AppEngineHttpRequest and
	// payload in PullMessage). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	//   "FULL" - All information is returned.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google
	// IAM](https://cloud.google.com/iam/)
	// permission on the Queue resource.
	ResponseView string `json:"responseView,omitempty"`

	// ScheduleTime: Required.
	//
	// The task's current schedule time, available in the
	// schedule_time returned by
	// LeaseTasks response or
	// RenewLease response. This restriction is
	// to ensure that your worker currently holds the lease.
	ScheduleTime string `json:"scheduleTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LeaseDuration") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LeaseDuration") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RenewLeaseRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RenewLeaseRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResumeQueueRequest: Request message for ResumeQueue.
type ResumeQueueRequest struct {
}

// RetryConfig: Retry config.
//
// These settings determine how a failed task attempt is retried.
type RetryConfig struct {
	// MaxAttempts: The maximum number of attempts for a task.
	//
	// Cloud Tasks will attempt the task `max_attempts` times (that
	// is, if the first attempt fails, then there will be
	// `max_attempts - 1` retries).  Must be > 0.
	MaxAttempts int64 `json:"maxAttempts,omitempty"`

	// MaxBackoff: A task will be [scheduled](Task.schedule_time) for retry
	// between
	// min_backoff and
	// max_backoff duration after it fails,
	// if the queue's RetryConfig specifies that the task should
	// be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the
	// default.
	//
	// This field is output only for
	// [pull queues](google.cloud.tasks.v2beta2.PullTarget).
	//
	//
	// `max_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [max_backoff_seconds in
	// queue.yaml/xml](/appengine/docs/standard/python/config/queueref#retry_
	// parameters).
	MaxBackoff string `json:"maxBackoff,omitempty"`

	// MaxDoublings: The time between retries will double `max_doublings`
	// times.
	//
	// A task's retry interval starts at
	// min_backoff, then doubles
	// `max_doublings` times, then increases linearly, and finally
	// retries retries at intervals of
	// max_backoff up to
	// max_attempts times.
	//
	// For example, if min_backoff is 10s,
	// max_backoff is 300s, and
	// `max_doublings` is 3, then the a task will first be retried in
	// 10s. The retry interval will double three times, and then
	// increase linearly by 2^3 * 10s.  Finally, the task will retry
	// at
	// intervals of max_backoff until the
	// task has been attempted max_attempts
	// times. Thus, the requests will retry at 10s, 20s, 40s, 80s,
	// 160s,
	// 240s, 300s, 300s, ....
	//
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the
	// default.
	//
	// This field is output only for
	// [pull queues](google.cloud.tasks.v2beta2.PullTarget).
	//
	//
	// This field has the same meaning as
	// [max_doublings in
	// queue.yaml/xml](/appengine/docs/standard/python/config/queueref#retry_
	// parameters).
	MaxDoublings int64 `json:"maxDoublings,omitempty"`

	// MaxRetryDuration: If positive, `max_retry_duration` specifies the
	// time limit for
	// retrying a failed task, measured from when the task was
	// first
	// attempted. Once `max_retry_duration` time has passed *and* the
	// task has been attempted max_attempts
	// times, no further attempts will be made and the task will
	// be
	// deleted.
	//
	// If zero, then the task age is unlimited.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the
	// default.
	//
	// This field is output only for
	// [pull
	// queues](google.cloud.tasks.v2beta2.PullTarget).
	//
	//
	// `max_retry_duration`
	//  will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [task_age_limit in
	// queue.yaml/xml](/appengine/docs/standard/python/config/queueref#retry_
	// parameters).
	MaxRetryDuration string `json:"maxRetryDuration,omitempty"`

	// MinBackoff: A task will be [scheduled](Task.schedule_time) for retry
	// between
	// min_backoff and
	// max_backoff duration after it fails,
	// if the queue's RetryConfig specifies that the task should
	// be
	// retried.
	//
	// If unspecified when the queue is created, Cloud Tasks will pick
	// the
	// default.
	//
	// This field is output only for
	// [pull queues](google.cloud.tasks.v2beta2.PullTarget).
	//
	//
	// `min_backoff` will be truncated to the nearest second.
	//
	// This field has the same meaning as
	// [min_backoff_seconds in
	// queue.yaml/xml](/appengine/docs/standard/python/config/queueref#retry_
	// parameters).
	MinBackoff string `json:"minBackoff,omitempty"`

	// UnlimitedAttempts: If true, then the number of attempts is unlimited.
	UnlimitedAttempts bool `json:"unlimitedAttempts,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaxAttempts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MaxAttempts") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RetryConfig) MarshalJSON() ([]byte, error) {
	type NoMethod RetryConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RunTaskRequest: Request message for forcing a task to run now
// using
// RunTask.
type RunTaskRequest struct {
	// ResponseView: The response_view specifies which subset of the Task
	// will be
	// returned.
	//
	// By default response_view is BASIC; not all
	// information is retrieved by default because some data, such
	// as
	// payloads, might be desirable to return only when needed because
	// of its large size or because of the sensitivity of data that
	// it
	// contains.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on
	// the
	// Task resource.
	//
	// Possible values:
	//   "VIEW_UNSPECIFIED" - Unspecified. Defaults to BASIC.
	//   "BASIC" - The basic view omits fields which can be large or can
	// contain
	// sensitive data.
	//
	// This view does not include the
	// (payload in AppEngineHttpRequest and
	// payload in PullMessage). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	//   "FULL" - All information is returned.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google
	// IAM](https://cloud.google.com/iam/)
	// permission on the Queue resource.
	ResponseView string `json:"responseView,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ResponseView") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResponseView") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RunTaskRequest) MarshalJSON() ([]byte, error) {
	type NoMethod RunTaskRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SetIamPolicyRequest: Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as
	// Projects)
	// might reject them.
	Policy *Policy `json:"policy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Policy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SetIamPolicyRequest) MarshalJSON() ([]byte, error) {
	type NoMethod SetIamPolicyRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` that can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There is a
	// common set of
	// message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Task: A unit of scheduled work.
type Task struct {
	// AppEngineHttpRequest: App Engine HTTP request that is sent to the
	// task's target. Can
	// be set only if
	// app_engine_http_target is set
	// on the queue.
	//
	// An App Engine task is a task that has AppEngineHttpRequest set.
	AppEngineHttpRequest *AppEngineHttpRequest `json:"appEngineHttpRequest,omitempty"`

	// CreateTime: Output only. The time that the task was
	// created.
	//
	// `create_time` will be truncated to the nearest second.
	CreateTime string `json:"createTime,omitempty"`

	// Name: The task name.
	//
	// The task name must have the following
	// format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tas
	// ks/TASK_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	// projects](/resource-manager/docs/creating-managing-projects#identifyin
	// g_projects)
	// * `LOCATION_ID` is the canonical ID for the task's location.
	//    The list of available locations can be obtained by calling
	//    ListLocations.
	//    For more information, see
	// https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	// * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//   hyphens (-), or underscores (_). The maximum length is 500
	// characters.
	//
	// Optionally caller-specified in CreateTask.
	Name string `json:"name,omitempty"`

	// PullMessage: LeaseTasks to process the task. Can be
	// set only if pull_target is set on the queue.
	//
	// A pull task is a task that has PullMessage set.
	PullMessage *PullMessage `json:"pullMessage,omitempty"`

	// ScheduleTime: The time when the task is scheduled to be
	// attempted.
	//
	// For App Engine queues, this is when the task will be attempted or
	// retried.
	//
	// For pull queues, this is the time when the task is available to
	// be leased; if a task is currently leased, this is the time when
	// the current lease expires, that is, the time that the task was
	// leased plus the lease_duration.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime string `json:"scheduleTime,omitempty"`

	// Status: Output only. The task status.
	Status *TaskStatus `json:"status,omitempty"`

	// View: Output only. The view specifies which subset of the Task
	// has
	// been returned.
	//
	// Possible values:
	//   "VIEW_UNSPECIFIED" - Unspecified. Defaults to BASIC.
	//   "BASIC" - The basic view omits fields which can be large or can
	// contain
	// sensitive data.
	//
	// This view does not include the
	// (payload in AppEngineHttpRequest and
	// payload in PullMessage). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	//   "FULL" - All information is returned.
	//
	// Authorization for FULL requires
	// `cloudtasks.tasks.fullView` [Google
	// IAM](https://cloud.google.com/iam/)
	// permission on the Queue resource.
	View string `json:"view,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AppEngineHttpRequest") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AppEngineHttpRequest") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Task) MarshalJSON() ([]byte, error) {
	type NoMethod Task
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TaskStatus: Status of the task.
type TaskStatus struct {
	// AttemptDispatchCount: Output only. The number of attempts
	// dispatched.
	//
	// This count includes tasks which have been dispatched but
	// haven't
	// received a response.
	AttemptDispatchCount int64 `json:"attemptDispatchCount,omitempty"`

	// AttemptResponseCount: Output only. The number of attempts which have
	// received a response.
	//
	// This field is not calculated for
	// [pull tasks](google.cloud.tasks.v2beta2.PullTaskTarget).
	AttemptResponseCount int64 `json:"attemptResponseCount,omitempty"`

	// FirstAttemptStatus: Output only. The status of the task's first
	// attempt.
	//
	// Only dispatch_time will be set.
	// The other AttemptStatus information is not retained by Cloud
	// Tasks.
	//
	// This field is not calculated for
	// [pull tasks](google.cloud.tasks.v2beta2.PullTaskTarget).
	FirstAttemptStatus *AttemptStatus `json:"firstAttemptStatus,omitempty"`

	// LastAttemptStatus: Output only. The status of the task's last
	// attempt.
	//
	// This field is not calculated for
	// [pull tasks](google.cloud.tasks.v2beta2.PullTaskTarget).
	LastAttemptStatus *AttemptStatus `json:"lastAttemptStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AttemptDispatchCount") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AttemptDispatchCount") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TaskStatus) MarshalJSON() ([]byte, error) {
	type NoMethod TaskStatus
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsRequest: Request message for `TestIamPermissions`
// method.
type TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the `resource`.
	// Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For
	// more
	// information see
	// [IAM
	// Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod TestIamPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TestIamPermissionsResponse: Response message for `TestIamPermissions`
// method.
type TestIamPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is
	// allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod TestIamPermissionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "cloudtasks.projects.locations.get":

type ProjectsLocationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get information about a location.
func (r *ProjectsLocationsService) Get(name string) *ProjectsLocationsGetCall {
	c := &ProjectsLocationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsGetCall) Context(ctx context.Context) *ProjectsLocationsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.get" call.
// Exactly one of *Location or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Location.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsLocationsGetCall) Do(opts ...googleapi.CallOption) (*Location, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Location{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get information about a location.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudtasks.projects.locations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Resource name for the location.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}",
	//   "response": {
	//     "$ref": "Location"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.list":

type ProjectsLocationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists information about the supported locations for this
// service.
func (r *ProjectsLocationsService) List(name string) *ProjectsLocationsListCall {
	c := &ProjectsLocationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsLocationsListCall) Filter(filter string) *ProjectsLocationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsLocationsListCall) PageSize(pageSize int64) *ProjectsLocationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsLocationsListCall) PageToken(pageToken string) *ProjectsLocationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsListCall) Context(ctx context.Context) *ProjectsLocationsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}/locations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.list" call.
// Exactly one of *ListLocationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLocationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsListCall) Do(opts ...googleapi.CallOption) (*ListLocationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListLocationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists information about the supported locations for this service.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations",
	//   "httpMethod": "GET",
	//   "id": "cloudtasks.projects.locations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The standard list filter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The resource that owns the locations collection, if applicable.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard list page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard list page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}/locations",
	//   "response": {
	//     "$ref": "ListLocationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsListCall) Pages(ctx context.Context, f func(*ListLocationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudtasks.projects.locations.queues.create":

type ProjectsLocationsQueuesCreateCall struct {
	s          *Service
	parent     string
	queue      *Queue
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a queue.
//
// Queues created with this method allow tasks to live for a maximum of
// 31
// days. After a task is 31 days old, the task will be deleted
// regardless of whether
// it was dispatched or not.
//
// WARNING: Using this method may have unintended side effects if you
// are
// using an App Engine `queue.yaml` or `queue.xml` file to manage your
// queues.
// Read
// [Overview of Queue Management and
// queue.yaml](/cloud-tasks/docs/queue-yaml)
// before using this method.
func (r *ProjectsLocationsQueuesService) Create(parent string, queue *Queue) *ProjectsLocationsQueuesCreateCall {
	c := &ProjectsLocationsQueuesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.queue = queue
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesCreateCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesCreateCall) Context(ctx context.Context) *ProjectsLocationsQueuesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.queue)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+parent}/queues")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.create" call.
// Exactly one of *Queue or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Queue.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesCreateCall) Do(opts ...googleapi.CallOption) (*Queue, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Queue{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a queue.\n\nQueues created with this method allow tasks to live for a maximum of 31\ndays. After a task is 31 days old, the task will be deleted regardless of whether\nit was dispatched or not.\n\nWARNING: Using this method may have unintended side effects if you are\nusing an App Engine `queue.yaml` or `queue.xml` file to manage your queues.\nRead\n[Overview of Queue Management and queue.yaml](/cloud-tasks/docs/queue-yaml)\nbefore using this method.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required.\n\nThe location name in which the queue will be created.\nFor example: `projects/PROJECT_ID/locations/LOCATION_ID`\n\nThe list of allowed locations can be obtained by calling Cloud\nTasks' implementation of\nListLocations.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+parent}/queues",
	//   "request": {
	//     "$ref": "Queue"
	//   },
	//   "response": {
	//     "$ref": "Queue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.delete":

type ProjectsLocationsQueuesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a queue.
//
// This command will delete the queue even if it has tasks in it.
//
// Note: If you delete a queue, a queue with the same name can't be
// created
// for 7 days.
//
// WARNING: Using this method may have unintended side effects if you
// are
// using an App Engine `queue.yaml` or `queue.xml` file to manage your
// queues.
// Read
// [Overview of Queue Management and
// queue.yaml](/cloud-tasks/docs/queue-yaml)
// before using this method.
func (r *ProjectsLocationsQueuesService) Delete(name string) *ProjectsLocationsQueuesDeleteCall {
	c := &ProjectsLocationsQueuesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesDeleteCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesDeleteCall) Context(ctx context.Context) *ProjectsLocationsQueuesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a queue.\n\nThis command will delete the queue even if it has tasks in it.\n\nNote: If you delete a queue, a queue with the same name can't be created\nfor 7 days.\n\nWARNING: Using this method may have unintended side effects if you are\nusing an App Engine `queue.yaml` or `queue.xml` file to manage your queues.\nRead\n[Overview of Queue Management and queue.yaml](/cloud-tasks/docs/queue-yaml)\nbefore using this method.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudtasks.projects.locations.queues.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe queue name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.get":

type ProjectsLocationsQueuesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a queue.
func (r *ProjectsLocationsQueuesService) Get(name string) *ProjectsLocationsQueuesGetCall {
	c := &ProjectsLocationsQueuesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsQueuesGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsQueuesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesGetCall) Context(ctx context.Context) *ProjectsLocationsQueuesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.get" call.
// Exactly one of *Queue or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Queue.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesGetCall) Do(opts ...googleapi.CallOption) (*Queue, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Queue{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a queue.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}",
	//   "httpMethod": "GET",
	//   "id": "cloudtasks.projects.locations.queues.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe resource name of the queue. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}",
	//   "response": {
	//     "$ref": "Queue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.getIamPolicy":

type ProjectsLocationsQueuesGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// GetIamPolicy: Gets the access control policy for a Queue.
// Returns an empty policy if the resource exists and does not have a
// policy
// set.
//
// Authorization requires the following [Google IAM](/iam) permission on
// the
// specified resource parent:
//
// * `cloudtasks.queues.getIamPolicy`
func (r *ProjectsLocationsQueuesService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsLocationsQueuesGetIamPolicyCall {
	c := &ProjectsLocationsQueuesGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesGetIamPolicyCall) Context(ctx context.Context) *ProjectsLocationsQueuesGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesGetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a Queue.\nReturns an empty policy if the resource exists and does not have a policy\nset.\n\nAuthorization requires the following [Google IAM](/iam) permission on the\nspecified resource parent:\n\n* `cloudtasks.queues.getIamPolicy`",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.list":

type ProjectsLocationsQueuesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists queues.
//
// Queues are returned in lexicographical order.
func (r *ProjectsLocationsQueuesService) List(parent string) *ProjectsLocationsQueuesListCall {
	c := &ProjectsLocationsQueuesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": `filter` can be used to
// specify a subset of queues. Any Queue
// field can be used as a filter and several operators as supported.
// For example: `<=, <, >=, >, !=, =, :`. The filter syntax is the same
// as
// described in
// [Stackdriver's Advanced Logs
// Filters](/logging/docs/view/advanced_filters).
//
// Sample filter "app_engine_http_target: *".
//
// Note that using filters might cause fewer queues than
// the
// requested_page size to be returned.
func (c *ProjectsLocationsQueuesListCall) Filter(filter string) *ProjectsLocationsQueuesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page
// size.
//
// The maximum page size is 9800. If unspecified, the page size will
// be the maximum. Fewer queues than requested might be returned,
// even if more queues exist; use the
// next_page_token in the
// response to determine if more queues exist.
func (c *ProjectsLocationsQueuesListCall) PageSize(pageSize int64) *ProjectsLocationsQueuesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying the page of results to return.
//
// To request the first page results, page_token must be empty.
// To
// request the next page of results, page_token must be the value
// of
// next_page_token returned
// from the previous call to ListQueues
// method. It is an error to switch the value of the
// filter while iterating through pages.
func (c *ProjectsLocationsQueuesListCall) PageToken(pageToken string) *ProjectsLocationsQueuesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesListCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsQueuesListCall) IfNoneMatch(entityTag string) *ProjectsLocationsQueuesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesListCall) Context(ctx context.Context) *ProjectsLocationsQueuesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+parent}/queues")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.list" call.
// Exactly one of *ListQueuesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListQueuesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsQueuesListCall) Do(opts ...googleapi.CallOption) (*ListQueuesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListQueuesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists queues.\n\nQueues are returned in lexicographical order.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues",
	//   "httpMethod": "GET",
	//   "id": "cloudtasks.projects.locations.queues.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "`filter` can be used to specify a subset of queues. Any Queue\nfield can be used as a filter and several operators as supported.\nFor example: `\u003c=, \u003c, \u003e=, \u003e, !=, =, :`. The filter syntax is the same as\ndescribed in\n[Stackdriver's Advanced Logs Filters](/logging/docs/view/advanced_filters).\n\nSample filter \"app_engine_http_target: *\".\n\nNote that using filters might cause fewer queues than the\nrequested_page size to be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size.\n\nThe maximum page size is 9800. If unspecified, the page size will\nbe the maximum. Fewer queues than requested might be returned,\neven if more queues exist; use the\nnext_page_token in the\nresponse to determine if more queues exist.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying the page of results to return.\n\nTo request the first page results, page_token must be empty. To\nrequest the next page of results, page_token must be the value of\nnext_page_token returned\nfrom the previous call to ListQueues\nmethod. It is an error to switch the value of the\nfilter while iterating through pages.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required.\n\nThe location name.\nFor example: `projects/PROJECT_ID/locations/LOCATION_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+parent}/queues",
	//   "response": {
	//     "$ref": "ListQueuesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsQueuesListCall) Pages(ctx context.Context, f func(*ListQueuesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudtasks.projects.locations.queues.patch":

type ProjectsLocationsQueuesPatchCall struct {
	s          *Service
	name       string
	queue      *Queue
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Updates a queue.
//
// This method creates the queue if it does not exist and updates
// the queue if it does exist.
//
// Queues created with this method allow tasks to live for a maximum of
// 31
// days. After a task is 31 days old, the task will be deleted
// regardless of whether
// it was dispatched or not.
//
// WARNING: Using this method may have unintended side effects if you
// are
// using an App Engine `queue.yaml` or `queue.xml` file to manage your
// queues.
// Read
// [Overview of Queue Management and
// queue.yaml](/cloud-tasks/docs/queue-yaml)
// before using this method.
func (r *ProjectsLocationsQueuesService) Patch(name string, queue *Queue) *ProjectsLocationsQueuesPatchCall {
	c := &ProjectsLocationsQueuesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.queue = queue
	return c
}

// UpdateMask sets the optional parameter "updateMask": A mask used to
// specify which fields of the queue are being updated.
//
// If empty, then all fields will be updated.
func (c *ProjectsLocationsQueuesPatchCall) UpdateMask(updateMask string) *ProjectsLocationsQueuesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesPatchCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesPatchCall) Context(ctx context.Context) *ProjectsLocationsQueuesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.queue)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.patch" call.
// Exactly one of *Queue or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Queue.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesPatchCall) Do(opts ...googleapi.CallOption) (*Queue, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Queue{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a queue.\n\nThis method creates the queue if it does not exist and updates\nthe queue if it does exist.\n\nQueues created with this method allow tasks to live for a maximum of 31\ndays. After a task is 31 days old, the task will be deleted regardless of whether\nit was dispatched or not.\n\nWARNING: Using this method may have unintended side effects if you are\nusing an App Engine `queue.yaml` or `queue.xml` file to manage your queues.\nRead\n[Overview of Queue Management and queue.yaml](/cloud-tasks/docs/queue-yaml)\nbefore using this method.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}",
	//   "httpMethod": "PATCH",
	//   "id": "cloudtasks.projects.locations.queues.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The queue name.\n\nThe queue name must have the following format:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`\n\n* `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),\n   hyphens (-), colons (:), or periods (.).\n   For more information, see\n   [Identifying projects](/resource-manager/docs/creating-managing-projects#identifying_projects)\n* `LOCATION_ID` is the canonical ID for the queue's location.\n   The list of available locations can be obtained by calling\n   ListLocations.\n   For more information, see https://cloud.google.com/about/locations/.\n* `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or\n  hyphens (-). The maximum length is 100 characters.\n\nCaller-specified and required in CreateQueue,\nafter which it becomes output only.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "A mask used to specify which fields of the queue are being updated.\n\nIf empty, then all fields will be updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}",
	//   "request": {
	//     "$ref": "Queue"
	//   },
	//   "response": {
	//     "$ref": "Queue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.pause":

type ProjectsLocationsQueuesPauseCall struct {
	s                 *Service
	name              string
	pausequeuerequest *PauseQueueRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Pause: Pauses the queue.
//
// If a queue is paused then the system will stop dispatching
// tasks
// until the queue is resumed via
// ResumeQueue. Tasks can still be added
// when the queue is paused. A queue is paused if its
// state is PAUSED.
func (r *ProjectsLocationsQueuesService) Pause(name string, pausequeuerequest *PauseQueueRequest) *ProjectsLocationsQueuesPauseCall {
	c := &ProjectsLocationsQueuesPauseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.pausequeuerequest = pausequeuerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesPauseCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesPauseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesPauseCall) Context(ctx context.Context) *ProjectsLocationsQueuesPauseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesPauseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesPauseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.pausequeuerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}:pause")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.pause" call.
// Exactly one of *Queue or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Queue.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesPauseCall) Do(opts ...googleapi.CallOption) (*Queue, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Queue{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Pauses the queue.\n\nIf a queue is paused then the system will stop dispatching tasks\nuntil the queue is resumed via\nResumeQueue. Tasks can still be added\nwhen the queue is paused. A queue is paused if its\nstate is PAUSED.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}:pause",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.pause",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe queue name. For example:\n`projects/PROJECT_ID/location/LOCATION_ID/queues/QUEUE_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}:pause",
	//   "request": {
	//     "$ref": "PauseQueueRequest"
	//   },
	//   "response": {
	//     "$ref": "Queue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.purge":

type ProjectsLocationsQueuesPurgeCall struct {
	s                 *Service
	name              string
	purgequeuerequest *PurgeQueueRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Purge: Purges a queue by deleting all of its tasks.
//
// All tasks created before this method is called are permanently
// deleted.
//
// Purge operations can take up to one minute to take effect.
// Tasks
// might be dispatched before the purge takes effect. A purge is
// irreversible.
func (r *ProjectsLocationsQueuesService) Purge(name string, purgequeuerequest *PurgeQueueRequest) *ProjectsLocationsQueuesPurgeCall {
	c := &ProjectsLocationsQueuesPurgeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.purgequeuerequest = purgequeuerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesPurgeCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesPurgeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesPurgeCall) Context(ctx context.Context) *ProjectsLocationsQueuesPurgeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesPurgeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesPurgeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.purgequeuerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}:purge")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.purge" call.
// Exactly one of *Queue or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Queue.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesPurgeCall) Do(opts ...googleapi.CallOption) (*Queue, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Queue{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Purges a queue by deleting all of its tasks.\n\nAll tasks created before this method is called are permanently deleted.\n\nPurge operations can take up to one minute to take effect. Tasks\nmight be dispatched before the purge takes effect. A purge is irreversible.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}:purge",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.purge",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe queue name. For example:\n`projects/PROJECT_ID/location/LOCATION_ID/queues/QUEUE_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}:purge",
	//   "request": {
	//     "$ref": "PurgeQueueRequest"
	//   },
	//   "response": {
	//     "$ref": "Queue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.resume":

type ProjectsLocationsQueuesResumeCall struct {
	s                  *Service
	name               string
	resumequeuerequest *ResumeQueueRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// Resume: Resume a queue.
//
// This method resumes a queue after it has been
// PAUSED or
// DISABLED. The state of a queue is stored
// in the queue's state; after calling this method it
// will be set to RUNNING.
//
// WARNING: Resuming many high-QPS queues at the same time can
// lead to target overloading. If you are resuming high-QPS
// queues, follow the 500/50/5 pattern described in
// [Managing Cloud Tasks Scaling
// Risks](/cloud-tasks/pdfs/managing-cloud-tasks-scaling-risks-2017-06-05
// .pdf).
func (r *ProjectsLocationsQueuesService) Resume(name string, resumequeuerequest *ResumeQueueRequest) *ProjectsLocationsQueuesResumeCall {
	c := &ProjectsLocationsQueuesResumeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.resumequeuerequest = resumequeuerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesResumeCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesResumeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesResumeCall) Context(ctx context.Context) *ProjectsLocationsQueuesResumeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesResumeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesResumeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.resumequeuerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}:resume")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.resume" call.
// Exactly one of *Queue or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Queue.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesResumeCall) Do(opts ...googleapi.CallOption) (*Queue, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Queue{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Resume a queue.\n\nThis method resumes a queue after it has been\nPAUSED or\nDISABLED. The state of a queue is stored\nin the queue's state; after calling this method it\nwill be set to RUNNING.\n\nWARNING: Resuming many high-QPS queues at the same time can\nlead to target overloading. If you are resuming high-QPS\nqueues, follow the 500/50/5 pattern described in\n[Managing Cloud Tasks Scaling Risks](/cloud-tasks/pdfs/managing-cloud-tasks-scaling-risks-2017-06-05.pdf).",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}:resume",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.resume",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe queue name. For example:\n`projects/PROJECT_ID/location/LOCATION_ID/queues/QUEUE_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}:resume",
	//   "request": {
	//     "$ref": "ResumeQueueRequest"
	//   },
	//   "response": {
	//     "$ref": "Queue"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.setIamPolicy":

type ProjectsLocationsQueuesSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// SetIamPolicy: Sets the access control policy for a Queue. Replaces
// any existing
// policy.
//
// Note: The Cloud Console does not check queue-level IAM permissions
// yet.
// Project-level permissions are required to use the Cloud
// Console.
//
// Authorization requires the following [Google IAM](/iam) permission on
// the
// specified resource parent:
//
// * `cloudtasks.queues.setIamPolicy`
func (r *ProjectsLocationsQueuesService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsLocationsQueuesSetIamPolicyCall {
	c := &ProjectsLocationsQueuesSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesSetIamPolicyCall) Context(ctx context.Context) *ProjectsLocationsQueuesSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesSetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy for a Queue. Replaces any existing\npolicy.\n\nNote: The Cloud Console does not check queue-level IAM permissions yet.\nProject-level permissions are required to use the Cloud Console.\n\nAuthorization requires the following [Google IAM](/iam) permission on the\nspecified resource parent:\n\n* `cloudtasks.queues.setIamPolicy`",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.testIamPermissions":

type ProjectsLocationsQueuesTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
	header_                   http.Header
}

// TestIamPermissions: Returns permissions that a caller has on a
// Queue.
// If the resource does not exist, this will return an empty set
// of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware
// UIs and command-line tools, not for authorization checking. This
// operation
// may "fail open" without warning.
func (r *ProjectsLocationsQueuesService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsLocationsQueuesTestIamPermissionsCall {
	c := &ProjectsLocationsQueuesTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTestIamPermissionsCall) Context(ctx context.Context) *ProjectsLocationsQueuesTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsQueuesTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on a Queue.\nIf the resource does not exist, this will return an empty set of\npermissions, not a NOT_FOUND error.\n\nNote: This operation is designed to be used for building permission-aware\nUIs and command-line tools, not for authorization checking. This operation\nmay \"fail open\" without warning.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.tasks.acknowledge":

type ProjectsLocationsQueuesTasksAcknowledgeCall struct {
	s                      *Service
	name                   string
	acknowledgetaskrequest *AcknowledgeTaskRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
	header_                http.Header
}

// Acknowledge: Acknowledges a pull task.
//
// The worker, that is, the entity that
// leased this task must call this method
// to indicate that the work associated with the task has finished.
//
// The worker must acknowledge a task within the
// lease_duration or the lease
// will expire and the task will become available to be leased
// again. After the task is acknowledged, it will not be returned
// by a later LeaseTasks,
// GetTask, or
// ListTasks.
//
// To acknowledge multiple tasks at the same time, use
// [HTTP batching](/storage/docs/json_api/v1/how-tos/batch)
// or the batching documentation for your client library, for
// example
// https://developers.google.com/api-client-library/python/guide/
// batch.
func (r *ProjectsLocationsQueuesTasksService) Acknowledge(name string, acknowledgetaskrequest *AcknowledgeTaskRequest) *ProjectsLocationsQueuesTasksAcknowledgeCall {
	c := &ProjectsLocationsQueuesTasksAcknowledgeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.acknowledgetaskrequest = acknowledgetaskrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksAcknowledgeCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksAcknowledgeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksAcknowledgeCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksAcknowledgeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksAcknowledgeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksAcknowledgeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.acknowledgetaskrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}:acknowledge")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.acknowledge" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesTasksAcknowledgeCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Acknowledges a pull task.\n\nThe worker, that is, the entity that\nleased this task must call this method\nto indicate that the work associated with the task has finished.\n\nThe worker must acknowledge a task within the\nlease_duration or the lease\nwill expire and the task will become available to be leased\nagain. After the task is acknowledged, it will not be returned\nby a later LeaseTasks,\nGetTask, or\nListTasks.\n\nTo acknowledge multiple tasks at the same time, use\n[HTTP batching](/storage/docs/json_api/v1/how-tos/batch)\nor the batching documentation for your client library, for example\nhttps://developers.google.com/api-client-library/python/guide/batch.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks/{tasksId}:acknowledge",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.tasks.acknowledge",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe task name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+/tasks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}:acknowledge",
	//   "request": {
	//     "$ref": "AcknowledgeTaskRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.tasks.cancelLease":

type ProjectsLocationsQueuesTasksCancelLeaseCall struct {
	s                  *Service
	name               string
	cancelleaserequest *CancelLeaseRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
	header_            http.Header
}

// CancelLease: Cancel a pull task's lease.
//
// The worker can use this method to cancel a task's lease by
// setting its schedule_time to now. This will
// make the task available to be leased to the next caller
// of
// LeaseTasks.
func (r *ProjectsLocationsQueuesTasksService) CancelLease(name string, cancelleaserequest *CancelLeaseRequest) *ProjectsLocationsQueuesTasksCancelLeaseCall {
	c := &ProjectsLocationsQueuesTasksCancelLeaseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.cancelleaserequest = cancelleaserequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksCancelLeaseCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksCancelLeaseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksCancelLeaseCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksCancelLeaseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksCancelLeaseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksCancelLeaseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cancelleaserequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}:cancelLease")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.cancelLease" call.
// Exactly one of *Task or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Task.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsQueuesTasksCancelLeaseCall) Do(opts ...googleapi.CallOption) (*Task, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Task{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancel a pull task's lease.\n\nThe worker can use this method to cancel a task's lease by\nsetting its schedule_time to now. This will\nmake the task available to be leased to the next caller of\nLeaseTasks.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks/{tasksId}:cancelLease",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.tasks.cancelLease",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe task name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+/tasks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}:cancelLease",
	//   "request": {
	//     "$ref": "CancelLeaseRequest"
	//   },
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.tasks.create":

type ProjectsLocationsQueuesTasksCreateCall struct {
	s                 *Service
	parent            string
	createtaskrequest *CreateTaskRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Create: Creates a task and adds it to a queue.
//
// To add multiple tasks at the same time, use
// [HTTP batching](/storage/docs/json_api/v1/how-tos/batch)
// or the batching documentation for your client library, for
// example
// https://developers.google.com/api-client-library/python/guide/
// batch.
//
// Tasks cannot be updated after creation; there is no UpdateTask
// command.
//
// * For [App Engine
// queues](google.cloud.tasks.v2beta2.AppEngineHttpTarget),
//   the maximum task size is 100KB.
// * For [pull queues](google.cloud.tasks.v2beta2.PullTarget), this
//   the maximum task size is 1MB.
func (r *ProjectsLocationsQueuesTasksService) Create(parent string, createtaskrequest *CreateTaskRequest) *ProjectsLocationsQueuesTasksCreateCall {
	c := &ProjectsLocationsQueuesTasksCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.createtaskrequest = createtaskrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksCreateCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksCreateCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createtaskrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+parent}/tasks")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.create" call.
// Exactly one of *Task or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Task.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsQueuesTasksCreateCall) Do(opts ...googleapi.CallOption) (*Task, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Task{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a task and adds it to a queue.\n\nTo add multiple tasks at the same time, use\n[HTTP batching](/storage/docs/json_api/v1/how-tos/batch)\nor the batching documentation for your client library, for example\nhttps://developers.google.com/api-client-library/python/guide/batch.\n\nTasks cannot be updated after creation; there is no UpdateTask command.\n\n* For [App Engine queues](google.cloud.tasks.v2beta2.AppEngineHttpTarget),\n  the maximum task size is 100KB.\n* For [pull queues](google.cloud.tasks.v2beta2.PullTarget), this\n  the maximum task size is 1MB.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.tasks.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required.\n\nThe queue name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`\n\nThe queue must already exist.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+parent}/tasks",
	//   "request": {
	//     "$ref": "CreateTaskRequest"
	//   },
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.tasks.delete":

type ProjectsLocationsQueuesTasksDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a task.
//
// A task can be deleted if it is scheduled or dispatched. A task
// cannot be deleted if it has completed successfully or
// permanently
// failed.
func (r *ProjectsLocationsQueuesTasksService) Delete(name string) *ProjectsLocationsQueuesTasksDeleteCall {
	c := &ProjectsLocationsQueuesTasksDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksDeleteCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLocationsQueuesTasksDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a task.\n\nA task can be deleted if it is scheduled or dispatched. A task\ncannot be deleted if it has completed successfully or permanently\nfailed.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks/{tasksId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudtasks.projects.locations.queues.tasks.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe task name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+/tasks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.tasks.get":

type ProjectsLocationsQueuesTasksGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a task.
func (r *ProjectsLocationsQueuesTasksService) Get(name string) *ProjectsLocationsQueuesTasksGetCall {
	c := &ProjectsLocationsQueuesTasksGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// ResponseView sets the optional parameter "responseView": The
// response_view specifies which subset of the Task will
// be
// returned.
//
// By default response_view is BASIC; not all
// information is retrieved by default because some data, such
// as
// payloads, might be desirable to return only when needed because
// of its large size or because of the sensitivity of data that
// it
// contains.
//
// Authorization for FULL requires
// `cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on
// the
// Task resource.
//
// Possible values:
//   "VIEW_UNSPECIFIED"
//   "BASIC"
//   "FULL"
func (c *ProjectsLocationsQueuesTasksGetCall) ResponseView(responseView string) *ProjectsLocationsQueuesTasksGetCall {
	c.urlParams_.Set("responseView", responseView)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksGetCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsQueuesTasksGetCall) IfNoneMatch(entityTag string) *ProjectsLocationsQueuesTasksGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksGetCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.get" call.
// Exactly one of *Task or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Task.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsQueuesTasksGetCall) Do(opts ...googleapi.CallOption) (*Task, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Task{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a task.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks/{tasksId}",
	//   "httpMethod": "GET",
	//   "id": "cloudtasks.projects.locations.queues.tasks.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe task name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+/tasks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "responseView": {
	//       "description": "The response_view specifies which subset of the Task will be\nreturned.\n\nBy default response_view is BASIC; not all\ninformation is retrieved by default because some data, such as\npayloads, might be desirable to return only when needed because\nof its large size or because of the sensitivity of data that it\ncontains.\n\nAuthorization for FULL requires\n`cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on the\nTask resource.",
	//       "enum": [
	//         "VIEW_UNSPECIFIED",
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}",
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.tasks.lease":

type ProjectsLocationsQueuesTasksLeaseCall struct {
	s                 *Service
	parent            string
	leasetasksrequest *LeaseTasksRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Lease: Leases tasks from a pull queue for
// lease_duration.
//
// This method is invoked by the worker to obtain a lease. The
// worker must acknowledge the task via
// AcknowledgeTask after they have
// performed the work associated with the task.
//
// The payload is intended to store data that
// the worker needs to perform the work associated with the task.
// To
// return the payloads in the response, set
// response_view to
// FULL.
//
// A maximum of 10 qps of LeaseTasks
// requests are allowed per
// queue. RESOURCE_EXHAUSTED
// is returned when this limit is
// exceeded. RESOURCE_EXHAUSTED
// is also returned when
// max_tasks_dispatched_per_second
// is exceeded.
func (r *ProjectsLocationsQueuesTasksService) Lease(parent string, leasetasksrequest *LeaseTasksRequest) *ProjectsLocationsQueuesTasksLeaseCall {
	c := &ProjectsLocationsQueuesTasksLeaseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.leasetasksrequest = leasetasksrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksLeaseCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksLeaseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksLeaseCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksLeaseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksLeaseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksLeaseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.leasetasksrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+parent}/tasks:lease")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.lease" call.
// Exactly one of *LeaseTasksResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *LeaseTasksResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsQueuesTasksLeaseCall) Do(opts ...googleapi.CallOption) (*LeaseTasksResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &LeaseTasksResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Leases tasks from a pull queue for\nlease_duration.\n\nThis method is invoked by the worker to obtain a lease. The\nworker must acknowledge the task via\nAcknowledgeTask after they have\nperformed the work associated with the task.\n\nThe payload is intended to store data that\nthe worker needs to perform the work associated with the task. To\nreturn the payloads in the response, set\nresponse_view to\nFULL.\n\nA maximum of 10 qps of LeaseTasks\nrequests are allowed per\nqueue. RESOURCE_EXHAUSTED\nis returned when this limit is\nexceeded. RESOURCE_EXHAUSTED\nis also returned when\nmax_tasks_dispatched_per_second\nis exceeded.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks:lease",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.tasks.lease",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "Required.\n\nThe queue name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+parent}/tasks:lease",
	//   "request": {
	//     "$ref": "LeaseTasksRequest"
	//   },
	//   "response": {
	//     "$ref": "LeaseTasksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.tasks.list":

type ProjectsLocationsQueuesTasksListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the tasks in a queue.
//
// By default, only the BASIC view is retrieved
// due to performance considerations;
// response_view controls the
// subset of information which is returned.
func (r *ProjectsLocationsQueuesTasksService) List(parent string) *ProjectsLocationsQueuesTasksListCall {
	c := &ProjectsLocationsQueuesTasksListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// OrderBy sets the optional parameter "orderBy": Sort order used for
// the query. The only fields supported for sorting
// are `schedule_time` and `pull_message.tag`. All results will
// be
// returned in approximately ascending order. The default ordering is
// by
// `schedule_time`.
func (c *ProjectsLocationsQueuesTasksListCall) OrderBy(orderBy string) *ProjectsLocationsQueuesTasksListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// Fewer tasks than requested might be returned.
//
// The maximum page size is 1000. If unspecified, the page size will
// be the maximum. Fewer tasks than requested might be returned,
// even if more tasks exist; use
// next_page_token in the
// response to determine if more tasks exist.
func (c *ProjectsLocationsQueuesTasksListCall) PageSize(pageSize int64) *ProjectsLocationsQueuesTasksListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying the page of results to return.
//
// To request the first page results, page_token must be empty.
// To
// request the next page of results, page_token must be the value
// of
// next_page_token returned
// from the previous call to ListTasks
// method.
//
// The page token is valid for only 2 hours.
func (c *ProjectsLocationsQueuesTasksListCall) PageToken(pageToken string) *ProjectsLocationsQueuesTasksListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// ResponseView sets the optional parameter "responseView": The
// response_view specifies which subset of the Task will
// be
// returned.
//
// By default response_view is BASIC; not all
// information is retrieved by default because some data, such
// as
// payloads, might be desirable to return only when needed because
// of its large size or because of the sensitivity of data that
// it
// contains.
//
// Authorization for FULL requires
// `cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on
// the
// Task resource.
//
// Possible values:
//   "VIEW_UNSPECIFIED"
//   "BASIC"
//   "FULL"
func (c *ProjectsLocationsQueuesTasksListCall) ResponseView(responseView string) *ProjectsLocationsQueuesTasksListCall {
	c.urlParams_.Set("responseView", responseView)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksListCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsQueuesTasksListCall) IfNoneMatch(entityTag string) *ProjectsLocationsQueuesTasksListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksListCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+parent}/tasks")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.list" call.
// Exactly one of *ListTasksResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListTasksResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsQueuesTasksListCall) Do(opts ...googleapi.CallOption) (*ListTasksResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListTasksResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the tasks in a queue.\n\nBy default, only the BASIC view is retrieved\ndue to performance considerations;\nresponse_view controls the\nsubset of information which is returned.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks",
	//   "httpMethod": "GET",
	//   "id": "cloudtasks.projects.locations.queues.tasks.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "orderBy": {
	//       "description": "Sort order used for the query. The only fields supported for sorting\nare `schedule_time` and `pull_message.tag`. All results will be\nreturned in approximately ascending order. The default ordering is by\n`schedule_time`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. Fewer tasks than requested might be returned.\n\nThe maximum page size is 1000. If unspecified, the page size will\nbe the maximum. Fewer tasks than requested might be returned,\neven if more tasks exist; use\nnext_page_token in the\nresponse to determine if more tasks exist.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying the page of results to return.\n\nTo request the first page results, page_token must be empty. To\nrequest the next page of results, page_token must be the value of\nnext_page_token returned\nfrom the previous call to ListTasks\nmethod.\n\nThe page token is valid for only 2 hours.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "Required.\n\nThe queue name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "responseView": {
	//       "description": "The response_view specifies which subset of the Task will be\nreturned.\n\nBy default response_view is BASIC; not all\ninformation is retrieved by default because some data, such as\npayloads, might be desirable to return only when needed because\nof its large size or because of the sensitivity of data that it\ncontains.\n\nAuthorization for FULL requires\n`cloudtasks.tasks.fullView` [Google IAM](/iam/) permission on the\nTask resource.",
	//       "enum": [
	//         "VIEW_UNSPECIFIED",
	//         "BASIC",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+parent}/tasks",
	//   "response": {
	//     "$ref": "ListTasksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsQueuesTasksListCall) Pages(ctx context.Context, f func(*ListTasksResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "cloudtasks.projects.locations.queues.tasks.renewLease":

type ProjectsLocationsQueuesTasksRenewLeaseCall struct {
	s                 *Service
	name              string
	renewleaserequest *RenewLeaseRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// RenewLease: Renew the current lease of a pull task.
//
// The worker can use this method to extend the lease by a new
// duration, starting from now. The new task lease will be
// returned in the task's schedule_time.
func (r *ProjectsLocationsQueuesTasksService) RenewLease(name string, renewleaserequest *RenewLeaseRequest) *ProjectsLocationsQueuesTasksRenewLeaseCall {
	c := &ProjectsLocationsQueuesTasksRenewLeaseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.renewleaserequest = renewleaserequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksRenewLeaseCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksRenewLeaseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksRenewLeaseCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksRenewLeaseCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksRenewLeaseCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksRenewLeaseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.renewleaserequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}:renewLease")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.renewLease" call.
// Exactly one of *Task or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Task.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsQueuesTasksRenewLeaseCall) Do(opts ...googleapi.CallOption) (*Task, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Task{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Renew the current lease of a pull task.\n\nThe worker can use this method to extend the lease by a new\nduration, starting from now. The new task lease will be\nreturned in the task's schedule_time.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks/{tasksId}:renewLease",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.tasks.renewLease",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe task name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+/tasks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}:renewLease",
	//   "request": {
	//     "$ref": "RenewLeaseRequest"
	//   },
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudtasks.projects.locations.queues.tasks.run":

type ProjectsLocationsQueuesTasksRunCall struct {
	s              *Service
	name           string
	runtaskrequest *RunTaskRequest
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Run: Forces a task to run now.
//
// This command is meant to be used for manual debugging. For
// example, RunTask can be used to retry a failed
// task after a fix has been made or to manually force a task to
// be
// dispatched now.
//
// When this method is called, Cloud Tasks will dispatch the task to
// its
// target, even if the queue is PAUSED.
//
// The dispatched task is returned. That is, the task that is
// returned
// contains the status after the task is dispatched but
// before the task is received by its target.
//
// If Cloud Tasks receives a successful response from the
// task's
// handler, then the task will be deleted; otherwise the
// task's
// schedule_time will be reset to the time that
// RunTask was called plus the retry delay specified
// in the queue and task's RetryConfig.
//
// RunTask returns
// NOT_FOUND when it is called on a
// task that has already succeeded or permanently
// failed. FAILED_PRECONDITION
// is returned when RunTask is called on task
// that is dispatched or already running.
//
// RunTask cannot be called on
// pull tasks.
func (r *ProjectsLocationsQueuesTasksService) Run(name string, runtaskrequest *RunTaskRequest) *ProjectsLocationsQueuesTasksRunCall {
	c := &ProjectsLocationsQueuesTasksRunCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.runtaskrequest = runtaskrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsQueuesTasksRunCall) Fields(s ...googleapi.Field) *ProjectsLocationsQueuesTasksRunCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsQueuesTasksRunCall) Context(ctx context.Context) *ProjectsLocationsQueuesTasksRunCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsLocationsQueuesTasksRunCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsLocationsQueuesTasksRunCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runtaskrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta2/{+name}:run")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "cloudtasks.projects.locations.queues.tasks.run" call.
// Exactly one of *Task or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Task.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsLocationsQueuesTasksRunCall) Do(opts ...googleapi.CallOption) (*Task, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Task{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Forces a task to run now.\n\nThis command is meant to be used for manual debugging. For\nexample, RunTask can be used to retry a failed\ntask after a fix has been made or to manually force a task to be\ndispatched now.\n\nWhen this method is called, Cloud Tasks will dispatch the task to its\ntarget, even if the queue is PAUSED.\n\nThe dispatched task is returned. That is, the task that is returned\ncontains the status after the task is dispatched but\nbefore the task is received by its target.\n\nIf Cloud Tasks receives a successful response from the task's\nhandler, then the task will be deleted; otherwise the task's\nschedule_time will be reset to the time that\nRunTask was called plus the retry delay specified\nin the queue and task's RetryConfig.\n\nRunTask returns\nNOT_FOUND when it is called on a\ntask that has already succeeded or permanently\nfailed. FAILED_PRECONDITION\nis returned when RunTask is called on task\nthat is dispatched or already running.\n\nRunTask cannot be called on\npull tasks.",
	//   "flatPath": "v2beta2/projects/{projectsId}/locations/{locationsId}/queues/{queuesId}/tasks/{tasksId}:run",
	//   "httpMethod": "POST",
	//   "id": "cloudtasks.projects.locations.queues.tasks.run",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required.\n\nThe task name. For example:\n`projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/locations/[^/]+/queues/[^/]+/tasks/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta2/{+name}:run",
	//   "request": {
	//     "$ref": "RunTaskRequest"
	//   },
	//   "response": {
	//     "$ref": "Task"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
