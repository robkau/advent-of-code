#### Day 1: Repeatedly applying frequency changes to a laser
https://adventofcode.com/2018/day/1

----

**How to run:**  
 
1\) Build binary file
 - go get -u github.com/robkau/advent-of-code/day01
 - go build  
  

2\) Set environment variable to retrieve your unique inputs:  
(Using Chrome for an example)
 - Browse to https://adventofcode.com/2018/day/1
 - Push F12 and go to the network tab
 - Push F5 and select the request named "1"
 - Switch to the Cookies tab
 - Copy the value of the **session** cookie

3\) Execute program
 - Set an environment variable named **aoc_session_cookie** equal to the cookie copied in step 2
   - [Linux](https://askubuntu.com/questions/58814/how-do-i-add-environment-variables) | [Windows](https://superuser.com/questions/79612/setting-and-getting-windows-environment-variables-from-the-command-prompt)
 - Run the compiled day01 binary file from step 1