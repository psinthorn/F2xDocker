# F2xDocker
Docker and docker-compose file libraries
with general server setting include general security, ssh, UFW and server infomation

## Docker Commands
## Docker Network คืออะไร
`หากต้องการ run 2 container พร้อมกันและต้อการให้ทั้ง 2 container สามารถเชื่อมต่อและสามารถทำรวมกันได้นั้นเราต้องทำอย่างไร`
ในตัวอย่างนี้เราจะทำการ run 2 container พร้อมกันและให้ทั้งสอง container เชื่อมต่อถึงกันผ่าน port ที่เรากำหนด

### Docker-Composer file 

### คำสั่งสำหรับการ run docker compose file แบบระบุชื่อไฟล์แบบเฉพาะเจาะจงและให้ทำงานที่ background 
`docker-compose -f mongodb_mongo-express.yaml up -d`

### คำสั่งสำหรับหยุดการทำงานของ Docker container
`docker-compose down`

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

### Generate public and private keys for mac and linux (this command is working at local computer not on server)
`ssh-keygen -t rsa`
### copy key to server (this command is working at local computer not on server)
`ssh-copy-id adminx@178.128.57.4`
### connect to server with ssh (this command is working at local computer not on server)
`ssh adminx@178.128.57.4`
### modify and backup sshd_config 
`sudo cp /etc/ssh/sshd_config /etc/ssh/sshd_config.dist`
`sudo vi /etc/ssh/sshd_config`

## Server infomation
### how to view current host name
`sudo cat /etc/hostname`
### to view server infomation (ubuntu)
`sudo hostnamectl`
### how to change host name
`sudo hostnamectl set-hostname new-server-name`
### how to view current date and time system will show you UTC time
`date`
### how to change server time zone after run command system will prompt you to select new timezone/city
`sudo dpkg-reconfigure tzdata`

## File Tranfer Protocal
### How to connect to server via sftp for file transfer
`sftp adminx@178.128.57.4`


## How to clone private git repository
### clone command
`git clone https://psinhtorn:gitPersonalToken@github.com/gituser/repository.git`
### how to get github personal token keys



