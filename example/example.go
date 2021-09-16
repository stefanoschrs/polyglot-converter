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

