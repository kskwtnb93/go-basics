package main

import (
	"fmt"
	"time"
)

type item struct {
	price float32
}

func main() {
	a := -1

	if a == 0 {
		fmt.Println("zero")
	} else if a > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("negative")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// for {
	// 	fmt.Println("working")
	// 	time.Sleep(2 * time.Second)
	// }
	var i int
	for {
		if i > 3 {
			break
		}
		fmt.Println(i)
		i += 1
		time.Sleep(300 * time.Millisecond)
	}

loop: // breakするときswitch文だけではなくfor文も抜けるために、for文に'loop'という名前を付けている
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			continue
		case 3:
			continue
		case 8:
			break loop // for文の名前を指定して処理を抜けている
		default:
			fmt.Printf("%v ", i)
		}
	}
	fmt.Printf("\n")

	items := []item{
		{price: 10.},
		{price: 20.},
		{price: 30.},
	}
	// コピーが生成されるので元のitemsの値が変わらない
	for _, i := range items {
		i.price *= 1.1
	}
	fmt.Printf("%+v\n", items)
	// インデックスを指定してアクセスすると元のitemsの値が変更される
	for i := range items {
		items[i].price *= 1.1
	}
	fmt.Printf("%+v\n", items)
}
