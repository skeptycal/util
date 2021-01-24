package diffuser

import "fmt"

func tryout(){
    var itemValue = "foo"
    var item  =  &itemValue

    var items = []*string{item}

    // Create slice of pointers to strings
    for i, thing := range items {
        tmp := fmt.Sprintf("%v-%d", *thing, i)
        items[i] = &tmp
    }

    mediumUpdate(item)

    loopBad(items)

    loopGood(items)

}
