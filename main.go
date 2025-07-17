/* MIT License

* Copyright (c) 2025 ᥇ꪶꪖᥴᛕᦓᥴꪗꪻꫝꫀ

* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:

* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.

* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE. */

package main

import ("fmt"
        "os")

        
func main(){

    commandArgs := os.Args[1:]
        
    for _, arg := range commandArgs{

        if arg == "-h"|| arg == "-help" || arg == "--help"{
            help()
            
    	} else if arg == "-all"{
    		fmt.Println("Arg used:", arg)

    	} else if arg == "-rs" {
    		fmt.Println("Arg used:", arg)
    		
    	} else if arg == "-ap"{
    		fmt.Println("Arg used:", arg)

    	} else {
    		fmt.Println("unknown arg:", arg)
    	}
    }
}

func help(){

//this function contains documentation, usage with examples, of the tool

fmt.Println( "USAGE:\n",

"--\n",
"Show all ip addresses available to investigate (in local network) with appropriate info.\n",

"Eg:\n",
" pJ -all\n",
"\n",

"Sniffing packets.\n",
"Eg:\n", 
"pJ -sniff <ip addr> -p <port no.>\n",
"\n",
"or\n",
"\n",
"Eg:\n",
" pJ -sp <ip addr> -p <port no.>\n",

"\n", 
"To sniff all packets it can sniff.\n",
"Eg:\n",
"pJ -sniff <ip addr> -all\n",
"\n",
"or\n",
"\n",
"Eg:\n",
"pJ -sp <ip addr> -all\n",
"\n",
"To check.\n",
"running services.\n",

"Eg:\n",
"pJ -rs <ip addr>\n",
"\n",
"available ports.\n",

"Eg:\n", 
"pJ -ap <ip addr>\n")
}
