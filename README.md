# GoBitTorrent

Building a BitTorrent client in GoLang while getting familiar with the BitTorrent protocol spec (version 0e08ddf84d8d3bf101cdf897fc312f2774588c9e).

## BitTorrent : 
BitTorrent is a peer-to-peer protocol (p2p) used for file distribution on the internet. Based on the protocol [specification](https://www.bittorrent.org/beps/bep_0003.html), it identifies content by URL and is designed to integrate seamlessly with the web.  Its advantage over plain HTTP is that when multiple downloads of the same file happen concurrently, the downloaders upload to each other, making it possible for the file source to support very large numbers of downloaders with only a modest increase in its load.

## Where to find peers
In order to keep the projet legal and stay away from piracy, I will be using a Linux Debian ISO as a test file to be downloaded. Therefore, to find peers that wille allow me to download this file, I will rely on official [Debian BitTorent Tracker](http://bttracker.debian.org:6969/)

## .torrent file
A . torrent file is afile that contains the description of a file that will be downloaded in addition to informations for connecting to a tracker.
This file is encoded in a format named **Bencode**. This format is used primarily by the BitTorrent protocol. It is designed to be simple, quick to parse, and easy to use in resource-constrained environments. Bencode supports four types of data: integers, strings, lists and dictionaries.

**Examples**
- Integers :
    - i100e is 100
    - i-25e is -25
 
- Strings :
    - 12:Hello World! is "Hello World!" (12 is for the number of characters)

- Lists :
    - l3:cat5:mousee is ["cat", "mouse"] (The list is between l and e)
 
- Dictionaries :
    - d6:animall3:cat5:mousee6:numberi2ee is {"animal": ["cat", "mouse"], "number": 2} (The dictionnary is between d and e)
