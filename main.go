
package main
import ("fmt";"os";"strconv";"time")
func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr,"Usage: countdown <seconds>"); os.Exit(1) }
	n,err:=strconv.Atoi(os.Args[1])
	if err!=nil || n<1 { fmt.Fprintln(os.Stderr,"Invalid number"); os.Exit(1) }
	for i:=n; i>0; i-- {
		fmt.Printf("\r%d... \033[K", i)
		time.Sleep(1*time.Second)
	}
	fmt.Println("\rDone!    ")
}

