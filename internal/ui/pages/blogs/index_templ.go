// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package blogs

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Index() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div hx-boost=\"true\" class=\"flex flex-col space-y-4\"><a href=\"/blogs/evergreen\" class=\"flex flex-col items-center rounded-lg border border-gray-200 bg-white shadow hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 md:max-w-xl md:flex-row\"><img class=\"h-96 w-full rounded-t-lg object-cover md:h-auto md:w-48 md:rounded-none md:rounded-s-lg\" src=\"/assets/img/blogs/evergreen.jpg\" alt=\"\"><div class=\"flex flex-col justify-between p-4 leading-normal\"><h5 class=\"mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white\">Evergreen Go</h5><p class=\"mb-3 font-normal text-gray-700 dark:text-gray-400\">Go is the only real option for building modern evergreen apps, without compromising access to talent and developer experience.</p></div></a> <a href=\"/blogs/simplicity\" class=\"flex flex-col items-center rounded-lg border border-gray-200 bg-white shadow hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 md:max-w-xl md:flex-row\"><img class=\"h-96 w-full rounded-t-lg object-cover md:h-auto md:w-48 md:rounded-none md:rounded-s-lg\" src=\"/assets/img/blogs/simplicity.jpg\" alt=\"\"><div class=\"flex flex-col justify-between p-4 leading-normal\"><h5 class=\"mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white\">Keeping Things Simple</h5><p class=\"mb-3 font-normal text-gray-700 dark:text-gray-400\">Simplicity applies to the functionality of the system, the internal structure and code, as well as the externals of the system such as the execution environment and dependencies.</p></div></a> <a href=\"/blogs/libraries\" class=\"flex flex-col items-center rounded-lg border border-gray-200 bg-white shadow hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 md:max-w-xl md:flex-row\"><img class=\"h-96 w-full rounded-t-lg object-cover md:h-auto md:w-48 md:rounded-none md:rounded-s-lg\" src=\"/assets/img/blogs/libraries.jpg\" alt=\"\"><div class=\"flex flex-col justify-between p-4 leading-normal\"><h5 class=\"mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white\">Libraries vs. Frameworks</h5><p class=\"mb-3 font-normal text-gray-700 dark:text-gray-400\">For skilled developers, the choice between libraries and frameworks hinges on the desire for creative freedom and expressiveness in code.</p></div></a> <a href=\"/blogs/nosql\" class=\"flex flex-col items-center rounded-lg border border-gray-200 bg-white shadow hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700 md:max-w-xl md:flex-row\"><img class=\"h-96 w-full rounded-t-lg object-cover md:h-auto md:w-48 md:rounded-none md:rounded-s-lg\" src=\"/assets/img/blogs/nosql.jpg\" alt=\"\"><div class=\"flex flex-col justify-between p-4 leading-normal\"><h5 class=\"mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white\">Sql DB, NoSql Schema</h5><p class=\"mb-3 font-normal text-gray-700 dark:text-gray-400\">The hybrid approach of using Sql databases with schemas where ACID requirements are stored in a relational manner and the remainder stored in a NoSql manner, is an attractive option for modern apps.</p></div></a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
