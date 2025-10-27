# Recursive Function

I tried to build a recursive function by composing the built in functions. This works with both approaches.

```bash
go run examples/recursive/main.go
```

```bash
FACTORIAL(5)
│   ├── FACTORIAL(4)
│   │   ├── FACTORIAL(3)
│   │   │   ├── FACTORIAL(2)
│   │   │   │   ├── FACTORIAL(1)
│   │   │   │   │   FACTORIAL = 1
│   │   │   │   ├── MULTIPLY(2), 1
│   │   │   │   │   MULTIPLY = 2
│   │   │   │   FACTORIAL = 2
│   │   │   ├── MULTIPLY(3), 2
│   │   │   │   MULTIPLY = 6
│   │   │   FACTORIAL = 6
│   │   ├── MULTIPLY(4), 6
│   │   │   MULTIPLY = 24
│   │   FACTORIAL = 24
│   ├── MULTIPLY(5), 24
│   │   MULTIPLY = 120
│   FACTORIAL = 120
Result = 120
```
