# WhereIs

Have you ever wondered in which AS or Public Cloud Region your current machine or any remote machines
is homed?
If so, `whereis` got you covered!

Capabilites:

- DNS lookups, if needed
- IP --> AS lookup via
   - [RADb](https://www.radb.net/)
   - RDAP / WHOIS information
- IP --> Public Cloud Region:
  - [AWS](https://ip-ranges.amazonaws.com/ip-ranges.json)
  - [Google Cloud](https://www.gstatic.com/ipranges/cloud.json)
  - [Azure](https://www.microsoft.com/en-us/download/details.aspx?id=56519)
- DeNATing of your own IP address\[es\] using [STUN](https://en.wikipedia.org/wiki/STUN)