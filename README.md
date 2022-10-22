# windows-app

Öncelikle go dilinde windows form ve diğer sistemsel işlemler için
bir C derleyicisine ihticayımız olacak. 
Aşağıdaki linklerde istediğiniz bir derleyiciyi kurabilirsiniz.
```
https://www.msys2.org/
```
```
https://jmeubank.github.io/tdm-gcc/download/
```
```
https://www.cygwin.com/
```
Program çalıştığında cmd ekranı da form ile birlikte açılır. 
cmd açmadan build etmek için aşağıdaki komut satırını kullanabilirsiniz.

```
go build -ldflags="-H windowsgui"
```

