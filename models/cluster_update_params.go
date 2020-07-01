// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterUpdateParams cluster update params
//
// swagger:model cluster-update-params
type ClusterUpdateParams struct {

	// Virtual IP used to reach the OpenShift cluster API.
	// Pattern: ^(([0-9]{1,3}\.){3}[0-9]{1,3})?$
	APIVip *string `json:"api_vip,omitempty"`

	// Base domain of the cluster. All DNS records must be sub-domains of this base and include the cluster name.
	BaseDNSDomain *string `json:"base_dns_domain,omitempty"`

	// IP address block from which Pod IPs are allocated This block must not overlap with existing physical networks. These IP addresses are used for the Pod network, and if you need to access the Pods from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^([0-9]{1,3}\.){3}[0-9]{1,3}\/[0-9]|[1-2][0-9]|3[0-2]?$
	ClusterNetworkCidr *string `json:"cluster_network_cidr,omitempty"`

	// The subnet prefix length to assign to each individual node. For example, if clusterNetworkHostPrefix is set to 23, then each node is assigned a /23 subnet out of the given cidr (clusterNetworkCIDR), which allows for 510 (2^(32 - 23) - 2) pod IPs addresses. If you are required to provide access to nodes from an external network, configure load balancers and routers to manage the traffic.
	// Maximum: 32
	// Minimum: 1
	ClusterNetworkHostPrefix *int64 `json:"cluster_network_host_prefix,omitempty"`

	// The desired hostname for hosts associated with the cluster.
	HostsNames []*ClusterUpdateParamsHostsNamesItems0 `json:"hosts_names" gorm:"type:varchar(64)[]"`

	// The desired role for hosts associated with the cluster.
	HostsRoles []*ClusterUpdateParamsHostsRolesItems0 `json:"hosts_roles" gorm:"type:varchar(64)[]"`

	// Virtual IP used for cluster ingress traffic.
	// Pattern: ^(([0-9]{1,3}\.){3}[0-9]{1,3})?$
	IngressVip *string `json:"ingress_vip,omitempty"`

	// OpenShift cluster name
	Name *string `json:"name,omitempty"`

	// The pull secret that obtained from the Pull Secret page on the Red Hat OpenShift Cluster Manager site.
	PullSecret *string `json:"pull_secret,omitempty"`

	// The IP address pool to use for service IP addresses. You can enter only one IP address pool. If you need to access the services from an external network, configure load balancers and routers to manage the traffic.
	// Pattern: ^([0-9]{1,3}\.){3}[0-9]{1,3}\/[0-9]|[1-2][0-9]|3[0-2]?$
	ServiceNetworkCidr *string `json:"service_network_cidr,omitempty"`

	// SSH public key for debugging OpenShift nodes.
	SSHPublicKey *string `json:"ssh_public_key,omitempty"`
}

// Validate validates this cluster update params
func (m *ClusterUpdateParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNetworkHostPrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostsNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostsRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressVip(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUpdateParams) validateAPIVip(formats strfmt.Registry) error {

	if swag.IsZero(m.APIVip) { // not required
		return nil
	}

	if err := validate.Pattern("api_vip", "body", string(*m.APIVip), `^(([0-9]{1,3}\.){3}[0-9]{1,3})?$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterUpdateParams) validateClusterNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("cluster_network_cidr", "body", string(*m.ClusterNetworkCidr), `^([0-9]{1,3}\.){3}[0-9]{1,3}\/[0-9]|[1-2][0-9]|3[0-2]?$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterUpdateParams) validateClusterNetworkHostPrefix(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNetworkHostPrefix) { // not required
		return nil
	}

	if err := validate.MinimumInt("cluster_network_host_prefix", "body", int64(*m.ClusterNetworkHostPrefix), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("cluster_network_host_prefix", "body", int64(*m.ClusterNetworkHostPrefix), 32, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterUpdateParams) validateHostsNames(formats strfmt.Registry) error {

	if swag.IsZero(m.HostsNames) { // not required
		return nil
	}

	for i := 0; i < len(m.HostsNames); i++ {
		if swag.IsZero(m.HostsNames[i]) { // not required
			continue
		}

		if m.HostsNames[i] != nil {
			if err := m.HostsNames[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts_names" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterUpdateParams) validateHostsRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.HostsRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.HostsRoles); i++ {
		if swag.IsZero(m.HostsRoles[i]) { // not required
			continue
		}

		if m.HostsRoles[i] != nil {
			if err := m.HostsRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterUpdateParams) validateIngressVip(formats strfmt.Registry) error {

	if swag.IsZero(m.IngressVip) { // not required
		return nil
	}

	if err := validate.Pattern("ingress_vip", "body", string(*m.IngressVip), `^(([0-9]{1,3}\.){3}[0-9]{1,3})?$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterUpdateParams) validateServiceNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceNetworkCidr) { // not required
		return nil
	}

	if err := validate.Pattern("service_network_cidr", "body", string(*m.ServiceNetworkCidr), `^([0-9]{1,3}\.){3}[0-9]{1,3}\/[0-9]|[1-2][0-9]|3[0-2]?$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUpdateParams) UnmarshalBinary(b []byte) error {
	var res ClusterUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterUpdateParamsHostsNamesItems0 cluster update params hosts names items0
//
// swagger:model ClusterUpdateParamsHostsNamesItems0
type ClusterUpdateParamsHostsNamesItems0 struct {

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`
}

// Validate validates this cluster update params hosts names items0
func (m *ClusterUpdateParamsHostsNamesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUpdateParamsHostsNamesItems0) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterUpdateParamsHostsNamesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUpdateParamsHostsNamesItems0) UnmarshalBinary(b []byte) error {
	var res ClusterUpdateParamsHostsNamesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterUpdateParamsHostsRolesItems0 cluster update params hosts roles items0
//
// swagger:model ClusterUpdateParamsHostsRolesItems0
type ClusterUpdateParamsHostsRolesItems0 struct {

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// role
	// Enum: [master worker]
	Role string `json:"role,omitempty"`
}

// Validate validates this cluster update params hosts roles items0
func (m *ClusterUpdateParamsHostsRolesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUpdateParamsHostsRolesItems0) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var clusterUpdateParamsHostsRolesItems0TypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["master","worker"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterUpdateParamsHostsRolesItems0TypeRolePropEnum = append(clusterUpdateParamsHostsRolesItems0TypeRolePropEnum, v)
	}
}

const (

	// ClusterUpdateParamsHostsRolesItems0RoleMaster captures enum value "master"
	ClusterUpdateParamsHostsRolesItems0RoleMaster string = "master"

	// ClusterUpdateParamsHostsRolesItems0RoleWorker captures enum value "worker"
	ClusterUpdateParamsHostsRolesItems0RoleWorker string = "worker"
)

// prop value enum
func (m *ClusterUpdateParamsHostsRolesItems0) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterUpdateParamsHostsRolesItems0TypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterUpdateParamsHostsRolesItems0) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterUpdateParamsHostsRolesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUpdateParamsHostsRolesItems0) UnmarshalBinary(b []byte) error {
	var res ClusterUpdateParamsHostsRolesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
