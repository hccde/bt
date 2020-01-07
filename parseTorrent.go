package main

import "container/list"

type TorrentInfo struct {
	fileList list.List
};

func ParseTorrent(str string) TorrentInfo{
	var t TorrentInfo;
	return t;
}
