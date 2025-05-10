package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
)

const (
	// AdmissionWebhook
	AdmissionWebhookAnnotationMutateKey = "lxcfs-admission-webhook.raids-lab.github.io/mutate"
	AdmissionWebhookAnnotationStatusKey = "lxcfs-admission-webhook.raids-lab.github.io/status"

	// AdmissionWebhook Values
	MutateValueTrue    = "true"
	MutateValueFalse   = "false"
	StatusValueMutated = "mutated"
)

var (
	// LXCFS VolumeMounts Template
	// See https://github.com/lxc/lxcfs for more details.
	VolumeMountsTemplate = []corev1.VolumeMount{
		{
			Name:      "lxcfs-proc-cpuinfo",
			MountPath: "/proc/cpuinfo",
			ReadOnly:  true,
		},
		{
			Name:      "lxcfs-proc-diskstats",
			MountPath: "/proc/diskstats",
			ReadOnly:  true,
		},
		{
			Name:      "lxcfs-proc-meminfo",
			MountPath: "/proc/meminfo",
			ReadOnly:  true,
		},
		{
			Name:      "lxcfs-proc-stat",
			MountPath: "/proc/stat",
			ReadOnly:  true,
		},
		{
			Name:      "lxcfs-proc-swaps",
			MountPath: "/proc/swaps",
			ReadOnly:  true,
		},
		{
			Name:      "lxcfs-proc-uptime",
			MountPath: "/proc/uptime",
			ReadOnly:  true,
		},
		{
			Name:      "lxcfs-proc-slabinfo",
			MountPath: "/proc/slabinfo",
			ReadOnly:  true,
		},
		// {
		// 	Name:      "lxcfs-proc-pressure-io",
		// 	MountPath: "/proc/pressure/io",
		// 	ReadOnly:  true,
		// },
		// {
		// 	Name:      "lxcfs-proc-pressure-cpu",
		// 	MountPath: "/proc/pressure/cpu",
		// 	ReadOnly:  true,
		// },
		// {
		// 	Name:      "lxcfs-proc-pressure-memory",
		// 	MountPath: "/proc/pressure/memory",
		// 	ReadOnly:  true,
		// },
		{
			Name:      "lxcfs-sys-devices-system-cpu-online",
			MountPath: "/sys/devices/system/cpu/online",
			ReadOnly:  true,
		},
		{
			Name:             "var-lib-lxc",
			MountPath:        "/var/lib/lxc",
			ReadOnly:         true,
			MountPropagation: ptr.To(corev1.MountPropagationHostToContainer),
		},
	}

	// LXCFS Volumes Template
	VolumesTemplate = []corev1.Volume{
		{
			Name: "lxcfs-proc-cpuinfo",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc/lxcfs/proc/cpuinfo",
					Type: ptr.To(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-diskstats",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc/lxcfs/proc/diskstats",
					Type: ptr.To(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-meminfo",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc/lxcfs/proc/meminfo",
					Type: ptr.To(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-stat",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc/lxcfs/proc/stat",
					Type: ptr.To(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-swaps",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc/lxcfs/proc/swaps",
					Type: ptr.To(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-uptime",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc/lxcfs/proc/uptime",
					Type: ptr.To(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "lxcfs-proc-slabinfo",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc/lxcfs/proc/slabinfo",
					Type: ptr.To(corev1.HostPathFile),
				},
			},
		},
		// {
		// 	Name: "lxcfs-proc-pressure-io",
		// 	VolumeSource: corev1.VolumeSource{
		// 		HostPath: &corev1.HostPathVolumeSource{
		// 			Path: "/var/lib/lxc/lxcfs/proc/pressure/io",
		// 			Type: ptr.To(corev1.HostPathFile),
		// 		},
		// 	},
		// },
		// {
		// 	Name: "lxcfs-proc-pressure-cpu",
		// 	VolumeSource: corev1.VolumeSource{
		// 		HostPath: &corev1.HostPathVolumeSource{
		// 			Path: "/var/lib/lxc/lxcfs/proc/pressure/cpu",
		// 			Type: ptr.To(corev1.HostPathFile),
		// 		},
		// 	},
		// },
		// {
		// 	Name: "lxcfs-proc-pressure-memory",
		// 	VolumeSource: corev1.VolumeSource{
		// 		HostPath: &corev1.HostPathVolumeSource{
		// 			Path: "/var/lib/lxc/lxcfs/proc/pressure/memory",
		// 			Type: ptr.To(corev1.HostPathFile),
		// 		},
		// 	},
		// },
		{
			Name: "lxcfs-sys-devices-system-cpu-online",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc/lxcfs/sys/devices/system/cpu/online",
					Type: ptr.To(corev1.HostPathFile),
				},
			},
		},
		{
			Name: "var-lib-lxc",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/lxc",
					Type: ptr.To(corev1.HostPathDirectory),
				},
			},
		},
	}

	// Ignored Namespaces
	IgnoredNamespaces = []string{
		metav1.NamespaceSystem,
		metav1.NamespacePublic,
		"kube-node-lease",
	}
)
