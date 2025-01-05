// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package frontend

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import . "github.com/gofs-cli/website/internal/ui/components"

func Handler() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col gap-y-2\"><h1 class=\"mb-4 text-4xl font-extrabold\">Adding a Handler</h1><p class=\"para\">Handlers are responsible to accepting incoming requests, acting on them  and constructing a response back to the client. Handlers should be called with relevant server objects they need, for example the database connection, storage bucket, etc. Handlers are responsible for wrapping operations in transactions and handling errors.</p><h2 class=\"mt-2 text-2xl font-extrabold\">Example</h2><p class=\"para\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(Clipboard, `func Index(db db.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var l []list.List
		var err error
		err = db.Transaction(r.Context(), func(tx *sql.Tx) error {
			l, err = list.GetLists(r.Context(), tx)
			return err
		})
		if err != nil {
			http.Error(w, "error getting lists", http.StatusInternalServerError)
			return
		}
		templ.Handler(ui.Index(l)).ServeHTTP(w, r)
	})
}
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
