class Color {
	static Red = 0
	static RedKey = "Red"
	static Green = 1
	static GreenKey = "Green"
	static Blue = 2
	static BlueKey = "Blue"
	static GetKey (value: number): string {
		if (value === this.Red) { return this.RedKey }
		if (value === this.Green) { return this.GreenKey }
		if (value === this.Blue) { return this.BlueKey }
		return ""
	}
	static GetValue (key: string): number {
		if (key === this.RedKey) { return this.Red }
		if (key === this.GreenKey) { return this.Green }
		if (key === this.BlueKey) { return this.Blue }
		return -1
	}
}

class Elements {
	static Fire = "fire"
	static FireKey = "Fire"
	static Earth = "earth"
	static EarthKey = "Earth"
	static Metal = "metal"
	static MetalKey = "Metal"
	static Water = "water"
	static WaterKey = "Water"
	static Wood = "wood"
	static WoodKey = "Wood"
	static GetKey (value: string): string {
		if (value === this.Fire) { return this.FireKey }
		if (value === this.Earth) { return this.EarthKey }
		if (value === this.Metal) { return this.MetalKey }
		if (value === this.Water) { return this.WaterKey }
		if (value === this.Wood) { return this.WoodKey }
		return ""
	}
	static GetValue (key: string): string {
		if (key === this.FireKey) { return this.Fire }
		if (key === this.EarthKey) { return this.Earth }
		if (key === this.MetalKey) { return this.Metal }
		if (key === this.WaterKey) { return this.Water }
		if (key === this.WoodKey) { return this.Wood }
		return ""
	}
}

