package main 
import (
	 "fmt";
);
 func main(){
	
	//   this is slice
	//   numbers := [] int{1,2,3,4,5};
	// //   this is array
	//     number := [5] int{1,2,3,4,5};
	//    fmt.Print(numbers);
	//   for i:= 0; i<len(numbers);i++{
	// 	   fmt.Println(numbers[i]);
	//   }

	  var n int;
	   fmt.Scan(&n);
	   ar := make([]int,n);
	    for i := 0;i<n;i++{
			  fmt.Scan(&ar[i]);
		};

        //     if n >= 2{
		//  sliceNUmber := ar[0:2]; 
		//  fmt.Print(sliceNUmber);
		// 	}else{
		// 		 fmt.Print("Not");
		// 	}
		 arr := [] int {1,2,3,45};
 		//  var number []int;
		number := make([]int,2);
		//    number = append(number, 1,2,3,4);
		//     fmt.Print(number);

		var  tmp  int = 0;
		for i := 1;i<3;i++{
			    number[tmp]  = arr[i];
				 tmp++;
		};
		 
		fmt.Print(number);

 }