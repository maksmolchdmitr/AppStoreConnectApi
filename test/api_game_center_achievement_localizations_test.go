/*
App Store Connect API

Testing GameCenterAchievementLocalizationsAPIService

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

func Test_openapi_GameCenterAchievementLocalizationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GameCenterAchievementLocalizationsAPIService GameCenterAchievementLocalizationsCreateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsCreateInstance(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterAchievementLocalizationsAPIService GameCenterAchievementLocalizationsDeleteInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsDeleteInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterAchievementLocalizationsAPIService GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterAchievementLocalizationsAPIService GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterAchievementLocalizationsAPIService GameCenterAchievementLocalizationsGetInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGetInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GameCenterAchievementLocalizationsAPIService GameCenterAchievementLocalizationsUpdateInstance", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsUpdateInstance(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
