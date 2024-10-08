//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationOrderGetByBillingAccount.json
func ExampleReservationOrdersClient_GetByBillingAccount_reservationOrderGetByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationOrdersClient().GetByBillingAccount(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", &armbilling.ReservationOrdersClientGetByBillingAccountOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReservationOrder = armbilling.ReservationOrder{
	// 	Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("microsoft.billing/billingAccounts/reservationOrders"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/reservationOrders/20000000-0000-0000-0000-000000000000"),
	// 	Etag: to.Ptr[int32](27),
	// 	Properties: &armbilling.ReservationOrderProperty{
	// 		BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:34:12.926Z"); return t}()),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbilling.ReservationBillingPlanMonthly),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-AAAA-AAA-AAA"),
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:34:13.973Z"); return t}()),
	// 		DisplayName: to.Ptr("VM_RI_11-24-2021_22-30"),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-25T06:34:12.926Z"); return t}()),
	// 		OriginalQuantity: to.Ptr[int32](1),
	// 		ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:31:18.392Z"); return t}()),
	// 		Reservations: []*armbilling.Reservation{
	// 			{
	// 				ID: to.Ptr("/providers/Microsoft.Capacity/reservationOrders/20000000-0000-0000-0000-000000000000/reservations/20000000-0000-0000-0000-000000000001"),
	// 		}},
	// 		Term: to.Ptr("P3Y"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationOrderGetByBillingAccountWithExpandPlanInformation.json
func ExampleReservationOrdersClient_GetByBillingAccount_reservationOrderGetByBillingAccountWithExpandPlanInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReservationOrdersClient().GetByBillingAccount(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", &armbilling.ReservationOrdersClientGetByBillingAccountOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReservationOrder = armbilling.ReservationOrder{
	// 	Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("microsoft.billing/billingAccounts/reservationOrders"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/reservationOrders/20000000-0000-0000-0000-000000000000"),
	// 	Etag: to.Ptr[int32](26),
	// 	Properties: &armbilling.ReservationOrderProperty{
	// 		BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:34:12.926Z"); return t}()),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbilling.ReservationBillingPlanMonthly),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-AAAA-AAA-AAA"),
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:34:13.973Z"); return t}()),
	// 		DisplayName: to.Ptr("VM_RI_11-24-2021_22-30"),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-25T06:34:12.926Z"); return t}()),
	// 		OriginalQuantity: to.Ptr[int32](1),
	// 		PlanInformation: &armbilling.ReservationOrderBillingPlanInformation{
	// 			NextPaymentDueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-01-25"); return t}()),
	// 			PricingCurrencyTotal: &armbilling.Price{
	// 				Amount: to.Ptr[float64](715.68),
	// 				CurrencyCode: to.Ptr("USD"),
	// 			},
	// 			StartDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-11-25"); return t}()),
	// 			Transactions: []*armbilling.ReservationPaymentDetail{
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-11-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-11-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-12-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2021-12-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-01-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-01-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-02-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-02-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-03-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-03-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-04-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-04-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-05-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-05-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-06-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-06-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-07-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-07-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-08-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-08-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-09-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-09-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-10-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-10-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-11-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-11-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					BillingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-12-25"); return t}()),
	// 					PaymentDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2022-12-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusCompleted),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-01-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-02-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-03-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-04-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-05-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-06-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-07-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-08-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-09-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-10-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-11-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2023-12-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-01-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-02-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-03-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-04-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-05-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-06-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-07-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-08-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-09-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 				},
	// 				{
	// 					DueDate: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-10-25"); return t}()),
	// 					PricingCurrencyTotal: &armbilling.Price{
	// 						Amount: to.Ptr[float64](19.88),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					Status: to.Ptr(armbilling.PaymentStatusScheduled),
	// 			}},
	// 		},
	// 		ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-25T06:31:18.392Z"); return t}()),
	// 		Reservations: []*armbilling.Reservation{
	// 			{
	// 				ID: to.Ptr("/providers/Microsoft.Capacity/reservationOrders/20000000-0000-0000-0000-000000000000/reservations/20000000-0000-0000-0000-000000000001"),
	// 		}},
	// 		Term: to.Ptr("P3Y"),
	// 	},
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/reservationOrdersListByBillingAccount.json
func ExampleReservationOrdersClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReservationOrdersClient().NewListByBillingAccountPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", &armbilling.ReservationOrdersClientListByBillingAccountOptions{Filter: nil,
		OrderBy:   nil,
		Skiptoken: nil,
	})
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
		// page.ReservationOrderList = armbilling.ReservationOrderList{
		// 	Value: []*armbilling.ReservationOrder{
		// 		{
		// 			Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("microsoft.billing/billingAccounts/reservationOrders"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/400000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/reservationOrders/20000000-0000-0000-0000-000000000000"),
		// 			Etag: to.Ptr[int32](10),
		// 			Properties: &armbilling.ReservationOrderProperty{
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-03T21:26:48.512Z"); return t}()),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/400000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingPlan: to.Ptr(armbilling.ReservationBillingPlanUpfront),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/400000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-AAAA-AAA-AAA"),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-03T21:26:50.778Z"); return t}()),
		// 				DisplayName: to.Ptr("SUSE_Plan_08-03-2021_14-22"),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-03T21:26:48.512Z"); return t}()),
		// 				OriginalQuantity: to.Ptr[int32](1),
		// 				ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				RequestDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-03T21:23:47.909Z"); return t}()),
		// 				Reservations: []*armbilling.Reservation{
		// 					{
		// 						ID: to.Ptr("/providers/Microsoft.Capacity/reservationOrders/20000000-0000-0000-0000-000000000000/reservations/20000000-0000-0000-0000-000000000001"),
		// 				}},
		// 				Term: to.Ptr("P3Y"),
		// 			},
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 	}},
		// }
	}
}
