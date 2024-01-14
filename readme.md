I use existing libs :

 - Chi Router
 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - testify for unit testing

 # For setup after cloning the repo:
- cd dating_app
- go mod tidy

# to do a unit test in usecase:
- go to the package you want to testing then run a command "go test"
- you can see the coverage testing in each package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package"

# for db table :
> I use Postgresql for DB
>> in folder db, there are some .sql files with the create table command. you can run the command in your sql editor page.

# to running the project
after clone and do some set up that explained before, do this following actions :
- make .env file by copying the .env.example and set database credential in the .env file
- go run main.go

# the endpoint
you can dorp these curl into postman etc

- curl --location --request POST 'http://localhost:8080/api/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"jody.almaida@gmail.com",
    "password" : "1234567"
}'

- curl --location --request POST 'http://localhost:8080/api/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"jody.almaida@gmail.com",
    "password" : "1234567"
}'