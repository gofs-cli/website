// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package backend

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import . "github.com/gofs-cli/website/internal/ui/components"

func Persistence() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col\"><h1 class=\"mb-4 text-4xl font-extrabold\">Persisting data</h1><h2 class=\"mt-2 text-2xl font-extrabold\">Json serialization</h2><p class=\"para\">We recommend persisting objects using a hybrid relational and nosql model where nosql data is stored as json blobs in a database. You can use different strategies such as pure-relational and you can also use an ORM. We have written a blog post on why we recommend using a hybrid approach <a class=\"link\" href=\"/blogs/nosql\">Sql DB, NoSql Schema</a> that you may find interesting.</p><p class=\"para\">In the file <code>/internal/app/list/list.go</code> lets prepare  the <code>List</code> struct for persistence by adding  json serialization annotations.  </p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(Clipboard, `type List struct {
	ID    string    `+"`"+`json:"id"`+"`"+`
	Name  string    `+"`"+`json:"name"`+"`"+`
	Items []string  `+"`"+`json:"items"`+"`"+`
}
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<h2 class=\"mt-2 text-2xl font-extrabold\">Create a table</h2><p class=\"para\">Add the create table sql statement to the database migration that is performed at server startup. The migrations can be found at <code>/internal/db/migrations/migrations.sql</code>.</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(Clipboard, `CREATE TABLE IF NOT EXISTS lists (
    id TEXT PRIMARY KEY,
    blob JSONB NOT NULL,
);
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<h2 class=\"mt-2 text-2xl font-extrabold\">Database CRUD functions</h2><p class=\"para\">Lets create a new file called <code>/internal/app/list/list_database.go</code> for the database CRUD functions.  </p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(NoClipboard, `mytodo
|--internal
|  |--app
|     |--list
|        |--list.go
|        |--list_database.go
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<p class=\"para\">The database CRUD functions will take a context and transaction, and execute sql commands to perform the functions. Here is an example of <code>/internal/app/list/list_database.go</code>. Note that these functions are private.</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(Clipboard, `package list

import (
	"context"
	"database/sql"
	"encoding/json"
)

func listCreate(ctx context.Context, tx *sql.Tx, list List) error {
	jsonb, err := json.Marshal(&list)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `+"`"+`
		INSERT INTO lists (
			id,
			blob
		) VALUES (
		 	$1, $2
		)`+"`"+`,
		list.ID,
		jsonb,
	)
	return err
}

func listGet(ctx context.Context, tx *sql.Tx, ID string) (*List, error) {
	var data []byte
	err := tx.QueryRowContext(ctx, `+"`"+`
		SELECT blob 
		FROM lists
		WHERE id = $1`+"`"+`,
		id).Scan(&data)
	if err != nil {
		return nil, err
	}
	var list List
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func listUpdate(ctx context.Context, tx *sql.Tx, list List) error {
	jsonb, err := json.Marshal(&list)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, `+"`"+`
		UPDATE lists
		SET 
			blob = $1
		WHERE id = $2`+"`"+`,
		jsonb,
		list.ID)
	return err
}
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<h2 class=\"mt-2 text-2xl font-extrabold\">Update the API</h2><p class=\"para\">Now we can update the API functions in <code>/internal/app/list/list.go</code> to call the database functions.</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(Clipboard, `func NewList(ctx context.Context, tx *sql.Tx, name string) (*List, error) {
	l := &List{
		ID:   uuid.NewString(),
		Name: name,
	}

	err := listCreate(ctx, tx, *l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func AddItem(ctx context.Context, tx *sql.Tx, listID string, item string) error {
	l, err := listGet(ctx, tx, listID)
	if err != nil {
		return err
	}

	// Check if the list is full
	if len(l.Items) >= MaxItems {
		return fmt.Errorf("Todo list is full")
	}

	l.Items = append(l.Items, item)

	err = listUpdate(ctx, tx, *l)
	if err != nil {
		return err
	}

	return nil
}
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
