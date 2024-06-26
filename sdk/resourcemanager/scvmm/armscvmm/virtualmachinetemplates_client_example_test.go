//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armscvmm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_ListBySubscription_MaximumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_NewListBySubscriptionPager_virtualMachineTemplatesListBySubscriptionMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineTemplatesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualMachineTemplateListResult = armscvmm.VirtualMachineTemplateListResult{
		// 	Value: []*armscvmm.VirtualMachineTemplate{
		// 		{
		// 			Name: to.Ptr("ioeuwaznkaayvhpqbnrwbr"),
		// 			Type: to.Ptr("egfzqiscydkyddksvsjujdlee"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
		// 			SystemData: &armscvmm.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
		// 				CreatedBy: to.Ptr("p"),
		// 				CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
		// 				LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("ayxsyduviotylbojh"),
		// 			Tags: map[string]*string{
		// 				"key9494": to.Ptr("kkbmfpwhmvlobm"),
		// 			},
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armscvmm.VirtualMachineTemplateProperties{
		// 				ComputerName: to.Ptr("asxghqngsojdsdptpirbz"),
		// 				CPUCount: to.Ptr[int32](23),
		// 				Disks: []*armscvmm.VirtualDisk{
		// 					{
		// 						Name: to.Ptr("fgnckfymwdsqnfxkdvexuaobe"),
		// 						Bus: to.Ptr[int32](8),
		// 						BusType: to.Ptr("zu"),
		// 						CreateDiffDisk: to.Ptr(armscvmm.CreateDiffDiskTrue),
		// 						DiskID: to.Ptr("ltdrwcfjklpsimhzqyh"),
		// 						DiskSizeGB: to.Ptr[int32](30),
		// 						DisplayName: to.Ptr("fgladknawlgjodo"),
		// 						Lun: to.Ptr[int32](10),
		// 						MaxDiskSizeGB: to.Ptr[int32](18),
		// 						StorageQosPolicy: &armscvmm.StorageQosPolicyDetails{
		// 							Name: to.Ptr("ceiyfrflu"),
		// 							ID: to.Ptr("o"),
		// 						},
		// 						TemplateDiskID: to.Ptr("lcdwrokpyvekqccclf"),
		// 						VhdFormatType: to.Ptr("vbcrrmhgahznifudvhxfagwoplcb"),
		// 						VhdType: to.Ptr("cnbeeeylrvopigdynvgpkfp"),
		// 						VolumeType: to.Ptr("ckkymkuekzzqhexyjueruzlfemoeln"),
		// 				}},
		// 				DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
		// 				DynamicMemoryMaxMB: to.Ptr[int32](21),
		// 				DynamicMemoryMinMB: to.Ptr[int32](21),
		// 				Generation: to.Ptr[int32](16),
		// 				InventoryItemID: to.Ptr("qjrykoogccwlgkd"),
		// 				IsCustomizable: to.Ptr(armscvmm.IsCustomizableTrue),
		// 				IsHighlyAvailable: to.Ptr(armscvmm.IsHighlyAvailableTrue),
		// 				LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
		// 				MemoryMB: to.Ptr[int32](24),
		// 				NetworkInterfaces: []*armscvmm.NetworkInterface{
		// 					{
		// 						Name: to.Ptr("kvofzqulbjlbtt"),
		// 						DisplayName: to.Ptr("yoayfd"),
		// 						IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 						IPv4Addresses: []*string{
		// 							to.Ptr("eeunirpkpqazzxhsqonkxcfuks")},
		// 							IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 							IPv6Addresses: []*string{
		// 								to.Ptr("pk")},
		// 								MacAddress: to.Ptr("oaeqqegt"),
		// 								MacAddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 								NetworkName: to.Ptr("lqbm"),
		// 								NicID: to.Ptr("roxpsvlo"),
		// 								VirtualNetworkID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName"),
		// 						}},
		// 						OSName: to.Ptr("qcbolnbisklo"),
		// 						OSType: to.Ptr(armscvmm.OsTypeWindows),
		// 						ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
		// 						UUID: to.Ptr("12345678-1234-1234-1234-12345678abcd"),
		// 						VmmServerID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_ListBySubscription_MinimumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_NewListBySubscriptionPager_virtualMachineTemplatesListBySubscriptionMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineTemplatesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualMachineTemplateListResult = armscvmm.VirtualMachineTemplateListResult{
		// 	Value: []*armscvmm.VirtualMachineTemplate{
		// 		{
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
		// 			Location: to.Ptr("ayxsyduviotylbojh"),
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_ListByResourceGroup_MaximumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_NewListByResourceGroupPager_virtualMachineTemplatesListByResourceGroupMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineTemplatesClient().NewListByResourceGroupPager("rgscvmm", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualMachineTemplateListResult = armscvmm.VirtualMachineTemplateListResult{
		// 	Value: []*armscvmm.VirtualMachineTemplate{
		// 		{
		// 			Name: to.Ptr("ioeuwaznkaayvhpqbnrwbr"),
		// 			Type: to.Ptr("egfzqiscydkyddksvsjujdlee"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
		// 			SystemData: &armscvmm.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
		// 				CreatedBy: to.Ptr("p"),
		// 				CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
		// 				LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("ayxsyduviotylbojh"),
		// 			Tags: map[string]*string{
		// 				"key9494": to.Ptr("kkbmfpwhmvlobm"),
		// 			},
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
		// 				Type: to.Ptr("customLocation"),
		// 			},
		// 			Properties: &armscvmm.VirtualMachineTemplateProperties{
		// 				ComputerName: to.Ptr("asxghqngsojdsdptpirbz"),
		// 				CPUCount: to.Ptr[int32](23),
		// 				Disks: []*armscvmm.VirtualDisk{
		// 					{
		// 						Name: to.Ptr("fgnckfymwdsqnfxkdvexuaobe"),
		// 						Bus: to.Ptr[int32](8),
		// 						BusType: to.Ptr("zu"),
		// 						CreateDiffDisk: to.Ptr(armscvmm.CreateDiffDiskTrue),
		// 						DiskID: to.Ptr("ltdrwcfjklpsimhzqyh"),
		// 						DiskSizeGB: to.Ptr[int32](30),
		// 						DisplayName: to.Ptr("fgladknawlgjodo"),
		// 						Lun: to.Ptr[int32](10),
		// 						MaxDiskSizeGB: to.Ptr[int32](18),
		// 						StorageQosPolicy: &armscvmm.StorageQosPolicyDetails{
		// 							Name: to.Ptr("ceiyfrflu"),
		// 							ID: to.Ptr("o"),
		// 						},
		// 						TemplateDiskID: to.Ptr("lcdwrokpyvekqccclf"),
		// 						VhdFormatType: to.Ptr("vbcrrmhgahznifudvhxfagwoplcb"),
		// 						VhdType: to.Ptr("cnbeeeylrvopigdynvgpkfp"),
		// 						VolumeType: to.Ptr("ckkymkuekzzqhexyjueruzlfemoeln"),
		// 				}},
		// 				DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
		// 				DynamicMemoryMaxMB: to.Ptr[int32](21),
		// 				DynamicMemoryMinMB: to.Ptr[int32](21),
		// 				Generation: to.Ptr[int32](16),
		// 				InventoryItemID: to.Ptr("qjrykoogccwlgkd"),
		// 				IsCustomizable: to.Ptr(armscvmm.IsCustomizableTrue),
		// 				IsHighlyAvailable: to.Ptr(armscvmm.IsHighlyAvailableTrue),
		// 				LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
		// 				MemoryMB: to.Ptr[int32](24),
		// 				NetworkInterfaces: []*armscvmm.NetworkInterface{
		// 					{
		// 						Name: to.Ptr("kvofzqulbjlbtt"),
		// 						DisplayName: to.Ptr("yoayfd"),
		// 						IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 						IPv4Addresses: []*string{
		// 							to.Ptr("eeunirpkpqazzxhsqonkxcfuks")},
		// 							IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 							IPv6Addresses: []*string{
		// 								to.Ptr("pk")},
		// 								MacAddress: to.Ptr("oaeqqegt"),
		// 								MacAddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
		// 								NetworkName: to.Ptr("lqbm"),
		// 								NicID: to.Ptr("roxpsvlo"),
		// 								VirtualNetworkID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName"),
		// 						}},
		// 						OSName: to.Ptr("qcbolnbisklo"),
		// 						OSType: to.Ptr(armscvmm.OsTypeWindows),
		// 						ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
		// 						UUID: to.Ptr("12345678-1234-1234-1234-12345678abcd"),
		// 						VmmServerID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_ListByResourceGroup_MinimumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_NewListByResourceGroupPager_virtualMachineTemplatesListByResourceGroupMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineTemplatesClient().NewListByResourceGroupPager("rgscvmm", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.VirtualMachineTemplateListResult = armscvmm.VirtualMachineTemplateListResult{
		// 	Value: []*armscvmm.VirtualMachineTemplate{
		// 		{
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
		// 			Location: to.Ptr("ayxsyduviotylbojh"),
		// 			ExtendedLocation: &armscvmm.ExtendedLocation{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_Get_MaximumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_Get_virtualMachineTemplatesGetMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineTemplatesClient().Get(ctx, "rgscvmm", "4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineTemplate = armscvmm.VirtualMachineTemplate{
	// 	Name: to.Ptr("ioeuwaznkaayvhpqbnrwbr"),
	// 	Type: to.Ptr("egfzqiscydkyddksvsjujdlee"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
	// 	SystemData: &armscvmm.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
	// 		CreatedBy: to.Ptr("p"),
	// 		CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
	// 		LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("ayxsyduviotylbojh"),
	// 	Tags: map[string]*string{
	// 		"key9494": to.Ptr("kkbmfpwhmvlobm"),
	// 	},
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VirtualMachineTemplateProperties{
	// 		ComputerName: to.Ptr("asxghqngsojdsdptpirbz"),
	// 		CPUCount: to.Ptr[int32](23),
	// 		Disks: []*armscvmm.VirtualDisk{
	// 			{
	// 				Name: to.Ptr("fgnckfymwdsqnfxkdvexuaobe"),
	// 				Bus: to.Ptr[int32](8),
	// 				BusType: to.Ptr("zu"),
	// 				CreateDiffDisk: to.Ptr(armscvmm.CreateDiffDiskTrue),
	// 				DiskID: to.Ptr("ltdrwcfjklpsimhzqyh"),
	// 				DiskSizeGB: to.Ptr[int32](30),
	// 				DisplayName: to.Ptr("fgladknawlgjodo"),
	// 				Lun: to.Ptr[int32](10),
	// 				MaxDiskSizeGB: to.Ptr[int32](18),
	// 				StorageQosPolicy: &armscvmm.StorageQosPolicyDetails{
	// 					Name: to.Ptr("ceiyfrflu"),
	// 					ID: to.Ptr("o"),
	// 				},
	// 				TemplateDiskID: to.Ptr("lcdwrokpyvekqccclf"),
	// 				VhdFormatType: to.Ptr("vbcrrmhgahznifudvhxfagwoplcb"),
	// 				VhdType: to.Ptr("cnbeeeylrvopigdynvgpkfp"),
	// 				VolumeType: to.Ptr("ckkymkuekzzqhexyjueruzlfemoeln"),
	// 		}},
	// 		DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
	// 		DynamicMemoryMaxMB: to.Ptr[int32](21),
	// 		DynamicMemoryMinMB: to.Ptr[int32](21),
	// 		Generation: to.Ptr[int32](16),
	// 		InventoryItemID: to.Ptr("qjrykoogccwlgkd"),
	// 		IsCustomizable: to.Ptr(armscvmm.IsCustomizableTrue),
	// 		IsHighlyAvailable: to.Ptr(armscvmm.IsHighlyAvailableTrue),
	// 		LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
	// 		MemoryMB: to.Ptr[int32](24),
	// 		NetworkInterfaces: []*armscvmm.NetworkInterface{
	// 			{
	// 				Name: to.Ptr("kvofzqulbjlbtt"),
	// 				DisplayName: to.Ptr("yoayfd"),
	// 				IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 				IPv4Addresses: []*string{
	// 					to.Ptr("eeunirpkpqazzxhsqonkxcfuks")},
	// 					IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 					IPv6Addresses: []*string{
	// 						to.Ptr("pk")},
	// 						MacAddress: to.Ptr("oaeqqegt"),
	// 						MacAddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 						NetworkName: to.Ptr("lqbm"),
	// 						NicID: to.Ptr("roxpsvlo"),
	// 						VirtualNetworkID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName"),
	// 				}},
	// 				OSName: to.Ptr("qcbolnbisklo"),
	// 				OSType: to.Ptr(armscvmm.OsTypeWindows),
	// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 				UUID: to.Ptr("12345678-1234-1234-1234-12345678abcd"),
	// 				VmmServerID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_Get_MinimumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_Get_virtualMachineTemplatesGetMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineTemplatesClient().Get(ctx, "rgscvmm", "m", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineTemplate = armscvmm.VirtualMachineTemplate{
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
	// 	Location: to.Ptr("ayxsyduviotylbojh"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_CreateOrUpdate_MaximumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_BeginCreateOrUpdate_virtualMachineTemplatesCreateOrUpdateMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineTemplatesClient().BeginCreateOrUpdate(ctx, "rgscvmm", "6", armscvmm.VirtualMachineTemplate{
		Location: to.Ptr("ayxsyduviotylbojh"),
		Tags: map[string]*string{
			"key9494": to.Ptr("kkbmfpwhmvlobm"),
		},
		ExtendedLocation: &armscvmm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
			Type: to.Ptr("customLocation"),
		},
		Properties: &armscvmm.VirtualMachineTemplateProperties{
			DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
			InventoryItemID:      to.Ptr("qjrykoogccwlgkd"),
			IsCustomizable:       to.Ptr(armscvmm.IsCustomizableTrue),
			IsHighlyAvailable:    to.Ptr(armscvmm.IsHighlyAvailableTrue),
			LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
			OSType:               to.Ptr(armscvmm.OsTypeWindows),
			UUID:                 to.Ptr("12345678-1234-1234-1234-12345678abcd"),
			VmmServerID:          to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineTemplate = armscvmm.VirtualMachineTemplate{
	// 	Name: to.Ptr("ioeuwaznkaayvhpqbnrwbr"),
	// 	Type: to.Ptr("egfzqiscydkyddksvsjujdlee"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
	// 	SystemData: &armscvmm.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
	// 		CreatedBy: to.Ptr("p"),
	// 		CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
	// 		LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("ayxsyduviotylbojh"),
	// 	Tags: map[string]*string{
	// 		"key9494": to.Ptr("kkbmfpwhmvlobm"),
	// 	},
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VirtualMachineTemplateProperties{
	// 		ComputerName: to.Ptr("asxghqngsojdsdptpirbz"),
	// 		CPUCount: to.Ptr[int32](23),
	// 		Disks: []*armscvmm.VirtualDisk{
	// 			{
	// 				Name: to.Ptr("fgnckfymwdsqnfxkdvexuaobe"),
	// 				Bus: to.Ptr[int32](8),
	// 				BusType: to.Ptr("zu"),
	// 				CreateDiffDisk: to.Ptr(armscvmm.CreateDiffDiskTrue),
	// 				DiskID: to.Ptr("ltdrwcfjklpsimhzqyh"),
	// 				DiskSizeGB: to.Ptr[int32](30),
	// 				DisplayName: to.Ptr("fgladknawlgjodo"),
	// 				Lun: to.Ptr[int32](10),
	// 				MaxDiskSizeGB: to.Ptr[int32](18),
	// 				StorageQosPolicy: &armscvmm.StorageQosPolicyDetails{
	// 					Name: to.Ptr("ceiyfrflu"),
	// 					ID: to.Ptr("o"),
	// 				},
	// 				TemplateDiskID: to.Ptr("lcdwrokpyvekqccclf"),
	// 				VhdFormatType: to.Ptr("vbcrrmhgahznifudvhxfagwoplcb"),
	// 				VhdType: to.Ptr("cnbeeeylrvopigdynvgpkfp"),
	// 				VolumeType: to.Ptr("ckkymkuekzzqhexyjueruzlfemoeln"),
	// 		}},
	// 		DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
	// 		DynamicMemoryMaxMB: to.Ptr[int32](21),
	// 		DynamicMemoryMinMB: to.Ptr[int32](21),
	// 		Generation: to.Ptr[int32](16),
	// 		InventoryItemID: to.Ptr("qjrykoogccwlgkd"),
	// 		IsCustomizable: to.Ptr(armscvmm.IsCustomizableTrue),
	// 		IsHighlyAvailable: to.Ptr(armscvmm.IsHighlyAvailableTrue),
	// 		LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
	// 		MemoryMB: to.Ptr[int32](24),
	// 		NetworkInterfaces: []*armscvmm.NetworkInterface{
	// 			{
	// 				Name: to.Ptr("kvofzqulbjlbtt"),
	// 				DisplayName: to.Ptr("yoayfd"),
	// 				IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 				IPv4Addresses: []*string{
	// 					to.Ptr("eeunirpkpqazzxhsqonkxcfuks")},
	// 					IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 					IPv6Addresses: []*string{
	// 						to.Ptr("pk")},
	// 						MacAddress: to.Ptr("oaeqqegt"),
	// 						MacAddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 						NetworkName: to.Ptr("lqbm"),
	// 						NicID: to.Ptr("roxpsvlo"),
	// 						VirtualNetworkID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName"),
	// 				}},
	// 				OSName: to.Ptr("qcbolnbisklo"),
	// 				OSType: to.Ptr(armscvmm.OsTypeWindows),
	// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 				UUID: to.Ptr("12345678-1234-1234-1234-12345678abcd"),
	// 				VmmServerID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_CreateOrUpdate_MinimumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_BeginCreateOrUpdate_virtualMachineTemplatesCreateOrUpdateMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineTemplatesClient().BeginCreateOrUpdate(ctx, "rgscvmm", "P", armscvmm.VirtualMachineTemplate{
		Location:         to.Ptr("ayxsyduviotylbojh"),
		ExtendedLocation: &armscvmm.ExtendedLocation{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineTemplate = armscvmm.VirtualMachineTemplate{
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
	// 	Location: to.Ptr("ayxsyduviotylbojh"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_Update_MaximumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_BeginUpdate_virtualMachineTemplatesUpdateMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineTemplatesClient().BeginUpdate(ctx, "rgscvmm", "g", armscvmm.VirtualMachineTemplateTagsUpdate{
		Tags: map[string]*string{
			"key6634": to.Ptr("wwfhrg"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineTemplate = armscvmm.VirtualMachineTemplate{
	// 	Name: to.Ptr("ioeuwaznkaayvhpqbnrwbr"),
	// 	Type: to.Ptr("egfzqiscydkyddksvsjujdlee"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualMachineTemplates/virtualMachineTemplateName"),
	// 	SystemData: &armscvmm.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.094Z"); return t}()),
	// 		CreatedBy: to.Ptr("p"),
	// 		CreatedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-29T22:28:00.095Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("goxcwpyyqlxndquly"),
	// 		LastModifiedByType: to.Ptr(armscvmm.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("ayxsyduviotylbojh"),
	// 	Tags: map[string]*string{
	// 		"key9494": to.Ptr("kkbmfpwhmvlobm"),
	// 	},
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/customLocationName"),
	// 		Type: to.Ptr("customLocation"),
	// 	},
	// 	Properties: &armscvmm.VirtualMachineTemplateProperties{
	// 		ComputerName: to.Ptr("asxghqngsojdsdptpirbz"),
	// 		CPUCount: to.Ptr[int32](23),
	// 		Disks: []*armscvmm.VirtualDisk{
	// 			{
	// 				Name: to.Ptr("fgnckfymwdsqnfxkdvexuaobe"),
	// 				Bus: to.Ptr[int32](8),
	// 				BusType: to.Ptr("zu"),
	// 				CreateDiffDisk: to.Ptr(armscvmm.CreateDiffDiskTrue),
	// 				DiskID: to.Ptr("ltdrwcfjklpsimhzqyh"),
	// 				DiskSizeGB: to.Ptr[int32](30),
	// 				DisplayName: to.Ptr("fgladknawlgjodo"),
	// 				Lun: to.Ptr[int32](10),
	// 				MaxDiskSizeGB: to.Ptr[int32](18),
	// 				StorageQosPolicy: &armscvmm.StorageQosPolicyDetails{
	// 					Name: to.Ptr("ceiyfrflu"),
	// 					ID: to.Ptr("o"),
	// 				},
	// 				TemplateDiskID: to.Ptr("lcdwrokpyvekqccclf"),
	// 				VhdFormatType: to.Ptr("vbcrrmhgahznifudvhxfagwoplcb"),
	// 				VhdType: to.Ptr("cnbeeeylrvopigdynvgpkfp"),
	// 				VolumeType: to.Ptr("ckkymkuekzzqhexyjueruzlfemoeln"),
	// 		}},
	// 		DynamicMemoryEnabled: to.Ptr(armscvmm.DynamicMemoryEnabledTrue),
	// 		DynamicMemoryMaxMB: to.Ptr[int32](21),
	// 		DynamicMemoryMinMB: to.Ptr[int32](21),
	// 		Generation: to.Ptr[int32](16),
	// 		InventoryItemID: to.Ptr("qjrykoogccwlgkd"),
	// 		IsCustomizable: to.Ptr(armscvmm.IsCustomizableTrue),
	// 		IsHighlyAvailable: to.Ptr(armscvmm.IsHighlyAvailableTrue),
	// 		LimitCPUForMigration: to.Ptr(armscvmm.LimitCPUForMigrationTrue),
	// 		MemoryMB: to.Ptr[int32](24),
	// 		NetworkInterfaces: []*armscvmm.NetworkInterface{
	// 			{
	// 				Name: to.Ptr("kvofzqulbjlbtt"),
	// 				DisplayName: to.Ptr("yoayfd"),
	// 				IPv4AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 				IPv4Addresses: []*string{
	// 					to.Ptr("eeunirpkpqazzxhsqonkxcfuks")},
	// 					IPv6AddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 					IPv6Addresses: []*string{
	// 						to.Ptr("pk")},
	// 						MacAddress: to.Ptr("oaeqqegt"),
	// 						MacAddressType: to.Ptr(armscvmm.AllocationMethodDynamic),
	// 						NetworkName: to.Ptr("lqbm"),
	// 						NicID: to.Ptr("roxpsvlo"),
	// 						VirtualNetworkID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/virtualNetworks/virtualNetworkName"),
	// 				}},
	// 				OSName: to.Ptr("qcbolnbisklo"),
	// 				OSType: to.Ptr(armscvmm.OsTypeWindows),
	// 				ProvisioningState: to.Ptr(armscvmm.ProvisioningStateSucceeded),
	// 				UUID: to.Ptr("12345678-1234-1234-1234-12345678abcd"),
	// 				VmmServerID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.ScVmm/vmmServers/vmmServerName"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_Update_MinimumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_BeginUpdate_virtualMachineTemplatesUpdateMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineTemplatesClient().BeginUpdate(ctx, "rgscvmm", "-", armscvmm.VirtualMachineTemplateTagsUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineTemplate = armscvmm.VirtualMachineTemplate{
	// 	Location: to.Ptr("ayxsyduviotylbojh"),
	// 	ExtendedLocation: &armscvmm.ExtendedLocation{
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_Delete_MaximumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_BeginDelete_virtualMachineTemplatesDeleteMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineTemplatesClient().BeginDelete(ctx, "rgscvmm", "6", &armscvmm.VirtualMachineTemplatesClientBeginDeleteOptions{Force: to.Ptr(armscvmm.ForceDeleteTrue)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7753cb8917f0968713c013a1f25875e8bd8dc492/specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineTemplates_Delete_MinimumSet_Gen.json
func ExampleVirtualMachineTemplatesClient_BeginDelete_virtualMachineTemplatesDeleteMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armscvmm.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineTemplatesClient().BeginDelete(ctx, "rgscvmm", "5", &armscvmm.VirtualMachineTemplatesClientBeginDeleteOptions{Force: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
