// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package frontend

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Noapi() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col gap-y-2\"><h1 class=\"mb-4 text-4xl font-extrabold\">No API</h1><p class=\"para\">Organizations contain many systems with useful data and functionality. These systems typically expose APIs to allow other systems to consume their data and functionality. They tend to be \"systems of record\" and are either COTS packages such as ERPs, CRMs, etc or are developed in-house as API-only to be consumed by other frontend \"systems of interaction\" such as websites or mobile apps.</p><p class=\"para\">Many \"systems of interaction\" apps today contain data and functionality that is only  consumed by the app itself. These apps may consume other APIs, but do not offer  data or functionality that is useful to other apps today. When designing these apps  it is tempting to think that the data and functionality in these apps <b>could</b> be  useful to other apps in <b>future</b>, but often this turns out not to be true. </p><p class=\"para\">APIs are complex to design and implement correctly. They require a lot of thought about use cases, edge cases, security, and performance under general use. It is a significant effort to design and implement an API for an app if its only consumer is the app's frontend. Unfortunately Single Page Apps (SPAs) require APIs in order to function, and hence today we have a proliferation of apps with poorly designed APIs that cater only for their frontends. Furthermore, a significant amount of time and effort goes into synchronizing the frontend and API in these apps, with no real benefit other than to make a SPA framework function.</p><h2 class=\"mt-2 text-2xl font-extrabold\">Standalone Apps without APIs</h2><p class=\"para\">Gofs apps have server-side rendered pages. The pages do not need an API to call the backend functionality, they can call it directly. This eliminates the need to design an API or the need to synchronize frontend code with the API. This eliminates a significant amount of complexity and effort in the app development.</p><p class=\"para\">Gofs apps can also expose API endpoints. We recommend that Gofs apps call the backend directly and that the API is exposes as a carefully design general purpose API.  If your app is exclusively an API then using vanilla Go may be a better fit than Gofs.</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
