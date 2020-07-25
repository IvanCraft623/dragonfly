package block

import (
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/world"
)

// GrassPath is a decorative block that can be created by using a shovel on a grass block.
type GrassPath struct {
	noNBT

	transparent
}

// NeighbourUpdateTick handles the turning from grass path into dirt if a block is placed on top of it.
func (p GrassPath) NeighbourUpdateTick(pos, _ world.BlockPos, w *world.World) {
	if _, air := w.Block(pos.Add(world.BlockPos{0, 1})).(Air); !air {
		// Technically vanilla doesn't always turn grass paths into dirt when a block is placed above it,
		// for example torches, but the logic doesn't make sense.
		w.SetBlock(pos, Dirt{})
	}
}

// BreakInfo ...
func (p GrassPath) BreakInfo() BreakInfo {
	return BreakInfo{
		Hardness:    0.6,
		Harvestable: alwaysHarvestable,
		Effective:   shovelEffective,
		Drops:       simpleDrops(item.NewStack(Dirt{}, 1)),
	}
}

// EncodeItem ...
func (p GrassPath) EncodeItem() (id int32, meta int16) {
	return 198, 0
}

// EncodeBlock ...
func (p GrassPath) EncodeBlock() (name string, properties map[string]interface{}) {
	return "minecraft:grass_path", nil
}

// Hash ...
func (p GrassPath) Hash() uint64 {
	return hashGrassPath
}
