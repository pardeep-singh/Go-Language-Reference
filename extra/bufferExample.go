package main

import (
 "bufio";
 "os";
 "fmt";
)

func main() {
	fmt.Fprintf(os.Stdout,"%s, ","Hello");
	buf := bufio.NewWriter(os.Stdout);
	fmt.Fprintf(buf,"%s\n","world");
	buf.Flush();
}