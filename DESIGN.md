# RTS Game Design Reference

## UI Structure

### Dashboard Layout

```
┌─────────────────────────────────────────────────────────────┐
│  Tabs: [Systems] [Colonies] [Ships] [Orders]                │
├──────────────────────────┬──────────────────────────────────┤
│                          │                                  │
│   Main View (60%)        │   Side Panel (40%)               │
│   - Planet List          │   - Selected Item Details        │
│   - System Map           │   - Quick Actions                │
│   - Fleet Overview       │   - Resource Summary             │
│                          │                                  │
├──────────────────────────┴──────────────────────────────────┤
│  Status Bar: Tick 1234 | Resources: 🌾 500 ⛏ 300 ⚡ 200     │
└─────────────────────────────────────────────────────────────┘
```

### Pane Navigation

- **Tab navigation**: Arrow keys / h,l to switch tabs
- **Enter**: Focus into selected pane
- **Esc**: Return to tab navigation
- **j/k**: Navigate within lists
- Panes can spawn child panes (modal-like behavior via focus stack)

### Pane Types

| Pane | Purpose |
|------|---------|
| SystemList | Browse star systems with search |
| SystemInfo | View planets in selected system |
| PlanetList | List all planets |
| PlanetInfo | Planet details, resources, actions |
| ColonyManager | Manage existing colonies |
| ShipManagement | Fleet overview and commands |
| OrderStatus | Pending/executing/completed orders |
| Dashboard | Grid container for multiple panes |

---

## Core Game Loop

### Player Action Flow

```
┌─────────────────────────────────────────────────────────────┐
│                    PLAYER ACTIONS                           │
│  Colonize → Build → Explore → Expand → Defend               │
└─────────────────────────────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                    ORDER QUEUE                              │
│  Orders have: Duration, Priority, Dependencies              │
└─────────────────────────────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                    TICK SYSTEMS                             │
│  1. Process Orders (pending → executing → complete)         │
│  2. Production (farms/mines/grids add resources)            │
│  3. Consumption (population eats food)                      │
│  4. Stability (unrest, happiness, corruption)               │
│  5. Events (threats spawn, discoveries, etc.)               │
└─────────────────────────────────────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────────┐
│                    STATE UPDATE                             │
│  Resources, Population, Ship Positions, Threats             │
└─────────────────────────────────────────────────────────────┘
```

### Order Status Flow

```
Pending → Executing → Complete
                   ↘ Failed
```

### Tick Rate

- 8 ticks per second (125ms per tick)
- Each tick processes: Orders → Actions → Constructions → Stability → Population

---

## Implementation Checklist

### Core Systems

- [x] Star system generation (procedural)
- [x] Planet/Ship/Colony entities
- [x] Resource system (Food, Minerals, Energy)
- [x] Production system (Farms, Mines, Solar Grids)
- [x] Order/Action queue with priority scheduling
- [x] Game tick loop (8 ticks/sec)
- [x] Population growth (based on food surplus)
- [ ] Resource scarcity consequences (starvation, power outages)
- [ ] Threat/Event system (attacks, discoveries, disasters)
- [ ] Win/Lose conditions
- [ ] Save/Load game state

### UI Panes

- [x] RootPane (sidebar navigation)
- [x] SystemList (with search)
- [x] SystemInfo (planets in system)
- [x] PlanetList
- [x] PlanetInfo (details + actions)
- [x] OrderStatus (with progress bars)
- [x] CreateColonyPane (form)
- [x] ShipManagement
- [x] Dashboard (grid layout container)
- [ ] ColonyManager (post-creation management)
- [ ] ResourceAllocation (sliders for distribution)
- [ ] ThreatAlert (incoming dangers)
- [ ] StatusBar (live resource counts)

### Gameplay Features

- [x] Colonize planets
- [x] Send scout ships
- [x] Build structures (Farm, Mine, SolarGrid)
- [ ] Upgrade structures to higher tiers
- [ ] Ship combat system
- [ ] Research/Tech tree
- [ ] Inter-planet trading
- [ ] Colony specialization
- [ ] Population migration

### Polish

- [ ] Sound/Audio feedback
- [ ] Notifications for completed orders
- [ ] Keyboard shortcut help overlay
- [ ] Tutorial/Onboarding
- [ ] Difficulty settings

---

## Data Models Reference

### Planet Resources

```go
Resources struct {
    Food     Resource  // Production from Farms
    Minerals Resource  // Production from Mines
    Energy   Resource  // Production from SolarGrids
}

Resource struct {
    Quantity        float64
    ConsumptionRate float64  // Per capita
}
```

### Planet Stability

```go
Stabilities struct {
    Corruption Stability  // Reduces efficiency
    Happiness  Stability  // Affects growth
    Unrest     Stability  // Can cause rebellion
}
```

### Construction Tiers

| Tier | Production |
|------|------------|
| 1    | 5/tick     |
| 2    | 10/tick    |
| 3    | 20/tick    |

---

## Event Ideas

| Event | Effect | Response |
|-------|--------|----------|
| Pirate Raid | Steals resources | Send fighters |
| Solar Flare | Disables energy | Wait it out |
| Resource Discovery | Bonus minerals | Claim with ship |
| Plague | Kills population | Build hospitals |
| Rebellion | Colony goes hostile | Military or diplomacy |
| Asteroid | Destroys structures | Intercept with ship |

---

## Priority Order for Implementation

1. **Population growth** - Makes colonies feel alive
2. **Resource scarcity** - Creates meaningful choices
3. **Random events** - Adds unpredictability
4. **Status bar** - Better player awareness
5. **Dashboard layout** - Better information density
6. **Win condition** - Gives the player a goal
