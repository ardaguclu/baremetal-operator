package empty

import (
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	metal3v1alpha1 "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"
	"github.com/metal3-io/baremetal-operator/pkg/bmc"
	"github.com/metal3-io/baremetal-operator/pkg/provisioner"
)

var log = logf.Log.WithName("provisioner").WithName("empty")

// Provisioner implements the provisioning.Provisioner interface
type emptyProvisioner struct {
}

// New returns a new Empty Provisioner
func New(host *metal3v1alpha1.BareMetalHost, bmcCreds bmc.Credentials, publisher provisioner.EventPublisher) (provisioner.Provisioner, error) {
	return &emptyProvisioner{}, nil
}

// ValidateManagementAccess tests the connection information for the
// host to verify that the location and credentials work.
func (p *emptyProvisioner) ValidateManagementAccess(credentialsChanged bool) (provisioner.Result, error) {
	return provisioner.Result{}, nil
}

// InspectHardware updates the HardwareDetails field of the host with
// details of devices discovered on the hardware. It may be called
// multiple times, and should return true for its dirty flag until the
// inspection is completed.
func (p *emptyProvisioner) InspectHardware() (provisioner.Result, *metal3v1alpha1.HardwareDetails, error) {
	return provisioner.Result{}, nil, nil
}

// UpdateHardwareState fetches the latest hardware state of the server
// and updates the HardwareDetails field of the host with details. It
// is expected to do this in the least expensive way possible, such as
// reading from a cache, and return dirty only if any state
// information has changed.
func (p *emptyProvisioner) UpdateHardwareState() (provisioner.Result, error) {
	return provisioner.Result{}, nil
}

// Adopt allows an externally-provisioned server to be adopted.
func (p *emptyProvisioner) Adopt(force bool) (provisioner.Result, error) {
	return provisioner.Result{}, nil
}

// Provision writes the image from the host spec to the host. It may
// be called multiple times, and should return true for its dirty flag
// until the deprovisioning operation is completed.
func (p *emptyProvisioner) Provision(hostConf provisioner.HostConfigData) (provisioner.Result, error) {
	return provisioner.Result{}, nil
}

// Deprovision removes the host from the image. It may be called
// multiple times, and should return true for its dirty flag until the
// deprovisioning operation is completed.
func (p *emptyProvisioner) Deprovision() (provisioner.Result, error) {
	return provisioner.Result{}, nil
}

// Delete removes the host from the provisioning system. It may be
// called multiple times, and should return true for its dirty flag
// until the deprovisioning operation is completed.
func (p *emptyProvisioner) Delete() (provisioner.Result, error) {
	return provisioner.Result{}, nil
}

// PowerOn ensures the server is powered on independently of any image
// provisioning operation.
func (p *emptyProvisioner) PowerOn() (provisioner.Result, error) {
	return provisioner.Result{}, nil
}

// PowerOff ensures the server is powered off independently of any image
// provisioning operation.
func (p *emptyProvisioner) PowerOff() (provisioner.Result, error) {
	return provisioner.Result{}, nil
}

// IsReady always returns true for the empty provisioner
func (p *emptyProvisioner) IsReady() (bool, error) {
	return true, nil
}
