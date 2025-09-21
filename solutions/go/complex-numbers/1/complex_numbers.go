package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	realPart      float64
	imaginaryPart float64
}

func (n Number) Real() float64 {
	return n.realPart
}

func (n Number) Imaginary() float64 {
	return n.imaginaryPart
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		realPart:      n1.realPart + n2.realPart,
		imaginaryPart: n1.imaginaryPart + n2.imaginaryPart,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		realPart:      n1.realPart - n2.realPart,
		imaginaryPart: n1.imaginaryPart - n2.imaginaryPart,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		realPart:      n1.realPart*n2.realPart - n1.imaginaryPart*n2.imaginaryPart,
		imaginaryPart: n1.imaginaryPart*n2.realPart + n1.realPart*n2.imaginaryPart,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		realPart:      n.realPart * factor,
		imaginaryPart: n.imaginaryPart * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
		realPart: (n1.realPart*n2.realPart + n1.imaginaryPart*n2.imaginaryPart) /
			(math.Pow(n2.realPart, 2) + math.Pow(n2.imaginaryPart, 2)),
		imaginaryPart: (n1.imaginaryPart*n2.realPart - n1.realPart*n2.imaginaryPart) /
			(math.Pow(n2.realPart, 2) + math.Pow(n2.imaginaryPart, 2)),
	}
}

func (n Number) Conjugate() Number {
	return Number{
		realPart:      n.realPart,
		imaginaryPart: -n.imaginaryPart,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.realPart, 2) + math.Pow(n.imaginaryPart, 2))
}

func (n Number) Exp() Number {
	return Number{
		realPart:      math.Exp(n.realPart) * math.Cos(n.imaginaryPart),
		imaginaryPart: math.Exp(n.realPart) * math.Sin(n.imaginaryPart),
	}
}
