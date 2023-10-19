# Constructors

If a struct is exported (starts with an uppercase letter), any other package can create an instance of it.

```go
twentyDollars := money.Money{
	Amount:   20, 
	Currency: "USD",
}
```

The cost of this freedom is that it's possible to create an incomplete or invalid structures.

```go
// Money with no currency
twenty := money.Money{
	Amount: 20,
}
```

You can introduce more safety by making the struct's fields unexported and creating a dedicated *constructor*.
A constructor is a function that returns an initialized struct.

```go
type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) Money {
	return Money{
		amount:   amount, 
		currency: currency,
	}
}
```

Now, whoever is using your package, needs to use the constructor to create a new struct instance.

```go
twentyDollars := money.NewMoney(20, "USD")
```

## Exercise

File: `11-structs-2/04-constructors/main.go`

The `DateRange` struct represents a date range, with a start and an end. The `Hours()` method returns the number of hours between the dates.

It's now possible to create a `DateRange` with empty dates or with the end date before the start date.
Using such `DateRange` can lead to bugs in our application.

* Add a `NewDateRange` constructor that returns a new `DateRange` and an `error`.
* Return an error if the provided start date or end date is empty. You can use the `IsZero()` method from [`time.Time`](https://pkg.go.dev/time#Time).
* Return an error if the provided end date happens before the start date. You can use the `Before()` method from [`time.Time`](https://pkg.go.dev/time#Time).
* Rework the `main` function to use the constructor.