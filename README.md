# Serverless Event-driven Architecture - AWS
A AWS SAM serverless HTTP-API which triggers an eventbridge event which triggers another lambda asyncronously

## Tech stack 
- Golang
- AWS SAM 
- HTTP Api (APIGateway)
- Lambda
- Event Bridge

## Architecture
`/hello` Http API endpoint using APIGateway triggers a synchronous api which in turn triggers another event-driven lambda function
