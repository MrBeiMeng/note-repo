# java 局域网聊天室

## 1 基础知识

### **Java 网络编程**

网络编程是指编写运行在多个设备（计算机）的程序，这些设备都通过网络连接起来。

java.net 包中 J2SE 的 API 包含有类和接口，它们提供低层次的通信细节。你可以直接使用这些类和接口，来专注于解决问题，而不用关注通信细节。

java.net 包中提供了两种常见的网络协议的支持：

- **TCP**：TCP（英语：Transmission Control Protocol，传输控制协议） 是一种面向连接的、可靠的、基于字节流的传输层通信协议，TCP 层是位于 IP 层之上，应用层之下的中间层。TCP 保障了两个应用程序之间的可靠通信。通常用于互联网协议，被称 TCP / IP。
- **UDP**：UDP （英语：User Datagram Protocol，用户数据报协议），位于 OSI 模型的传输层。一个无连接的协议。提供了应用程序之间要发送数据的数据报。由于UDP缺乏可靠性且属于无连接协议，所以应用程序通常必须容许一些丢失、错误或重复的数据包。

### Socket 编程

套接字使用TCP提供了两台计算机之间的通信机制。 客户端程序创建一个套接字，并尝试连接服务器的套接字。

当连接建立时，服务器会创建一个 Socket 对象。客户端和服务器现在可以通过对 Socket 对象的写入和读取来进行通信。

java.net.Socket 类代表一个套接字，并且 java.net.ServerSocket 类为服务器程序提供了一种来监听客户端，并与他们建立连接的机制。

以下步骤在两台计算机之间使用套接字建立TCP连接时会出现：

- 服务器实例化一个 ServerSocket 对象，表示通过服务器上的端口通信。
- 服务器调用 ServerSocket 类的 accept() 方法，该方法将一直等待，直到客户端连接到服务器上给定的端口。
- 服务器正在等待时，一个客户端实例化一个 Socket 对象，指定服务器名称和端口号来请求连接。
- Socket 类的构造函数试图将客户端连接到指定的服务器和端口号。如果通信被建立，则在客户端创建一个 Socket 对象能够与服务器进行通信。
- 在服务器端，accept() 方法返回服务器上一个新的 socket 引用，该 socket 连接到客户端的 socket。

连接建立后，通过使用 I/O 流在进行通信，每一个socket都有一个输出流和一个输入流，客户端的输出流连接到服务器端的输入流，而客户端的输入流连接到服务器端的输出流。

TCP 是一个双向的通信协议，因此数据可以通过两个数据流在同一时间发送.以下是一些类提供的一套完整的有用的方法来实现 socket。



**一个简单的例子**

客户端

```java
import java.io.IOException;
import java.io.OutputStream;
import java.net.Socket;

class A {
    public A() throws IOException {
        //建立连接
        Socket socket = new Socket("127.0.0.1",6000);

        //发送数据
        OutputStream outputStream = socket.getOutputStream();
        outputStream.write("Hello\ndfhoahgoihakfhakfahgf]\ndfghajgdfjad".getBytes());

        //关闭流、socket
        outputStream.close();
        socket.close();
    }

    public static void main(String[] args) throws IOException {
        new A();
    }
}
```

服务端

```java
import java.io.IOException;
import java.io.InputStream;
import java.net.ServerSocket;
import java.net.Socket;

class B {
    public B() throws IOException {
        String str = null;
        //建立连接
        ServerSocket serverSocket = new ServerSocket(6000);
        Socket socket = serverSocket.accept();

        System.out.println("连接成功！！！");

        //接收数据
        InputStream inputStream = socket.getInputStream();
        byte[] bytes = new byte[1024];
        int len;

        while ((len = inputStream.read(bytes))!=-1){
            str = new String(bytes,0,len);
            System.out.println(str);
        }


        //关闭流、socket
        inputStream.close();
        socket.close();
    }

    public static void main(String[] args) throws IOException {
        new B();
    }
}
```

## 2 设计原理

> 每个客户端连接到服务端都会生成一个socket对象。通过这个对象我们可以得到对应的一对输入输出流。
>
> 所以如果我们想创建一个聊天室，我们可以一次性连接多个客户端。将所有生成的socket对象保存起来。当某个人想要向这个房间中的所有人说话的时候。我们可以让这个过程等同为【这个客户端向服务端发起了一次对话。而服务端将他发出的这句话转发给了所有人】
>
> 如果你对这个想法还是不太理解。看这里👉[聊天室原理](https://www.processon.com/view/link/618664775653bb14c2657b94)
>
> 

### 2.1 设计类图

![image-20211106193756863](https://gitee.com/xiao-ai-beimeng/beimeng/raw/master/img/202111061937354.png)

#### **Server**

```java

import java.io.*;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Date;
import java.util.concurrent.CopyOnWriteArrayList;

public class Server {
    //用该容器存储与用户建立的连接
    static CopyOnWriteArrayList<Channel> allClient = new CopyOnWriteArrayList<>();
    //以下两个常量用来区分 系统和用户
    private final static boolean system = true;
    private final static boolean user = false;

    public static void main(String[] args) throws IOException {
        //创建服务端
        ServerSocket server = new ServerSocket(5521);
        System.out.println("服务端初始化完毕 ip地址: "+server.getLocalSocketAddress() + "端口号:" + server.getLocalPort());
        //准备连接
        Channel c;
        //循环等待客户端连接
        while (true){
            //监听到连接后,为该客户端开辟一个线程
            Socket client = server.accept();
            c = new Channel(client);
            new Thread(c).start();
            //加入到容器中
            allClient.add(c);
            System.out.println("用户 "+c.name+ " 加入聊天室...");
        }
    }
    //通信管道,一个用户对应一个管道(内部类),封装了流和用户信息
    static class Channel implements Runnable{
        private Socket client;
        private DataInputStream dis;
        private DataOutputStream dos;
        private boolean isRunning;

        private String name;

        //此构造方法用来初始化流和用户昵称
        public Channel(Socket client){
            this.client = client;
            try {
                dis = new DataInputStream(client.getInputStream());
                dos = new DataOutputStream(client.getOutputStream());
                //默认处于连接状态,用于收发信息
                isRunning = true;
                //初始化姓名
                for (Server.Channel channel : allClient) {
                    if (channel.name.equals(name)){
                        this.send("这个昵称已经被人占用了。请尝试别的昵称");
                        release();
                    }
                }
                //+1为包括自身
                this.send("欢迎进入聊天室，当前在线用户数: "+(allClient.size()+1)+"\r\n"+
                        "指令:ls(查询当前在线用户),@XX:YY(向用户XX发送消息YY),exit(退出聊天室)");
                sendOthers("用户"+this.name+" 加入聊天室"+new Date(),system);
            }catch (IOException e) {
                System.out.println("初始化失败");
                //释放资源
                release();
            }
        }
        //run方法循环转发用户消息
        @Override
        public void run() {
            while (isRunning){
                String msg = receive(client);
                if(!msg.equals("")){
                    if (msg.equals("exit")){//断开连接
                        System.out.println("用户:"+this.name+" 已断开连接");
                        sendOthers(name+"退出了聊天室",system);//告诉所有人 此人离开了
                        release();
                        break;
                    }else if (msg.startsWith("@")){//私聊
                        int index = msg.indexOf(":");
                        if (index!= -1){
                            String targetName = msg.substring(1,index);
                            sendTarget(msg,targetName);
                        }else {
                            sendTarget(msg,"error");
                        }
                    }else if (msg.equals("ls")){//查询在线玩家
                        StringBuilder sb = searchOnline();
                        if (!sb.equals("]")) {
                            this.send("当前在线的玩家:"+sb.toString());
                        }
                    }
                    else {
                        sendOthers(name+"说:"+msg,user);//向所有人广播此消息,除了他自己
                    }
                }
            }
        }

        //接受信息
        public String receive(Socket client) {
            String msg = "";
            try {
                //阻塞,只有用户发送了信息,服务器才会接受并广播给其他用户
                msg = dis.readUTF();
            } catch (IOException e) {
                System.out.println("玩家:"+this.name+" 异常中断");
                //非输入exit退出,同样需要广播此用户的退出
                sendOthers(name+"退出了群聊",system);//告诉所有人 此人离开了
                //客户端强制退出时,释放资源
                release();
            }
            return msg;
        }

        //广播的最底层操作是调用Channel中的send方法
        public void send(String msg){
            try {
                dos.writeUTF(msg);
                dos.flush();
            }catch (IOException e){
                System.out.println("信息传出异常,连接已断开");
                release();
            }
        }

        //广播,调用时确认是系统还是用户
        public void sendOthers(String msg,boolean who){
            if (msg!=""){
                for (Channel c: allClient){
                    //任何消息不回反馈给发送者本身
                    if (c!=this) {
                        if (who == system) {
                            c.send("***系统消息: " + msg + "***");
                        } else {//用户
                            c.send(msg);
                        }
                    }
                }
            }
        }

        //私聊
        public void sendTarget(String msg,String name){
            boolean flag = true;
            for (Channel c: allClient){
                if (c.name.equals(name)){
                    c.send(this.name+"悄悄地对你说:"+msg.substring(msg.indexOf(":")+1));
                    flag = false;
                    break;
                }
            }
            if (flag){
                System.out.println(flag);
                if (name.equals("error")){
                    this.send("指令错误,请输入英文状态下的冒号");
                }else {
                    this.send("***找不到该用户***");
                }
            }
        }

        //查询其他在线的玩家
        public StringBuilder searchOnline(){
            StringBuilder sb = new StringBuilder("[");
            for (Channel c: allClient){
                sb.append(c.name+",");
            }
            sb.setCharAt(sb.length()-1,']');
            return sb;
        }

        //释放资源时从容器中剔除自身的引用(此Channel因各种原因断开),并关闭转发器(停止run方法中的while循环)
        private void release(){
            this.isRunning = false;
            allClient.remove(this);
            System.out.println("剩余玩家"+ allClient.size());
            try {
                dis.close();
                dos.close();
                client.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
```

#### Client

```java
import java.io.*;
import java.net.Socket;
import java.util.Scanner;

public class Client {
    public static void main(String[] args) throws IOException {
        //输入昵称后,再向服务器建立连接
        System.out.print("请输入昵称:");
        String name = new Scanner(System.in).nextLine();
        System.out.print("请输入目标ip地址:");
        String targetIp = new Scanner(System.in).nextLine();
        System.out.print("请输入目标端口:");
        int targetPort = new Scanner(System.in).nextInt();
        //创建客户端
        Socket client = new Socket(targetIp,targetPort);
        /**
         * 用户应该能同时收发消息,所以需要用到多线程
         * Sender为发送者线程
         * Receiver为接收者线程
         */
        new Thread(new Sender(client,name)).start();
        new Thread(new Receiver(client)).start();
    }
}

```

#### Receiver

```java

import java.io.DataInputStream;
import java.io.IOException;
import java.net.Socket;

public class Receiver implements Runnable{
    private Socket client;
    private DataInputStream dis = null;

    //初始化流,用来接收服务器广播的消息
    public Receiver(Socket client){
        this.client = client;
        try {
            dis = new DataInputStream(client.getInputStream());
        } catch (IOException e) {
            System.out.println("初始化失败");
            release();
        }
    }
    @Override
    public void run() {
        while (true){
            try {
                //当输入exit时,此操作异常,run方法终止
                System.out.println(dis.readUTF());
            } catch (IOException e) {
                System.out.println("与服务器断开连接");
                break;
            }
        }
        release();
    }

    //释放资源
    private void release(){
        try {
            client.close();
            dis.close();
        }catch (Exception e){
            e.printStackTrace();
        }
    }
}

```

#### Sender

```java

import java.io.BufferedReader;
import java.io.DataOutputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.Socket;

public class Sender implements Runnable{
    private Socket client;
    private DataOutputStream dos;
    private BufferedReader reader;

    //初始化流,并向服务器传递用户昵称,服务器在初始化链接(Channel)时接收昵称
    public Sender(Socket client,String name){
        this.client = client;
        try {
            dos = new DataOutputStream(client.getOutputStream());
            reader = new BufferedReader(new InputStreamReader(System.in));
            dos.writeUTF(name);
        } catch (IOException e) {
            System.out.println("初始化流失败");
            release();
        }
    }

    @Override
    public void run() {
        while (true){
            try {
                String msg = reader.readLine();
                if (msg.equals("exit")){
                    dos.writeUTF(msg);//输入exit主动断开连接
                    dos.flush();
                    break;
                }
                dos.writeUTF(msg);
                dos.flush();
            } catch (IOException e) {
                System.out.println("发送异常");
                release();
            }
        }
    }

    //释放资源
    private void release(){
        try {
            reader.close();
            client.close();
            dos.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

```

