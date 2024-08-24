package main

import (
	"os"
	"strconv"
)

var STRINGS = os.Getenv("STRINGS")
var SIZE, _ = strconv.Atoi(os.Getenv("SIZE"))
