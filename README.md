Lately, I am exploring about Neo4J graph database. 
I found some ways to insert data to Neo4J. 
First using driver and the second using http request.
I found a problem when use http request, when I want to create a node with some dynamic entities.
But I don't know how to implement the dynamic entities.
We can't pass json string to it, and after I read and analyze the data need is not json in general.
The json need is without quotes, like :
- Json in general
  {
     "name":"tjandrayana
  } 

- Json without Quotes
  {
     name:"tjandrayana" 
  }


I try to search library that can handle my use case (json without quotes) and I still can't find (CMIIW).
So, I decide to create custom marshal json.
