// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func TestCreatePublicIPAddress(t *testing.T) {
	expected := PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('masterPublicIPAddressName')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				DNSSettings: &network.PublicIPAddressDNSSettings{
					DomainNameLabel: to.StringPtr("[variables('masterFqdnPrefix')]"),
				},
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}

	actual := CreatePublicIPAddress()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateJumpboxPublicIPAddress(t *testing.T) {
	expected := PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("[variables('jumpboxPublicIpAddressName')]"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				DNSSettings: &network.PublicIPAddressDNSSettings{
					DomainNameLabel: to.StringPtr("[variables('masterFqdnPrefix')]"),
				},
				PublicIPAllocationMethod: network.Dynamic,
			},
			Sku: &network.PublicIPAddressSku{
				Name: network.PublicIPAddressSkuNameBasic,
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}

	actual := createJumpboxPublicIPAddress()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateClusterPublicIPAddress(t *testing.T) {
	expected := PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("fee-ipv4"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.Static,
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}

	actual := CreateClusterPublicIPAddress()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}

func TestCreateClusterPublicIPv6Address(t *testing.T) {
	expected := PublicIPAddressARM{
		ARMResource: ARMResource{
			APIVersion: "[variables('apiVersionNetwork')]",
		},
		PublicIPAddress: network.PublicIPAddress{
			Location: to.StringPtr("[variables('location')]"),
			Name:     to.StringPtr("fee-ipv6"),
			PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.Dynamic,
				PublicIPAddressVersion:   "IPv6",
			},
			Sku: &network.PublicIPAddressSku{
				Name: "[variables('loadBalancerSku')]",
			},
			Type: to.StringPtr("Microsoft.Network/publicIPAddresses"),
		},
	}

	actual := CreateClusterPublicIPv6Address()

	diff := cmp.Diff(actual, expected)

	if diff != "" {
		t.Errorf("unexpected diff while expecting equal structs: %s", diff)
	}
}
