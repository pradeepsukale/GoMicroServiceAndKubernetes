File : container.go
  Container uses defualt gin-gonic engine. The router registers http GET handler for "/aboutme".
  
File: myDeployment.yaml
  This is kubernetes deployment file for deploying docker image created from 'container.go' with replica set to 1(single instance)
  
File: myService.yaml
  This is kubernetes service with which we can access pod.
    Ex: http://minicubeip:nodeport/aboutme -- Replace minicube with minicube ip and port with confgired nodeport!
   
File: Dockerfile
  1.Copies source code into container. 2. gets gin-gonic from github  3. Builds image   4. Starts image
