//
// Code generated by grafana-app-sdk. DO NOT EDIT.
//

package v1alpha1

import (
	"fmt"
	"github.com/grafana/grafana-app-sdk/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

// +k8s:openapi-gen=true
type Dashboard struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata" yaml:"metadata"`

	// Spec is the spec of the Dashboard
	Spec DashboardSpec `json:"spec" yaml:"spec"`

	Status DashboardStatus `json:"status" yaml:"status"`
}

func (o *Dashboard) GetSpec() any {
	return o.Spec
}

func (o *Dashboard) SetSpec(spec any) error {
	cast, ok := spec.(DashboardSpec)
	if !ok {
		return fmt.Errorf("cannot set spec type %#v, not of type Spec", spec)
	}
	o.Spec = cast
	return nil
}

func (o *Dashboard) GetSubresources() map[string]any {
	return map[string]any{
		"status": o.Status,
	}
}

func (o *Dashboard) GetSubresource(name string) (any, bool) {
	switch name {
	case "status":
		return o.Status, true
	default:
		return nil, false
	}
}

func (o *Dashboard) SetSubresource(name string, value any) error {
	switch name {
	case "status":
		cast, ok := value.(DashboardStatus)
		if !ok {
			return fmt.Errorf("cannot set status type %#v, not of type DashboardStatus", value)
		}
		o.Status = cast
		return nil
	default:
		return fmt.Errorf("subresource '%s' does not exist", name)
	}
}

func (o *Dashboard) GetStaticMetadata() resource.StaticMetadata {
	gvk := o.GroupVersionKind()
	return resource.StaticMetadata{
		Name:      o.ObjectMeta.Name,
		Namespace: o.ObjectMeta.Namespace,
		Group:     gvk.Group,
		Version:   gvk.Version,
		Kind:      gvk.Kind,
	}
}

func (o *Dashboard) SetStaticMetadata(metadata resource.StaticMetadata) {
	o.Name = metadata.Name
	o.Namespace = metadata.Namespace
	o.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   metadata.Group,
		Version: metadata.Version,
		Kind:    metadata.Kind,
	})
}

func (o *Dashboard) GetCommonMetadata() resource.CommonMetadata {
	dt := o.DeletionTimestamp
	var deletionTimestamp *time.Time
	if dt != nil {
		deletionTimestamp = &dt.Time
	}
	// Legacy ExtraFields support
	extraFields := make(map[string]any)
	if o.Annotations != nil {
		extraFields["annotations"] = o.Annotations
	}
	if o.ManagedFields != nil {
		extraFields["managedFields"] = o.ManagedFields
	}
	if o.OwnerReferences != nil {
		extraFields["ownerReferences"] = o.OwnerReferences
	}
	return resource.CommonMetadata{
		UID:               string(o.UID),
		ResourceVersion:   o.ResourceVersion,
		Generation:        o.Generation,
		Labels:            o.Labels,
		CreationTimestamp: o.CreationTimestamp.Time,
		DeletionTimestamp: deletionTimestamp,
		Finalizers:        o.Finalizers,
		UpdateTimestamp:   o.GetUpdateTimestamp(),
		CreatedBy:         o.GetCreatedBy(),
		UpdatedBy:         o.GetUpdatedBy(),
		ExtraFields:       extraFields,
	}
}

func (o *Dashboard) SetCommonMetadata(metadata resource.CommonMetadata) {
	o.UID = types.UID(metadata.UID)
	o.ResourceVersion = metadata.ResourceVersion
	o.Generation = metadata.Generation
	o.Labels = metadata.Labels
	o.CreationTimestamp = metav1.NewTime(metadata.CreationTimestamp)
	if metadata.DeletionTimestamp != nil {
		dt := metav1.NewTime(*metadata.DeletionTimestamp)
		o.DeletionTimestamp = &dt
	} else {
		o.DeletionTimestamp = nil
	}
	o.Finalizers = metadata.Finalizers
	if o.Annotations == nil {
		o.Annotations = make(map[string]string)
	}
	if !metadata.UpdateTimestamp.IsZero() {
		o.SetUpdateTimestamp(metadata.UpdateTimestamp)
	}
	if metadata.CreatedBy != "" {
		o.SetCreatedBy(metadata.CreatedBy)
	}
	if metadata.UpdatedBy != "" {
		o.SetUpdatedBy(metadata.UpdatedBy)
	}
	// Legacy support for setting Annotations, ManagedFields, and OwnerReferences via ExtraFields
	if metadata.ExtraFields != nil {
		if annotations, ok := metadata.ExtraFields["annotations"]; ok {
			if cast, ok := annotations.(map[string]string); ok {
				o.Annotations = cast
			}
		}
		if managedFields, ok := metadata.ExtraFields["managedFields"]; ok {
			if cast, ok := managedFields.([]metav1.ManagedFieldsEntry); ok {
				o.ManagedFields = cast
			}
		}
		if ownerReferences, ok := metadata.ExtraFields["ownerReferences"]; ok {
			if cast, ok := ownerReferences.([]metav1.OwnerReference); ok {
				o.OwnerReferences = cast
			}
		}
	}
}

func (o *Dashboard) GetCreatedBy() string {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	return o.ObjectMeta.Annotations["grafana.com/createdBy"]
}

func (o *Dashboard) SetCreatedBy(createdBy string) {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	o.ObjectMeta.Annotations["grafana.com/createdBy"] = createdBy
}

func (o *Dashboard) GetUpdateTimestamp() time.Time {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	parsed, _ := time.Parse(time.RFC3339, o.ObjectMeta.Annotations["grafana.com/updateTimestamp"])
	return parsed
}

func (o *Dashboard) SetUpdateTimestamp(updateTimestamp time.Time) {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	o.ObjectMeta.Annotations["grafana.com/updateTimestamp"] = updateTimestamp.Format(time.RFC3339)
}

func (o *Dashboard) GetUpdatedBy() string {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	return o.ObjectMeta.Annotations["grafana.com/updatedBy"]
}

func (o *Dashboard) SetUpdatedBy(updatedBy string) {
	if o.ObjectMeta.Annotations == nil {
		o.ObjectMeta.Annotations = make(map[string]string)
	}

	o.ObjectMeta.Annotations["grafana.com/updatedBy"] = updatedBy
}

func (o *Dashboard) Copy() resource.Object {
	return resource.CopyObject(o)
}

func (o *Dashboard) DeepCopyObject() runtime.Object {
	return o.Copy()
}

func (o *Dashboard) DeepCopy() *Dashboard {
	cpy := &Dashboard{}
	o.DeepCopyInto(cpy)
	return cpy
}

func (o *Dashboard) DeepCopyInto(dst *Dashboard) {
	dst.TypeMeta.APIVersion = o.TypeMeta.APIVersion
	dst.TypeMeta.Kind = o.TypeMeta.Kind
	o.ObjectMeta.DeepCopyInto(&dst.ObjectMeta)
	o.Spec.DeepCopyInto(&dst.Spec)
	o.Status.DeepCopyInto(&dst.Status)
}

// Interface compliance compile-time check
var _ resource.Object = &Dashboard{}

// +k8s:openapi-gen=true
type DashboardList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []Dashboard `json:"items" yaml:"items"`
}

func (o *DashboardList) DeepCopyObject() runtime.Object {
	return o.Copy()
}

func (o *DashboardList) Copy() resource.ListObject {
	cpy := &DashboardList{
		TypeMeta: o.TypeMeta,
		Items:    make([]Dashboard, len(o.Items)),
	}
	o.ListMeta.DeepCopyInto(&cpy.ListMeta)
	for i := 0; i < len(o.Items); i++ {
		if item, ok := o.Items[i].Copy().(*Dashboard); ok {
			cpy.Items[i] = *item
		}
	}
	return cpy
}

func (o *DashboardList) GetItems() []resource.Object {
	items := make([]resource.Object, len(o.Items))
	for i := 0; i < len(o.Items); i++ {
		items[i] = &o.Items[i]
	}
	return items
}

func (o *DashboardList) SetItems(items []resource.Object) {
	o.Items = make([]Dashboard, len(items))
	for i := 0; i < len(items); i++ {
		o.Items[i] = *items[i].(*Dashboard)
	}
}

func (o *DashboardList) DeepCopy() *DashboardList {
	cpy := &DashboardList{}
	o.DeepCopyInto(cpy)
	return cpy
}

func (o *DashboardList) DeepCopyInto(dst *DashboardList) {
	resource.CopyObjectInto(dst, o)
}

// Interface compliance compile-time check
var _ resource.ListObject = &DashboardList{}

// Copy methods for all subresource types

// DeepCopy creates a full deep copy of DashboardStatus
func (s *DashboardStatus) DeepCopy() *DashboardStatus {
	cpy := &DashboardStatus{}
	s.DeepCopyInto(cpy)
	return cpy
}

// DeepCopyInto deep copies DashboardStatus into another DashboardStatus object
func (s *DashboardStatus) DeepCopyInto(dst *DashboardStatus) {
	resource.CopyObjectInto(dst, s)
}
