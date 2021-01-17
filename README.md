# Roman Numbers

### Instalation :
* Read Go installation steps from [here](golang.org/doc/install).

* Clone the repository using command 
```sh
go get github.com/ajengpu/roman-numbers
```

Or go to your project directory and clone from git using command 
```sh
git clone https://github.com/ajengpu/juno.git
```

### Run Project
* Go to the repository directory 

* Type command below to run the project
```sh
go build && ./roman-numbers
```

### Usage
* To declare the parameter type 
```sh
[parameter_name] is [roman_number]
```
then press `Enter` or `return`.
parameter_name : the name of the parameter to save the roman number
roman_number   : the roman number which will be saved in paramater

Example :
```sh
blob is I
```

* To declare item type
```sh
[parameter_name1] [parameter_name2] [item_name] is [value] Credits
```
then press `Enter` or `return`.
parameter_name1, parameter_name1... : the name of parameters represent roman number, the number of parameter is not limited as long as represent valid roman numeric
item_name   : the name of the declared item which has Credits value
value       : the value of parameters values * item value

Example :
```sh
blob pesh Silver is 34000 Credits
```

* To get the decimal value of roman number represented by paramaters type
```sh
how much is [parameter_name1] [parameter_name2] ?
```
then press `Enter` or `return`.
parameter_name1, parameter_name1... : the name of parameters represent roman number, the number of parameter is not limited as long as represent valid roman numeric

Example :
```sh
how much is blob pesh ?
```

* To get the decimal value of Credits
```sh
how many Credits is [parameter_name1] [parameter_name2] [item_name] ?
```
then press `Enter` or `return`.
parameter_name1, parameter_name1... : the name of parameters represent roman number, the number of parameter is not limited as long as represent valid roman numeric
item_name   : the name of the declared item which has Credits value

Example :
```sh
how many Credits is blob pesh Silver ?
```

* To get the result press `Enter` or `return` without typing anything

### Run Test
* Go to repository directory `/test/test-case` 

* Open the `test-data.json` 
Modify `test_data` array to add or remove test case
Modify `input` array to add or remove inputs of a test case
Modify `output` array to add or remove the EXPECTED outputs of a test case

* Go to the repository directory 

* Type command below to run the project
```sh
cd test && go build && ./test [json_path]
```
json_path : the path to test data json which will be used to test the program

Example :
```sh
cd test && go build && ./test test-case/test-data.json
```
