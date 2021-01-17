package model

var (
	params = make(map[string]string)
	items  = make(map[string]float32)
)

func UpsertParam(name string, value string) {
	params[name] = value
}

func UpsertItem(name string, value float32) {
	items[name] = value
}

func GetParamValue(name string) (string, error) {
	if params[name] == "" {
		return "", ErrParamNotFound
	}
	return params[name], nil
}

func GetItemValue(name string) (float32, error) {
	if items[name] == 0 {
		return 0, ErrItemNotFound
	}
	return items[name], nil
}

func TransalateParamToRoman(params []string) (string, error) {
	res := ""
	for _, p := range params {
		if value, err := GetParamValue(p); err == nil {
			res += value
		} else {
			return "", ErrParamNotFound
		}
	}
	return res, nil
}

func GetItemValueFromStatement(roman string, total float32) (float32, error) {
	n, err := GetDecimalFromRoman(roman)
	if err != nil {
		return float32(0), err
	}

	return total/float32(n), nil
}
