Simple CLI password manager

Command	Description	Example Usage
add	   |  Add a new password	           |   go run main.go add --name="ExampleName" --username="user123" --password="mypassword123" --key="yoursecretkey"
list   |  Show all saved passwords	       |   go run main.go list
find   |  Find a password by password name |   go run main.go find ExampleName --key="yoursecretkey"
remove |  Remove a password for a resource |   go run main.go remove --name="ExampleName"
update |  Update a password for a resource |   go run main.go update --name="ExampleName" --password="newPassword123" --key="yoursecretkey"