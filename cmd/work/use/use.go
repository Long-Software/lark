package use

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"github.com/Long-Software/Sonality/internal/consts"
	"github.com/Long-Software/Sonality/internal/utils"
	"strings"
)

var (
	out = flag.String("out", "work.json", "Config file for syncing with the workspace")
	mod = flag.String("mod", "", "module to add to the packages to be added to the workspace")
)

type UseCommand struct{}

func (u *UseCommand) Run() {
	flag.Usage = u.Help
	flag.CommandLine.Parse(os.Args[2:])
	if *mod == "" {
		flag.Usage()
		os.Exit(2)
	}

	var config consts.WorkConfig
	err := utils.LoadJsonConfig(&config, *out)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
		return
	}
	module_dir := filepath.Dir(*mod)
	// base, isPkg, isService, isApp := handleModuleDir(module_dir, config.Packages.Path, config.Services.Path)

	if !utils.HasDir(module_dir) {
		err = os.MkdirAll(module_dir, os.ModePerm)
		if err != nil {
			log.Fatalf("faild to create directory: %v", err)
		}
	}
	// if !isPkg && !isService {
	// 	log.Fatalf("faild to create mod file in packages or services directory: %v", err)
	// }
	var module string
	// if isPkg {
	// 	module = fmt.Sprintf("%s/%s/%s", config.Name, config.Packages.Path, base)
	// }
	// if isService {
	// 	module = fmt.Sprintf("%s/%s/%s", config.Name, config.Services.Path, base)
	// }
	// if isApp {
	// 	module = fmt.Sprintf("%s/%s/%s", config.Name, config.Apps.Path, base)
	// }
	if !utils.HasModFile(module_dir) {
		err = utils.MakeModFile(module_dir, module)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	}
	// if isPkg {
	// 	utils.AddPackageToActive(*out, base)
	// } else if isService {
	// 	utils.AddServicesToActive(*out, base)
	// }
}

func (u *UseCommand) Help() {
	fmt.Println("usage: work use [options]")
	flag.PrintDefaults()
}

func trim(path string) string {
	prefixes := []string{"./", "/", ".\\", "\\"}
	for _, prefix := range prefixes {
		path = strings.TrimPrefix(path, prefix)
	}
	suffixes := []string{"/", "\\"}
	for _, suffix := range suffixes {
		path = strings.TrimSuffix(path, suffix)
	}
	return path
}

func createModuleName(name, module_dir string) string {
	return path.Join(append([]string{name}, strings.Split(module_dir, string(filepath.Separator))...)...)
}

func handleModuleDir(module_dir, package_dir, service_dir string) (string, bool, bool, bool) {
	module := trim(module_dir)
	if strings.HasPrefix(module, package_dir) {
		return filepath.Base(module), true, false, false
	}
	if strings.HasPrefix(module, service_dir) {
		return filepath.Base(module), false, true, false
	}
	return "", false, false, false
}
