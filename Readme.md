# GoFLParceXML

### Script to optimize the search for orders on FL.RU

#

## Algorithm v1.0

- [X] Parsing proxy from site

- [X] Send proxy request to FL and get XML

- [X] Parsing XML

- [X] Send new order to us in Telegram 

## Algorithm v2.0

- [ ] Cover everything with tests

- [ ] Put it in the docker

- [ ] Check proxy


  
## Problems
- [X] If the proxy has any problems then our program fails, find a solution 
  - At the moment we decided that when an error we send a request from our ip, later planning to implement an automatic change of proxy, but now I am lazy :)

