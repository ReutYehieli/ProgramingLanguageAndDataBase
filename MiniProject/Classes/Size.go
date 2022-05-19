package Classes

type Size int

const (
	small Size = iota
	medium
	large
	unknownSize = -1
)

func CreateSize(size string) Size {
	switch size {
	case "small":
		return small
	case "medium":
		return medium
	case "large":
		return large
	}
	return unknownSize
}
