package main 
import "fmt";

type Clacolat func(int , int) int;

type Area struct{
	 length int
	 width int
	 rect Clacolat
};

 func main (){
	//    type parson struct{
	// 	 name string;
	// 	  age int;
	//    };
	//     var parson1 parson;
	// 	 parson1 = parson{
	// 		 name:"saif",
	// 		  age: 24,
	// 	 };
	// 	 parson2 := parson{"said",23};
	// 	  fmt.Println(parson2.name);
    //     //  fmt.Println(parson2);
	// 	  fmt.Println(parson1);

	//   type ractangle struct{
	// 	 lenght int;
	// 	 breadth int;
	// 	 wieght int;
	//   };
	//   ret := ractangle{12,23,2};
	//   fmt.Println(ret.lenght);
	//   fmt.Println(ret.breadth);

	//   area := ret.lenght * ret.breadth;
	//   fmt.Println(area);
	   
	//   area := ractangle{23,34,3};
	//   fmt.Println(area.wieght);
  
	   
	 type Parson struct {
		 fastName  string;
		  lastName  string;
		   age  int;
		    cgpa float64;
	 };
	    // var parson Parson;
		//  parson = Parson{
		// 	 fastName: "saifur",
		// 	 lastName: "rahman",
		// 	 age: 23,
		// 	 cgpa: 3.45,
		//  };
		//   for   name,detail := range parson{
		// 	   fmt.Println(name,":",detail);
		//   }
		// fmt.Println(parson);



	var ar  Area;
	 ar = Area{
		 length: 10,
		 width: 30,
		//   rect: func(len int , wi int) int {
		// 	  return len*wi;
		//   },
		rect: func(length int, wi int) int {
			return length * wi;
		  },
	 };
	//  fmt.Println(ar);
	//   fmt.Println(ar.rect(ar.length,ar.width));
	fmt.Println(ar.rect(ar.length,ar.width));

 }


// Program to use function as a field  of struct

