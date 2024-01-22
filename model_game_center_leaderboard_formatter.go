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

// GameCenterLeaderboardFormatter the model 'GameCenterLeaderboardFormatter'
type GameCenterLeaderboardFormatter string

// List of GameCenterLeaderboardFormatter
const (
	INTEGER                  GameCenterLeaderboardFormatter = "INTEGER"
	DECIMAL_POINT_1_PLACE    GameCenterLeaderboardFormatter = "DECIMAL_POINT_1_PLACE"
	DECIMAL_POINT_2_PLACE    GameCenterLeaderboardFormatter = "DECIMAL_POINT_2_PLACE"
	DECIMAL_POINT_3_PLACE    GameCenterLeaderboardFormatter = "DECIMAL_POINT_3_PLACE"
	ELAPSED_TIME_MILLISECOND GameCenterLeaderboardFormatter = "ELAPSED_TIME_MILLISECOND"
	ELAPSED_TIME_MINUTE      GameCenterLeaderboardFormatter = "ELAPSED_TIME_MINUTE"
	ELAPSED_TIME_SECOND      GameCenterLeaderboardFormatter = "ELAPSED_TIME_SECOND"
	MONEY_POUND_DECIMAL      GameCenterLeaderboardFormatter = "MONEY_POUND_DECIMAL"
	MONEY_POUND              GameCenterLeaderboardFormatter = "MONEY_POUND"
	MONEY_DOLLAR_DECIMAL     GameCenterLeaderboardFormatter = "MONEY_DOLLAR_DECIMAL"
	MONEY_DOLLAR             GameCenterLeaderboardFormatter = "MONEY_DOLLAR"
	MONEY_EURO_DECIMAL       GameCenterLeaderboardFormatter = "MONEY_EURO_DECIMAL"
	MONEY_EURO               GameCenterLeaderboardFormatter = "MONEY_EURO"
	MONEY_FRANC_DECIMAL      GameCenterLeaderboardFormatter = "MONEY_FRANC_DECIMAL"
	MONEY_FRANC              GameCenterLeaderboardFormatter = "MONEY_FRANC"
	MONEY_KRONER_DECIMAL     GameCenterLeaderboardFormatter = "MONEY_KRONER_DECIMAL"
	MONEY_KRONER             GameCenterLeaderboardFormatter = "MONEY_KRONER"
	MONEY_YEN                GameCenterLeaderboardFormatter = "MONEY_YEN"
)

// All allowed values of GameCenterLeaderboardFormatter enum
var AllowedGameCenterLeaderboardFormatterEnumValues = []GameCenterLeaderboardFormatter{
	"INTEGER",
	"DECIMAL_POINT_1_PLACE",
	"DECIMAL_POINT_2_PLACE",
	"DECIMAL_POINT_3_PLACE",
	"ELAPSED_TIME_MILLISECOND",
	"ELAPSED_TIME_MINUTE",
	"ELAPSED_TIME_SECOND",
	"MONEY_POUND_DECIMAL",
	"MONEY_POUND",
	"MONEY_DOLLAR_DECIMAL",
	"MONEY_DOLLAR",
	"MONEY_EURO_DECIMAL",
	"MONEY_EURO",
	"MONEY_FRANC_DECIMAL",
	"MONEY_FRANC",
	"MONEY_KRONER_DECIMAL",
	"MONEY_KRONER",
	"MONEY_YEN",
}

func (v *GameCenterLeaderboardFormatter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GameCenterLeaderboardFormatter(value)
	for _, existing := range AllowedGameCenterLeaderboardFormatterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GameCenterLeaderboardFormatter", value)
}

// NewGameCenterLeaderboardFormatterFromValue returns a pointer to a valid GameCenterLeaderboardFormatter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGameCenterLeaderboardFormatterFromValue(v string) (*GameCenterLeaderboardFormatter, error) {
	ev := GameCenterLeaderboardFormatter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GameCenterLeaderboardFormatter: valid values are %v", v, AllowedGameCenterLeaderboardFormatterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GameCenterLeaderboardFormatter) IsValid() bool {
	for _, existing := range AllowedGameCenterLeaderboardFormatterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GameCenterLeaderboardFormatter value
func (v GameCenterLeaderboardFormatter) Ptr() *GameCenterLeaderboardFormatter {
	return &v
}

type NullableGameCenterLeaderboardFormatter struct {
	value *GameCenterLeaderboardFormatter
	isSet bool
}

func (v NullableGameCenterLeaderboardFormatter) Get() *GameCenterLeaderboardFormatter {
	return v.value
}

func (v *NullableGameCenterLeaderboardFormatter) Set(val *GameCenterLeaderboardFormatter) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardFormatter) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardFormatter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardFormatter(val *GameCenterLeaderboardFormatter) *NullableGameCenterLeaderboardFormatter {
	return &NullableGameCenterLeaderboardFormatter{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardFormatter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardFormatter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
