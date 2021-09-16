# Full Example
## YAML
```yml
Color:
  Red: 0
  Green: 1
  Blue: 2

Elements:
  Wood: "wood"
  Fire: "fire"
  Earth: "earth"
  Metal: "metal"
  Water: "water"
```
## Go
```go
package main
type Color struct {}
func (a Color) Red() int { return 0 }
func (a Color) RedKey() string { return "Red" }
func (a Color) Green() int { return 1 }
func (a Color) GreenKey() string { return "Green" }
func (a Color) Blue() int { return 2 }
func (a Color) BlueKey() string { return "Blue" }
func (a Color) GetKey(value int) string {
	if value == a.Red() { return a.RedKey() }
	if value == a.Green() { return a.GreenKey() }
	if value == a.Blue() { return a.BlueKey() }
	return ""
}
func (a Color) GetValue(key string) int {
	if key == a.RedKey() { return a.Red() }
	if key == a.GreenKey() { return a.Green() }
	if key == a.BlueKey() { return a.Blue() }
	return -1
}

type Elements struct {}
func (a Elements) Water() string { return "water" }
func (a Elements) WaterKey() string { return "Water" }
func (a Elements) Wood() string { return "wood" }
func (a Elements) WoodKey() string { return "Wood" }
func (a Elements) Fire() string { return "fire" }
func (a Elements) FireKey() string { return "Fire" }
func (a Elements) Earth() string { return "earth" }
func (a Elements) EarthKey() string { return "Earth" }
func (a Elements) Metal() string { return "metal" }
func (a Elements) MetalKey() string { return "Metal" }
func (a Elements) GetKey(value string) string {
	if value == a.Water() { return a.WaterKey() }
	if value == a.Wood() { return a.WoodKey() }
	if value == a.Fire() { return a.FireKey() }
	if value == a.Earth() { return a.EarthKey() }
	if value == a.Metal() { return a.MetalKey() }
	return ""
}
func (a Elements) GetValue(key string) string {
	if key == a.WaterKey() { return a.Water() }
	if key == a.WoodKey() { return a.Wood() }
	if key == a.FireKey() { return a.Fire() }
	if key == a.EarthKey() { return a.Earth() }
	if key == a.MetalKey() { return a.Metal() }
	return ""
}

```
## TypeScript
```typescript
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

```
## Python
```python3
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

```
