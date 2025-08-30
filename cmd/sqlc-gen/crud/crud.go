package crud

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/Long-Software/Sonality/internal/utils"
)

var (
	name = flag.String("name", "user", "Model name for generating the SQLC sql file")
	dir  = flag.String("dir", ".", "Directory to store the SQLC sql file")
)

type CRUDCommand struct{}

func (c *CRUDCommand) Run() {
	flag.Usage = c.Help
	flag.CommandLine.Parse(os.Args[2:])

	file, err := os.Create(fmt.Sprintf("%s/%s", *dir, *name+".sql"))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	tmpl, data := generateTemplate(*name)
	t, err := template.New("sqlc").Parse(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	err = t.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	fmt.Printf("%s created inside %s successfully\n", *name, *dir)
}

func (c *CRUDCommand) Help() {
	fmt.Println("usage: sqlc-gen crud [options]")
	flag.PrintDefaults()
}

type TemplateData struct {
	Capital       string
	Plural        string
	CapitalPlural string
}

func generateTemplate(model string) (string, TemplateData) {
	tmpl := `-- name: Create{{.Capital}} :one
INSERT INTO {{.Plural}} (column1, column2)
VALUES ($1, $2) RETURNING *;

-- name: Delete{{.Capital}} :exec
DELETE 
FROM {{.Plural}}
WHERE id = $1;

-- name: Get{{.Capital}} :one
SELECT *
FROM {{.Plural}}
WHERE id = $1 
LIMIT 1;

-- name: List{{.CapitalPlural}} :many
SELECT *
FROM {{.Plural}};

-- name: Update{{.Capital}} :one
UPDATE {{.Plural}}
SET column1 = $2, column2 = $3
WHERE id = $1 RETURNING *;
`
	return tmpl, TemplateData{
		Capital:       utils.Capitalize(model),
		Plural:        utils.Pluralize(model),
		CapitalPlural: utils.Capitalize(utils.Pluralize(model)),
	}
}
