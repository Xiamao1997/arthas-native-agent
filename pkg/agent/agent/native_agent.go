package agent

import (
	"github.com/Xiamao1997/arthas-native-agent/pkg/utils"
	"github.com/sirupsen/logrus"
	"os"
	"os/user"
	"path"
)

// @author Xiamao1997 2022-06-24

// NativeAgent communicates with the work node and the tunnel server.
type NativeAgent struct {
	tunnelServerIp   string
	tunnelServerPort string
	nativeAgentName  string
	arthasHomeDir    string
	arthasVersion    string
	stop             chan error
}

func NewNativeAgent(ip string, port string, name string, home string, version string) *NativeAgent {
	return &NativeAgent{
		tunnelServerIp:   ip,
		tunnelServerPort: port,
		nativeAgentName:  name,
		arthasHomeDir:    home,
		arthasVersion:    version,
		stop:             nil,
	}
}

// Start the NativeAgent object
func (na *NativeAgent) Start() error {
	logrus.Infoln("Starting the NativeAgent:", na.nativeAgentName)

	logrus.Infoln("Checking arthas home.")
	if err := na.checkArthasHome(); err != nil {
		return err
	}

	// Wait for stop event
	select {
	case ch := <-na.stop:
		na.Stop()
		return ch
	}
}

// Stop the NativeAgent object
func (na *NativeAgent) Stop() error {
	logrus.Infoln("Stopping the NativeAgent:", na.nativeAgentName)
	return nil
}

// Check the arthas home
func (na *NativeAgent) checkArthasHome() error {
	if na.arthasHomeDir != "" && na.verifyArthasJar() == nil {
		return nil
	}

	// Create arthas home directory and download arthas
	current, err := user.Current()
	if err != nil {
		return err
	}
	userHomeDir := current.HomeDir

	arthasLibDir := path.Join(userHomeDir, ".arthas", "lib")
	if _, err := os.Stat(arthasLibDir); os.IsNotExist(err) {
		err := os.MkdirAll(arthasLibDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	if na.arthasVersion != "" {
		versions := []string{}
		if versions, err = utils.GetRemoteVersions(); err != nil {
			return err
		}
		find := false
		for _, v := range versions {
			if v == na.arthasVersion {
				find = true
				break
			}
		}
		if !find {
			logrus.Errorln("The specified arthas version is invalid, use the latest version.")
			if na.arthasVersion, err = utils.GetLatestReleaseVersion(); err != nil {
				return err
			}
		}
	} else {
		logrus.Infoln("There is no specified arthas version, use the latest version.")
		if na.arthasVersion, err = utils.GetLatestReleaseVersion(); err != nil {
			return err
		}
	}

	na.arthasHomeDir = path.Join(arthasLibDir, na.arthasVersion, "arthas")
	if _, err := os.Stat(na.arthasHomeDir); os.IsNotExist(err) {
		err := os.MkdirAll(na.arthasHomeDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if na.verifyArthasJar() == nil {
		return nil
	}

	logrus.Infoln("Downloading arthas from remote server, version:", na.arthasVersion)
	if err := utils.DownloadArthasAndExtract(na.arthasVersion, na.arthasHomeDir); err != nil {
		return err
	}

	return nil
}

// Verify the arthas jar package
func (na *NativeAgent) verifyArthasJar() error {
	fileList := []string{"arthas-core.jar", "arthas-agent.jar", "arthas-spy.jar"}
	for _, file := range fileList {
		filePath := path.Join(na.arthasHomeDir, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
