package main

import (
  "fmt"
//  "strings"
//  "os"
//  "bufio"
)

// At start of program, print hello message to user.
// Wait for user input (add song, delete song, list songs) with a for loop.
// Execute input and print message.
// If user inputs "exit" then end for loop and print goodbye message.

func main() {
  // dataStore will store the songs that users provide
  // A nested map structure with the following format:
  // {
  //   "We Will Rock You": {
  //                         "artist": "Queen",
  //                         "genre":  "Rock"
  //                       }
  // }
  //var dataStore = map[string]map[string]string{}

  fmt.Println("Welcome to the Music Store!")
  fmt.Println("1) add a song")
  fmt.Println("2) delete a song")
  fmt.Println("3) list songs\n")

}
