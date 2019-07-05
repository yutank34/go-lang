package tempconv

// 摂氏を華氏に変換
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// 華氏を摂氏に変換
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// 摂氏を絶対温度に変換
func CToK(c Celsius) Kelvin { return Kelvin(c + 273) }

// 絶対温度を摂氏に変換
func KToC(k Kelvin) Celsius { return Celsius(k - 273) }

// 華氏を絶対温度に変換
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// 絶対温度を華氏に変換
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
