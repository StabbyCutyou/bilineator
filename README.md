Bilineator
===========

Provide X1,Y1 and X2,Y2, along with the results of an assumed function F for each
combination of inputs Q11, Q12, Q21, Q22, you can call Bilineate and provide an
X and Y value, for which you will get the result of the function F

The API is subject to change as it's probably not perfect

Also, pretty sure it isn't working

Example
=======

```go
b := &Bilineator{
  // Inputs
  X1: 5.0,
  X2: 10.0,

  Y1: 6.0,
  Y2: 11.0,

  //Results of function f over the combination of inputs
  Q11: 20.0, // f(X1,Y1)
  Q12: 50.0, // f(X1,Y2)
  Q21: 2.0,  // f(X2,Y1)
  Q22: 40.0, // f(X2,Y2)
}

log.Println(b.Bilineate(14.0, 20.0))
```

LICENSE
=========
Apache v2 - See LICENSE
