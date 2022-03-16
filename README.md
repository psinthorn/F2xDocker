# F2xDocker
Docker and docker-compose file libraries

## Docker Commands
## Docker Network คืออะไร
`หากต้องการ run 2 container พร้อมกันและต้อการให้ทั้ง 2 container สามารถเชื่อมต่อและสามารถทำรวมกันได้นั้นเราต้องทำอย่างไร`
ในตัวอย่างนี้เราจะทำการ run 2 container พร้อมกันและให้ทั้งสอง container เชื่อมต่อถึงกันผ่าน port ที่เรากำหนด

### Docker-Composer file 

### คำสั่งสำหรับการ run docker compose file แบบระบุชื่อไฟล์แบบเฉพาะเจาะจง
`docker-compose -f mongodb_mongo-express.yaml up`

## Server and Security command (Linux)
### how to add user to linux server
`adduser adminx`
### How to add user to super user group (sudo)
`usermod -aG sudo adminx`
### how to update and upgrade server
`sudo apt update`
`sudo apt upgrade`
### then reboot server refresh update 
`sudo reboot`

### how to switch to another user
`su - username`

### how to delete user and user home directory 
#### delete user 
`sudo userdel username`
#### remove directory, sub-directory and all files that stored in all directory (-rf recurcive force) 
`sudo rmdir -rf directory name`

### Generate public and private keys for mac and linux
`ssh-keygen -t rsa`
### copy key to server
`ssh-copy-id adminx@178.128.57.4`
### connect to server with ssh
`ssh adminx@178.128.57.4`

### how to clone private git repository
`git clone https://psinhtorn:gitPersonalToken@github.com/gituser/repository.git`



