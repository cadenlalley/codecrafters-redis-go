# Future Plans


## Expiry
* Expiry should be better handled. Right now, a time.AfterFunc is spun off in a new go routine for every single item that is set with an expiry.

