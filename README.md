# hips - Hiding In Plain Sight

A command line tool for hiding a file (payload) onto a carrier (image).

## How does it work?

Running `hips --hide testdata/secret.txt` attaches the bytes of the `secret.txt` onto the embedded carrier image file:

| Before                         | After                                            |
| ------------------------------ | ------------------------------------------------ |
| Carrier **without** secret.txt | Carrier **with** secret.txt                      |
| ![carrier](forest.jpg)         | ![carrier with payload](testdata/secret.txt.jpg) |

Can you see the difference? I bet you can't. Only if you look at the filesize.

If you want to get the original data (payload) back, you can run `hips --uncover testdata/secret.txt.jpg`

## How to use it?

If you have Go, you can easily build the project by running `go build .`.
