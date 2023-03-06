# point-manage-api
point manage demo app api

# endpoints
| HTTP method | Path                             | Overview                                                 | 
| ----------- | -------------------------------- | -------------------------------------------------------- | 
| POST        | /register                        | Register a new user                                      | 
| POST        | /login                           | Obtain an access token with registered user information  | 
| POST        | /users/{user_id}/points/add      | add up the number of points using an access token        | 
| POST        | /users/{user_id}/points/subtract | Subtract the number of points using an access token      | 
| GET         | /users/{user_id}/points          | reference the number of points using an access token     | 
| GET         | /users/{user_id}/points/history  | Refer to the list of point history using an access token | 
