# Go webserver template

Writing a webserver in go gets really repetitive after a while, this is a boilerplate to 
get started quickly.

Dependencies:
1. gin-gonic/gin for routing
2. go-pg/pg for postgres orm
3. spf13/viper for config management
4. (todo) sirupsen/logurs for logging
5. (todo) jarcoal/httpmock for http mocking unit tests

## Getting started
1. Rename go module to your needs: https://www.jetbrains.com/go/guide/tips/rename-go-module-name/
2. docker-compose up -d to run as it is.

### TODO
1. Logging
2. Interfaces for services + gomock
3. Interfaces for go-pg + gomock
4. Sample http mock test for services' external api calls
5. point 2 to 4 => more tests