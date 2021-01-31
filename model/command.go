package model

type CommandInput struct {
	Param       string
	Value       float32
	RomanNumber string
	ItemName    string
	Variables   [][]string
}

type CommandOutput struct {
	Value interface{}
}

var commands = map[string]func(Dictionary, CommandInput) (CommandOutput, error){
	"SaveParam":    SaveParamCommand,
	"SaveItem":     SaveItemCommand,
	"GetValue":     GetValueCommad,
	"CompareValue": CompareValueCommand,
}

func CallCommand(d Dictionary, input CommandInput, f func(Dictionary, CommandInput) (CommandOutput, error)) (CommandOutput, error) {
	return f(d, input)
}

func SaveParamCommand(d Dictionary, input CommandInput) (CommandOutput, error) {
	UpsertParam(input.Param, input.RomanNumber)
	return CommandOutput{}, nil
}

func SaveItemCommand(d Dictionary, input CommandInput) (CommandOutput, error) {
	romanNum, err := TransalateParamToRoman(input.Variables[0])
	if err != nil {
		return CommandOutput{}, err
	}

	value, err := GetItemValueFromStatement(romanNum, input.Value)
	if err != nil {
		return CommandOutput{}, err
	}

	UpsertItem(d.ItemType, input.ItemName, value)
	return CommandOutput{}, nil
}

func GetValueCommad(d Dictionary, input CommandInput) (CommandOutput, error) {
	if len(d.ItemType) > 0 {
		return GetItemValueCommand(d, input)
	}

	value, err := GetDecimalFromParam(input.Variables[0])
	if err != nil {
		return CommandOutput{}, err
	}

	return CommandOutput{Value: value}, nil
}

func GetItemValueCommand(d Dictionary, input CommandInput) (CommandOutput, error) {
	value, err := GetValueFromStatement(input.Variables[0][0:len(input.Variables[0])-1],
		d.ItemType,
		input.Variables[0][len(input.Variables[0])-1])
	if err != nil {
		return CommandOutput{}, err
	}

	return CommandOutput{Value: value}, nil
}

func CompareValueCommand(d Dictionary, input CommandInput) (CommandOutput, error) {
	if len(d.ItemType) > 0 {
		return CompareItemValueCommand(d, input)
	}
	var val []int

	for i := 0; i < d.VariableNumber; i++ {
		v, err := GetDecimalFromParam(input.Variables[i])
		if err != nil {
			return CommandOutput{}, err
		}

		val = append(val, v)
	}

	if val[0] > val[1] {
		return CommandOutput{Value: "larger"}, nil
	}

	return CommandOutput{Value: "smaller"}, nil
}

func CompareItemValueCommand(d Dictionary, input CommandInput) (CommandOutput, error) {
	var val []float32
	for i := 0; i < d.VariableNumber; i++ {
		v, err := GetValueFromStatement(input.Variables[i][0:len(input.Variables[i])-1],
			d.ItemType,
			input.Variables[i][len(input.Variables[i])-1])
		if err != nil {
			return CommandOutput{}, err
		}

		val = append(val, v)
	}

	if val[0] > val[1] {
		return CommandOutput{Value: "more"}, nil
	}

	return CommandOutput{Value: "less"}, nil
}
