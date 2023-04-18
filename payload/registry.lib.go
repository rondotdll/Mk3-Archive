package main

import (
	"golang.org/x/sys/windows/registry"
)

type PRODUCT_ID struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// self explanatory, gets the system's activation key and if it is OEM or Personal
func GetProductKey() *PRODUCT_ID {
	output := new(PRODUCT_ID)

	sir, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		println(err)
	}
	defer sir.Close()

	Organization, _, err := sir.GetStringValue("RegisteredOrganization")
	if err != nil {
		Organization = ""
	}

	pkr, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\SoftwareProtectionPlatform`, registry.QUERY_VALUE)
	if err != nil {
		println(err)
	}
	defer pkr.Close()

	ActivationKey, _, err := pkr.GetStringValue("BackupProductKeyDefault")
	if err != nil {
		ActivationKey = ""
	}

	KMSClient, _, err := pkr.GetStringValue("KeyManagementServiceName")
	if err != nil {
		KMSClient = ""
	}

	if KMSClient == "" || Organization == "" {
		output.Type = "KMS / OEM"
	} else {
		output.Type = "Retail"
	}

	output.Value = ActivationKey

	return output
}
