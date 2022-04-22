# Rocket Service API
This application is written / build in GoLang 1.18

## Start the application
Run the application with the following command:
### Using default port (8088) and Production mode
```.\rocketservice.exe```.

### Using custom port
```.\rocketservice.exe port 9090```.

### Using custom port and debug mode
```.\rocketservice.exe port 9090 debug```.

## OpenAPI / Swagger
Assuming the application runs on default port `8088`:
http://localhost:8088/swagger/index.html

## Notes
I never worked with events in this fashion before, so I'm sure it could be solved in a better way :D
For the sport of it and contradicting the task description, I used GoLang over C# to develop the solution.

there for I also spend quite a bit more than 6 hours, but I really enjoyed the process.

At the end, when time became an issue, I stopped doing TDD and adding tests, bad practise I know.

Please note that the application serves the API via http and not https as described in the readme.