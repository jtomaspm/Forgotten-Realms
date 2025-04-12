package realm_settings

import (
	"backend/lib/game_server/dal/services/settings/buildings"
	"backend/pkg/sdk/game/enum"
	"encoding/json"
	"fmt"
	"strconv"
)

func ParseFarmLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]buildings.FarmLevel, error) {
	var rawData struct {
		Levels map[string]buildings.FarmLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []buildings.FarmLevel
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
		result = append(result, obj)
	}

	return result, nil
}

func ParseForestLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]buildings.ForestLevel, error) {
	var rawData struct {
		Levels map[string]buildings.ForestLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []buildings.ForestLevel
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
		result = append(result, obj)
	}

	return result, nil
}

func ParseHeadquartersLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]buildings.HeadquartersLevel, error) {
	var rawData struct {
		Levels map[string]buildings.HeadquartersLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []buildings.HeadquartersLevel
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
		result = append(result, obj)
	}

	return result, nil
}

func ParseMineLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]buildings.MineLevel, error) {
	var rawData struct {
		Levels map[string]buildings.MineLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []buildings.MineLevel
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
		result = append(result, obj)
	}

	return result, nil
}

func ParseQuarryLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]buildings.QuarryLevel, error) {
	var rawData struct {
		Levels map[string]buildings.QuarryLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []buildings.QuarryLevel
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
		result = append(result, obj)
	}

	return result, nil
}

func ParseWarehouseLevelsFromJSON(faction enum.Faction, jsonData []byte) ([]buildings.WarehouseLevel, error) {
	var rawData struct {
		Levels map[string]buildings.WarehouseLevelDto `json:"levels"`
	}

	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var result []buildings.WarehouseLevel
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
		result = append(result, obj)
	}

	return result, nil
}
