# Columbus

## Name

*columbus* - insert domains to Columbus.

## Description

Using *columbus* insert valid domain (`RCODE` is `0`) from question to [Columbus](https://columbus/elmasy.com).
This plugin uses a channel to buffer the domains to not degrade the performance.

If `cache` is enabled, cached records are not inserted (*columbus* comes after *cache*).

## Syntax

~~~ txt
columbus [BUFF] [WORKERS]
~~~

- **BUFF**: Optional argument to set the underlying channel size (default: *10000*).
- **WORKERS**: Optional argument to set the number of insert workers (default: *number of proc*).

The **BUFF** and **WORKER** needs to be set once.

## Examples

Enable with default settings:

~~~ corefile
. {
    columbus
    forward . 1.1.1.1
}
~~~

Set buffer size to 100:

~~~ corefile
. {
    columbus 100
    forward . 1.1.1.1
}
~~~

Set buffer size to 100 with 4 workers:

~~~ corefile
. {
    columbus 100 4
    forward . 1.1.1.1
}
~~~