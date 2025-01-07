// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package prereqs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import . "github.com/gofs-cli/website/internal/ui/components"

func Prereqs() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col gap-y-2\"><h1 class=\"mb-4 text-4xl font-extrabold\">Prerequisites</h1><h2 class=\"mt-2 text-2xl font-extrabold\">Familiarity with Go</h2><p class=\"para\">You should have some experience with Go. There are some good tutorials <a class=\"link\" href=\"https://go.dev/doc/tutorial/getting-started\">here</a> and <a class=\"link\" href=\"https://tour.golang.org/welcome/1\">here</a> which you can use to get started if you are new to Go. These docs assume you are familiar with Go.</p><h2 class=\"mt-2 text-2xl font-extrabold\">Go toolchain</h2><p class=\"para\"><a class=\"link\" href=\"https://go.dev/dl/\">Download go</a> for your OS and install it. The official instructions are <a class=\"link\" href=\"https://go.dev/doc/install\">here</a>.</p><p class=\"para mt-4\">If you are using Windows, the installer should setup the path for you. Ensure that `go` is on your path, you should be able to type the command below in `terminal` or `cmd/powershell` and see the installed go version.")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Cmdline("go version").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</p><p class=\"para mt-4\">Once go is installed, you can add the GOPATH to your system path. GOPATH is where go will install any tools installed using `go install`, This is typically your `$HOME/go/bin`, this should be added to your system PATH to allow you to run installed tools.</p><h2 class=\"mt-2 text-2xl font-extrabold\">Gofs cli</h2><p class=\"para\">Once go is installed, you can install the gofs cli by running:")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Cmdline("go install github.com/gofs-cli/gofs@latest").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</p><h2 class=\"mt-2 text-2xl font-extrabold\">Local Development</h2><p class=\"para\">The Gofs template includes a Makefile for convenience. You will need to install a make tool for example  <a class=\"link\" href=\"https://www.gnu.org/software/make/\">GNU Make</a></p><p class=\"para\">Once make is setup and you have initialized a project you can install all required dependencies with:")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Cmdline("make deps").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "This will install air, templ, htmx and alpine.</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate