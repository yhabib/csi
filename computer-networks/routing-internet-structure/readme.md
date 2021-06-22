# Routing, and the Structure of the Internet

## Traceroute

> Program used to determine the path between two hosts on the Internet.

### How it works?

It exploits the TTL field in the IP header by incrementing it from 1 till it reaches the destination.
It uses ICMP package(error message protocol) to let the sender know what happened

### Questions

* Sometimes nothing happens. It triggers the initial PING but no response??
* It gets stucked after the sixth hop
* If I trigger traceroute some traffic comes in ?

### Resources

* [rfc792](https://datatracker.ietf.org/doc/html/rfc792)
* [Video](https://www.youtube.com/watch?v=75yKT3OuE44)
