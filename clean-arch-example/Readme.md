# The constraints in the Clean Architecture are

Independent of Frameworks
                            The architecture does not depend on the existence of some library of feature laden software.

Testable 
                            The business rules can be tested without the UI, Database, Web Server, or any other external element.


Independent of UI 
                            The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.

Independent of Database 
                            You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.                            


# Architecture we can divide our code in 4 layers :

Entities: 
        encapsulate enterprise wide business rules. An entity in Go is a set of data structures and functions.

Use Cases: 
        the software in this layer contains application specific business rules. It encapsulates and implements all of the use cases of the system.        

Controller: 
        the software in this layer is a set of adapters that convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web.

Framework & Driver: 
                this layer is generally composed of frameworks and tools such as the Database, the Web Framework, etc.                