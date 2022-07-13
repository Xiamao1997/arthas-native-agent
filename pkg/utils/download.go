package utils

import (
	"archive/zip"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

// @author Xiamao1997 2022-07-04

const (
	ARTHAS_LATEST_VERSIONS_URL = "https://arthas.aliyun.com/api/latest_version"
	ARTHAS_VERSIONS_URL        = "https://arthas.aliyun.com/api/versions"
	ARTHAS_DOWNLOAD_URL        = "https://arthas.aliyun.com/download"
)

// Get arthas latest release version
func GetLatestReleaseVersion() (string, error) {
	response, err := http.Get(ARTHAS_LATEST_VERSIONS_URL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	version := string(body)
	if len(version) > 100 {
		return "", errors.New("failed to get arthas latest release version")
	}
	return version, nil
}

// Get arthas remote versions
func GetRemoteVersions() ([]string, error) {
	response, err := http.Get(ARTHAS_VERSIONS_URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	versionsStr := string(body)
	if len(versionsStr) > 100 {
		return nil, errors.New("failed to get arthas remote versions")
	}
	versions := strings.Split(versionsStr, "\r\n")
	if versions[len(versions)-1] == "" {
		versions = versions[0 : len(versions)-1]
	}
	return versions, nil
}

// Download the arthas package file and extract the components
func DownloadArthasAndExtract(version string, savePath string) error {
	tempFile, err := ioutil.TempFile(savePath, "arthas_temp")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	url := ARTHAS_DOWNLOAD_URL + "/" + version + "?mirror=aliyun"
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	contentLength := response.Header.Get("Content-Length")
	fileSize, err := strconv.ParseInt(contentLength, 10, 64)
	if err != nil {
		return err
	}

	size, err := io.Copy(tempFile, response.Body)
	if err != nil {
		return err
	}
	if size != fileSize {
		return errors.New("failed to download arthas")
	}
	logrus.Infoln("Download success, total size:", size, "Byte.")

	reader, err := zip.OpenReader(tempFile.Name())
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, f := range reader.File {
		filePath := path.Join(savePath, f.Name)
		info := f.FileInfo()
		if info.IsDir() {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}
		open, err := f.Open()
		if err != nil {
			return err
		}
		create, err := os.Create(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(create, open)
		if err != nil {
			return err
		}
		logrus.Infoln("Extract file:", f.Name)
	}
	logrus.Infoln("Extract success, total file numbers:", len(reader.File))

	return nil
}
