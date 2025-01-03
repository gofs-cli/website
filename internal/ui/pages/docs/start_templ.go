// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package docs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import . "github.com/gofs-cli/website/internal/ui/components"

func Start() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col gap-y-2\"><h1 class=\"mb-4 text-4xl font-extrabold\">Start a Project</h1><h2 class=\"mt-2 text-2xl font-extrabold\">Initialize</h2><p class=\"para\">Provide gofs with a go module name such as <code>github.com/myorg/myapp</code></p><p class=\"para\">Navigate to the folder where you want to initialize your project and run</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Cmdline("gofs init github.com/[your org]/[your app]").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<p class=\"para\">Or if you have not created the folder then provide a folder name to gofs and it will be created</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Cmdline("gofs init github.com/[your org]/[your app] [your folder]").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<h2 class=\"mt-2 text-2xl font-extrabold\">Folder Structure</h2><p class=\"para\">Go is not opinionated about folder structure however by convention go code should be kept close to where it is used. In Gofs we use the default folder structure below.</p><p class=\"para\">NB: Gofs is not a framework so you can use or modify this as you see fit. </p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(NoClipboard, `root
|--.github              Github actions configuration
|--.gofs                Gofs configuration
|--.vscode              Vscode configuration
|--cmd                  Gofs command
|--docker               Docker folder including default Dockerfile
|--internal
|  |--app               Backend (see backend docs)
|  |--auth              Authentication functions
|  |--config            App config
|  |--db                Database functions
|  |--server            
|  |  |--assets         Static assets
|  |  |--handlers       Generic handlers
|  |  |--logging        Logging functions
|  |  |--telemetry      Telemetry functions
|  |  |--middleware.go  Server middleware
|  |  |--routes.go      Server routes
|  |  |--server.go      Server code
|  |--ui                Frontend (see frontend docs)
|--scripts              Build Scripts
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<h2 class=\"mt-2 text-2xl font-extrabold\">Dependencies</h2><p class=\"para\">Run make for the first time to get all dependencies and build the default app</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Cmdline("make").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
