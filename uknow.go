package main

import (
  "fmt"
  "os"
  "bufio"
  "io/ioutil"
  "strings"

)

func main() {

  data, err := ioutil.ReadFile("./the_book_of_questions.txt")

  if err != nil {
    panic(err)
  }

  questions := strings.Split(string(data), "\n")

  reader := bufio.NewReader(os.Stdin)

  for _, question := range questions {

    fmt.Println(question)

    fmt.Print("Input sth: ")
    data, _, _ := reader.ReadLine()
    fmt.Println(string(data))
  }

}
