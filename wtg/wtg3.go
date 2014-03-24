package main

import "fmt"

func main() {
        printMapDays()
        printStruct()
        vcard()
        printRec()
        printAnonymous()
        em := &employee{3232.10}
        em.giveRaise(1.30)
        fmt.Println(em)
        c1 := &car{wheelCount: 22}
        fmt.Println(c1.numberofWheels())
        m1 := &mercedes{car{wheelCount: 4}}
        fmt.Println(m1.numberofWheels())
        m1.sayHiToMerkel()
}

func mapDays() map[int]string {
        days := map[int]string{
                1: "Monday", 2: "Tuesday", 3: "Wednesday",
                4: "Thursday", 5: "Friday", 6: "Saturday",
                7: "Sunday"}
        return days
}
func printMapDays() {
        days := mapDays()
        fmt.Println(days)
        for _, value := range days {
                if value == "Thursday" {
                        fmt.Println("Thursday is a day in the week")
                } else if value == "Hollyday" {
                        fmt.Println("Hollyday is a day in the week!")
                } else {
                        fmt.Println("I can't find what you are looking for.")
                }
        }
}

type struct1 struct {
        i1  int
        f1  float32
        str string
}

func printStruct() {
        ms := new(struct1)
        ms.i1 = 10
        ms.f1 = 15.5
        ms.str = "severin"
        fmt.Printf("The int is : %d\n", ms.i1)
        fmt.Printf("The float is: %f\n", ms.f1)
        fmt.Printf("The string is: %s\n", ms.str)
        fmt.Printf("the struct is : %v\n", ms)
        ms2 := &struct1{20, 33.3, "martin"}
        fmt.Println(ms2)

}

type address struct {
        firstname string
        lastname  string
        street    string
        city      string
        birthdate string
        photo     string
}

func vcard() {
        v1 := &address{firstname: "martin", lastname: "severin",
                street: "somestreet 3", city: "toxicity",
                birthdate: "1/2/3", photo: "photo.jpg"}
        fmt.Println(v1)
}

type rectangle struct {
        length float32
        width  float32
}

func (r *rectangle) area() float32 {
        return r.length * r.width
}
func (r *rectangle) perimeter() float32 {
        return 2*r.length + 2*r.width
}
func printRec() {
        r1 := &rectangle{3.0, 4.5}
        r2 := &rectangle{2.3, 4.1}
        fmt.Printf("area of r1 is: %.2f\n", r1.area())
        fmt.Printf("perimeter of r2 is: %.2f\n", r2.perimeter())

}

type anonymous struct {
        f1 float32
        int
        string
}

func printAnonymous() {
        a1 := &anonymous{3.2, 4, "hallo"}
        fmt.Println(a1)
}

type employee struct {
        salary float32
}

func (emp *employee) giveRaise(raise float32) {
        emp.salary *= raise
}

type engine interface {
        Start()
        Stop()
}
type car struct {
        engine
        wheelCount int
}
type mercedes struct {
        car
}

func (c *car) numberofWheels() int {
        return c.wheelCount
}

func (m *mercedes) sayHiToMerkel() {
        fmt.Println("hi, angie")
}
