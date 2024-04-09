package main

type Enum int
type PaymentType string

const (
	Zero Enum = iota
	One
	Two
	Three
)

const (
	PENDING PaymentType = "PENDING"
	CANCEL  PaymentType = "CANCEL"
	SUCCESS PaymentType = "SUCCESS"
)

func main() {
	var a Enum = Zero
	var b Enum = Three
	var c Enum = 3
	enumSwitch(a)
	enumSwitch(b)
	enumSwitch(c)

	var status = PENDING
	paymentSwitch(status)

}

func enumSwitch(e Enum) {
	switch e {
	case Zero:
		println("Zero")
	case One:
		println("One")
	case Two:
		println("Two")
	case Three:
		println("Three")
	default:
		println("Unknown")
	}

}

func paymentSwitch(p PaymentType) {
	switch p {
	case PENDING, CANCEL:
		println("Not Success")
	case SUCCESS:
		println("SUCCESS")
	default:
		println("Unknown")
	}
}
