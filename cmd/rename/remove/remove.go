package remove

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	dir   = flag.String("dir", "", "Directory path containing the files to rename")
	start = flag.Int("start", 0, "Number of characters to remove from the start of each file name")
	end   = flag.Int("end", 0, "Number of characters to remove from the end of each file name")
)

type RemoveCommand struct{}

func (r *RemoveCommand) Run() {
	flag.Usage = r.Help
	flag.CommandLine.Parse(os.Args[2:])
	if *dir == "" {
		log.Fatal("Please provide a directory path using the -dir flag. Use -h for help.")
	}

	err := filepath.Walk(*dir, func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if file.IsDir() {
			return nil
		}

		filename := file.Name()
		ext := filepath.Ext(filename)
		name := filename[:len(filename)-len(ext)]
		if len(name) <= *start+*end {
			fmt.Printf("Skipping '%s' (name is too short to remove %d characters from start and %d characters from end)\n",
				filename, *start, *end)
			return nil
		}

		newNameWithoutExt := name[*start : len(name)-*end]
		newName := newNameWithoutExt + ext
		oldPath := filepath.Join(*dir, filename)
		newPath := filepath.Join(*dir, newName)

		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Error '%s' => '%s': %v\n", filename, newName, err)
		} else {
			fmt.Printf("'%s' => '%s'\n", filename, newName)
		}
		return nil
	})
	if err != nil {
		fmt.Print("failed to read folder")
	}
}

func (r *RemoveCommand) Help() {
	fmt.Println("usage: rename remove [options]")
	flag.PrintDefaults()
}
