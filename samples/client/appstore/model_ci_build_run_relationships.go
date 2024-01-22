/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CiBuildRunRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunRelationships{}

// CiBuildRunRelationships struct for CiBuildRunRelationships
type CiBuildRunRelationships struct {
	Builds *AppRelationshipsBuilds `json:"builds,omitempty"`
	Workflow *CiBuildRunRelationshipsWorkflow `json:"workflow,omitempty"`
	Product *AppRelationshipsCiProduct `json:"product,omitempty"`
	SourceBranchOrTag *CiBuildRunRelationshipsSourceBranchOrTag `json:"sourceBranchOrTag,omitempty"`
	DestinationBranch *CiBuildRunRelationshipsSourceBranchOrTag `json:"destinationBranch,omitempty"`
	PullRequest *CiBuildRunRelationshipsPullRequest `json:"pullRequest,omitempty"`
}

// NewCiBuildRunRelationships instantiates a new CiBuildRunRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunRelationships() *CiBuildRunRelationships {
	this := CiBuildRunRelationships{}
	return &this
}

// NewCiBuildRunRelationshipsWithDefaults instantiates a new CiBuildRunRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunRelationshipsWithDefaults() *CiBuildRunRelationships {
	this := CiBuildRunRelationships{}
	return &this
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *CiBuildRunRelationships) GetBuilds() AppRelationshipsBuilds {
	if o == nil || IsNil(o.Builds) {
		var ret AppRelationshipsBuilds
		return ret
	}
	return *o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationships) GetBuildsOk() (*AppRelationshipsBuilds, bool) {
	if o == nil || IsNil(o.Builds) {
		return nil, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *CiBuildRunRelationships) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given AppRelationshipsBuilds and assigns it to the Builds field.
func (o *CiBuildRunRelationships) SetBuilds(v AppRelationshipsBuilds) {
	o.Builds = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *CiBuildRunRelationships) GetWorkflow() CiBuildRunRelationshipsWorkflow {
	if o == nil || IsNil(o.Workflow) {
		var ret CiBuildRunRelationshipsWorkflow
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationships) GetWorkflowOk() (*CiBuildRunRelationshipsWorkflow, bool) {
	if o == nil || IsNil(o.Workflow) {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *CiBuildRunRelationships) HasWorkflow() bool {
	if o != nil && !IsNil(o.Workflow) {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given CiBuildRunRelationshipsWorkflow and assigns it to the Workflow field.
func (o *CiBuildRunRelationships) SetWorkflow(v CiBuildRunRelationshipsWorkflow) {
	o.Workflow = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *CiBuildRunRelationships) GetProduct() AppRelationshipsCiProduct {
	if o == nil || IsNil(o.Product) {
		var ret AppRelationshipsCiProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationships) GetProductOk() (*AppRelationshipsCiProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *CiBuildRunRelationships) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given AppRelationshipsCiProduct and assigns it to the Product field.
func (o *CiBuildRunRelationships) SetProduct(v AppRelationshipsCiProduct) {
	o.Product = &v
}

// GetSourceBranchOrTag returns the SourceBranchOrTag field value if set, zero value otherwise.
func (o *CiBuildRunRelationships) GetSourceBranchOrTag() CiBuildRunRelationshipsSourceBranchOrTag {
	if o == nil || IsNil(o.SourceBranchOrTag) {
		var ret CiBuildRunRelationshipsSourceBranchOrTag
		return ret
	}
	return *o.SourceBranchOrTag
}

// GetSourceBranchOrTagOk returns a tuple with the SourceBranchOrTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationships) GetSourceBranchOrTagOk() (*CiBuildRunRelationshipsSourceBranchOrTag, bool) {
	if o == nil || IsNil(o.SourceBranchOrTag) {
		return nil, false
	}
	return o.SourceBranchOrTag, true
}

// HasSourceBranchOrTag returns a boolean if a field has been set.
func (o *CiBuildRunRelationships) HasSourceBranchOrTag() bool {
	if o != nil && !IsNil(o.SourceBranchOrTag) {
		return true
	}

	return false
}

// SetSourceBranchOrTag gets a reference to the given CiBuildRunRelationshipsSourceBranchOrTag and assigns it to the SourceBranchOrTag field.
func (o *CiBuildRunRelationships) SetSourceBranchOrTag(v CiBuildRunRelationshipsSourceBranchOrTag) {
	o.SourceBranchOrTag = &v
}

// GetDestinationBranch returns the DestinationBranch field value if set, zero value otherwise.
func (o *CiBuildRunRelationships) GetDestinationBranch() CiBuildRunRelationshipsSourceBranchOrTag {
	if o == nil || IsNil(o.DestinationBranch) {
		var ret CiBuildRunRelationshipsSourceBranchOrTag
		return ret
	}
	return *o.DestinationBranch
}

// GetDestinationBranchOk returns a tuple with the DestinationBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationships) GetDestinationBranchOk() (*CiBuildRunRelationshipsSourceBranchOrTag, bool) {
	if o == nil || IsNil(o.DestinationBranch) {
		return nil, false
	}
	return o.DestinationBranch, true
}

// HasDestinationBranch returns a boolean if a field has been set.
func (o *CiBuildRunRelationships) HasDestinationBranch() bool {
	if o != nil && !IsNil(o.DestinationBranch) {
		return true
	}

	return false
}

// SetDestinationBranch gets a reference to the given CiBuildRunRelationshipsSourceBranchOrTag and assigns it to the DestinationBranch field.
func (o *CiBuildRunRelationships) SetDestinationBranch(v CiBuildRunRelationshipsSourceBranchOrTag) {
	o.DestinationBranch = &v
}

// GetPullRequest returns the PullRequest field value if set, zero value otherwise.
func (o *CiBuildRunRelationships) GetPullRequest() CiBuildRunRelationshipsPullRequest {
	if o == nil || IsNil(o.PullRequest) {
		var ret CiBuildRunRelationshipsPullRequest
		return ret
	}
	return *o.PullRequest
}

// GetPullRequestOk returns a tuple with the PullRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationships) GetPullRequestOk() (*CiBuildRunRelationshipsPullRequest, bool) {
	if o == nil || IsNil(o.PullRequest) {
		return nil, false
	}
	return o.PullRequest, true
}

// HasPullRequest returns a boolean if a field has been set.
func (o *CiBuildRunRelationships) HasPullRequest() bool {
	if o != nil && !IsNil(o.PullRequest) {
		return true
	}

	return false
}

// SetPullRequest gets a reference to the given CiBuildRunRelationshipsPullRequest and assigns it to the PullRequest field.
func (o *CiBuildRunRelationships) SetPullRequest(v CiBuildRunRelationshipsPullRequest) {
	o.PullRequest = &v
}

func (o CiBuildRunRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	if !IsNil(o.Workflow) {
		toSerialize["workflow"] = o.Workflow
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.SourceBranchOrTag) {
		toSerialize["sourceBranchOrTag"] = o.SourceBranchOrTag
	}
	if !IsNil(o.DestinationBranch) {
		toSerialize["destinationBranch"] = o.DestinationBranch
	}
	if !IsNil(o.PullRequest) {
		toSerialize["pullRequest"] = o.PullRequest
	}
	return toSerialize, nil
}

type NullableCiBuildRunRelationships struct {
	value *CiBuildRunRelationships
	isSet bool
}

func (v NullableCiBuildRunRelationships) Get() *CiBuildRunRelationships {
	return v.value
}

func (v *NullableCiBuildRunRelationships) Set(val *CiBuildRunRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunRelationships(val *CiBuildRunRelationships) *NullableCiBuildRunRelationships {
	return &NullableCiBuildRunRelationships{value: val, isSet: true}
}

func (v NullableCiBuildRunRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


