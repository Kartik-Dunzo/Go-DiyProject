# Go-DiyProject

Created according to the requirements mentioned in https://docs.google.com/document/d/1cl4KRB-BHDou9DgMMHMplwodLDL0inyI4CNR1etOC9g/edit


Added few more services like(user) with the services mentioned in thw document.

A very basic class diagram is created to explain the functionality ![](https://github.com/anejakartik/Go-DiyProject/blob/cb224c83491346d0403ff8aa1f9c979fa0aac11a/extras/product_class%20diagram.jpeg)

Instruction for setup and running the code first time:

edit the diy_project/config.json for configurations like Database, Product and Server details

now execute following commands in terminal at diy_project directory

go get .

go build main.go

go run main.go

A postman collection created for all the services in extras directory for utilizing available services.

After successfully running the service you can perform unit test case at diy_project/testing.
Execute "go test -v" for running all the cases.
