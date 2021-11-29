File :container.go
  Container uses defualt gin-gonic engine. The router registers http GET handler for "/aboutme".
  
File: myDeployment.yaml
  This is kubernetes deployment file for deploying docker image created from 'container.go' with replica set to 1(single instance)
