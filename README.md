# Period

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/cyrildever/period)
![GitHub last commit](https://img.shields.io/github/last-commit/cyrildever/period)
![GitHub issues](https://img.shields.io/github/issues/cyrildever/period)
![GitHub](https://img.shields.io/github/license/cyrildever/period)

A _Period_ is a predefined range of milliseconds during which operations may happen.
It has an `ID`, ie. the number of _Periods_ elapsed since the original _Period_ set at initialization.
You can also get its start and end timestamps.

It's mainly used in immutable shared ledgers that implements the [**_DRD<sup>2</sup>_** consensus protocol](https://github.com/cyrildever/drd2).


### Usage

```console
go get github.com/cyrildever/period
```

You first have to initialize the main parameters of the _Period_ environment: _(in milliseconds)_
```golang
import "github.com/cyrildever/period"

originTimestamp := 1574252822201
periodSpan := 10000

err := period.Init(originTimestamp, periodSpan)
```

Then, you can start using _Periods_:
```golang
import (
    "github.com/cyrildever/period"
    "github.com/cyrildever/period/timestamp"
)

currentTimestamp, err := timestamp.CurrentMillis()
p := period.Get(currentTimestamp)

if currentTimestamp != p.StartTimestampMillis() {
    // It's an error
}
if p.EndTimestampMillis() != p.StartTimestampMillis() + period.Span() - 1 {
    // It's another error
}
if p.Next().ID != p.ID + 1 {
    // Yet another error
}
```

**IMPORTANT:** You might want to notice that, as there is no year zero in the Gregorian calendar, the first _Period_ ID is `1` (not `0` as we, computer engineers, generally like to start our arrays).


NB: You should set a NTP server somewhere but, if you don't, you still can get the current timestamp of the machine running your program by passing `true` as an argument to the `CurrentMillis` function, eg.
```golang
currentTimestamp, _ := timestamp.CurrentMillis(true)
```


### API

```golang
// Mandatory initialization
err := period.Init(originTimestamp, periodSpan, isTestEnvironment)
if _, ok := err.(*period.AlreadyInitializedError); !ok {
    // Do something with any error other than period already initialized
}

// To get the current Period
currentPeriod, err := period.Current()

// To get a period at timestamp ts
periodAtTS, err = period.Get(ts)

// To get the span of the Periods either in milliseconds or as a time.Duration
ms := period.Span()
duration := period.GetDuration()

// To get the start of a Period
start := currentPeriod.StartTimestampMillis()

// To get the end of a Period, ie. the last millisecond before moving to the next Period
end := currentPeriod.EndTimestampMillis()

// To get one's following Period
nextPeriod := currentPeriod.Next()

// Unset test environment
isTest := period.SetTestEnvironment(false)
```


### License

This module is distributed under a MIT license.
See the [LICENSE](LICENSE) file.


<hr />
&copy; 2019-2022 Cyril Dever. All rights reserved.