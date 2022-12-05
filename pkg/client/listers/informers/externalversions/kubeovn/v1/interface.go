/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/kubeovn/kube-ovn/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// HtbQoses returns a HtbQosInformer.
	HtbQoses() HtbQosInformer
	// IPs returns a IPInformer.
	IPs() IPInformer
	// IptablesDnatRules returns a IptablesDnatRuleInformer.
	IptablesDnatRules() IptablesDnatRuleInformer
	// IptablesEIPs returns a IptablesEIPInformer.
	IptablesEIPs() IptablesEIPInformer
	// IptablesFIPRules returns a IptablesFIPRuleInformer.
	IptablesFIPRules() IptablesFIPRuleInformer
	// IptablesSnatRules returns a IptablesSnatRuleInformer.
	IptablesSnatRules() IptablesSnatRuleInformer
	// OvnEips returns a OvnEipInformer.
	OvnEips() OvnEipInformer
	// OvnFips returns a OvnFipInformer.
	OvnFips() OvnFipInformer
	// OvnSnatRules returns a OvnSnatRuleInformer.
	OvnSnatRules() OvnSnatRuleInformer
	// ProviderNetworks returns a ProviderNetworkInformer.
	ProviderNetworks() ProviderNetworkInformer
	// SecurityGroups returns a SecurityGroupInformer.
	SecurityGroups() SecurityGroupInformer
	// Subnets returns a SubnetInformer.
	Subnets() SubnetInformer
	// SwitchLBRules returns a SwitchLBRuleInformer.
	SwitchLBRules() SwitchLBRuleInformer
	// Vips returns a VipInformer.
	Vips() VipInformer
	// Vlans returns a VlanInformer.
	Vlans() VlanInformer
	// Vpcs returns a VpcInformer.
	Vpcs() VpcInformer
	// VpcDnses returns a VpcDnsInformer.
	VpcDnses() VpcDnsInformer
	// VpcNatGateways returns a VpcNatGatewayInformer.
	VpcNatGateways() VpcNatGatewayInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// HtbQoses returns a HtbQosInformer.
func (v *version) HtbQoses() HtbQosInformer {
	return &htbQosInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// IPs returns a IPInformer.
func (v *version) IPs() IPInformer {
	return &iPInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// IptablesDnatRules returns a IptablesDnatRuleInformer.
func (v *version) IptablesDnatRules() IptablesDnatRuleInformer {
	return &iptablesDnatRuleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// IptablesEIPs returns a IptablesEIPInformer.
func (v *version) IptablesEIPs() IptablesEIPInformer {
	return &iptablesEIPInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// IptablesFIPRules returns a IptablesFIPRuleInformer.
func (v *version) IptablesFIPRules() IptablesFIPRuleInformer {
	return &iptablesFIPRuleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// IptablesSnatRules returns a IptablesSnatRuleInformer.
func (v *version) IptablesSnatRules() IptablesSnatRuleInformer {
	return &iptablesSnatRuleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OvnEips returns a OvnEipInformer.
func (v *version) OvnEips() OvnEipInformer {
	return &ovnEipInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OvnFips returns a OvnFipInformer.
func (v *version) OvnFips() OvnFipInformer {
	return &ovnFipInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// OvnSnatRules returns a OvnSnatRuleInformer.
func (v *version) OvnSnatRules() OvnSnatRuleInformer {
	return &ovnSnatRuleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ProviderNetworks returns a ProviderNetworkInformer.
func (v *version) ProviderNetworks() ProviderNetworkInformer {
	return &providerNetworkInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// SecurityGroups returns a SecurityGroupInformer.
func (v *version) SecurityGroups() SecurityGroupInformer {
	return &securityGroupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Subnets returns a SubnetInformer.
func (v *version) Subnets() SubnetInformer {
	return &subnetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// SwitchLBRules returns a SwitchLBRuleInformer.
func (v *version) SwitchLBRules() SwitchLBRuleInformer {
	return &switchLBRuleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Vips returns a VipInformer.
func (v *version) Vips() VipInformer {
	return &vipInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Vlans returns a VlanInformer.
func (v *version) Vlans() VlanInformer {
	return &vlanInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Vpcs returns a VpcInformer.
func (v *version) Vpcs() VpcInformer {
	return &vpcInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VpcDnses returns a VpcDnsInformer.
func (v *version) VpcDnses() VpcDnsInformer {
	return &vpcDnsInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VpcNatGateways returns a VpcNatGatewayInformer.
func (v *version) VpcNatGateways() VpcNatGatewayInformer {
	return &vpcNatGatewayInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
