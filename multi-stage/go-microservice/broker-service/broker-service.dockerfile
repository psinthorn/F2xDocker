# base go image
# สั่ง go ให้สร้าง app เป็น docker image 
FROM golang:1.18-alpine as builder

RUN mkdir /app
WORKDIR /app
COPY . /app 
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
RUN chmod +x /app/brokerApp

# build a tiny docker image
# สร้าง docker image app จาก docker image จากด้านบน เพื่อให้มีขนาดไฟล์่ที่เล็กลง

# create new alpine docker image
# สร้าง docker image ขึ้นมาใหม่โดยใช้ alpine image 
FROM alpine:latest

# create app folder
# สร้าง folder ชื่อ app 
RUN mkdir /app
WORKDIR /app

# copy ไฟล์จาก brokerApp ที่สร้างจาก builder stage มาไว้ที่ /app folder
COPY --from=builder /app/brokerApp /app

# คำสั่งให้ run app 
CMD ["/app/brokerApp"]