package template_method

import "fmt"

type Downloader interface {
	Download(url string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("Prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("Finish downloading\n")
}

func (t *template) save() {
	fmt.Print("Default save\n")
}

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("Download %s via HTTP\n", d.uri)
}

func (d *HTTPDownloader) save() {
	fmt.Print("HTTP save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template

	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("Download %s via FTP\n", d.uri)
}
