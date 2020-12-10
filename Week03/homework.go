package Week03

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func httpServer() error {
	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/user/login", UserLoginHandler)

	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if err != nil {
		fmt.Println("服务器错误")
		return err
	}
	return nil
}

func signhandle(stop chan os.Signal) {
	stop = make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}

func main() {
	var stop chan os.Signal
	g := new(errgroup.Group)

	go signhandle(stop)

	for len(stop) < 1 {
		g.Go(httpServer)
	}

	if err := g.Wait(); err != nil {

	}

}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloWorld!")
}

func UserLoginHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Handler Hello")
	fmt.Fprintf(response, "Login Success")
}
