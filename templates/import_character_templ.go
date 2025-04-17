// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func ImportNewCharacter() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<h2 class=\"text-2xl font-bold text-gray-900 md:text-3xl dark:text-white\">Import new character</h2><form class=\"mt-8\" hx-post=\"/character/add\" hx-target=\"#characters\"><label for=\"name\" class=\"block overflow-hidden rounded-md border border-gray-200 px-3 py-2 shadow-sm focus-within:border-orange-600 focus-within:ring-1 focus-within:ring-orange-600 dark:border-gray-700 dark:bg-gray-800\"><span class=\"text-xs font-medium text-gray-700 dark:text-gray-200\">Name </span> <input type=\"text\" name=\"name\" placeholder=\"enter character name\" id=\"char-name\" class=\"mt-1 w-full border-none bg-transparent p-0 focus:border-transparent focus:outline-none focus:ring-0 sm:text-sm dark:text-white\"></label> <label for=\"name\" class=\"block overflow-hidden rounded-md border border-gray-200 px-3 py-2 shadow-sm focus-within:border-orange-600 focus-within:ring-1 focus-within:ring-orange-600 dark:border-gray-700 dark:bg-gray-800 mt-4\"><span class=\"text-xs font-medium text-gray-700 dark:text-gray-200\">Realm </span> <input type=\"text\" name=\"realm\" placeholder=\"enter character realm\" id=\"char-realm\" class=\"mt-1 w-full border-none bg-transparent p-0 focus:border-transparent focus:outline-none focus:ring-0 sm:text-sm dark:text-white\"></label><button type=\"submit\" class=\"block mt-4 w-full rounded-md px-5 py-2.5 text-sm font-medium text-white bg-orange-600 transition hover:bg-orange-700 dark:hover:bg-orange-500\">Add character</button></form><div id=\"characters\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
