
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
    a := []byte("a")  
    
    for _, arg := range commandArgs{

    	if arg == "-all"{
    		fmt.Println("Arg used:", arg)

    	} else if arg == "-rs" {
    		fmt.Println("Arg used:", arg)
    		
    	} else if arg == "-ap"{
    		fmt.Println("Arg used:", arg)
    	} else {
    		fmt.Println("unknown arg:", arg)
    	}
    }
    fmt.Println(commandArgs, a)
}

func help(){

// this function contains documentation, usage with examples, of the tool.

/*

USAGE:
--
**Show all ip addresses available to investigate (in local network) with appropriate info.**


Eg:
 pJ -all

**Sniffing packets.**
Eg: 
pJ -sniff <ip addr> -p <port no.>

or
 
Eg:
 pJ -sp <ip addr> -p <port no.>

 
**To sniff all packets it can sniff.**
Eg:
pJ -sniff <ip addr> -all

or 
  
Eg:
pJ -sp <ip addr> -all

**To check.**
running services.

Eg:
pJ -rs <ip addr>

available ports.

Eg: 
pJ -ap <ip addr>

*/
}
