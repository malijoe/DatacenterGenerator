package datacenter

import "github.com/malijoe/DatacenterGenerator/pkg/internal/units"

// Datacenter represents a datacenter.
type Datacenter struct {
	// the unique identifier for the datacenter
	ID string

	// the site name of the datacenter.
	Site string
	// the building where the datacenter is located.
	Building string
	// the room of the building where the datacenter is located.
	Room string
	// the providers and their transfer speeds for this datacenter.
	Providers map[string]units.Value

	// the racks that comprise the datacenter.
	Racks []*Rack
	// the pods that exist in the datacenter
	Pods []*Pod

	// used to track the number of instances by function of pods in the datacenter.
	podMetadata map[Function]int
	// used to track the number of device configuration instances in the datacenter.
	// map[hardwareModelPID]map[deviceTemplateVariant]int
	deviceMetadata map[string]map[string]int
}

func NewDatacenter() *Datacenter {
	return &Datacenter{
		podMetadata:    make(map[Function]int),
		deviceMetadata: make(map[string]map[string]int),
	}
}

// NumPodInstances returns the number of pod instances with the passed function in the datacenter.
func (d *Datacenter) NumPodInstances(function Function) int {
	return d.podMetadata[function]
}

// NumDeviceInstances returns the number of device instances with the passed modelPID and variant in the datacenter.
func (d *Datacenter) NumDeviceInstances(modelPID string, variant string) int {
	return d.deviceMetadata[modelPID][variant]
}

// CountPod iterates the datacenter's instance counter for pods of the passed function.
func (d *Datacenter) CountPod(function Function) {
	d.podMetadata[function]++
}

// CountDevice iterates the datacenter's instance counter for devices of the passed modelPID and variant.
func (d *Datacenter) CountDevice(modelPID string, variant string) {
	d.deviceMetadata[modelPID][variant]++
}
