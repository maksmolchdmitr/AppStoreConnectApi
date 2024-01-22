/*
App Store Connect API

Testing EndAppAvailabilityPreOrdersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_EndAppAvailabilityPreOrdersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EndAppAvailabilityPreOrdersAPIService EndAppAvailabilityPreOrdersCreateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndAppAvailabilityPreOrdersAPI.EndAppAvailabilityPreOrdersCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
