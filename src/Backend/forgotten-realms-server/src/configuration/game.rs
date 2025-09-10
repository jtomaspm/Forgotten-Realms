pub struct GameConfig {
    pub BuildingUpgrades: BuildingUpgradeConfig,
}

pub struct BuildingUpgradeConfig {
    pub costs: HashMap<BuildingType, Vec<>>,
}

