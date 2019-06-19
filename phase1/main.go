package main

import (
  "fmt"
  "strings"
  "os"
  "bufio"
)

func main() {
  // dataStore will store the songs that users provide
  // A nested map structure with the following format:
  // {
  //   "We Will Rock You": {
  //                         "artist": "Queen",
  //                         "genre":  "Rock",
  //                         "year":   "1977"
  //                       }
  // }
  var dataStore = map[string]map[string]string{}

}
