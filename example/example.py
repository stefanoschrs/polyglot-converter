class Color:
	Red = 0
	RedKey = "Red"
	Green = 1
	GreenKey = "Green"
	Blue = 2
	BlueKey = "Blue"
	@staticmethod
	def GetKey(value: int) -> str:
		if value == Color.Red: return Color.RedKey
		if value == Color.Green: return Color.GreenKey
		if value == Color.Blue: return Color.BlueKey
		return ""
	@staticmethod
	def GetValue(key: str) -> int:
		if key == Color.RedKey: return Color.Red
		if key == Color.GreenKey: return Color.Green
		if key == Color.BlueKey: return Color.Blue
		return -1

class Elements:
	Wood = "wood"
	WoodKey = "Wood"
	Fire = "fire"
	FireKey = "Fire"
	Earth = "earth"
	EarthKey = "Earth"
	Metal = "metal"
	MetalKey = "Metal"
	Water = "water"
	WaterKey = "Water"
	@staticmethod
	def GetKey(value: str) -> str:
		if value == Elements.Wood: return Elements.WoodKey
		if value == Elements.Fire: return Elements.FireKey
		if value == Elements.Earth: return Elements.EarthKey
		if value == Elements.Metal: return Elements.MetalKey
		if value == Elements.Water: return Elements.WaterKey
		return ""
	@staticmethod
	def GetValue(key: str) -> str:
		if key == Elements.WoodKey: return Elements.Wood
		if key == Elements.FireKey: return Elements.Fire
		if key == Elements.EarthKey: return Elements.Earth
		if key == Elements.MetalKey: return Elements.Metal
		if key == Elements.WaterKey: return Elements.Water
		return ""

