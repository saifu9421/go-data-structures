package main 
 import "fmt";
  func main(){
	 var ar = [5] int {1,2,3,4,5};
    fmt.Print(ar,"\n");
	var st = [...] string{"Hello saif"};
	fmt.Print(st,"\n");

	languages := [...] string{"go","c++","java"};
	fmt.Println(languages[0]);
	fmt.Println(languages[1]);
	 fmt.Println(languages[2]);
	//  fmt.Println(languages[]);
	  var n int;
	     fmt.Scan(&n);
	lan := make([]string,n);
	  for i:=0;i<n;i++{
		 fmt.Scan(&lan[i]); 
	  };
	fmt.Println(lan);
    number := [] int {1,2,3,4};
	 fmt.Print(number);
  }