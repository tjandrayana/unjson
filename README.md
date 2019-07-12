# Unjson Package

Lately, I am exploring about cypher on Neo4J graph database. 
I found some ways to insert data to Neo4J. 
First using driver and the second using http request.
I found a problem when use http request, when I want to create a node with some dynamic entities.
But I don't know how to implement the dynamic entities.
We can't pass json string to it, and after I read and analyze the data need is not json in general.
The json need is without quotes, like :


**Json in general**
  
  ```
  {
     "name" : "tjandrayana
  }
  ```

**Json without Quotes**

  ```
  {
     name : "tjandrayana" 
  }
  ```


I try to search library that can handle my use case (json without quotes) and I still can't find (CMIIW).
So, I decide to create custom marshal json.





## How To Use

```
package main

import (	
	"time"
	"log"
	"fmt"
	"github.com/tjandrayana/unjson"
)


type User struct{
	UserID 		int64 		`json:"user_id"`
	Username 	string		`json:"username"`
	Age		int		`json:"age"`
	Birthday	time.Time	`json:"birthday"`

}

func main() {
	user := User{
		UserID 		: 1111,
		Username 	: "Mr. Alucard",
		Age		: 25,
		Birthday	: time.Now(),	
	}
	
	result,err := unjson.MarshalWithoutQuotes(user)  // call function custom marshal
	if err != nil{
		log.Println(err)
	}
	
	fmt.Println(result)		
}


```

```
Output : {user_id: 1111, username: "Mr. Alucard", age: 25, birthday: "2009-11-10 23:00:00 +0000 UTC m=+0.000000001"}

```

