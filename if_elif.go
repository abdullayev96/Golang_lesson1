//////////////   if and elif 

// package main

// import "fmt"

// func main() {
	
// // taking a local variable
//  var v int = 700

// if v < 1000 {

// 	fmt.Printf("v is less than 1000\n")
// }
	
// fmt.Printf("Value of v is : %d\n", v)
	
// }

////////////

// Go program to illustrate the 
// use of if...else statement
// package main

// import "fmt"

// func main() {
	
// // taking a local variable
// var v int = 1200
// var b int 

// if v < b {
		
// 	fmt.Printf("v is less than  b \n")
//     fmt.Scan(&b)
	
// } else {
		
	
// 	fmt.Printf("v is greater than b \n")
// }
	
// }

///////////////////////////

// package main

// import "fmt"

// func main() {
// 	// taking a local variable
// 	var v int = 1200
// 	var b int

// 	// Prompt user for input
// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&b) // Read user input

// 	if v < b {
// 		fmt.Printf("v kichik  b dan ")
// 	} else {
// 		fmt.Printf("v kotta  b dan ")
// 	}
// }

//////////////////////////

// package main
// import "fmt"

// func main() {
	
// // taking two local variable
// var v1 int = 400
// var v2 int = 700
// var b int

// fmt.Print("Enter a number: ")
// fmt.Scan(&b) // Read user input

// // using if statement
// if( v1 > b  ) {
		
// 	fmt.Printf("Value of v1 kotta b dan " );

// }else if( v1 < b && b< v2) {
	
// 	fmt.Printf("Value of v1 kotta v2 kotta b " );

// }else {
        
//     fmt.Printf("bunday amal topilmadi !!")
//  }
// }

/////////  Hafta kunini aniqlash 

// package main
// import "fmt"

// func main() {
	

// var b int

// fmt.Print("Enter a number: ")
// fmt.Scan(&b) // Read user input

// // using if statement
// if( b ==1  ) {
		
// 	fmt.Printf("Dushanba" );

// }else if(b==2) {
	
// 	fmt.Printf("Seshanba");



// }else if(b==3) {
	
//     fmt.Printf("Chorshanba");


// }else if(b==4) {
	
//     fmt.Printf("Payshanba");
    
    
// }else {
        
//     fmt.Printf("bunday amal topilmadi !!")
//  }
// }

///////////////   kakulyator 
// package main
// import "fmt"

// func main() {
	

// var b,a  int
// var c string


// fmt.Print("Qanday amal bajarasiz: ")
// fmt.Scan(&c) // Read user input
// fmt.Print("Enter a number1: ")
// fmt.Scan(&b) // Read user input
// fmt.Print("Enter a number2: ")
// fmt.Scan(&a) // Read user input

// // using if statement
// if( c == "+") {	
// 	fmt.Println(a+b);

// }else if(c == "-") {
	
// 	fmt.Println(a-b);

// }else if(c == "*") {
	
//     fmt.Println(a*b);


// // }else if(b==4) {
	
// //     fmt.Printf("Payshanba");
    
    
// }else {
        
//     fmt.Printf("bunday amal topilmadi !!")
//  }
// }


package main
import "fmt"

func main() {
	

var b int

fmt.Print("Enter a number1: ")
fmt.Scan(&b) // Read user input


// using if statement
if( b > 0 ) {

	fmt.Print("Musbat");

}else if(b < 0) {
	
	fmt.Print("Manfiy");

}else if(b == 0) {
	
    fmt.Print("Musbat");

    
}else {
        
    fmt.Print("bunday amal topilmadi !!")
 }
}