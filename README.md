# Mini_Project

in this project I am using Gin framework and GORM for database Query

For initiation of the database connection which is written in the init function in inits package

Now to store the Log file in the database I have created the structure of the required field for searching and also one column
nmaed Declare to store the extra information of the log file
for storing use the route http://localhost:8080/store and pass the parameter path="path of the file to store"

since I am operating with the unstructured data I used map[string]interface{} to get the objects of the JSON and store them into 
the database

There are 2 routes for showing the data to the user and the admin

for admin 
          I showed the entire data fields using http://localhost:8080/admin/show
          
          
for user
          I showed entire data except password and Mobile no which is Obscured using http://localhost:8080/user/show
          
          
