package realm_settings

import (
	"backend/lib/game_server/dal/services/settings"
	"backend/lib/game_server/dal/services/settings/buildings"
	"backend/pkg/sdk/game/enum"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func GetJsonHandler(building string) (func(faction enum.Faction, jsonData []byte) ([]settings.Setting, error), error) {
	switch strings.ToLower(building) {
	case "headquarters":
		return ParseHeadquartersLevelsFromJSON, nil
	case "warehouse":
		return ParseWarehouseLevelsFromJSON, nil
	case "farm":
		return ParseFarmLevelsFromJSON, nil
	case "forest":
		return ParseForestLevelsFromJSON, nil
	case "quarry":
		return ParseQuarryLevelsFromJSON, nil
	case "mine":
		return ParseMineLevelsFromJSON, nil
	default:
		return nil, fmt.Errorf("invalid building: %s", building)
	}
}

func ParseFarmLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]settings.Setting, error) {
	var rawData struct {
		Levels map[string]buildings.FarmLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []settings.Setting
	for levelStr, dto := range rawData.Levels {
		level, err := strconv.Atoi(levelStr)
		if err != nil {
			return nil, fmt.Errorf("invalid level key: %s", levelStr)
		}

		dto.Level = level
		dto.Faction = faction.String()
		obj, err := dto.ToObj()
		if err != nil {
			return nil, fmt.Errorf("invalid dto: %v", dto)
		}
		result = append(result, &obj)
	}

	return result, nil
}

func ParseForestLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]settings.Setting, error) {
	var rawData struct {
		Levels map[string]buildings.ForestLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []settings.Setting
	for levelStr, dto := range rawData.Levels {
		level, err := strconv.Atoi(levelStr)
		if err != nil {
			return nil, fmt.Errorf("invalid level key: %s", levelStr)
		}
		dto.Level = level
		dto.Faction = faction.String()

		obj, err := dto.ToObj()
		if err != nil {
			return nil, fmt.Errorf("invalid dto: %v", dto)
		}
		result = append(result, &obj)
	}

	return result, nil
}

func ParseHeadquartersLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]settings.Setting, error) {
	var rawData struct {
		Levels map[string]buildings.HeadquartersLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []settings.Setting
	for levelStr, dto := range rawData.Levels {
		level, err := strconv.Atoi(levelStr)
		if err != nil {
			return nil, fmt.Errorf("invalid level key: %s", levelStr)
		}
		dto.Level = level
		dto.Faction = faction.String()

		obj, err := dto.ToObj()
		if err != nil {
			return nil, fmt.Errorf("invalid dto: %v", dto)
		}
		result = append(result, &obj)
	}

	return result, nil
}

func ParseMineLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]settings.Setting, error) {
	var rawData struct {
		Levels map[string]buildings.MineLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []settings.Setting
	for levelStr, dto := range rawData.Levels {
		level, err := strconv.Atoi(levelStr)
		if err != nil {
			return nil, fmt.Errorf("invalid level key: %s", levelStr)
		}
		dto.Level = level
		dto.Faction = faction.String()

		obj, err := dto.ToObj()
		if err != nil {
			return nil, fmt.Errorf("invalid dto: %v", dto)
		}
		result = append(result, &obj)
	}

	return result, nil
}

func ParseQuarryLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]settings.Setting, error) {
	var rawData struct {
		Levels map[string]buildings.QuarryLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []settings.Setting
	for levelStr, dto := range rawData.Levels {
		level, err := strconv.Atoi(levelStr)
		if err != nil {
			return nil, fmt.Errorf("invalid level key: %s", levelStr)
		}
		dto.Level = level
		dto.Faction = faction.String()

		obj, err := dto.ToObj()
		if err != nil {
			return nil, fmt.Errorf("invalid dto: %v", dto)
		}
		result = append(result, &obj)
	}

	return result, nil
}

func ParseWarehouseLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]settings.Setting, error) {
	var rawData struct {
		Levels map[string]buildings.WarehouseLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []settings.Setting
	for levelStr, dto := range rawData.Levels {
		level, err := strconv.Atoi(levelStr)
		if err != nil {
			return nil, fmt.Errorf("invalid level key: %s", levelStr)
		}
		dto.Level = level
		dto.Faction = faction.String()

		obj, err := dto.ToObj()
		if err != nil {
			return nil, fmt.Errorf("invalid dto: %v", dto)
		}
		result = append(result, &obj)
	}

	return result, nil
}
