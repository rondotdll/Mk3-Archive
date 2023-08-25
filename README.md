# Liveton Mk3 [WIP]
### Liveton Mk3 payload generator application source
This payload is most likely detected now, but it may still prove useful for research purposes.

<br>

## What is it?
#### In short, it's Malware for obtaining sensitive information about a target host.

**In full,** Liveton is a custom payload creation tool designed for use in a penetration testing environment to obtain potentially useful information about a target, such as advanced location data, saved chromium data (passwords & cookies), Windows activation information, and a couple other useful bits of general system info. There are also a few extra bells and whistles that were meant to help potentially lower suspicion / make the payload more difficult to analyze like killing the explorer (desktop) process, displaying fake error pop-ups, triggering a Hardstop [BSoD] via NTAPI, and wiping personal file directories. This is the 3rd major evolution of this software, hence Mk3
<br><br><br>
I had plans to implement a GUI & CLI that would make generating payloads significantly easier, however this project ended up getting abandoned mid development in favor of [Shield](https://www.studioseven.dev/)

Feel free to take a peek at the source, below are some screenshots of the non-implemented GUI, and I'll work on adding comments over the coming months.

![image](https://user-images.githubusercontent.com/47403033/226260747-2f5c5843-1a2d-4416-9e7f-9b0bec954f15.png)
![image](https://user-images.githubusercontent.com/47403033/226260894-9e7483ca-d16c-41ae-8dc8-c9dc33a88b7d.png)
![image](https://user-images.githubusercontent.com/47403033/226260982-bf025bc8-d548-45ab-9ddc-73e051129f39.png)

## Usage
> #### Note to self: don't forget to add this part!

## [!] Disclaimer
#### I take NO RESPONSIBILITY for the potential malicious misuse of this tool.
> This software was designed for usage in a secure penetration testing research environment, where all parties are fully aware of the functionality / intent of this software's user.
