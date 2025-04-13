package realm_settings

import (
	"backend/lib/game_server/dal/services/settings"
	"backend/pkg/database"
	"backend/pkg/sdk/game/enum"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func SyncRealmSettings(basePath string, db *database.Database) (err error) {
	ctx := context.Background()
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	body, err := os.ReadFile(basePath + "realm.json")
	if err != nil {
		return fmt.Errorf("unable to read file: %v", err)
	}
	realmSettings, err := ParseRealmFromJSON(body)
	if err != nil {
		return err
	}
	err = realmSettings.Sync(ctx, tx)
	if err != nil {
		return err
	}

	default_buildings, err := os.ReadDir(basePath + "buildings/")
	if err != nil {
		return err
	}
	for _, file := range default_buildings {
		if file.IsDir() {
			continue
		}
		building_name := strings.TrimSuffix(file.Name(), ".json")

		body, err := os.ReadFile(basePath + "buildings/" + file.Name())
		if err != nil {
			return fmt.Errorf("unable to read file: %v", err)
		}
		factions := []enum.Faction{enum.Caldari, enum.Varnak, enum.Dawnhold, enum.Forgotten}
		for _, faction := range factions {
			handler, err := GetJsonHandler(building_name)
			if err != nil {
				return err
			}
			settings, err := handler(faction, body)
			if err != nil {
				return err
			}
			for _, setting := range settings {
				err = setting.Sync(ctx, tx)
				if err != nil {
					return err
				}
			}
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	log.Println("All json settings synced successfully.")
	return nil
}

func ParseRealmFromJSON(jsonData []byte) (result settings.Realm, err error) {
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
