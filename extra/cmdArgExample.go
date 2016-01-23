package main

import "os"
import "flag"

var nFlag =  flag.Bool("n",false, `no \n` )

func main(){
	flag.Parse()
	s := ""
	for i:=0 ; i<flag.NArg();i++{
		if(i>0){
			s+=" "
		}
		s+=flag.Arg(i)
	}
	if !*nFlag{
		s+="\n"
	}
	os.Stdout.WriteString(s)
}