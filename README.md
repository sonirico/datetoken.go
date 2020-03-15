# datetoken

## Featuring

- **Token representation of relative dates**. (dates whose value depends upon when they are evaluated)
- **Time zone support**. Time zones tend to be a major pending subject to many devs.
- **Configure when week start**. Cause not every country cries on Mondays.
- **Business weeks**. The previous point allows this.

## Motivation

This package aims to solve a set of needs present in applications where
dates need to be represented in a relative fashion, like background periodic
tasks, datetime range pickers... in a compact and stringified format. This
enables the programmer to persist these tokens during the lifetime of a
process or even longer, since calculations are performed in the moment of
evaluation. Theses tokens are also useful when caching URLs as replacement
of timestamps, which would break caching given their mutability nature.

Some common examples of relative tokens:

|                                | From           | To            |
|--------------------------------|----------------|---------------|
| Today                          | `now/d`        | `now`         |
| Yesterday                      | `now-d/d`      | `now-d@d`     |
| Last 24 hours                  | `now-24h`      | `now`         |
| Last business week             | `now-w/bw`     | `now-w@bw`    |
| This business week             | `now/bw`       | `now@bw`      |
| Last month                     | `now-1M/M`     | `now-1M@M`    |
| Next week                      | `now+w/w`      | `now+w@w`     |
| Custom range                   | `now+w-2d/h`   | `now+2M-10h`  |
| Last month first business week | `now-M/M+w/bw` | `now-M/+w@bw` |

As you may have noticed, token follow a pattern:

- The word `now`. It means the point in the future timeline when tokens are
  parsed to their datetime form.
- Optionally, modifiers to add and/or subtract the future value of `now` can
  be used. Unsurprisingly, additions are set via `+`, while `-` mean
  subtractions. These modifiers can be chained as many times as needed.
  E.g: `now-1M+3d+2h`. Along with the arithmetical sign and the amount, the
  unit of time the amount refers to must be specified. Currently, the supported
  units are:
  - `s` seconds
  - `m` minutes
  - `h` hours
  - `d` days
  - `w` weeks
  - `M` months
- Optionally, there exist two extra modifiers to snap dates to the start or the
  end of any given snapshot unit. Those are:
  - `/` Snap the date to the start of the snapshot unit.
  - `@` Snap the date to the end of the snapshot unit.

  Snapshot units are the same as arithmetical modifiers, plus `bw`, meaning
  _business week_. With this, we achieve a simple way to define canonical
  relative date ranges, such as _Today_ or _Last month_. As an example of
  the later:

  - String representation: `now-1M/M`, `now-1M@M`
  - Being today _15 Jan 2018_, the result range should be:
    _2018-01-01 00:00:00 / 2018-01-31 23:59:59_

## Compatibility

This library has been developed under version 1.14 but you can expect the lib
to work under lower versions too.

## Installing

```shell
go get github.com/sonirico/datetoken.go

go mod install github.com/sonirico/datetoken.go
```

## Examples

Most probably you will be dealing with simple presets such as _yesterday_ or
the _last 24 hours_.

```golang
package main

import (
    "fmt"
    "time"

    "github.com/sonirico/datetoken.go"
)

func main() {
    Madrid, _ := time.LoadLocation("Europe/Madrid")
    time, err := datetoken.Eval("now/d", Madrid, time.Monday)
    if err != nil {
        panic(err)
    }
    fmt.Println(time.String())
}
```

For more examples you can refer to https://github.com/sonirico/datetoken.go/tree/master/examples    

