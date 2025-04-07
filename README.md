# Forgotten Realms

## Project Overview

**Game Title**: Forgotten Realms  
**Genre**:  Browser-Based MMO-RTS  
**Platform**: Web Browser  
**Target Audience**: Strategy gamers, fans of persistent world games, ages 13+   
**Inspirations**: Tribalwars, Travian, Starcraft, Warcraft  

Forgotten Realms is an immersive MMO-RTS game set in a newly discovered world filled with ancient secrets and contested territories. Players choose between three unique factions, each with its own strengths, units, and heroes, and compete for dominance in a realm full of both PvE and PvP content.

## Key Features

* **Dynamic Factions**: Players can choose between Caldari, Varnak, and Dawnhold, each with distinct playstyles, units, and heroes.
* **PvE and PvP Gameplay**: Battle NPC-controlled factions, conquer villages, engage in large-scale sieges, and form alliances with other players.
* **Hero System**: Each faction has unique heroes with abilities that can influence battles and help to level up troops.
* **Strategic Combat**: Utilize ground and flying units in combat. Develop strategies and counter enemy forces.
* **World Map**: A hex-based layout with strategic resource management.

## Technical Details

### Frontend
* **Framework**: Svelte
* **Real-Time Updates**: WebSockets for real-time communication (in-game events, battles, etc.)
* **Responsive Design**: Optimized for desktop browsers, with plans for a mobile version in the future.

### Backend
* **Language**: Go (Golang)
* **Database**: Postgres (for persistence), Redis (for caching and session management)
* **Real-Time**: WebSockets for live player interactions, battle outcomes, and world updates.

### Server Architecture
* **Persistent World Simulation**: Backend simulates the game world, including player actions, world progression, and NPC behavior.

## Installation and Setup

### Prerequisites

* [Git](https://git-scm.com/)
* [Bash](https://www.gnu.org/software/bash/)
* [Docker Compose](https://docs.docker.com/compose/)

### Run

1. **Clone the repository:**
    ```bash
    git clone https://github.com/jtomaspm/Forgotten-Realms.git 
    cd Forgotten-Realms
    ```

2. **Setup .env**

Check the [test workflow](https://github.com/jtomaspm/Forgotten-Realms/blob/main/.github/workflows/docker-image.yml) to see how to create an environment(.env) file.

3. **Run**
    * For a persintent run
    ```bash
    ./scripts/run.sh <env>
    ```
    * For a clean run
    ```bash
    ./scripts/run.sh <env> clean
    ```

## Documentation

* [Game Design Document](https://github.com/jtomaspm/Forgotten-Realms/blob/main/documentation/game_design/game_design.md)
* [Workflows](https://github.com/jtomaspm/SimplifiedCrafter/blob/main/documentation/workflows.md)
* [Infrastructure](https://github.com/jtomaspm/SimplifiedCrafter/blob/main/documentation/infrastructure.md)
* [Scripts](https://github.com/jtomaspm/SimplifiedCrafter/blob/main/documentation/scripts.md)