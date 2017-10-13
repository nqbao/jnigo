# jnigo
JNI wrapper for Go

## Installation

### for Debian/Ubuntu

~~~
apt install default-jre-headless

go get github.com/juntaki/jnigo
~~~

### for Mac

~~~
brew cask install java

export JAVA_HOME=`/usr/libexec/java_home` 
export CGO_CFLAGS="-I$JAVA_HOME/include -I$JAVA_HOME/include/darwin" 
export CGO_LDFLAGS="-L$JAVA_HOME/jre/lib/server/ -ljvm -lpthread" 
go get github.com/juntaki/jnigo
~~~
