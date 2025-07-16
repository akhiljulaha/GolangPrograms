ype OrderStatus int      //1-a   // Define a new custom type called OrderStatus based on int
// // Decla			re a list of constant values of type OrderStatus using iota
// const(                             // 1-b
// 	Received OrderStatus = iota   // iota=0 it's untyped integer and i will increment
// 	Confirmed
// 	Prepared
// 	Delivered
// )