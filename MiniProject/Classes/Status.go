package Classes

type Status int

const (
	done Status = iota
	active
	unknown = -1
)

func (s Status) String() string {
	switch s {
	case done:
		return "done"
	case active:
		return "active"
	}
	return "unknown"
}

func CreateStatus(sString string) Status {
	switch sString {
	case "active":
		return active
	case "done":
		return done
	case "":
		return active
	}
	return unknown
}
