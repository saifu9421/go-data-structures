package main 
 import "fmt";
  func main (){
	  numbers := [] int {2,3,4,5};
	//   append added  number to the slice.
	   numbers = append(numbers, 10,20,30);
	//    fmt.Println(numbers);
	   evenNumber := [] int {2,4};
	    oddNumber := [] int {1,3};
		 evenNumber = append(evenNumber, oddNumber...);
		//  fmt.Print(evenNumber);
		//  copy() copy function copy the another slice.
		// copy primeNumber slice  theke  number slice.
		   primeNumbers := [] int{2,3,5,7};
		    numbers2 :=  [] int {2,3,5,5};
			copy(numbers2,primeNumbers);
			fmt.Println(numbers2);
			     var flag bool;
				  flag = false;
			  if   len(primeNumbers) == len(numbers2){
				    for i:= 0;i<len(primeNumbers);i++{
						 if primeNumbers[i] != numbers2[i]{
							  flag =  false;
							//  fmt.Println("Two slice Not same/equal");
							 break;
						 }else{
							 flag = true;
						 }
					}
			  };

			  if(flag == true){
				 fmt.Println("Tow slice Equal");
			  }else{
				 fmt.Println("NOT Equal");
			  }
  }