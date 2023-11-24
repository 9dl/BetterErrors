# BetterErrors

## Overview

**BetterErrors** is a GoLang package that provides a more beautiful and effective way to handle and display errors in your applications. It includes functions for logging detailed error information, recovering from panics, and printing general user information.

## Features

- **LogError:** Log detailed error information with runtime details.
- **RecoverFromPanic:** Recover from panics and log panic details.
- **PrintInfo:** Print general information about the user.

## Usage

### Installation

```bash
go get -u github.com/9dl/BetterErrors
```

### Example

```go
package main

import (
	"errors"
	"github.com/9dl/BetterErrors"
)

func main() {
	err := someFunction()
	BetterErrors.LogError(err)

	// ... rest of your code ...
}

func someFunction() error {
	return errors.New("something went wrong")
}
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
