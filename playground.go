package main

  import (
    "os"
    "io"
    "fmt"
    "log"
    "strconv"
    "encoding/csv"
  )

type Person struct {
  name string
  age int
  next, prev *Person
}

type List struct {
  head, tail *Person
}

func (l *List) Push(p Person) *List { //ref: http://l3x.github.io/golang-code-examples/2014/07/23/doubly-linked-list.html (have some minor edits like putting next and prev in the value structure. Seemed weird to have them seperate)
  peep := &p
  if l.head == nil {
    fmt.Println("in if")
    fmt.Println(peep)
    l.head = peep
    l.tail = peep
  } else { // if{} and then putting else on the next line does not compile. Has to be } else {}
    temp := *l.head
    //hold := *l.head // a lot nicer than c
    l.tail.next = peep //list has head and tail of type person so can access elements of the head person or tail person
    peep.prev = l.tail
    fmt.Println(temp)
    fmt.Println(peep)
  }
  l.tail = peep
  return l
}

func main() {
  fmt.Println("Welcome")
  file, err := os.Open("test.csv")
  if err != nil {
    fmt.Println(err) // this should not cause program to exit, unlike .Fatal? confirmed
  }

  fmt.Println("still here")

  defer file.Close() // will close once main ends
  data := csv.NewReader(file)
  //fmt.Println(data) //prints weird stuff, based off go documentation "Read always returns either a non-nil record or a non-nil error, but not both"
  // allData, err := data.ReadAll()
  // if err != nil {
  //   log.Fatal(err) //"Fatal is equivalent to Print() followed by a call to os.Exit(1). "
  // }
  // fmt.Println(allData) //in format [[] [] []] so list of lists or approx
  // fmt.Println(allData[1])
  m := make(map[string]int)
  l := new(List)
  header, _ := data.Read()
  fmt.Println(header) //"removes" the header. Might be a nicer way to do this for range [1:] esque thing would be nice but not sure if it exists
  for {
    entry, err := data.Read()
    if err == io.EOF {
      break //standard practice for golang seemingly: https://www.dotnetperls.com/csv-go and https://github.com/thbar/golang-playground/blob/master/csv-parsing.go
    } else if err != nil {
      log.Fatal(err)
    }
    x, _ := strconv.Atoi(entry[1]) //I know the data, normally should check the err
    m[entry[0]] = x
    p := Person{entry[0],x,nil,nil}
    l.Push(p)

    fmt.Println(entry)
  }
  fmt.Println(m)
  fmt.Println(l)
}
