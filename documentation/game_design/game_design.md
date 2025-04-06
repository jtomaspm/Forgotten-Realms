# Game design document: Forgotten Realms

**Game Title**: Forgotten Realms  
**Genre**: Browser-Based MMO-RTS  
**Inspirations**: Tribalwars, Travian, Starcraft, Warcraft  

- [Game design document: Forgotten Realms](#game-design-document-forgotten-realms)
- [1. Game Overview](#1-game-overview)
- [2. Lore and Setting](#2-lore-and-setting)
- [3. Core Gameplay](#3-core-gameplay)
    - [Village Management](#village-management)
    - [Resource System](#resource-system)
    - [Combat](#combat)
    - [Hero System](#hero-system)
    - [Alliances](#alliances)
    - [PvE Content](#pve-content)
    - [Victory Conditions](#victory-conditions)
- [4. Factions Overview](#4-factions-overview)
  - [4.1 Caldari](#41-caldari)
    - [Troops](#troops)
    - [Buildings](#buildings)
    - [Heroes](#heroes)
  - [4.2 Varnak](#42-varnak)
    - [Troops](#troops-1)
    - [Buildings](#buildings-1)
    - [Heroes](#heroes-1)
  - [4.3 Dawnhold](#43-dawnhold)
    - [Troops](#troops-2)
    - [Buildings](#buildings-2)
    - [Heroes](#heroes-2)
  - [4.4 The Forgotten (NPC)](#44-the-forgotten-npc)
- [5. World Map and Progression](#5-world-map-and-progression)
- [6. Monetization](#6-monetization)
- [7. Technical Requirements](#7-technical-requirements)
- [8. Future Plans](#8-future-plans)


# 1. Game Overview

**Description:** Forgotten Realms is a massively multiplayer online real-time strategy (MMO-RTS) game set in a newly discovered world filled with ancient secrets and contested territories. Players choose from one of three playable factions, each with unique units, buildings, and heroes, and compete to dominate the realm. Alliances are central to gameplay, allowing players to support each other, engage in coordinated battles, and strive for realm-wide domination. A fourth, non-playable faction native to the Forgotten Realms serves as a common enemy, offering challenges and rewards.

**Platform:** Web browser

**Target Audience:** Strategy gamers, fans of persistent world games, ages 13+

# 2. Lore and Setting

The Forgotten Realms were long hidden from the known world until a rift was discovered. Now, empires from distant lands have ventured through, seeking to claim this bountiful and mysterious land. However, the realm is not uninhabited. An ancient native faction fiercely defends their home.

**Playable Factions:**

1. **Caldari** - A highly advanced magical-tech faction (inspired by Protoss/High Elves)

2. **Varnak** - Brutal, war-driven creatures with strong melee units (inspired by Orcs/Zerg)

3. **Dawnhold** - Strategic human-like settlers focused on balance and economics (inspired by Travian Romans)

**Non-Playable Faction:**  

1. **The Forgotten** - Ancient native defenders of the realm with powerful NPC-controlled villages and elite troops

# 3. Core Gameplay

### Village Management

Players start with a single village.  
Construct and upgrade buildings for resource production, troop training, research, and defense.  
Expand by founding new villages or conquering others.  

### Resource System

Wood, Stone, Iron

### Combat

Train diverse units with counters (infantry, cavalry, siege, magic).  
Conduct raids, full-scale battles, and sieges.  
Heroes lead armies and provide buffs.  

### Hero System

Each faction has unique hero classes.  
Heroes gain experience, level up, and equip gear.  

### Alliances

Form alliances for mutual defense and coordinated attacks.  
Share resources, send reinforcements, and communicate via alliance chat.  

### PvE Content

Attack villages controlled by The Forgotten.  
PvE villages offer loot, rare gear, and special resources.  
Seasonal events.  

### Victory Conditions

An alliance wins the game by controlling over 60% of all villages in the realm.  

# 4. Factions Overview

## 4.1 Caldari

Magic and tech  synergy.  

### Troops

Energy.infused warriors, flying units.  

* Ground_troop_slot1
  * Arcblade
    * Basic infantry with energy-forged blades; fast training, low cost.
  * Etherblade
    * Young warrior fused with energy conduits; basic frontline unit.
* Ground_troop_slot2
  * Sentinel
    * Heavy infantry using barrier-shields and pulse cannons; strong vs. cavalry.
  * Vanguard
    * Heavily armored guardian powered by enchanted gems; slow but tanky.
* spy (spy troops and buildings in other villages)
  * Seeker Drone 
    * Magical-mechanical drone that scans enemy villages undetected.   
  * Whispershard
    * A floating arcane crystal that records everything it sees.
* Air_troop_slot1
  * Warden
    * Light flying unit, fast and evasive; great for scouting and harassment.   
  * Prism
    * Hovering construct that refracts beams of energy; good harassment unit.
* Air_troop_slot2
  * Void Striker
    * Mid-tier flying unit with strong AoE attacks; weak to anti-air.
  * Celestial Raptor
    * Elegant winged construct with fast dive attacks; fragile but mobile. 
* Air_troop_slot3
  * Celestial Engine
    * Massive air unit; a floating war-machine with devastating long-range plasma bursts.
  * Storm Herald
    * Massive flying battleship that channels lightning storms; high cost, big boom.
* Ram (destroy walls)
  * Phase Ram
    * Dematerializes into walls, destabilizing structures from within.
* Catapult (destroy buildings)
  * Mana Cannon
    * Long-range siege weapon that channels arcane blasts.
  * Arcflare Engine
    * Launches unstable arcane fireballs that explode into chaos energy.
* Hero (1 per village max)
* Noble (lower village loyalty)
  * Ego Warden
    * Ethereal unit that corrupts loyalty using psychic domination.
  * Mindcaller
    * Uses psychic whispers to destabilize loyalty and sow rebellion.
* ? Militia
  * Etherguard
  * Tinkersworn

### Buildings

Arcane Foundries.  

### Heroes

Archmages, War-Tech Engineers.  

## 4.2 Varnak

Brute force and numbers.  

### Troops

Berserkers, War Bewasts, Insects.  

* Ground_troop_slot1
  * Bonecrusher
    * Cheap melee infantry; good swarmers, weak armor.
  * Grub
    * Rapid-breeding insectoid with serrated limbs; individually weak, swarm-capable.
* Ground_troop_slot2
  * Spineguard
    * Medium infantry with bone-plate armor and shields; counters archers and light cavalry.
  * Bonecleaver
    * Warrior clad in scavenged armor and wielding cleavers made from femurs.
* Ground_troop_slot3
  * Dreadbeast Rider
    * Savage mounted unit on mutated war beasts; extremely fast.
  * Ravager
    * Fire-resistant mutant that charges into enemy lines screaming (probably on fire).
* spy (spy troops and buildings in other villages)
  * Shadow Crawler
    * Camouflaged insectoid spy; hard to detect.
  * Skulker
    * Semi-living totem that sneaks into villages and records troop movement by smell.
* Air_troop_slot1
  * Winged Skar
    * Harpy-like creature; light flier for quick strikes.
  * Spinewing
    * Flying vermin that divebomb enemies with acidic spines.
* Air_troop_slot2
  * Carrion Maw
    * Mid-tier flying unit that eats corpses to heal; terrifying presence.
  * Wretch Angel
    * Horrific winged monstrosity that weeps corrosive bile and fear auras.
* Ram (destroy walls)
  * Siegebreaker
    * Giant beast bred to ram and tear down walls.
  * Behemoth
    * Huge insectoid with reinforced plates and rage-fueled strength.
* Catapult (destroy buildings)
  * Flesh Lobber
    * Lobs alchemical flesh sacs that explode on impact. Horrifying but effective.
  * Meat Mortar
    * Lobs gory pods filled with stinging larvae and explosive bile.
* Hero (1 per village max)
* Noble (lower village loyalty)
  * Skull Chanter
    * Uses blood rituals to demoralize and convert enemy villagers.
  * Skinspeaker
    * Wears faces of enemies to demoralize and corrupt defenders from within.
* ? Militia
  * Tribe Kin
  * Fleshspawn

### Buildings

Blood Altars.  

### Heroes

Warlords, Beastmasters.  

## 4.3 Dawnhold

Balanced development.  

### Troops

Swordsmen, Cavalry, Siege Engineers.  

* Ground_troop_slot1
  * Iron Militia
    * Standard infantry, good defense; cost-efficient.
  * Halberdier
    * Versatile foot soldier trained in holding lines and piercing charges.
* Ground_troop_slot2
  * Templar
    * Heavy cavalry, strong against infantry and archers.
  * Knight 
    * Gallant horseman with a penchant for solo glory and dramatic timing.
* Ground_troop_slot3
  * Crossbowmen
    * Ranged infantry; excellent vs. flyers and light troops.
  * Gunner
    * Experimental troop using a crude rapid-fire bolt launcher. Tends to jam.
* spy (spy troops and buildings in other villages)
  * Scoutmaster
    * Skilled recon expert; reveals troop numbers and types.
  * Scribe
    * Masters of infiltration disguised as tax collectors or inspectors.
* Air_troop_slot1
  * Balloon
    * Basic flying unit with thrown spears or bombs; low health.
  * Glider 
    * Primitive hang-glider unit that drops down with lances; hard to control, hilarious when it works.
* Air_troop_slot2
  * Galleon
    * Heavily armored flying warship; slow but deadly.
  * Garrison
    * Flying platform with multiple archers and rotating ballistae.
* Ram (destroy walls)
  * Ironclad Battering Ram
    * Standard siege ram; upgraded through tech research.
  * Siege Mule
    * A stubborn beast pulling a reinforced siege log. Surprisingly effective and impossible to stop once it gets moving.
* Catapult (destroy buildings)
  * Stone Mangonel
    * Traditional catapult; accurate with research.
  * Trebuchet
    * Advanced siege machine built by guild engineers—reliable, accurate, and smug.
* Hero (1 per village max)
* Noble (lower village loyalty)
  * Diplomat
    * Uses cunning and gold to sway villages to your cause.
  * Commissioner
    * Uses law, bribes, and subtle threats to bring villages into the fold.
* ? Militia
  * Village Watch

### Buildings

Town Halls.  

### Heroes

Commanders, Strategists.  

## 4.4 The Forgotten (NPC)

* Weak at the start, so players can farm. Progress with the realm.  
* Mixed unit types with unique mechanics(the unique mechanics things is questionable)
* Strong defenses and elite units
* Key OvE objectives on the map

# 5. World Map and Progression

* Hex-based map layout with player and NPC villages, resource points, and strategic locations.
* Fog of war mechanics until areas are scouted or explored(spy/scout like units).
* Progression: Through research(upgrades), territory acquisition, and hero development(level/items).

# 6. Monetization

* Preminum account with time-saving perks (not pay-to-win)
    * Better account overview
    * Farm list (similar to travian gold club)
    * Mass action utils (like send support from multiple villages at the same time)
    * Account wide, not per realm basis
* Not allowed to have a premium currency
* Players should not be able to buy resources or reduce building time with real money
* No payed loot boxes (loot boxes in game events that don't envolve real money are allowed)

# 7. Technical Requirements

* Svelte frontend
* Golang backend with a scalable database (PostgreSQL, Redis)
* WebSocket support for real-time updates
* Persistent world simulation

# 8. Future Plans

* Seasonal events
* Mobile
* Tournament realms