package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"github.com/Long-Software/Sonality/internal/consts"
)

func HasModFile(dir string) bool {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == consts.MOD_FILE {
			return filepath.SkipDir
		}
		return nil
	})
	return err == filepath.SkipDir
}

func MakeModFile(dir string, module string) error {
	path := filepath.Join(dir, consts.MOD_FILE)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	output := fmt.Sprintf("module %s\n\ngo %s", module, GetGOVersion())
	_, err = file.WriteString(output)
	return err
}

func AddPackageToActive(work_file, module string) {
	// var config  consts.WorkConfig
	// err := LoadJsonConfig(&config, work_file)
	// if err != nil {
	// 	log.Fatalf("fail to load config file: %v", err)
	// 	return
	// }
	// if stringInSlice(module, config.Packages.Active) {
    //     return
    // }
	// for i, v := range config.Packages.InActive {
    //     if v == module {
    //         config.Packages.InActive = append(config.Packages.InActive[:i], config.Packages.InActive[i+1:]...)
    //     }
    // }
	// config.Packages.Active = append(config.Packages.Active, module)
	// err = WriteJsonConfig(config, work_file)
	// if err != nil {
	// 	log.Fatalf("unable to write to file: %v", err)
	// }
}

// TODO: implement the add to service 
func AddServicesToActive(work_file, module string){}
