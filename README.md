# Columbus

## Name

*columbus* - insert domains to Columbus.

## Description

Using *columbus* insert valid domain (`RCODE` is `0`) from question to [Columbus](https://columbus/elmasy.com).
This plugin uses a channel to buffer the domains to not degrade the performance.

Setting `COREDNS_COLUMBUS_BUFFSIZE` environment variable controls how many domain can the underlying channel hold. Default: `100000`.

Setting `COREDNS_COLUMBUS_WORKER` environment variable controls the number workers to concurrent insert domains. Default: `number of cpu`.

## Syntax

~~~ txt
columbus
~~~

## Examples

Enable plugin:

~~~ corefile
. {
    columbus
    whoami
}
~~~
