# Ledger File Format

Maintaining your Transaction Record in `ledger` format.

Tracking your transactions for analysis with `ledger` is as easy as writing some
text to a file in a very human-readable format.
The format is _structured_ but appears _unstructured_ to many because it doesn't
use curly brackets, key-value pairs, or other special characters to model
transaction data.
Instead, the things that matter are just having enough whitespace between
certain elements in order for the `ledger` parser to understand the difference
between dates, amounts, and so on.

Start your favorite text editor and you'll get started on the path to personal
finance greatness.

## Terminology

* Transaction - Series of consecutive lines that represent the move of money
from one account to one (*or more*) other accounts.
* Transaction Date - Date the transaction occured.
* Payee - Description following on the same line as the Date. Usually the place
of business or person the transaction occured at/with.
* Posting - Line containing account and (*optionally*) amount.

## Basic transaction format

The basic format of a `ledger` transaction, shown below.

```ledger
2017-06-26 Commonplace Coffee
  Assets:Cash:Wallet           -3.00
  Expenses:Restaurants:Coffee   3.00

```

In the example, line 1 shows the _transaction date_ and _payee_.
Lines 2 and 3 are two _postings_ comprised of an _account_ and an _amount_.

All transactions must balance. That is, the amount credited must
equal the amount debited: credits minus debits must equal zero.
In other words, the sum of all _postings_ must equal zero.

A transaction must have at least two _postings_. There is not limit on the
number of _postings_ per transaction.

Note the _accounts_ used in this example.
One begins with `Expenses` and the other begins with `Assets`.
Expenses are _credited_ because the money flows _toward_ them.
Assets are credited when you add funds and debited when you move money to
something else.
In this transaction, you're deducting money from an account representing your
wallet and adding it to an expense representing your coffee spending.

`ledger` has some great conveniences that ease entry.
One such convenience is that `ledger` allows transactions to omit the _amount_
on a single _posting_.
The missing amount is calculated and is equal to whatever amount is necessary
to balance the transaction.

```ledger
2017-06-26 Commonplace Coffee
  Assets:Cash:Wallet
  Expenses:Restaurants:Coffee   3.00

```

You can also supply comments for a transaction or posting.
Postings can only have one comment line but transactions can have as many as
you want.

```ledger
; cold brew
; morning
2017-06-26 Commonplace Coffee
  Expenses:Restaurants:Coffee   3.00   ; Grande
  Assets:Cash:Wallet           -3.00

```

## Ledger file

A ledger file is a series of transactions separated by blank lines in between
them. Here's an example.

```ledger
2013/01/02 McDonald's #24233 HOUSTON TX
    Expenses:Dining Out:Fast Food        5.60
    Assets:Cash:Wallet

2013/01/02 Burger King
    Expenses:Dining Out:Fast Food        15.60
    Assets:Cash:Wallet

2013/01/02 Purchase 100 IVV
    Assets:Bank:Checking       -15000
    Assets:Investments:IVV
    Expenses:Investments:Commissions    4.99

```

You may be wondering how we track stocks, currencies, commodities, etc.
All of those are reporting considerations, transactions are all that's contained
in a ledger file. Simplicity of file formats for possible use with other tools
is a guiding principle of ledger.

Reporting functions available in ledger are very powerful, and will be introduced
in later chapters.

## Differences from Other Ledger

The file format supported by this version of ledger is heavily inspired by the
format defined by the other [ledger](https://www.ledger-cli.org/). However,
this version supports only the most basic of features in the ledger file itself.

### No Currencies or Commodities

There is no support for prepending a currency token (such as $) to a number. Nor
is there support for appending a token or string to signify a commodity (such
as "APPL", "BTC", or "USD").

All amounts must be numbers. Ledger balances and moves amounts (numbers)
between accounts in transactions. The significance of what a number means or
represents is entirely up to the user.

**Note:** Even though there is no support for commodities in the ledger file
format, support for commidities exists in the web reporting features.

### Minimal Command Directive Support

The other ledger supports many [Command Directives](https://www.ledger-cli.org/3.0/doc/ledger3.html#Command-Directives).

The only supported directives are:

* include - to import/include transactions of another ledger file.
* account - parsed but ignored.

All other directives will cause errors in this application as they will be
assumed to be a line starting a transaction.

### Transactions are basic

* No metadata support
* No "state" (pending, cleared, ...)
* No virtual postings
* No balance assertions

Postings are account and an optional amount.
