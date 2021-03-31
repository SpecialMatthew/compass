/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"git.ypt.dameng.com/dmcca/compass/tools"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cpv1 "git.ypt.dameng.com/dmcca/compass/api/v1"
)

const FieldManager string = "autonomy-apps-operator"
const Finalizer string = "finalizer.autonomy.operator.dameng.com"

// AutonomyReconciler reconciles a Autonomy object
type AutonomyReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.dameng.com,resources=autonomies,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.dameng.com,resources=autonomies/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.dameng.com,resources=autonomies/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments;statefulsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services;persistentvolumeclaims;configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=get;list;watch;create;update;patch;delete

func (r *AutonomyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("autonomy", req.NamespacedName)

	var origin cpv1.Autonomy
	if err := r.Get(ctx, req.NamespacedName, &origin); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	klog.V(8).Infof("origin: %v", origin)

	workbench := origin.DeepCopy()
	workbench.Status.Phase = cpv1.Reconciling

	if origin.ObjectMeta.DeletionTimestamp.IsZero() {
		appendFinalizers(workbench, []string{Finalizer})
		runtimeObjects, err := r.generateRuntimeObjects(&origin)
		if err != nil {
			logger.Error(err, "decode error")
			return ctrl.Result{}, err
		}

		for _, object := range runtimeObjects {
			(*object).(*unstructured.Unstructured).SetNamespace(origin.Namespace)
			if err := controllerutil.SetControllerReference(&origin, (*object).(*unstructured.Unstructured), r.Scheme); err != nil {
				logger.Error(err, "maintain controller reference error")
				return ctrl.Result{}, err
			}
			if err := r.Patch(ctx, (*object).(*unstructured.Unstructured), client.Merge, &client.PatchOptions{FieldManager: FieldManager}); err != nil {
				if errors.IsNotFound(err) {
					if err := r.Create(ctx, (*object).(*unstructured.Unstructured), &client.CreateOptions{FieldManager: FieldManager}); err != nil {
						logger.Error(err, "object create error", "object", object)
						return ctrl.Result{}, err
					}
				} else {
					logger.Error(err, "object patch error", "object", object)
					return ctrl.Result{}, err
				}
			}
		}
	} else {
		removeFinalizers(workbench, []string{Finalizer})
	}

	if !reflect.DeepEqual(workbench.TypeMeta, origin.TypeMeta) || !reflect.DeepEqual(workbench.ObjectMeta, origin.ObjectMeta) || !reflect.DeepEqual(workbench.Spec, origin.Spec) {
		if err := r.Update(ctx, workbench); err != nil {
			logger.Error(err, "update without status error")
			return ctrl.Result{}, err
		}
	}

	if !reflect.DeepEqual(workbench.Status, origin.Status) {
		if err := r.Status().Update(ctx, workbench); err != nil {
			logger.Error(err, "update status error")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AutonomyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &appsv1.Deployment{}, ".metadata.controller", func(object client.Object) []string {
		owner := metav1.GetControllerOf(object.(*appsv1.Deployment))
		if owner == nil || owner.APIVersion != cpv1.GroupVersion.String() || owner.Kind != "Autonomy" {
			return nil
		}
		return []string{owner.Name}
	}); err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).For(&cpv1.Autonomy{}).Owns(&cpv1.Autonomy{}).Complete(r)
}

func (r *AutonomyReconciler) generateRuntimeObjects(autonomy *cpv1.Autonomy) (runtimeObjects []*runtime.Object, err error) {
	template, err := tools.ParseTemplate("apps.dameng.com_autonomies.gotmpl", autonomy)
	if err != nil {
		klog.Errorf("generate runtime error: %v", err)
		return runtimeObjects, err
	}
	klog.V(8).Infof("template: %s", template)
	return tools.Decode(template)
}

func appendFinalizers(autonomy *cpv1.Autonomy, finalizers []string) runtime.Object {
	for _, finalizer := range finalizers {
		if !tools.ArrayContains(autonomy.GetFinalizers(), finalizer) {
			autonomy.SetFinalizers(append(autonomy.GetFinalizers(), finalizer))
		}
	}
	return autonomy
}

func removeFinalizers(autonomy *cpv1.Autonomy, finalizers []string) runtime.Object {
	for _, finalizer := range finalizers {
		if tools.ArrayContains(autonomy.GetFinalizers(), finalizer) {
			autonomy.SetFinalizers(tools.ArrayRemove(autonomy.GetFinalizers(), finalizer))
		}
	}
	return autonomy
}
