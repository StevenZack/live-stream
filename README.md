# go-live-stream-example
A simple example of video live stream on HTML5 browser
一个简单的Go语言Web直播示例代码

## Usage 使用方法

Windows: 
```just double click the main.exe ```

Linux:
```chmod +x main-linux-amd64.run &&./main-linux-amd64.run```


Then browse http://YourIP:9090/camera.html to share your camera ,

browse http://YourIP:9090  to watch live content.

## How it works 代码解析

#### main.go

Start a websocket server to broadcast live frame to every end.

#### index.html

connect to websocket server , once received image data , show it.

#### camera.html

get camera data from local device , and send it to Web server
