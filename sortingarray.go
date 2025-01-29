package main 
import(
	 "fmt";
);

 func sort (ar[] int, n int) [] int{
	    for i:=0;i<n;i++{
			 for j:= i+1;j<n;j++{
				 if ar[i]  > ar[j]{
					 var tmp int;
					  tmp = ar[i];
					   ar[i] = ar[j];
					    ar[j] = tmp;
				 };
			 }
		};
		return ar;
 };
  
 func main (){
	   var n int;
	    fmt.Scan(&n);
		ar :=  make([]int,n);
		for i:= 0; i<n;i++{ 
			fmt.Scan(&ar[i]);
		};
		 
          arso := sort(ar,n);
		  for i:= 0; i<len(arso);i++{
			    fmt.Println(arso[i]);
		  }
 }