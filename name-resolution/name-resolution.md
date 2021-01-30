---
title: "Server Name Resolution"
---

Every Harmony homeserver is identified by a "server name." A server name uniquely identifies a server, regardless of whether the resolved IP address changes or not. Two server names pointing to the same IP address would be considered two different homeservers.

A server name is resolved to an IP address and port using the following process:

## When IP Address (`xxx.xxx.xxx.xxx` + optional `:port`)

When the server name is an IP address, this is the resolved IP of the server that requests should be sent to. If not specified, `:port` is 2289.

The `Host` header in requests should be set to the IP literal including the port if specified in the server name.

## When domain and port (`example.com:port`)

When the server name is a hostname with a port, the hostname should be resolved to an IP address using AAAA or A record lookups in the DNS. This IP address will be the one requests are sent to.

The `Host` header in requests should be set to the domain name and given port, not the resolved IP address.

## When domain without port (`example.com`)

When the server name is a host without a port, a request should be made to `https://hostname/_harmony/server`. If the request fails with a 404, resolution of the server name should be done as if it had port 2289 specified using the above method for resolving domains with ports. Otherwise, the request is expected to return a JSON object with the following schema:

```json
{
    "h.server": "",
}
```

`h.server` should be either an IP address or a domain + port. It may not be another domain without a port. `h.server` should be resolved using the above two methods as per usual to an IP address. This IP address will be the one requests are sent to.

The `Host` header in requests should be set to the domain name without port.
