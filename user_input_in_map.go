package main 
import (
	 "fmt";
);
 func main(){
	//  var n int;
	//  fmt.Scan(&n);
	//   parson := make(map[int]string);
	//    for i:= 1; i<=n;i++{
	// 	   var number int;
	// 	      var name  string;
	// 		   fmt.Scan(&number);
	// 		    fmt.Scan(&name);
	// 			 parson[number] = name;
	//    };
	//      for number ,name := range parson{
	// 		 fmt.Println(number,":",name);
	// 	 }
	 var n int;
	  fmt.Scan(&n);
	   parson := make(map[int]string);
	    for i:= 1; i<=n;i++{
			 var number int;
			 var name string;
			  fmt.Scan(&number,&name);
			  parson[number] = name;
		};
		  for number ,name := range parson{
			 fmt.Println(number, ":" , name);
		  };
		   for _,number :=  range parson{
			 fmt.Println(number);
		   }
 }