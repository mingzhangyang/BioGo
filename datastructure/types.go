package datastructure

// Kind of a value
type Kind int

// GetType return the type of the value
func (k Kind) GetType() string {
	switch k {
	case 1:
		return "number"
	case 2:
		return "string"
	case 3:
		return "boolean"
	case 4:
		return "array"
	case 5:
		return "object"
	default:
		return "null"
	}
}
