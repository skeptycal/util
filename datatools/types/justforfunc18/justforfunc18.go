package justforfunc18

import "fmt"

const (
	cf float64 = 9.0 / 5.0 // (100°C × 9/5) + 32 = 212°F
	fc float64 = 5.0 / 9.0 // (212°F − 32) × 5/9 = 100°C
)

type celsius float64

func (c celsius) String() string    { return fmt.Sprintf("%.2f °C", c) }
func (c celsius) Value() float64    { return float64(c) }
func (c celsius) Farenheit() string { return fmt.Sprintf("%.2f °F", c.Value()*cf+32) }
func NewCelsius(f float64) celsius  { return celsius((f - 32.0) * fc) }

type temperature struct{ celsius }

type anotherTemp temperature

//Temperature Tests commented out ...

func TempTests() {
	// based on justforfunc #18

	c := celsius(10.0)
	fmt.Println(c)
	fmt.Println(c.Farenheit())
	// Output:
	// 10.00 °C
	// 50.00 °F

	j := NewCelsius(95 + 12) // 107 °F
	fmt.Println(j)
	fmt.Println(j.Farenheit())
	// Output:
	// 41.67 °C
	// 107.00 °F

	// c2 := temperature(c)
	// error: cannot convert c (variable of type celsius) to temperaturecompiler

	t := temperature{c}
	fmt.Println(t)
	a := temperature{c + 5}
	fmt.Println(a)
	// Output:
	// 10.00 °C
	// 15.00 °C

	// var a anotherTemp = t
	// error: cannot use t (variable of type temperature) as anotherTemp value in variable declarationcompiler

	var b anotherTemp = anotherTemp(t)
	fmt.Println(b)
	// Output:
	// 10.00 °C

	// cannot convert 5 (untyped int constant) to temperaturecompiler
	// var d anotherTemp = anotherTemp(t + 5)

	// expected operand, found '{'syntax
	// var c anotherTemp = {t + 5}

	// ! passing an anonymous struct is an important detail
	var e anotherTemp = struct {
		celsius
	}{c + 12}
	fmt.Println(e)
	// Output:
	// 22.00 °C

	// cannot use t (variable of type temperature) as celsius value in struct literalcompiler
	// var f anotherTemp = struct {
	// 	celsius
	// }{t}

	var f = celsius(33)
	fmt.Println(f)
	// Output:
	// 33.00 °C
}

// const temptests = 1
