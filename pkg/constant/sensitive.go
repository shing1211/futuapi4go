package constant

import "fmt"

type SensitiveString string

func (s SensitiveString) String() string {
	return "[REDACTED]"
}

func (s SensitiveString) GoString() string {
	return "[REDACTED]"
}

func (s SensitiveString) Raw() string {
	return string(s)
}

func (s SensitiveString) IsEmpty() bool {
	return len(s) == 0
}

func (s SensitiveString) Format(f fmt.State, verb rune) {
	switch verb {
	case 's', 'v', 'q', 'x', 'X', 'd', 'b', 'o', 'U':
		f.Write([]byte("[REDACTED]"))
	default:
		f.Write([]byte("[REDACTED]"))
	}
}
