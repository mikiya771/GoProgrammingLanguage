package length

type Meter float64
type Inch float64

func Units() map[string]string {
	ret := map[string]string{}
	ret["Meter"] = "m"
	ret["Inch"] = "in"
	return ret
}

func MtoI(m Meter) Inch {
	return Inch(m * 100 / 2.55)
}
func ItoM(i Inch) Meter {
	return Meter(i * 2.54 / 100)
}
