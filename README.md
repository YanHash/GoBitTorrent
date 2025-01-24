# GoBitTorrent

Building a BitTorrent client in GoLang while getting familiar with the BitTorrent protocol spec (version 0e08ddf84d8d3bf101cdf897fc312f2774588c9e).

## BitTorrent : 
BitTorrent is a peer-to-peer protocol (p2p) used for file distribution on the internet. Based on the protocol [specification](https://www.bittorrent.org/beps/bep_0003.html), it identifies content by URL and is designed to integrate seamlessly with the web.  Its advantage over plain HTTP is that when multiple downloads of the same file happen concurrently, the downloaders upload to each other, making it possible for the file source to support very large numbers of downloaders with only a modest increase in its load.
