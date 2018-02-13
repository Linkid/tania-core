package storage

import (
	"fmt"
	"time"

	"github.com/Tanibox/tania-server/src/assets/domain"
	deadlock "github.com/sasha-s/go-deadlock"
	uuid "github.com/satori/go.uuid"
)

type FarmStorage struct {
	Lock    *deadlock.RWMutex
	FarmMap map[uuid.UUID]domain.Farm
}

func CreateFarmStorage() *FarmStorage {
	rwMutex := deadlock.RWMutex{}
	deadlock.Opts.DeadlockTimeout = time.Second * 10
	deadlock.Opts.OnPotentialDeadlock = func() {
		fmt.Println("FARM STORAGE DEADLOCK!")
	}

	return &FarmStorage{FarmMap: make(map[uuid.UUID]domain.Farm), Lock: &rwMutex}
}

type FarmEventStorage struct {
	Lock       *deadlock.RWMutex
	FarmEvents []FarmEvent
}

type FarmEvent struct {
	FarmUID uuid.UUID
	Version int
	Event   interface{}
}

func CreateFarmEventStorage() *FarmEventStorage {
	rwMutex := deadlock.RWMutex{}
	deadlock.Opts.DeadlockTimeout = time.Second * 10
	deadlock.Opts.OnPotentialDeadlock = func() {
		fmt.Println("FARM EVENT STORAGE DEADLOCK!")
	}

	return &FarmEventStorage{Lock: &rwMutex}
}

type FarmRead struct {
	UID         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	Type        string    `json:"type"`
	CountryCode string    `json:"country_code"`
	CityCode    string    `json:"city_code"`
	IsActive    bool      `json:"is_active"`
}

type FarmReadStorage struct {
	Lock        *deadlock.RWMutex
	FarmReadMap map[uuid.UUID]FarmRead
}

func CreateFarmReadStorage() *FarmReadStorage {
	rwMutex := deadlock.RWMutex{}
	deadlock.Opts.DeadlockTimeout = time.Second * 10
	deadlock.Opts.OnPotentialDeadlock = func() {
		fmt.Println("FARM READ STORAGE DEADLOCK!")
	}

	return &FarmReadStorage{FarmReadMap: make(map[uuid.UUID]FarmRead), Lock: &rwMutex}
}

type AreaStorage struct {
	Lock    *deadlock.RWMutex
	AreaMap map[uuid.UUID]domain.Area
}

func CreateAreaStorage() *AreaStorage {
	rwMutex := deadlock.RWMutex{}
	deadlock.Opts.DeadlockTimeout = time.Second * 10
	deadlock.Opts.OnPotentialDeadlock = func() {
		fmt.Println("AREA STORAGE DEADLOCK!")
	}

	return &AreaStorage{AreaMap: make(map[uuid.UUID]domain.Area), Lock: &rwMutex}
}

type ReservoirStorage struct {
	Lock         *deadlock.RWMutex
	ReservoirMap map[uuid.UUID]domain.Reservoir
}

func CreateReservoirStorage() *ReservoirStorage {
	rwMutex := deadlock.RWMutex{}
	deadlock.Opts.DeadlockTimeout = time.Second * 10
	deadlock.Opts.OnPotentialDeadlock = func() {
		fmt.Println("RESERVOIR STORAGE DEADLOCK!")
	}

	return &ReservoirStorage{ReservoirMap: make(map[uuid.UUID]domain.Reservoir), Lock: &rwMutex}
}

type MaterialStorage struct {
	Lock        *deadlock.RWMutex
	MaterialMap map[uuid.UUID]domain.Material
}

func CreateMaterialStorage() *MaterialStorage {
	rwMutex := deadlock.RWMutex{}
	deadlock.Opts.DeadlockTimeout = time.Second * 10
	deadlock.Opts.OnPotentialDeadlock = func() {
		fmt.Println("MATERIAL STORAGE DEADLOCK!")
	}

	return &MaterialStorage{MaterialMap: make(map[uuid.UUID]domain.Material), Lock: &rwMutex}
}
