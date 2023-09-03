package agung

type MicroValue int

const micro = 100

func Add(first int, second int) int {
	return first + second
}

func Sub(first int, second int) int {
	return first - second
}

func Time(first int, second int) int {
	return first * second
}

func Div(first int, second int) int {
	return first / second
}

func Calc(first MicroValue, operator string, second MicroValue) MicroValue {
	var total MicroValue

	switch operator {
	case add:
		return MicroValue(Add(int(first), int(second)))
	case sub:
		return MicroValue(Sub(int(first), int(second)))
	case time:
		return ToMicro(Time(ToMacro(first), ToMacro(second)))
	case div:
		return MicroValue(Div(int(first), ToMacro(second)))
	default:
		return total
	}
}

func DisplayMacro(value MicroValue) float32 {
	return float32(value) / micro
}

func ToMacro(value MicroValue) int {
	return Div(int(value), micro)
}

func ToMicro(value int) MicroValue {
	return MicroValue(Time(value, micro))
}
