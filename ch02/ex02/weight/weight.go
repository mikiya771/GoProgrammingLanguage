package weight

type Gram float64
type Pond float64

func Units() map[string]string {
	ret := map[string]string{}
	ret["Gram"] = "g"
	ret["Pond"] = "lb"
	return ret
}
func GtoP(g Gram) Pond {
	return Pond(g * 2.20462)
}
func PtoG(p Pond) Gram {
	return Gram(p / 2.20462)
}
