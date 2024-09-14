package order

// Status - Custom type to hold value for order status
type Status int

const (
	StatusNew        Status = iota + 1 // EnumIndex = 1
	StatusProcessing                   // EnumIndex = 2
	StatusInvalid                      // EnumIndex = 3
	StatusProcessed                    // EnumIndex = 4
)

// String - Status string representation
func (s Status) String() string {
	return [...]string{"NEW", "PROCESSING", "INVALID", "PROCESSED"}[s-1]
}

// EnumIndex - Status int representation
func (s Status) EnumIndex() int {
	return int(s)
}
