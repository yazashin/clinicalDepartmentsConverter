# clinicalDepartmentsConverter
診療科名について一般名から略称を取得したり、その逆ができます。

![対応表](https://github.com/yazashin/clinicalDepartmentsConverter/blob/main/clinical-department-map.png)

# Installation

```
go get github.com/yazashin/clinicalDepartmentsConverter
```

# Usage

```go
package main

import (
	"fmt"

	converter "github.com/yazashin/clinicalDepartmentsConverter"
)

func main() {
	fmt.Println(GetShortName("リハビリテーション科"))

	fmt.Println(GetLongName("心外"))

	fmt.Println(strings.Join(LongNameList(), ", "))

	fmt.Println(strings.Join(ShortNameList(), ", "))
}

```

# Licenses

This software is released under the MIT License, see LICENSE