// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package blogs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Evergreen() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col gap-y-2 md:w-1/2\"><img class=\"h-96 w-full rounded-lg object-cover md:h-auto md:rounded-none\" src=\"/assets/img/blogs/evergreen.jpg\" alt=\"\"><h1 class=\"mb-4 text-4xl font-extrabold\">Evergreen Go</h1><p class=\"para\">Large organizations have hundreds or even thousands of custom systems.  Many of these systems were built decades ago and codify the business as it was decades ago.  Often these systems have been underinvested in for years and have become an obstacle for the  organization to change.</p><h2 class=\"text-2xl font-extrabold\">Technology Instability</h2><p class=\"para\">Custom systems are not typically kept up to date because as time passes, the technology  on which these systems are built also goes out-of-date. Custom systems need periodic  refactoring or rewriting as technology changes. Not keeping up with technology changes  leads to lower developer productivity and eventually an inability to attract talent  (developers want to build their CVs with the latest technologies), and vendors eventually  withdraw support of the old technology, blocking developers from working on the systems  at all. For example, Microsoft is requires customers with .Net Framework applications to  rewrite them on .Net Core. Updating a Python 2 application to Python 3 to access new  features of Python 3 involves a significant rewrite. Java versions are historically  incompatible, and so on.</p><h2 class=\"text-2xl font-extrabold\">Upgrade Cycles</h2><p class=\"para\">With most technologies this happens in a one-to-two-year cycle and organizations fail  to plan and budget for these cycles properly. This is partially because this refactoring  effort does not add new features and can be seen as a large spend for which there appears  to be no benefit. Organizations that do budget for these upgrades can often end up  allocating significant amounts of their IT budget just to ‘stand still’ – we have seen  organizations spend 80% of their budget on ‘standing still’ leaving only 20% for new  features. Often, finance teams ‘kick the can’ down the road when it comes to these  large ‘standing still’ cost items leading to legacy debt problems that eventually become  insurmountable. Eventually developers are unable to work on custom systems and they drift  away from how the business works. Considering these systems were built to support the  unique and differentiated parts of a business, these systems eventually become an anchor  on the business’ ability to turn up in unique and differentiated ways in the market.</p><h2 class=\"text-2xl font-extrabold\">Evergreen Is Possible</h2><p class=\"para\">The idea of evergreen code has been around for decades. The C community have been  writing evergreen apps for nearly 40 years – a C program from 40 years ago will compile on a  modern C compiler and run as it was designed on modern hardware. C is a systems  programming language, and the C ecosystem is not suitable for modern business application  development.  </p><h2 class=\"text-2xl font-extrabold\">Evergreen Go</h2><p class=\"para\">Go (designed by the creators of C) was designed to be a modern business programming  language, and the Go team have made a commitment to backward compatibility. Go has  gained wide adoption, and the tooling ecosystem is rich and meets the expectations of a  modern developer. This is why Go is the only real option for building modern evergreen  apps. </p><p class=\"para\">Using Go for frontend and backend brings go's backward compatibility commitment to the entire app. This means an app that is written today will compile in the future with no  changes. Go evergreen apps should be carefully crafted to minimize direct dependencies,  which means these apps have minimal supply chain issues such as obsolescence, patching,  or supply chain vulnerabilities. The few dependences used in gofs apps have been stable  for many years and have no transitive dependencies. </p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
