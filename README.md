# Currency-Converter

Command-line currency converter app written in Go.

How to Use
Clone the repository:
```
git clone https://github.com/Adam0Brien/currency-converter.git
cd currency-converter
```
Build the executable:
```
go build
```
Run the application:
```
./currency-converter
```
Follow the prompts to enter the 3-letter currency symbols and the amount you want to convert.

Example:

```
Enter the 3 letter currency symbol you wish to exchange: USD
Enter the 3 letter currency symbol you wish to receive: EUR
Enter the amount: 100
```

# Dependencies
This app relies on the API Ninjas Currency Conversion API for currency conversion data.

The currency symbols should be valid 3-letter codes. Check the currencies.csv file for a list of supported currencies.


