Bilineator
===========

Provide X1,Y1 and X2,Y2, along with the results of an assumed function F for each
combination of inputs Q11, Q12, Q21, Q22, you can call Bilineate and provide an
X and Y value, for which you will get the result of the function F

The API is subject to change as it's probably not perfect

Example
=======

```go
// Straw function f - assume you don't know this ahead of time
f := func(x float64, y float64) float64 {
  return x + y*2
}
b := &Bilineator{
  X1: 5.0,
  X2: 10.0,

  Y1: 6.0,
  Y2: 11.0,

  Q11: f(5.0, 6.0)    //X1,Y1,
  Q12: f(5.0, 11.0)   //X1,Y2,
  Q21: f(10.0, 6.0)   //X2,Y1,
  Q22: f(10.0, 11.0)  //X2,Y2,
}

log.Println(b.Bilineate(15.0, 16.0))
```

LICENSE
=========
Apache v2 - See LICENSE
