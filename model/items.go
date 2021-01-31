package model

var (
	params = make(map[string]string)
	items  = map[string]map[string]float32{}
)

func UpsertParam(name string, value string) {
	params[name] = value
}

func UpsertItem(itemType string, name string, value float32) {
	if items[itemType] == nil {
		items[itemType] = make(map[string]float32)
	}
	items[itemType][name] = value
}

func GetParamValue(name string) (string, error) {
	if params[name] == "" {
		return "", ErrParamNotFound
	}
	return params[name], nil
}

func GetItemValue(itemType string, name string) (float32, error) {
	if items[itemType] == nil {
		return 0, ErrItemNotFound
	}

	if items[itemType][name] == 0 {
		return 0, ErrItemNotFound
	}

	return items[itemType][name], nil
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

	return total / float32(n), nil
}

func GetDecimalFromParam(params []string) (int, error) {
	romanNum, err := TransalateParamToRoman(params)
	if err != nil {
		return 0, err
	}

	value, err := GetDecimalFromRoman(romanNum)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func GetValueFromStatement(params []string, itemType string, item string) (float32, error) {
	paramVal, err := GetDecimalFromParam(params)
	if err != nil {
		return float32(0), err
	}

	itemVal, err := GetItemValue(itemType, item)
	if err != nil {
		return float32(0), err
	}

	return float32(paramVal) * itemVal, nil
}
