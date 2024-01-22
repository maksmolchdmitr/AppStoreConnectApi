package main

import (
	"context"
	"fmt"
	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
	"os"
)

func main() {
	openapi.NewBundleIdCapabilityCreateRequest()
	bundleIdCapabilityCreateRequest := *openapiclient.NewBundleIdCapabilityCreateRequest(
		*openapiclient.NewBundleIdCapabilityCreateRequestData(
			"Type_example",
			*openapiclient.NewBundleIdCapabilityCreateRequestDataAttributes("ICLOUD"),
			*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationships(*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("Type_example", "Id_example"))),
		),
	) // BundleIdCapabilityCreateRequest | BundleIdCapability representation

	configuration := openapiclient.NewConfiguration()
	//configuration.AddDefaultHeader("", "")
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance(context.Background()).BundleIdCapabilityCreateRequest(bundleIdCapabilityCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BundleIdCapabilitiesCreateInstance`: BundleIdCapabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `BundleIdCapabilitiesAPI.BundleIdCapabilitiesCreateInstance`: %v\n", resp)
}
