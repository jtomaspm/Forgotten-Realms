package realm_settings

import (
	"backend/lib/game_server/dal/services/settings"
	"encoding/json"
)

func LoadRealmSettings(basePath string) (err error) {

	return err
}

func ParseRealmFromJSON(jsonData []byte) (result settings.Realm, err error) {
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
