// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func header(title string) templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, inital-scale=1.0\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/layout.templ`, Line: 8, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</title><script src=\"static/script/htmx.min.js\"></script><link rel=\"stylesheet\" href=\"static/css/style.css\"><script>const whTooltips = {colorLinks: true, iconizeLinks: true, renameLinks: true};</script><script src=\"https://wow.zamimg.com/js/tooltips.js\"></script></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func footer() templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<footer class=\"bg-white dark:bg-gray-900\"><div class=\"px-4 pb-6 pt-16 sm:px-6 lg:px-8 lg:pt-24\"><div class=\"divider divider-accent\"></div><div class=\"flex justify-between\"><div class=\"text-center sm:flex sm:justify-between sm:text-left\"><p class=\"text-sm text-gray-500 dark:text-gray-400\"><span class=\"block sm:inline\">All rights reserved.</span> <a class=\"inline-block text-gray-600 underline transition hover:text-emerald-600/75 dark:text-emerald-500 dark:hover:text-emerald-500/75\" href=\"#\">Terms & Conditions</a> <span>&middot;</span> <a class=\"inline-block text-gray-600 underline transition hover:text-emerald-600/75 dark:text-emerald-500 dark:hover:text-emerald-500/75\" href=\"#\">Privacy Policy</a></p></div><div><ul class=\"flex justify-center gap-6 sm:justify-start md:gap-8\"><li><a href=\"https://github.com/Leo-Mart/Exam-Goth-Stack\" rel=\"noreferrer\" target=\"_blank\" class=\"text-gray-500 transition hover:text-emerald-900 \">Github</a></li></ul></div></div></div></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func nav(activeLink string, isLoggedIn bool) templ.Component {
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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<header class=\"bg-white dark:bg-gray-900 shadow\"><div class=\"mx-auto flex h-16 max-w-screen-xl items-center gap-8 px-4 sm:px-6 lg:px-8\"><nav class=\"flex flex-1 items-center justify-between\"><ul class=\"flex items-center gap-6 text-sm\"><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if activeLink == "/" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-emerald-600 transition\" href=\"/\">Home</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-white transition hover:bg-emerald-700 dark:hover:bg-emerald-500 dark:hover:text-white\" href=\"/\">Home</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</li><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if activeLink == "/import" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-emerald-600 transition\" href=\"/\">Import new character</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-white transition hover:bg-emerald-700 dark:hover:bg-emerald-500 dark:hover:text-white\" href=\"/import\">Import new character</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</li><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if activeLink == "/characters" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-emerald-600 transition\" href=\"/\">Imported Characters</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-white transition hover:bg-emerald-700 dark:hover:bg-emerald-500 dark:hover:text-white\" href=\"/characters\">Imported Characters</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</li><li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if activeLink == "/about" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-emerald-600 transition\" href=\"/\">About</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-white transition hover:bg-emerald-700 dark:hover:bg-emerald-500 dark:hover:text-white\" href=\"/about\">About</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</li></ul><ul class=\"flex items-center gap-6 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if isLoggedIn {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if activeLink == "/userpage" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<a id=\"sign-up-link\" class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-emerald-600 transition\" href=\"/\">Userpage</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-white transition hover:bg-emerald-700 dark:hover:bg-emerald-500 dark:hover:text-white\" href=\"/userpage\">Userpage</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "</li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "<li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if activeLink == "/sign-up" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "<a id=\"sign-up-link\" class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-emerald-600 transition\" href=\"/\">Create Account</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "<a id=\"sign-up-link\" class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-white transition hover:bg-emerald-700 dark:hover:bg-emerald-500 dark:hover:text-white\" href=\"/sign-up\">Create account</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "</li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if isLoggedIn {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "<li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if activeLink == "/login" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "<a class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-emerald-600 transition\" href=\"/\">Logout</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "<a id=\"login-link\" class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-white transition hover:bg-emerald-700 dark:hover:bg-emerald-500 dark:hover:text-white\" href=\"/logout\">Logout</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "</li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "<li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if activeLink == "/login" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "<a id=\"login-link\" class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-emerald-600 transition\" href=\"/\">Login</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "<a id=\"login-link\" class=\"block rounded-md px-5 py-2.5 text-sm font-medium text-white transition hover:bg-emerald-700 dark:hover:bg-emerald-500 dark:hover:text-white\" href=\"/login\">Login</a>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "</li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 33, "</ul></nav></div></header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func Layout(contents templ.Component, title string, activeLink string, isLoggedIn bool) templ.Component {
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
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 34, "<!doctype HTML><html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = header(title).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 35, "<body class=\"flex flex-col h-full bg-slate-700\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = nav(activeLink, isLoggedIn).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 36, "<main class=\"flex-1 \">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = contents.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 37, "</main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = footer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 38, "<script>\r\n    function refreshWowheadLinks(evt) {\r\n      window.$WowheadPower.refreshLinks()\r\n    }\r\n  </script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
