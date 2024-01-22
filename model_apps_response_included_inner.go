/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AppsResponseIncludedInner - struct for AppsResponseIncludedInner
type AppsResponseIncludedInner struct {
	AppClip                     *AppClip
	AppCustomProductPage        *AppCustomProductPage
	AppEncryptionDeclaration    *AppEncryptionDeclaration
	AppEvent                    *AppEvent
	AppInfo                     *AppInfo
	AppPreOrder                 *AppPreOrder
	AppPrice                    *AppPrice
	AppStoreVersion             *AppStoreVersion
	AppStoreVersionExperimentV2 *AppStoreVersionExperimentV2
	BetaAppLocalization         *BetaAppLocalization
	BetaAppReviewDetail         *BetaAppReviewDetail
	BetaGroup                   *BetaGroup
	BetaLicenseAgreement        *BetaLicenseAgreement
	Build                       *Build
	CiProduct                   *CiProduct
	EndUserLicenseAgreement     *EndUserLicenseAgreement
	GameCenterDetail            *GameCenterDetail
	GameCenterEnabledVersion    *GameCenterEnabledVersion
	InAppPurchase               *InAppPurchase
	InAppPurchaseV2             *InAppPurchaseV2
	PrereleaseVersion           *PrereleaseVersion
	PromotedPurchase            *PromotedPurchase
	ReviewSubmission            *ReviewSubmission
	SubscriptionGracePeriod     *SubscriptionGracePeriod
	SubscriptionGroup           *SubscriptionGroup
	Territory                   *Territory
}

// AppClipAsAppsResponseIncludedInner is a convenience function that returns AppClip wrapped in AppsResponseIncludedInner
func AppClipAsAppsResponseIncludedInner(v *AppClip) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppClip: v,
	}
}

// AppCustomProductPageAsAppsResponseIncludedInner is a convenience function that returns AppCustomProductPage wrapped in AppsResponseIncludedInner
func AppCustomProductPageAsAppsResponseIncludedInner(v *AppCustomProductPage) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppCustomProductPage: v,
	}
}

// AppEncryptionDeclarationAsAppsResponseIncludedInner is a convenience function that returns AppEncryptionDeclaration wrapped in AppsResponseIncludedInner
func AppEncryptionDeclarationAsAppsResponseIncludedInner(v *AppEncryptionDeclaration) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppEncryptionDeclaration: v,
	}
}

// AppEventAsAppsResponseIncludedInner is a convenience function that returns AppEvent wrapped in AppsResponseIncludedInner
func AppEventAsAppsResponseIncludedInner(v *AppEvent) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppEvent: v,
	}
}

// AppInfoAsAppsResponseIncludedInner is a convenience function that returns AppInfo wrapped in AppsResponseIncludedInner
func AppInfoAsAppsResponseIncludedInner(v *AppInfo) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppInfo: v,
	}
}

// AppPreOrderAsAppsResponseIncludedInner is a convenience function that returns AppPreOrder wrapped in AppsResponseIncludedInner
func AppPreOrderAsAppsResponseIncludedInner(v *AppPreOrder) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppPreOrder: v,
	}
}

// AppPriceAsAppsResponseIncludedInner is a convenience function that returns AppPrice wrapped in AppsResponseIncludedInner
func AppPriceAsAppsResponseIncludedInner(v *AppPrice) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppPrice: v,
	}
}

// AppStoreVersionAsAppsResponseIncludedInner is a convenience function that returns AppStoreVersion wrapped in AppsResponseIncludedInner
func AppStoreVersionAsAppsResponseIncludedInner(v *AppStoreVersion) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppStoreVersion: v,
	}
}

// AppStoreVersionExperimentV2AsAppsResponseIncludedInner is a convenience function that returns AppStoreVersionExperimentV2 wrapped in AppsResponseIncludedInner
func AppStoreVersionExperimentV2AsAppsResponseIncludedInner(v *AppStoreVersionExperimentV2) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		AppStoreVersionExperimentV2: v,
	}
}

// BetaAppLocalizationAsAppsResponseIncludedInner is a convenience function that returns BetaAppLocalization wrapped in AppsResponseIncludedInner
func BetaAppLocalizationAsAppsResponseIncludedInner(v *BetaAppLocalization) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		BetaAppLocalization: v,
	}
}

// BetaAppReviewDetailAsAppsResponseIncludedInner is a convenience function that returns BetaAppReviewDetail wrapped in AppsResponseIncludedInner
func BetaAppReviewDetailAsAppsResponseIncludedInner(v *BetaAppReviewDetail) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		BetaAppReviewDetail: v,
	}
}

// BetaGroupAsAppsResponseIncludedInner is a convenience function that returns BetaGroup wrapped in AppsResponseIncludedInner
func BetaGroupAsAppsResponseIncludedInner(v *BetaGroup) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		BetaGroup: v,
	}
}

// BetaLicenseAgreementAsAppsResponseIncludedInner is a convenience function that returns BetaLicenseAgreement wrapped in AppsResponseIncludedInner
func BetaLicenseAgreementAsAppsResponseIncludedInner(v *BetaLicenseAgreement) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		BetaLicenseAgreement: v,
	}
}

// BuildAsAppsResponseIncludedInner is a convenience function that returns Build wrapped in AppsResponseIncludedInner
func BuildAsAppsResponseIncludedInner(v *Build) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		Build: v,
	}
}

// CiProductAsAppsResponseIncludedInner is a convenience function that returns CiProduct wrapped in AppsResponseIncludedInner
func CiProductAsAppsResponseIncludedInner(v *CiProduct) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		CiProduct: v,
	}
}

// EndUserLicenseAgreementAsAppsResponseIncludedInner is a convenience function that returns EndUserLicenseAgreement wrapped in AppsResponseIncludedInner
func EndUserLicenseAgreementAsAppsResponseIncludedInner(v *EndUserLicenseAgreement) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		EndUserLicenseAgreement: v,
	}
}

// GameCenterDetailAsAppsResponseIncludedInner is a convenience function that returns GameCenterDetail wrapped in AppsResponseIncludedInner
func GameCenterDetailAsAppsResponseIncludedInner(v *GameCenterDetail) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		GameCenterDetail: v,
	}
}

// GameCenterEnabledVersionAsAppsResponseIncludedInner is a convenience function that returns GameCenterEnabledVersion wrapped in AppsResponseIncludedInner
func GameCenterEnabledVersionAsAppsResponseIncludedInner(v *GameCenterEnabledVersion) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		GameCenterEnabledVersion: v,
	}
}

// InAppPurchaseAsAppsResponseIncludedInner is a convenience function that returns InAppPurchase wrapped in AppsResponseIncludedInner
func InAppPurchaseAsAppsResponseIncludedInner(v *InAppPurchase) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		InAppPurchase: v,
	}
}

// InAppPurchaseV2AsAppsResponseIncludedInner is a convenience function that returns InAppPurchaseV2 wrapped in AppsResponseIncludedInner
func InAppPurchaseV2AsAppsResponseIncludedInner(v *InAppPurchaseV2) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		InAppPurchaseV2: v,
	}
}

// PrereleaseVersionAsAppsResponseIncludedInner is a convenience function that returns PrereleaseVersion wrapped in AppsResponseIncludedInner
func PrereleaseVersionAsAppsResponseIncludedInner(v *PrereleaseVersion) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		PrereleaseVersion: v,
	}
}

// PromotedPurchaseAsAppsResponseIncludedInner is a convenience function that returns PromotedPurchase wrapped in AppsResponseIncludedInner
func PromotedPurchaseAsAppsResponseIncludedInner(v *PromotedPurchase) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		PromotedPurchase: v,
	}
}

// ReviewSubmissionAsAppsResponseIncludedInner is a convenience function that returns ReviewSubmission wrapped in AppsResponseIncludedInner
func ReviewSubmissionAsAppsResponseIncludedInner(v *ReviewSubmission) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		ReviewSubmission: v,
	}
}

// SubscriptionGracePeriodAsAppsResponseIncludedInner is a convenience function that returns SubscriptionGracePeriod wrapped in AppsResponseIncludedInner
func SubscriptionGracePeriodAsAppsResponseIncludedInner(v *SubscriptionGracePeriod) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		SubscriptionGracePeriod: v,
	}
}

// SubscriptionGroupAsAppsResponseIncludedInner is a convenience function that returns SubscriptionGroup wrapped in AppsResponseIncludedInner
func SubscriptionGroupAsAppsResponseIncludedInner(v *SubscriptionGroup) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		SubscriptionGroup: v,
	}
}

// TerritoryAsAppsResponseIncludedInner is a convenience function that returns Territory wrapped in AppsResponseIncludedInner
func TerritoryAsAppsResponseIncludedInner(v *Territory) AppsResponseIncludedInner {
	return AppsResponseIncludedInner{
		Territory: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppClip
	err = newStrictDecoder(data).Decode(&dst.AppClip)
	if err == nil {
		jsonAppClip, _ := json.Marshal(dst.AppClip)
		if string(jsonAppClip) == "{}" { // empty struct
			dst.AppClip = nil
		} else {
			match++
		}
	} else {
		dst.AppClip = nil
	}

	// try to unmarshal data into AppCustomProductPage
	err = newStrictDecoder(data).Decode(&dst.AppCustomProductPage)
	if err == nil {
		jsonAppCustomProductPage, _ := json.Marshal(dst.AppCustomProductPage)
		if string(jsonAppCustomProductPage) == "{}" { // empty struct
			dst.AppCustomProductPage = nil
		} else {
			match++
		}
	} else {
		dst.AppCustomProductPage = nil
	}

	// try to unmarshal data into AppEncryptionDeclaration
	err = newStrictDecoder(data).Decode(&dst.AppEncryptionDeclaration)
	if err == nil {
		jsonAppEncryptionDeclaration, _ := json.Marshal(dst.AppEncryptionDeclaration)
		if string(jsonAppEncryptionDeclaration) == "{}" { // empty struct
			dst.AppEncryptionDeclaration = nil
		} else {
			match++
		}
	} else {
		dst.AppEncryptionDeclaration = nil
	}

	// try to unmarshal data into AppEvent
	err = newStrictDecoder(data).Decode(&dst.AppEvent)
	if err == nil {
		jsonAppEvent, _ := json.Marshal(dst.AppEvent)
		if string(jsonAppEvent) == "{}" { // empty struct
			dst.AppEvent = nil
		} else {
			match++
		}
	} else {
		dst.AppEvent = nil
	}

	// try to unmarshal data into AppInfo
	err = newStrictDecoder(data).Decode(&dst.AppInfo)
	if err == nil {
		jsonAppInfo, _ := json.Marshal(dst.AppInfo)
		if string(jsonAppInfo) == "{}" { // empty struct
			dst.AppInfo = nil
		} else {
			match++
		}
	} else {
		dst.AppInfo = nil
	}

	// try to unmarshal data into AppPreOrder
	err = newStrictDecoder(data).Decode(&dst.AppPreOrder)
	if err == nil {
		jsonAppPreOrder, _ := json.Marshal(dst.AppPreOrder)
		if string(jsonAppPreOrder) == "{}" { // empty struct
			dst.AppPreOrder = nil
		} else {
			match++
		}
	} else {
		dst.AppPreOrder = nil
	}

	// try to unmarshal data into AppPrice
	err = newStrictDecoder(data).Decode(&dst.AppPrice)
	if err == nil {
		jsonAppPrice, _ := json.Marshal(dst.AppPrice)
		if string(jsonAppPrice) == "{}" { // empty struct
			dst.AppPrice = nil
		} else {
			match++
		}
	} else {
		dst.AppPrice = nil
	}

	// try to unmarshal data into AppStoreVersion
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersion)
	if err == nil {
		jsonAppStoreVersion, _ := json.Marshal(dst.AppStoreVersion)
		if string(jsonAppStoreVersion) == "{}" { // empty struct
			dst.AppStoreVersion = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersion = nil
	}

	// try to unmarshal data into AppStoreVersionExperimentV2
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperimentV2)
	if err == nil {
		jsonAppStoreVersionExperimentV2, _ := json.Marshal(dst.AppStoreVersionExperimentV2)
		if string(jsonAppStoreVersionExperimentV2) == "{}" { // empty struct
			dst.AppStoreVersionExperimentV2 = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionExperimentV2 = nil
	}

	// try to unmarshal data into BetaAppLocalization
	err = newStrictDecoder(data).Decode(&dst.BetaAppLocalization)
	if err == nil {
		jsonBetaAppLocalization, _ := json.Marshal(dst.BetaAppLocalization)
		if string(jsonBetaAppLocalization) == "{}" { // empty struct
			dst.BetaAppLocalization = nil
		} else {
			match++
		}
	} else {
		dst.BetaAppLocalization = nil
	}

	// try to unmarshal data into BetaAppReviewDetail
	err = newStrictDecoder(data).Decode(&dst.BetaAppReviewDetail)
	if err == nil {
		jsonBetaAppReviewDetail, _ := json.Marshal(dst.BetaAppReviewDetail)
		if string(jsonBetaAppReviewDetail) == "{}" { // empty struct
			dst.BetaAppReviewDetail = nil
		} else {
			match++
		}
	} else {
		dst.BetaAppReviewDetail = nil
	}

	// try to unmarshal data into BetaGroup
	err = newStrictDecoder(data).Decode(&dst.BetaGroup)
	if err == nil {
		jsonBetaGroup, _ := json.Marshal(dst.BetaGroup)
		if string(jsonBetaGroup) == "{}" { // empty struct
			dst.BetaGroup = nil
		} else {
			match++
		}
	} else {
		dst.BetaGroup = nil
	}

	// try to unmarshal data into BetaLicenseAgreement
	err = newStrictDecoder(data).Decode(&dst.BetaLicenseAgreement)
	if err == nil {
		jsonBetaLicenseAgreement, _ := json.Marshal(dst.BetaLicenseAgreement)
		if string(jsonBetaLicenseAgreement) == "{}" { // empty struct
			dst.BetaLicenseAgreement = nil
		} else {
			match++
		}
	} else {
		dst.BetaLicenseAgreement = nil
	}

	// try to unmarshal data into Build
	err = newStrictDecoder(data).Decode(&dst.Build)
	if err == nil {
		jsonBuild, _ := json.Marshal(dst.Build)
		if string(jsonBuild) == "{}" { // empty struct
			dst.Build = nil
		} else {
			match++
		}
	} else {
		dst.Build = nil
	}

	// try to unmarshal data into CiProduct
	err = newStrictDecoder(data).Decode(&dst.CiProduct)
	if err == nil {
		jsonCiProduct, _ := json.Marshal(dst.CiProduct)
		if string(jsonCiProduct) == "{}" { // empty struct
			dst.CiProduct = nil
		} else {
			match++
		}
	} else {
		dst.CiProduct = nil
	}

	// try to unmarshal data into EndUserLicenseAgreement
	err = newStrictDecoder(data).Decode(&dst.EndUserLicenseAgreement)
	if err == nil {
		jsonEndUserLicenseAgreement, _ := json.Marshal(dst.EndUserLicenseAgreement)
		if string(jsonEndUserLicenseAgreement) == "{}" { // empty struct
			dst.EndUserLicenseAgreement = nil
		} else {
			match++
		}
	} else {
		dst.EndUserLicenseAgreement = nil
	}

	// try to unmarshal data into GameCenterDetail
	err = newStrictDecoder(data).Decode(&dst.GameCenterDetail)
	if err == nil {
		jsonGameCenterDetail, _ := json.Marshal(dst.GameCenterDetail)
		if string(jsonGameCenterDetail) == "{}" { // empty struct
			dst.GameCenterDetail = nil
		} else {
			match++
		}
	} else {
		dst.GameCenterDetail = nil
	}

	// try to unmarshal data into GameCenterEnabledVersion
	err = newStrictDecoder(data).Decode(&dst.GameCenterEnabledVersion)
	if err == nil {
		jsonGameCenterEnabledVersion, _ := json.Marshal(dst.GameCenterEnabledVersion)
		if string(jsonGameCenterEnabledVersion) == "{}" { // empty struct
			dst.GameCenterEnabledVersion = nil
		} else {
			match++
		}
	} else {
		dst.GameCenterEnabledVersion = nil
	}

	// try to unmarshal data into InAppPurchase
	err = newStrictDecoder(data).Decode(&dst.InAppPurchase)
	if err == nil {
		jsonInAppPurchase, _ := json.Marshal(dst.InAppPurchase)
		if string(jsonInAppPurchase) == "{}" { // empty struct
			dst.InAppPurchase = nil
		} else {
			match++
		}
	} else {
		dst.InAppPurchase = nil
	}

	// try to unmarshal data into InAppPurchaseV2
	err = newStrictDecoder(data).Decode(&dst.InAppPurchaseV2)
	if err == nil {
		jsonInAppPurchaseV2, _ := json.Marshal(dst.InAppPurchaseV2)
		if string(jsonInAppPurchaseV2) == "{}" { // empty struct
			dst.InAppPurchaseV2 = nil
		} else {
			match++
		}
	} else {
		dst.InAppPurchaseV2 = nil
	}

	// try to unmarshal data into PrereleaseVersion
	err = newStrictDecoder(data).Decode(&dst.PrereleaseVersion)
	if err == nil {
		jsonPrereleaseVersion, _ := json.Marshal(dst.PrereleaseVersion)
		if string(jsonPrereleaseVersion) == "{}" { // empty struct
			dst.PrereleaseVersion = nil
		} else {
			match++
		}
	} else {
		dst.PrereleaseVersion = nil
	}

	// try to unmarshal data into PromotedPurchase
	err = newStrictDecoder(data).Decode(&dst.PromotedPurchase)
	if err == nil {
		jsonPromotedPurchase, _ := json.Marshal(dst.PromotedPurchase)
		if string(jsonPromotedPurchase) == "{}" { // empty struct
			dst.PromotedPurchase = nil
		} else {
			match++
		}
	} else {
		dst.PromotedPurchase = nil
	}

	// try to unmarshal data into ReviewSubmission
	err = newStrictDecoder(data).Decode(&dst.ReviewSubmission)
	if err == nil {
		jsonReviewSubmission, _ := json.Marshal(dst.ReviewSubmission)
		if string(jsonReviewSubmission) == "{}" { // empty struct
			dst.ReviewSubmission = nil
		} else {
			match++
		}
	} else {
		dst.ReviewSubmission = nil
	}

	// try to unmarshal data into SubscriptionGracePeriod
	err = newStrictDecoder(data).Decode(&dst.SubscriptionGracePeriod)
	if err == nil {
		jsonSubscriptionGracePeriod, _ := json.Marshal(dst.SubscriptionGracePeriod)
		if string(jsonSubscriptionGracePeriod) == "{}" { // empty struct
			dst.SubscriptionGracePeriod = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionGracePeriod = nil
	}

	// try to unmarshal data into SubscriptionGroup
	err = newStrictDecoder(data).Decode(&dst.SubscriptionGroup)
	if err == nil {
		jsonSubscriptionGroup, _ := json.Marshal(dst.SubscriptionGroup)
		if string(jsonSubscriptionGroup) == "{}" { // empty struct
			dst.SubscriptionGroup = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionGroup = nil
	}

	// try to unmarshal data into Territory
	err = newStrictDecoder(data).Decode(&dst.Territory)
	if err == nil {
		jsonTerritory, _ := json.Marshal(dst.Territory)
		if string(jsonTerritory) == "{}" { // empty struct
			dst.Territory = nil
		} else {
			match++
		}
	} else {
		dst.Territory = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppClip = nil
		dst.AppCustomProductPage = nil
		dst.AppEncryptionDeclaration = nil
		dst.AppEvent = nil
		dst.AppInfo = nil
		dst.AppPreOrder = nil
		dst.AppPrice = nil
		dst.AppStoreVersion = nil
		dst.AppStoreVersionExperimentV2 = nil
		dst.BetaAppLocalization = nil
		dst.BetaAppReviewDetail = nil
		dst.BetaGroup = nil
		dst.BetaLicenseAgreement = nil
		dst.Build = nil
		dst.CiProduct = nil
		dst.EndUserLicenseAgreement = nil
		dst.GameCenterDetail = nil
		dst.GameCenterEnabledVersion = nil
		dst.InAppPurchase = nil
		dst.InAppPurchaseV2 = nil
		dst.PrereleaseVersion = nil
		dst.PromotedPurchase = nil
		dst.ReviewSubmission = nil
		dst.SubscriptionGracePeriod = nil
		dst.SubscriptionGroup = nil
		dst.Territory = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppClip != nil {
		return json.Marshal(&src.AppClip)
	}

	if src.AppCustomProductPage != nil {
		return json.Marshal(&src.AppCustomProductPage)
	}

	if src.AppEncryptionDeclaration != nil {
		return json.Marshal(&src.AppEncryptionDeclaration)
	}

	if src.AppEvent != nil {
		return json.Marshal(&src.AppEvent)
	}

	if src.AppInfo != nil {
		return json.Marshal(&src.AppInfo)
	}

	if src.AppPreOrder != nil {
		return json.Marshal(&src.AppPreOrder)
	}

	if src.AppPrice != nil {
		return json.Marshal(&src.AppPrice)
	}

	if src.AppStoreVersion != nil {
		return json.Marshal(&src.AppStoreVersion)
	}

	if src.AppStoreVersionExperimentV2 != nil {
		return json.Marshal(&src.AppStoreVersionExperimentV2)
	}

	if src.BetaAppLocalization != nil {
		return json.Marshal(&src.BetaAppLocalization)
	}

	if src.BetaAppReviewDetail != nil {
		return json.Marshal(&src.BetaAppReviewDetail)
	}

	if src.BetaGroup != nil {
		return json.Marshal(&src.BetaGroup)
	}

	if src.BetaLicenseAgreement != nil {
		return json.Marshal(&src.BetaLicenseAgreement)
	}

	if src.Build != nil {
		return json.Marshal(&src.Build)
	}

	if src.CiProduct != nil {
		return json.Marshal(&src.CiProduct)
	}

	if src.EndUserLicenseAgreement != nil {
		return json.Marshal(&src.EndUserLicenseAgreement)
	}

	if src.GameCenterDetail != nil {
		return json.Marshal(&src.GameCenterDetail)
	}

	if src.GameCenterEnabledVersion != nil {
		return json.Marshal(&src.GameCenterEnabledVersion)
	}

	if src.InAppPurchase != nil {
		return json.Marshal(&src.InAppPurchase)
	}

	if src.InAppPurchaseV2 != nil {
		return json.Marshal(&src.InAppPurchaseV2)
	}

	if src.PrereleaseVersion != nil {
		return json.Marshal(&src.PrereleaseVersion)
	}

	if src.PromotedPurchase != nil {
		return json.Marshal(&src.PromotedPurchase)
	}

	if src.ReviewSubmission != nil {
		return json.Marshal(&src.ReviewSubmission)
	}

	if src.SubscriptionGracePeriod != nil {
		return json.Marshal(&src.SubscriptionGracePeriod)
	}

	if src.SubscriptionGroup != nil {
		return json.Marshal(&src.SubscriptionGroup)
	}

	if src.Territory != nil {
		return json.Marshal(&src.Territory)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppsResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AppClip != nil {
		return obj.AppClip
	}

	if obj.AppCustomProductPage != nil {
		return obj.AppCustomProductPage
	}

	if obj.AppEncryptionDeclaration != nil {
		return obj.AppEncryptionDeclaration
	}

	if obj.AppEvent != nil {
		return obj.AppEvent
	}

	if obj.AppInfo != nil {
		return obj.AppInfo
	}

	if obj.AppPreOrder != nil {
		return obj.AppPreOrder
	}

	if obj.AppPrice != nil {
		return obj.AppPrice
	}

	if obj.AppStoreVersion != nil {
		return obj.AppStoreVersion
	}

	if obj.AppStoreVersionExperimentV2 != nil {
		return obj.AppStoreVersionExperimentV2
	}

	if obj.BetaAppLocalization != nil {
		return obj.BetaAppLocalization
	}

	if obj.BetaAppReviewDetail != nil {
		return obj.BetaAppReviewDetail
	}

	if obj.BetaGroup != nil {
		return obj.BetaGroup
	}

	if obj.BetaLicenseAgreement != nil {
		return obj.BetaLicenseAgreement
	}

	if obj.Build != nil {
		return obj.Build
	}

	if obj.CiProduct != nil {
		return obj.CiProduct
	}

	if obj.EndUserLicenseAgreement != nil {
		return obj.EndUserLicenseAgreement
	}

	if obj.GameCenterDetail != nil {
		return obj.GameCenterDetail
	}

	if obj.GameCenterEnabledVersion != nil {
		return obj.GameCenterEnabledVersion
	}

	if obj.InAppPurchase != nil {
		return obj.InAppPurchase
	}

	if obj.InAppPurchaseV2 != nil {
		return obj.InAppPurchaseV2
	}

	if obj.PrereleaseVersion != nil {
		return obj.PrereleaseVersion
	}

	if obj.PromotedPurchase != nil {
		return obj.PromotedPurchase
	}

	if obj.ReviewSubmission != nil {
		return obj.ReviewSubmission
	}

	if obj.SubscriptionGracePeriod != nil {
		return obj.SubscriptionGracePeriod
	}

	if obj.SubscriptionGroup != nil {
		return obj.SubscriptionGroup
	}

	if obj.Territory != nil {
		return obj.Territory
	}

	// all schemas are nil
	return nil
}

type NullableAppsResponseIncludedInner struct {
	value *AppsResponseIncludedInner
	isSet bool
}

func (v NullableAppsResponseIncludedInner) Get() *AppsResponseIncludedInner {
	return v.value
}

func (v *NullableAppsResponseIncludedInner) Set(val *AppsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsResponseIncludedInner(val *AppsResponseIncludedInner) *NullableAppsResponseIncludedInner {
	return &NullableAppsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
