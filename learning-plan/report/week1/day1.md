# Day1 Report
## Go installation
First We need to download go 1.26.1 from go official website.
Then we need to delete the old go environment in our local environment.
And then, uncompress it to /usr/local directory.
After this we successfully install go in our linux environment.
```bash
wget https://go.dev/dl/go1.26.1.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.26.1.linux-amd64.tar.gz
```

But we can't use go at this time. 
Becuase we didn't add environment variable to our bashrc file.
So then we need to add the path /usr/local/go/bin to $PATH in the file.
After done this we finally successfully finish installation.

```bash
echo 'export PATH=/usr/local/go/bin:$HOME/go/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
```

## Check environment
```bash
go version
```

And its result will be like
> go version go1.26.1 linux/amd64

## First Application
Use go mod to init a new app.
```bash
go mod init example/hello
```

Paste the following code.
```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, World")
}

```

- Declare a main package (a package is a way to group functions, 
and it's made up of all the files in the same directory).
- Import the popular fmt package, which contains functions for formatting text,
including printing to the console. 
- This package is one of the standard library packages you got when you installed Go.
Implement a main function to print a message to the console. 
A main function executes by default when you run the main package.

The result will be like:
```go
dreamking@ZenlessPC:~/study-for-work/go-backend-learning/week1/day1-hello$ go run .
Hello, World!
```

## Smoke Test
Write an easiest smoke test.
```go
package main
import "testing"

func TestSmoke(t *testing.T) {
    // Todoa
    // arrange: 准备输入
    // act: 调用函数
    // assert: 校验结果
}
```

Run the test.
```bash
dreamking@ZenlessPC:~/study-for-work/go-backend-learning/week1/day1-hello$ go test .
ok      example/hello   0.003s
dreamking@ZenlessPC:~/study-for-work/go-backend-learning/week1/day1-hello$ go test ./...
ok      example/hello   (cached)
```

The variable t can do four things.
  1. 报错失败

  - t.Errorf(...)：记录错误，继续执行
  - t.Fatalf(...)：记录错误并立刻结束当前测试

  2. 打日志

  - t.Log(...) / t.Logf(...)

  3. 组织子测试

  - t.Run("case1", func(t *testing.T){ ... })

  4. 并行/辅助控制

  - t.Parallel() 等
