# Roman Numbers

### Instalation :
* Read Go installation steps from [here](golang.org/doc/install).

* Clone the repository using command 
```sh
go get github.com/ajengpu/roman-numbers
```

* This project are using [GoModules](https://github.com/golang/go/wiki/Modules) , hit command bellow to get all dependenciess and create `./vendor` dir.
```sh
go mod vendor
```

### Run Project
* Go to the repository directory 

* Type command below to build the project
```sh
go build
```

* Type command below to run the project
```sh
roman-numbers
```

### Usage
* To declare the parameter type 
```sh
[parameter_name] is [roman_number]
```
then press enter.
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
then press enter.
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
then press enter.
parameter_name1, parameter_name1... : the name of parameters represent roman number, the number of parameter is not limited as long as represent valid roman numeric

Example :
```sh
how much is blob pesh ?
```

* To get the decimal value of Credits
```sh
how many Credits is [parameter_name1] [parameter_name2] [item_name] ?
```
then press enter.
parameter_name1, parameter_name1... : the name of parameters represent roman number, the number of parameter is not limited as long as represent valid roman numeric
item_name   : the name of the declared item which has Credits value

Example :
```sh
how many Credits is blob pesh Silver ?
```

* To get the result press `Enter` or `return` without typing anything