// Copyright (c) 2016 Intel Corporation
// Copyright (c) 2019 Huawei Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package persistapi

import (
	"github.com/opencontainers/runc/libcontainer/configs"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// HypervisorConfig saves configurations of sandbox hypervisor
type HypervisorConfig struct {
	// NumVCPUs specifies default number of vCPUs for the VM.
	NumVCPUs uint32

	//DefaultMaxVCPUs specifies the maximum number of vCPUs for the VM.
	DefaultMaxVCPUs uint32

	// DefaultMem specifies default memory size in MiB for the VM.
	MemorySize uint32

	// DefaultBridges specifies default number of bridges for the VM.
	// Bridges can be used to hot plug devices
	DefaultBridges uint32

	// Msize9p is used as the msize for 9p shares
	Msize9p uint32

	// MemSlots specifies default memory slots the VM.
	MemSlots uint32

	// MemOffset specifies memory space for nvdimm device
	MemOffset uint32

	// VirtioFSCacheSize is the DAX cache size in MiB
	VirtioFSCacheSize uint32

	// KernelPath is the guest kernel host path.
	KernelPath string

	// ImagePath is the guest image host path.
	ImagePath string

	// InitrdPath is the guest initrd image host path.
	// ImagePath and InitrdPath cannot be set at the same time.
	InitrdPath string

	// FirmwarePath is the bios host path
	FirmwarePath string

	// MachineAccelerators are machine specific accelerators
	MachineAccelerators string

	// HypervisorPath is the hypervisor executable host path.
	HypervisorPath string

	// HypervisorCtlPath is the hypervisor ctl executable host path.
	HypervisorCtlPath string

	// JailerPath is the jailer executable host path.
	JailerPath string

	// BlockDeviceDriver specifies the driver to be used for block device
	// either VirtioSCSI or VirtioBlock with the default driver being defaultBlockDriver
	BlockDeviceDriver string

	// HypervisorMachineType specifies the type of machine being
	// emulated.
	HypervisorMachineType string

	// MemoryPath is the memory file path of VM memory. Used when either BootToBeTemplate or
	// BootFromTemplate is true.
	MemoryPath string

	// DevicesStatePath is the VM device state file path. Used when either BootToBeTemplate or
	// BootFromTemplate is true.
	DevicesStatePath string

	// EntropySource is the path to a host source of
	// entropy (/dev/random, /dev/urandom or real hardware RNG device)
	EntropySource string

	// Shared file system type:
	//   - virtio-9p (default)
	//   - virtio-fs
	SharedFS string

	// VirtioFSDaemon is the virtio-fs vhost-user daemon path
	VirtioFSDaemon string

	// VirtioFSCache cache mode for fs version cache or "none"
	VirtioFSCache string

	// VirtioFSExtraArgs passes options to virtiofsd daemon
	VirtioFSExtraArgs []string

	// File based memory backend root directory
	FileBackedMemRootDir string

	// BlockDeviceCacheSet specifies cache-related options will be set to block devices or not.
	BlockDeviceCacheSet bool

	// BlockDeviceCacheDirect specifies cache-related options for block devices.
	// Denotes whether use of O_DIRECT (bypass the host page cache) is enabled.
	BlockDeviceCacheDirect bool

	// BlockDeviceCacheNoflush specifies cache-related options for block devices.
	// Denotes whether flush requests for the device are ignored.
	BlockDeviceCacheNoflush bool

	// DisableBlockDeviceUse disallows a block device from being used.
	DisableBlockDeviceUse bool

	// EnableIOThreads enables IO to be processed in a separate thread.
	// Supported currently for virtio-scsi driver.
	EnableIOThreads bool

	// Debug changes the default hypervisor and kernel parameters to
	// enable debug output where available.
	Debug bool

	// MemPrealloc specifies if the memory should be pre-allocated
	MemPrealloc bool

	// HugePages specifies if the memory should be pre-allocated from huge pages
	HugePages bool

	// VirtioMem is used to enable/disable virtio-mem
	VirtioMem bool

	// Realtime Used to enable/disable realtime
	Realtime bool

	// Mlock is used to control memory locking when Realtime is enabled
	// Realtime=true and Mlock=false, allows for swapping out of VM memory
	// enabling higher density
	Mlock bool

	// DisableNestingChecks is used to override customizations performed
	// when running on top of another VMM.
	DisableNestingChecks bool

	// UseVSock use a vsock for agent communication
	UseVSock bool

	// DisableImageNvdimm disables nvdimm for guest rootfs image
	DisableImageNvdimm bool

	// HotplugVFIOOnRootBus is used to indicate if devices need to be hotplugged on the
	// root bus instead of a bridge.
	HotplugVFIOOnRootBus bool

	// PCIeRootPort is used to indicate the number of PCIe Root Port devices
	// The PCIe Root Port device is used to hot-plug the PCIe device
	PCIeRootPort uint32

	// BootToBeTemplate used to indicate if the VM is created to be a template VM
	BootToBeTemplate bool

	// BootFromTemplate used to indicate if the VM should be created from a template VM
	BootFromTemplate bool

	// DisableVhostNet is used to indicate if host supports vhost_net
	DisableVhostNet bool

	// GuestHookPath is the path within the VM that will be used for 'drop-in' hooks
	GuestHookPath string

	// VMid is the id of the VM that create the hypervisor if the VM is created by the factory.
	// VMid is "" if the hypervisor is not created by the factory.
	VMid string
}

// KataAgentConfig is a structure storing information needed
// to reach the Kata Containers agent.
type KataAgentConfig struct {
	LongLiveConn bool
	UseVSock     bool
}

// ProxyConfig is a structure storing information needed from any
// proxy in order to be properly initialized.
type ProxyConfig struct {
	Path  string
	Debug bool
}

// ShimConfig is the structure providing specific configuration
// for shim implementation.
type ShimConfig struct {
	Path  string
	Debug bool
}

// NetworkConfig is the network configuration related to a network.
type NetworkConfig struct {
	NetNSPath         string
	NetNsCreated      bool
	DisableNewNetNs   bool
	InterworkingModel int
}

type ContainerConfig struct {
	ID          string
	Annotations map[string]string
	RootFs      string
	// Resources for recoding update
	Resources specs.LinuxResources
}

// SandboxConfig is a sandbox configuration.
// Refs: virtcontainers/sandbox.go:SandboxConfig
type SandboxConfig struct {
	HypervisorType   string
	HypervisorConfig HypervisorConfig

	// only one agent config can be non-nil according to agent type
	AgentType       string
	KataAgentConfig *KataAgentConfig `json:",omitempty"`

	ProxyType   string
	ProxyConfig ProxyConfig

	ShimType       string
	KataShimConfig *ShimConfig

	NetworkConfig NetworkConfig

	ShmSize uint64

	// SharePidNs sets all containers to share the same sandbox level pid namespace.
	SharePidNs bool

	// Stateful keeps sandbox resources in memory across APIs. Users will be responsible
	// for calling Release() to release the memory resources.
	Stateful bool

	// SystemdCgroup enables systemd cgroup support
	SystemdCgroup bool

	// SandboxCgroupOnly enables cgroup only at podlevel in the host
	SandboxCgroupOnly bool

	DisableGuestSeccomp bool

	// Experimental enables experimental features
	Experimental []string

	// Information for fields not saved:
	// * Annotation: this is kind of casual data, we don't need casual data in persist file,
	// 				if you know this data needs to persist, please gives it
	//				a specific field

	ContainerConfigs []ContainerConfig

	// Cgroups specifies specific cgroup settings for the various subsystems that the container is
	// placed into to limit the resources the container has available
	Cgroups *configs.Cgroup `json:"cgroups"`
}
