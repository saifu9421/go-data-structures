package main 
 
import "fmt";
// key->value.
 func main(){

	 m:= map[string]int{
		  "age":23,
		  "cgpa":4,
		  "Golang":34,
		  "java":34,
		  "c++":45,
	 };

	 m1 := map[string] bool{
		   "abc" : true,
		   "xyz": false,
	 };
	  
         value := m["age"];
		  fmt.Println(value);
		   bol ,ok := m1["abc"];
		       if ok == true{
				 fmt.Println("value axgised successfully",bol);
			   }
		    delete(m,"age");
			fmt.Println(m);
			    var m2 map[string] int; 
				 fmt.Println(m2["name"]);
				//   m2["name"] = 233;
				//    fmt.Println(m2);

				 nameAge := map[string] int{
					   "saif" :12,
					    "asif" :34,
						"alif" :45,
						 "kamal" :45,
				 };
				  
				 fmt.Println(nameAge);
				 nameAge["saif"] = 45;
				  fmt.Println(nameAge);

				  parson :=  map[int] string{
					 1:"saif",
					 2:"Kamal",
					 3:"jamal",
					 4:"sofiya",
				  };

				//   for number ,name :=  range parson{
				// 	fmt.Println(number,":",name);
				//   }

				for number , name := range parson {
					 fmt.Println(number,":",name);
				};

				parson2 := make(map[int]string);
				    
                    for number ,name :=  range parson2{
						 fmt.Println(number,":",name);
					}
 }