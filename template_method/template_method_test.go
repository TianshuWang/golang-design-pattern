package template_method

import "testing"

func TestHTTPDownloader(t *testing.T) {
	downloader := NewHTTPDownloader()
	downloader.Download("http://example.com/abc.zip")
}

func TestFTPDownloader(t *testing.T) {
	downloader := NewFTPDownloader()
	downloader.Download("http://example.com/abc.zip")
}
