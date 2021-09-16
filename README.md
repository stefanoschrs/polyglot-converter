# Polyglot Converter

> Motivation:  
> In the example scenario where you have a frontend (**TypeScript**) and 2 backend services (**Go**, **Python**), 
> you want to use the same enums throughout, 
> for **consistency**, **less errors** and letting your IDE help you with **autocomplete**.  

## Example
[Full example here](./example)

- Input
```yaml
Color:
  Red: 0
```

- Output
```go
// Go
type Color struct {}
func (a Color) Red() int { return 0 }
func (a Color) RedKey() string { return "Red" }
func (a Color) GetKey(value int) string {
    if value == a.Red() { return a.RedKey() }
    return ""
}
func (a Color) GetValue(key string) int {
    if key == a.RedKey() { return a.Red() }
    return -1
}
```
```typescript
// TypeScript
class Color {
	static Red = 0
	static RedKey = "Red"
	static GetKey (value: number): string {
		if (value === this.Red) { return this.RedKey }
		return ""
	}
	static GetValue (key: string): number {
		if (key === this.RedKey) { return this.Red }
		return -1
	}
}
```
```python3
# Python
class Color:
	Red = 0
	RedKey = "Red"
	@staticmethod
	def GetKey(value: int) -> str:
		if value == Color.Red: return Color.RedKey
		return ""
	@staticmethod
	def GetValue(key: str) -> int:
		if key == Color.RedKey: return Color.Red
		return -1
```

## Features
- [x] Enum String
- [x] Enum Number
- [ ] Switch args to flags [cobra](https://github.com/spf13/cobra)
- [ ] Package option for Go (default main)

## Language support
### Go
- Using `type` and `methods`
- [Playground](https://play.golang.org/)

### TypeScript
- Using `class`
- [Playground](https://www.typescriptlang.org/play?#code/Q)

### Python 3
- Using `class`
- [Playground](https://trinket.io/python3)
