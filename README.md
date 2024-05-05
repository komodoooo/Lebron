# Lebron
![image](https://github.com/komodoooo/Lebron/assets/68278515/85f25bab-f83b-401b-ada3-8c3afc11c32f)

A Chromium-based info stealer for windows, sends credentials via a discord webhook

The program currently supports the following broswers:
* Chrome
* Edge
* Brave
* Opera / GX

###### * only the stable versions of these broswers are actually supported
# Compile
First, edit the code and add your discord webhook inside che constant **webhook**.

Your webhook must be hex encoded. _(without 0x)_<br>
This is just a easy information hiding step to not let someone find your webhook analyzing the string in the executable at first.

Before compiling you must setup the dependencies, run _`make setup`_<br>
After that you can finally run  _`make`_

## Notes
I based myself on various detailed reads i found online.<br>
I just wrote this for learning & fun purposes.<br>
