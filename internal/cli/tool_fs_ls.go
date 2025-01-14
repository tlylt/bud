package cli

import (
	"context"
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/livebud/bud/framework"
)

type ToolFsLs struct {
	Flag *framework.Flag
	Path string
}

func (c *CLI) ToolFsLs(ctx context.Context, in *ToolFsLs) error {
	// Generate bud files
	generate := &Generate{Flag: in.Flag}
	if err := c.Generate(ctx, generate); err != nil {
		return err
	}

	// Read the directory out
	des, err := os.ReadDir(path.Clean(in.Path))
	if err != nil {
		return err
	}
	// Directories come first
	sort.Slice(des, func(i, j int) bool {
		if des[i].IsDir() && !des[j].IsDir() {
			return true
		} else if !des[i].IsDir() && des[j].IsDir() {
			return false
		}
		return des[i].Name() < des[j].Name()
	})
	// Print out list
	for _, de := range des {
		name := de.Name()
		if de.IsDir() {
			name += "/"
		}
		fmt.Fprintln(c.Stdout, name)
	}
	return nil
}
