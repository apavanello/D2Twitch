package main

import (
	"encoding/json"
)

type Heroes map[string]HeroesValue

type HeroesValue struct {
	ID              int64       `json:"id"`
	Name            string      `json:"name"`
	LocalizedName   string      `json:"localized_name"`
	PrimaryAttr     PrimaryAttr `json:"primary_attr"`
	AttackType      AttackType  `json:"attack_type"`
	Roles           []Role      `json:"roles"`
	Img             string      `json:"img"`
	Icon            string      `json:"icon"`
	BaseHealth      int64       `json:"base_health"`
	BaseHealthRegen *float64    `json:"base_health_regen"`
	BaseMana        int64       `json:"base_mana"`
	BaseManaRegen   float64     `json:"base_mana_regen"`
	BaseArmor       float64     `json:"base_armor"`
	BaseMr          int64       `json:"base_mr"`
	BaseAttackMin   int64       `json:"base_attack_min"`
	BaseAttackMax   int64       `json:"base_attack_max"`
	BaseStr         int64       `json:"base_str"`
	BaseAgi         int64       `json:"base_agi"`
	BaseInt         int64       `json:"base_int"`
	StrGain         float64     `json:"str_gain"`
	AgiGain         float64     `json:"agi_gain"`
	IntGain         float64     `json:"int_gain"`
	AttackRange     int64       `json:"attack_range"`
	ProjectileSpeed int64       `json:"projectile_speed"`
	AttackRate      float64     `json:"attack_rate"`
	MoveSpeed       int64       `json:"move_speed"`
	TurnRate        *float64    `json:"turn_rate"`
	CMEnabled       bool        `json:"cm_enabled"`
	Legs            int64       `json:"legs"`
}

type AttackType string

const (
	Melee  AttackType = "Melee"
	Ranged AttackType = "Ranged"
)

type PrimaryAttr string

const (
	Agi PrimaryAttr = "agi"
	Int PrimaryAttr = "int"
	Str PrimaryAttr = "str"
)

type Role string

const (
	Carry     Role = "Carry"
	Disabler  Role = "Disabler"
	Durable   Role = "Durable"
	Escape    Role = "Escape"
	Initiator Role = "Initiator"
	Jungler   Role = "Jungler"
	Nuker     Role = "Nuker"
	Pusher    Role = "Pusher"
	Support   Role = "Support"
)

func (h Heroes) loadConst() {
	f := readFile("data/dotaConst/heroes.json")
	if err := json.Unmarshal(f, &heroes); err != nil {
		processError(err)
	}
}
