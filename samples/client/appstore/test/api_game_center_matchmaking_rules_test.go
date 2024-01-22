/*
App Store Connect API

Testing GameCenterMatchmakingRulesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_GameCenterMatchmakingRulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GameCenterMatchmakingRulesAPIService GameCenterMatchmakingRulesCreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRulesAPIService GameCenterMatchmakingRulesDeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRulesAPIService GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRulesAPIService GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRulesAPIService GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterMatchmakingRulesAPIService GameCenterMatchmakingRulesUpdateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
