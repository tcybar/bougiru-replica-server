/*
 * ギルド戦記API
 *
 * ゲーム「ギルド戦記」で使用するAPIの仕様書です.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package guildapi

type GetQuestsResponse struct {

	// クエストリスト
	QuestList []int32 `json:"quest_list"`
}

// AssertGetQuestsResponseRequired checks if the required fields are not zero-ed
func AssertGetQuestsResponseRequired(obj GetQuestsResponse) error {
	elements := map[string]interface{}{
		"quest_list": obj.QuestList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertGetQuestsResponseConstraints checks if the values respects the defined constraints
func AssertGetQuestsResponseConstraints(obj GetQuestsResponse) error {
	return nil
}
