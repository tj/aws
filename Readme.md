# go-cloudwatch

Higher level CloudWatch operations.

## Metrics

```go
package metrics_test

import (
	"time"

	"github.com/tj/go-cloudwatch/metrics"
)

func Example() {
	m := metrics.New().
		Namespace("AWS/ApiGateway").
		Metric("Count").
		Stat("Sum").
		Dimension("ApiName", "app").
		Period(5).
		TimeRange(time.Now().Add(-time.Hour), time.Now())

	res, err := metrics.Get(m)
	_ = res
	_ = err
}
```

---

[![GoDoc](https://godoc.org/github.com/tj/go-cloudwatch?status.svg)](https://godoc.org/github.com/tj/go-cloudwatch)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-stable-green.svg)

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
