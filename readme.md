## Advent of Code 2018

These are my solutions to the problems at https://adventofcode.com/2018/

----
**How to run solutions:**  
 
 Each 'day' problem compiles to a standalone binary executable. Here is an example for how to compile and run day01, each other day is similar:
 
1\) Build binary file
 - go get github.com/robkau/advent-of-code/
 - cd ~/go/src/github.com/robkau/advent-of-code/day01
 - go build

2\) Find session cookie to retrieve your unique inputs. You should only need to do this once:
 - Open Chrome
 - Browse to https://adventofcode.com/2018/ and log in
 - Push F12 and go to the network tab
 - Push F5 and select the request named "2018/"
 - Switch to the Cookies tab
 - Copy the value of the **session** cookie

3\) Execute program
 - Set an environment variable named **aoc_session_cookie** equal to the cookie copied in step 2
   - [Linux](https://askubuntu.com/questions/58814/how-do-i-add-environment-variables) | [Windows](https://superuser.com/questions/79612/setting-and-getting-windows-environment-variables-from-the-command-prompt)
 - Run the compiled day01 binary file from step 1
 


*Example Usage:*
 ````cgo
export aoc_session_cookie=xxxxxxxxxxxxxxx
./day01
Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied? :  437
What is the first frequency your device reaches twice? :  655
````