Tests if packages imported by subpackages/external packages use different instances.

Package output:

```
Interval A: 8
Interval B: 8
Interval External: 8

Interval A set to 15
Interval A: 15
Interval B: 15
Interval External: 15

Interval B set to 20
Interval A: 20
Interval B: 20
Interval External: 20

Interval External set to 25
Interval A: 25
Interval B: 25
Interval External: 25
```