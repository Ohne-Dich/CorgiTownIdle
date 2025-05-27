package main

import (
	"fmt"
	"math"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func commandBuild(m *model, args ...string) (tea.Cmd, error) {
	if len(args) < 1 {
		m.log = append(m.log, "Was möchtest du bauen?")
		return nil, nil
	}
	target := strings.ToLower(args[0])

	meta, ok := buildingDefs[target]
	if !ok {
		m.log = append(m.log, fmt.Sprintf("Unbekanntes Gebäude: %s", target))
		return nil, nil
	}

	// Anzahl bisher gebauter Gebäude
	var count int
	switch target {
	case "house":
		count = m.build.Houses
		// evtl. mehr Gebäude
	}

	// Kosten berechnen
	cost := applyCost(meta.BaseCost, meta.Multiplier, count)

	// Ressourcen prüfen
	if m.res.Wood < cost.Wood || m.res.Stone < cost.Stone || m.res.Gold < cost.Gold {
		m.log = append(m.log, fmt.Sprintf(
			"Nicht genug Ressourcen für %s (%d Holz, %d Stein, %d Gold benötigt)",
			meta.Name, cost.Wood, cost.Stone, cost.Gold,
		))
		return nil, nil
	}

	// Ressourcen abziehen
	m.res.Wood -= cost.Wood
	m.res.Stone -= cost.Stone
	m.res.Gold -= cost.Gold

	// Effekt anwenden (z. B. Gebäude hochzählen)
	meta.Effect(m)

	// Logeintrag
	m.log = append(m.log, fmt.Sprintf(
		"%s gebaut! -%d Holz, -%d Stein, -%d Gold",
		meta.Name, cost.Wood, cost.Stone, cost.Gold,
	))

	return nil, nil
}

func applyCost(base ResourceSet, multiplier float64, count int) ResourceSet {
	scale := math.Pow(multiplier, float64(count))
	return ResourceSet{
		Wood:  int(math.Ceil(float64(base.Wood) * scale)),
		Stone: int(math.Ceil(float64(base.Stone) * scale)),
		Gold:  int(math.Ceil(float64(base.Gold) * scale)),
	}
}
