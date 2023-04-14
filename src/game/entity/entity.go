package entity

type Entity struct {
	Type entityType
}

type entityType string

const PlayerEntity entityType = "player"
const NPCEntity entityType = "npc"
const MobEntity entityType = "mob"
