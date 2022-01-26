package v1alpha1

import (
	v1beta1 "github.com/mercari/spanner-autoscaler/api/v1beta1"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"

	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var log = ctrllog.Log.WithName("spannerautoscaler-v1alpha1.converter")

func (src *SpannerAutoscaler) ConvertTo(dstRaw conversion.Hub) error {
	log.V(2).Info("begin conversion from v1alpha1 to v1beta1", "src", src)

	dst := dstRaw.(*v1beta1.SpannerAutoscaler)
	dst.Spec.TargetInstance = v1beta1.TargetInstance{
		ProjectID:  *src.Spec.ScaleTargetRef.ProjectID,
		InstanceID: *src.Spec.ScaleTargetRef.InstanceID,
	}

	auth := v1beta1.Authentication{}

	if src.Spec.ImpersonateConfig != nil {
		auth.Type = v1beta1.AuthenticationTypeImpersonation
		auth.ImpersonateConfig = v1beta1.ImpersonateConfig{
			TargetServiceAccount: src.Spec.ImpersonateConfig.TargetServiceAccount,
			Delegates:            src.Spec.ImpersonateConfig.Delegates,
		}
	}

	if src.Spec.ServiceAccountSecretRef != nil {
		auth.Type = v1beta1.AuthenticationTypeSA
		auth.IAMKeySecret = v1beta1.IAMKeySecret{
			Name: *src.Spec.ServiceAccountSecretRef.Name,
			Key:  *src.Spec.ServiceAccountSecretRef.Key,
		}

		if src.Spec.ServiceAccountSecretRef.Namespace != nil && *src.Spec.ServiceAccountSecretRef.Namespace != "" {
			auth.IAMKeySecret.Namespace = *src.Spec.ServiceAccountSecretRef.Namespace
		}
	}

	dst.Spec.Authentication = auth

	scaleConfig := v1beta1.ScaleConfig{}
	if src.Spec.MinNodes != nil && *src.Spec.MinNodes >= 1 && src.Spec.MaxNodes != nil && *src.Spec.MaxNodes >= 1 {
		scaleConfig.ComputeType = v1beta1.ComputeTypeNode
		scaleConfig.Nodes = v1beta1.ScaleConfigNodes{
			Min: int(*src.Spec.MinNodes),
			Max: int(*src.Spec.MaxNodes),
		}
	}
	if src.Spec.MinProcessingUnits != nil && *src.Spec.MinProcessingUnits >= 100 && src.Spec.MaxProcessingUnits != nil && *src.Spec.MaxProcessingUnits >= 100 {
		scaleConfig.ComputeType = v1beta1.ComputeTypePU
		scaleConfig.ProcessingUnits = v1beta1.ScaleConfigPUs{
			Min: int(*src.Spec.MinProcessingUnits),
			Max: int(*src.Spec.MaxProcessingUnits),
		}
	}
	scaleConfig.ScaledownStepSize = int(*src.Spec.MaxScaleDownNodes)
	scaleConfig.TargetCPUUtilization = v1beta1.TargetCPUUtilization{
		HighPriority: int(*src.Spec.TargetCPUUtilization.HighPriority),
	}

	dst.Spec.ScaleConfig = scaleConfig

	dst.ObjectMeta = src.ObjectMeta

	log.V(2).Info("finished conversion from v1alpha1 to v1beta1", "src", src, "dst", dst)

	return nil
}

//nolint:stylecheck
func (dst *SpannerAutoscaler) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.SpannerAutoscaler)
	log.V(2).Info("begin conversion from v1beta1 to v1alpha1", "src", src)

	dst.Spec.ScaleTargetRef = ScaleTargetRef{
		ProjectID:  pointer.String(src.Spec.TargetInstance.ProjectID),
		InstanceID: pointer.String(src.Spec.TargetInstance.InstanceID),
	}

	switch src.Spec.Authentication.Type {
	case v1beta1.AuthenticationTypeSA:
		dst.Spec.ServiceAccountSecretRef = &ServiceAccountSecretRef{
			Name:      pointer.String(src.Spec.Authentication.IAMKeySecret.Name),
			Namespace: pointer.String(src.Spec.Authentication.IAMKeySecret.Namespace),
			Key:       pointer.String(src.Spec.Authentication.IAMKeySecret.Key),
		}
	case v1beta1.AuthenticationTypeImpersonation:
		dst.Spec.ImpersonateConfig = &ImpersonateConfig{
			TargetServiceAccount: src.Spec.Authentication.ImpersonateConfig.TargetServiceAccount,
			Delegates:            src.Spec.Authentication.ImpersonateConfig.Delegates,
		}
	}

	switch src.Spec.ScaleConfig.ComputeType {
	case v1beta1.ComputeTypeNode:
		dst.Spec.MinNodes = pointer.Int32(int32(src.Spec.ScaleConfig.Nodes.Min))
		dst.Spec.MaxNodes = pointer.Int32(int32(src.Spec.ScaleConfig.Nodes.Max))

	case v1beta1.ComputeTypePU:
		dst.Spec.MinProcessingUnits = pointer.Int32(int32(src.Spec.ScaleConfig.ProcessingUnits.Min))
		dst.Spec.MaxProcessingUnits = pointer.Int32(int32(src.Spec.ScaleConfig.ProcessingUnits.Max))
	}

	dst.Spec.MaxScaleDownNodes = pointer.Int32(int32(src.Spec.ScaleConfig.ScaledownStepSize))
	dst.Spec.TargetCPUUtilization = TargetCPUUtilization{
		HighPriority: pointer.Int32(int32(src.Spec.ScaleConfig.TargetCPUUtilization.HighPriority)),
	}

	dst.ObjectMeta = src.ObjectMeta

	log.V(2).Info("finished conversion from v1beta1 to v1alpha1", "src", src, "dst", dst)

	return nil
}