# client-bgs

This is the client part of a service that can change your background. 
In this repo gets the image bits via a TCP connection and creates a temp file and saves it.
This requires you to install [feh](https://feh.finalrewind.org/) as a prerequest

passed to the [server side](https://github.com/praveenmahasena647/server-bgs).
Also this server is bind to my [vim-config](https://github.com/praveenmahasena647/nvim) by a simple key bind you could trigger this app.

## for dev purposes
```sh
https://github.com/praveenmahasena647/client-bgs.git . && make run
```
