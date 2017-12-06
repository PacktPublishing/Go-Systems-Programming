# Go Systems Programming
This is the code repository for [Go Systems Programming](https://www.packtpub.com/networking-and-servers/go-systems-programming?utm_source=github&utm_medium=repository&utm_campaign=9781787125643), published by [Packt](https://www.packtpub.com/?utm_source=github). It contains all the supporting project files necessary to work through the book from start to finish.
## About the Book


Go is the new systems programming language for Linux and UNIX systems. It is also the language in which some of the most prominent Cloud-level systems have been written, for instance Docker et al. Where C programmers used to rule, Go programmers are gaining considerable traction to write highly optimized systems programming code.

Created by the original designers of C and UNIX, it expands the systems programmers toolkit to add a mature, clear programming language. Threads and networking code, as well as traditional systems application become easier to write since pointers are not relevant and garbage collection has been taken away the most problematic area for low-level systems code: memory management.

This book opens up the world of high-performant systems to the beginning Go programmer. It does not get stuck on single systems or even system types, but tries to expand the original teachings from Unix system's level programming to all types of servers, the cloud and the web.

## Instructions and Navigation
All of the code is organized into folders. Each folder starts with a number followed by the application name. For example, Chapter02.



The code will look like the following:
```
package main 
 
import "fmt" 
import "os" 
 
func main() { 
   arguments := os.Args 
   for i := 0; i < len(arguments); i++ { 
         fmt.Println(arguments[i]) 
   } 
} 
```

This book requires a computer running a Unix variant with a relatively recent Go version, which includes any machine running Mac OS X, macOS, or Linux.

Apple used to call its operating system as Mac OS X followed by the version number; however, after Mac OS X 10.11 (El Capitan), Apple changed that, and Mac OS X 10.12 is now called macOS 10.12 (Sierra) – in this book, the terms Mac OS X and macOS are used interchangeably. Additionally, there is a big chance that by the time you read this book, the latest version of macOS will be macOS 10.13 (High Sierra). You can learn more about the various versions of macOS by visiting https://en.wikipedia.org/wiki/MacOS.

All of the Go code in this book has been tested with Go 1.8.x running on a iMac using macOS 10.12 Sierra and with Go version 1.3.3 running on a Debian Linux machine. Most of the code can run on both Go versions without any code changes. However, when newer Go features are used, the code will fail to compile with Go 1.3.3—the book states the Go programs that will not compile with Go version 1.3.3 or require Go version 1.8 or newer.

Please note that at the time of writing this text, the latest Go version is 1.9. Given the way Go works, you will be able to compile all the Go code of this book in newer Go versions without any changes.

## Related Products
* [Go: Design Patterns for Real-World Projects](https://www.packtpub.com/application-development/go-design-patterns-real-world-projects?utm_source=github&utm_medium=repository&utm_campaign=9781788390552)

* [Go: Building Web Applications](https://www.packtpub.com/application-development/go-building-web-applications?utm_source=github&utm_medium=repository&utm_campaign=9781787123496)

* [Building Microservices with Go](https://www.packtpub.com/application-development/building-microservices-go?utm_source=github&utm_medium=repository&utm_campaign=9781786468666)

