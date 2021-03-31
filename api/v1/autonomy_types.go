/*
     Company: 达梦数据库有限公司
  Department: 达梦技术公司/产品研发中心
      Author: 陈磊
      E-mail: chenlei@dameng.com
      Create: 2021/2/5 09:21
     Project: compass
     Package: v1
    Describe: Todo
*/

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Label struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Image struct {
	Repository      string            `json:"repository,omitempty"`
	Tag             string            `json:"tag,omitempty"`
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
}

type Environment struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type VolumeType string

const (
	Mounted               VolumeType = "Mounted"
	PersistentVolumeClaim VolumeType = "PersistentVolumeClaim"
	HostPath              VolumeType = "HostPath"
	EmptyDir              VolumeType = "EmptyDir"
)

type Volume struct {
	ID               string                            `json:"id,omitempty"`
	Type             VolumeType                        `json:"type,omitempty"`
	Mount            string                            `json:"mount,omitempty"`
	AccessMode       corev1.PersistentVolumeAccessMode `json:"accessMode,omitempty"`
	Capacity         string                            `json:"capacity,omitempty"`
	StorageClassName string                            `json:"storageClassName,omitempty"`
	Location         string                            `json:"location,omitempty"`
	LocationType     corev1.HostPathType               `json:"locationType,omitempty"`
}

type Config struct {
	ID      string `json:"id,omitempty"`
	Mount   string `json:"mount,omitempty"`
	Content string `json:"content,omitempty"`
}

type Port struct {
	ID            string          `json:"id,omitempty"`
	Protocol      corev1.Protocol `json:"protocol,omitempty"`
	ContainerPort int32           `json:"containerPort,omitempty"`
	ServerPort    int32           `json:"serverPort,omitempty"`
	NodePort      int32           `json:"nodePort,omitempty"`
	Ingress       bool            `json:"ingress,omitempty"`
	Host          string          `json:"host,omitempty"`
	Path          string          `json:"path,omitempty"`
	PathType      netv1.PathType  `json:"pathType,omitempty"`
}

type Action string

const (
	HttpGet   Action = "HTTPGet"
	Exec      Action = "Exec"
	TCPSocket Action = "TCPSocket"
)

type Handler struct {
	Action  Action               `json:"action,omitempty"`
	Scheme  corev1.URIScheme     `json:"scheme,omitempty"`
	Host    string               `json:"host,omitempty"`
	Port    int32                `json:"port,omitempty"`
	Path    string               `json:"path,omitempty"`
	Headers []*corev1.HTTPHeader `json:"headers,omitempty"`
	Command string               `json:"command,omitempty"`
}

type ProbeType string

type Probe struct {
	Handler             *Handler `json:"handler,omitempty"`
	InitialDelaySeconds int32    `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      int32    `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       int32    `json:"periodSeconds,omitempty"`
	SuccessThreshold    int32    `json:"successThreshold,omitempty"`
	FailureThreshold    int32    `json:"failureThreshold,omitempty"`
}

type Resource struct {
	Share   bool  `json:"share,omitempty"`
	Request int64 `json:"request,omitempty"`
	Limit   int64 `json:"limit,omitempty"`
}

type Parameter struct {
	ID    string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

type Terminator struct {
	Handler Handler `json:"handler,omitempty"`
	Grace   int32   `json:"grace,omitempty"`
}

type UpgradeStrategy struct {
	Strategy                appsv1.DeploymentStrategyType `json:"strategy,omitempty"`
	MaxUnavailable          string                        `json:"maxUnavailable,omitempty"`
	MaxUnavailableUnit      string                        `json:"maxUnavailableUnit,omitempty"`
	MaxSurge                string                        `json:"maxSurge,omitempty"`
	MaxSurgeUnit            string                        `json:"maxSurgeUnit,omitempty"`
	RevisionHistoryLimit    *int32                        `json:"revisionHistoryLimit,omitempty"`
	MinReadySeconds         *int32                        `json:"minReadySeconds,omitempty"`
	ProgressDeadlineSeconds *int32                        `json:"progressDeadlineSeconds,omitempty"`
}

type Log struct {
	ID        string `json:"id,omitempty"`
	Directory string `json:"directory,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
}

type Metric struct {
	MetricPath string `json:"metricPath,omitempty"`
	MetricPort int32  `json:"metricPort,omitempty"`
}

type Autoscaler struct {
	Min    int32 `json:"min,omitempty"`
	Max    int32 `json:"max,omitempty"`
	CPU    int32 `json:"cpu,omitempty"`
	Memory int32 `json:"memory,omitempty"`
}

// AutonomySpec defines the desired state of Autonomy
type AutonomySpec struct {
	ID               string                        `json:"id"`
	Title            string                        `json:"title,omitempty"`
	Describe         string                        `json:"describe,omitempty"`
	Labels           []*Label                      `json:"labels,omitempty"`
	Image            Image                         `json:"image,omitempty"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	Replicas         *int32                        `json:"replicas,omitempty"`
	ServiceName      string                        `json:"serviceName,omitempty"`
	Ports            []*Port                       `json:"ports,omitempty"`
	Memory           *Resource                     `json:"memory,omitempty"`
	CPU              *Resource                     `json:"cpu,omitempty"`
	Initial          string                        `json:"initial,omitempty"`
	Command          string                        `json:"command,omitempty"`
	Parameters       []*Parameter                  `json:"parameters,omitempty"`
	Environments     []*Environment                `json:"environments,omitempty"`
	Volumes          []*Volume                     `json:"volumes,omitempty"`
	Configs          []*Config                     `json:"configs,omitempty"`
	RestartPolicy    corev1.RestartPolicy          `json:"restartPolicy,omitempty"`
	Terminator       *Terminator                   `json:"terminator,omitempty"`
	Readiness        *Probe                        `json:"readiness,omitempty"`
	Liveness         *Probe                        `json:"liveness,omitempty"`
	Startup          *Probe                        `json:"startup,omitempty"`
	HostPID          bool                          `json:"hostPID,omitempty"`
	HostNetwork      bool                          `json:"hostNetwork,omitempty"`
	HostIPC          bool                          `json:"hostIPC,omitempty"`
	HostAliases      []*corev1.HostAlias           `json:"hostAliases,omitempty"`
	SecurityContext  *corev1.PodSecurityContext    `json:"securityContext,omitempty"`
	UpgradeStrategy  *UpgradeStrategy              `json:"upgradeStrategy,omitempty"`
	Logs             []*Log                        `json:"logs,omitempty"`
	Metric           *Metric                       `json:"metric,omitempty"`
	Autoscaler       *Autoscaler                   `json:"autoscaler,omitempty"`
}

type Phase string

const (
	Reconciling Phase = "Reconciling"
	Running     Phase = "Running"
	Deleting    Phase = "Deleting"
)

// AutonomyStatus defines the observed state of Autonomy
type AutonomyStatus struct {
	Phase Phase `json:"phase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:printcolumn:name="Creator",type=string,JSONPath=`.metadata.annotations['operator\.dameng\.com/creator']`
// +kubebuilder:printcolumn:name="Department",type=string,JSONPath=`.metadata.annotations['operator\.dameng\.com/department']`

// Autonomy is the Schema for the autonomies API
type Autonomy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AutonomySpec   `json:"spec,omitempty"`
	Status AutonomyStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AutonomyList contains a list of Autonomy
type AutonomyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Autonomy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Autonomy{}, &AutonomyList{})
}

type HistoryType string

const (
	Create   HistoryType = "Create"
	Rollback HistoryType = "Rollback"
	Patch    HistoryType = "Patch"
	Update   HistoryType = "Update"
)

type History struct {
	Time       metav1.Time  `json:"time"`
	Type       HistoryType  `json:"type"`
	Spec       AutonomySpec `json:"spec"`
	Describe   string       `json:"describe"`
	Principles string       `json:"principles"`
}
