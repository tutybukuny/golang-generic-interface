# Generic Interface Use Case
This repo is an example for using generic interface of golang

## Structure
- cmd: contains running function
- entity: define all entity models that represent tables inside database
- repository: define repository interface that service layer will interact directly
- storage: define all storages and implement of them that we would use like mysql, postgresql, mongodb,etc