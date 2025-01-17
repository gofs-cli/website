// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package frontend

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import . "github.com/gofs-cli/website/internal/ui/components"

func Page() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"flex flex-col gap-y-2\"><h1 class=\"mb-4 text-4xl font-extrabold\">Adding a Page</h1><h2 class=\"mt-2 text-2xl font-extrabold\">The internal/ui folder</h2><p class=\"para\">Frontend pages are located in the <code>/internal/ui/pages</code> folder.</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(NoClipboard, `root
|--internal
|  |--ui
|     |--components     // shared components
|     |--pages          // frontend pages
|     |--index.templ    // the app's main template page 
|     |--handlers.go    // root handlers
|     |--style.css      // custom style sheet (if you installed tailwind)
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<p class=\"para\">The pages folder structure should mirror the structure of the website. Keep handler and page components close to where they are used. </p><h2 class=\"mt-2 text-2xl font-extrabold\">Indirection in code</h2><p class=\"para\">We recommend keeping indirection in pages to a minimum. This means that sometimes it is better to repeat code in pages than to create a new component that hides the code on a page. There are many tools that will help you manage/refactor repeated code such as find-and-replace,  multi-cursors, and so on. The goal is to make the code as easy to read and maintain as possible. Sometimes this principle of simplicity goes against the  principle of Do-Not-Repeat-Yourself (DRY), but on balance you gain more than you loose.</p><p class=\"para\">For example, this code is much easier to read and maintain than the same code with indirection shown below it.")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(NoClipboard, `templ Page() {
    <h2 class="mt-2 text-2xl font-extrabold">New section</h2>
}
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(NoClipboard, `templ Page() {
    @SomeHeadingComponent(SomePropsComponent{
        Text: "New section",
        Size: "2xl",
        Font: "extrabold",
    })
}
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</p><h2 class=\"mt-2 text-2xl font-extrabold\">Example</h2><p class=\"para\">Lets continue the simple todo list app (mytodo) we started in backend <a class=\"link\" href=\"/docs/backend/module\">here</a>. The design of mytodo frontend is a single page showing all the todo lists and their items. </p><h3 class=\"text-2xl font-extrabold\">Display todo lists</h3><p class=\"para\">Lets start by modifying the <code>/internal/ui/pages/index.templ</code> to  take the todo lists as a parameter and display them.")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(NoClipboard, `mytodo
|--internal
|  |--ui
|     |--pages
|        |--index.templ
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(Clipboard, `package ui

import "github.com/myorg/mytodo/internal/app/list"

// add lists as a parameter to the Index templ
templ Index(lists []list.List) {
	<!DOCTYPE html>
	<html lang="en">
        ... header and script tags
		<body hx-ext="response-targets" hx-target-error="this">
			<main class="flex w-full flex-col">
				// delete the example code in the main tag and add the following
				for _, l := range lists {
					<div class="flex flex-col gap-y-2">
						<h2 class="text-2xl font-extrabold">{ l.Name }</h2>
						<div class="flex flex-col gap-y-2">
							for _, item := range list.Items {
								<div class="flex items-center gap-x-2">
									<span>{ item }</span>
								</div>
							}
						</div>
					</div>
				}
			</main>
		</body>
	</html>
}
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</p><h3 class=\"text-2xl font-extrabold\">Add buttons</h3><p class=\"para\">Lets include buttons to add todo lists and items. Note that these use <code>hx-post</code> to end points that we will create in the next section and return the HTML to be swapped in.")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CodeBlock(Clipboard, `package ui

import "github.com/myorg/mytodo/internal/app/lists"

templ Index(lists []lists.List) {
	<!DOCTYPE html>
	<html lang="en">
        ... header and script tags
		<body hx-ext="response-targets" hx-target-error="this">
			<main class="flex w-full flex-col">
				for _, l := range lists {
					<div class="flex flex-col gap-y-2">
						<h2 class="text-2xl font-extrabold">{ l.Name }</h2>
						<div class="flex flex-col gap-y-2">
							for _, item := range list.Items {
								<div class="flex items-center gap-x-2">
									<span>{ item }</span>
								</div>
							}
							<input id="todo-list-item" type="text"/>
							<button 
								type="button" 
								hx-post={"/list/" + l.ID} 
								hx-include="todo-list-item">
								Add Item
							</button>
						</div>
					</div>
				}
				<input id="todo-list-name" type="text"/>
				<button 
					type="button" 
					hx-post="/list" 
					hx-include="todo-list-name">
					Add List
				</button>
			</main>
		</body>
	</html>
}
`).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
