# ProxyIp
****
### Author:mikeweiwei
****
* [get :fetch ip]
* [check: check ip]
* [put :mysql]
* [model: ip struct]
* [web :some api]
    * /v1/all get all proxy ip
    * /v1/type get proxy ip by type（parameter ？type=http/https/all）
* [main]
    * connection DB and create table
    * check in DB and delete (schedule task)
    * get spider to chan (schedule task)
    * chan to DB
    
