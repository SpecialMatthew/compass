/*
Copyright 2021.
*/

package v1

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var logger = logf.Log.WithName("autonomy-resource")
var kubernetes client.Client

func (r *Autonomy) SetupWebhookWithManager(mgr ctrl.Manager) error {
	kubernetes = mgr.GetClient()
	return ctrl.NewWebhookManagedBy(mgr).For(r).Complete()
}

//+kubebuilder:webhook:path=/mutate-apps-dameng-com-v1-autonomy,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.dameng.com,resources=autonomies,verbs=create;update,versions=v1,name=mautonomy.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &Autonomy{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Autonomy) Default() {
	logger.Info("default", "name", r.Name, "namespace", r.Namespace)
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-apps-dameng-com-v1-autonomy,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.dameng.com,resources=autonomies,verbs=create;update,versions=v1,name=vautonomy.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &Autonomy{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Autonomy) ValidateCreate() error {
	logger.Info("validate create", "name", r.Name, "namespace", r.Namespace)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Autonomy) ValidateUpdate(old runtime.Object) error {
	logger.Info("validate update", "name", r.Name, "namespace", r.Namespace)

	pvcs := getPersistentVolumeClaims(r)
	oldPvcs := getPersistentVolumeClaims(old.(*Autonomy))

	if len(oldPvcs) != 0 {
		var shouldDelete = false
		if len(pvcs) == len(oldPvcs) {
			for _, newer := range pvcs {
				for _, older := range oldPvcs {
					if newer.ID != older.ID {
						shouldDelete = true
					}
				}
			}
		} else {
			shouldDelete = true
		}
		logger.Info("should delete", "name", r.Name, "namespace", r.Namespace, "should", shouldDelete)
		if shouldDelete {
			logger.Info("delete statefulSet and create deployment", "name", r.Name, "namespace", r.Namespace)
			var statefulSet appsv1.StatefulSet
			if err := kubernetes.Get(context.Background(), client.ObjectKey{Namespace: r.Namespace, Name: r.Name}, &statefulSet); err != nil && !errors.IsNotFound(err) {
				logger.Error(err, "get statefulSet error", "name", r.Name, "namespace", r.Namespace)
				return err
			}
			if &statefulSet != nil && statefulSet.DeletionTimestamp.IsZero() {
				if err := kubernetes.Delete(context.Background(), &statefulSet); err != nil {
					logger.Error(err, "delete statefulSet error", "name", r.Name, "namespace", r.Namespace)
					return err
				}
				logger.Info("delete statefulSet...", "name", r.Name, "namespace", r.Namespace)
			}
			return nil
		}
	}

	if len(oldPvcs) == 0 && len(pvcs) != 0 {
		logger.Info("delete deployment and create statefulSet", "name", r.Name, "namespace", r.Namespace)
		var deployment appsv1.Deployment
		if err := kubernetes.Get(context.Background(), client.ObjectKey{Namespace: r.Namespace, Name: r.Name}, &deployment); err != nil && !errors.IsNotFound(err) {
			logger.Error(err, "get deployment error", "name", r.Name, "namespace", r.Namespace)
			return err
		}
		if &deployment != nil && deployment.DeletionTimestamp.IsZero() {
			if err := kubernetes.Delete(context.Background(), &deployment); err != nil {
				logger.Error(err, "delete deployment error", "name", r.Name, "namespace", r.Namespace)
				return err
			}
			logger.Info("delete deployment...", "name", r.Name, "namespace", r.Namespace)
		}
		return nil
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Autonomy) ValidateDelete() error {
	logger.Info("validate delete", "name", r.Name, "namespace", r.Namespace)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func getPersistentVolumeClaims(autonomy *Autonomy) (volumes []*Volume) {
	for _, volume := range autonomy.Spec.Volumes {
		if volume.Type == PersistentVolumeClaim || volume.Type == Mounted {
			volumes = append(volumes, volume)
		}
	}
	return volumes
}
