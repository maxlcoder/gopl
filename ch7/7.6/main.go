package __6

// func (f *celsiusFlag) Set(s string) error {
// 	var unit string
// 	var value float64
//
// 	fmt.Sscanf(s, "%f%s", &value, &unit) // 无须检查错误
// 	switch unit {
// 	case "C", "ºC":
// 		f.Celsius = Celsius(value)
// 		return nil
// 	case "F", "ºF":
// 		f.Celsius = FToC(Fahrenheit(value))
// 		return nil
// 	case "K", "ºK":
// 		f.Celsius = KToC(Kelvin(value))
// 		return nil
//
//
// 	}
// 	return fmt.Errorf("invalid temperature %q", s)
//
// }