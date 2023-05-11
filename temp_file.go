//create a channel in main function, pass any type of data into the channel, create a go routine, pass this channel, receive the data in Go Routine, print the received data in go routine.

// SELECT
//     customers.customer_id,
//     customers.customer_name,
//     order_details.total_price
// FROM
//     customers
//     JOIN (
//         SELECT
//             orders.customer_id,
//             SUM(order_items.quantity * products.price) AS total_price
//         FROM
//             orders
//             JOIN order_items ON orders.order_id = order_items.order_id
//             JOIN products ON order_items.product_id = products.product_id
//         GROUP BY
//             orders.customer_id
//     ) AS order_details ON customers.customer_id = order_details.customer_id
// WHERE
//     order_details.total_price > 1000

package main

import "fmt"

func main() {
	// declaring array
	a := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println("Array after creation:", a)

	var b []string = a[1:4] //created a slice named b
	fmt.Println("Slice after creation:", b)

	b[0] = "changed" // changed the slice data
	fmt.Println("Slice after modifying:", b)
	fmt.Println("Array after slice modification:", a)
}
